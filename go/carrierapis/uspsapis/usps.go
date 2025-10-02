package uspsapis

import (
	"bytes"
	"context"
	"delivrio.io/go/carrierapis/common"
	uspsrequest2 "delivrio.io/go/carrierapis/uspsapis/uspsrequest"
	uspsresponse2 "delivrio.io/go/carrierapis/uspsapis/uspsresponse"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierserviceusps"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/deliveryoptionusps"
	"delivrio.io/go/ent/packaginguspsprocessingcategory"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PackageConfig struct {
	PublicOrderID      string
	DelivrioShipmentID pulid.ID
	DelivrioColliID    pulid.ID
	Packaging          *ent.Packaging
	Items              []*ent.OrderLine
}

type USPSAPIAuth struct {
	ConsumerKey      string
	ConsumerSecret   string
	MID              string
	CRID             string
	EPSAccountNumber string
}

type ShipmentConfig struct {
	USPSAPIAuth
	ConsignorAddress   *ent.Address
	ConsigneeAddress   *ent.Address
	ParcelShopAddress  *ent.ParcelShop
	Packages           []PackageConfig
	AdditionalServices []AdditionalService
	IsReturn           bool
}

type FetchLabelOutput struct {
	Package          PackageConfig
	ResponseMetadata *uspsresponse2.Label
	Responseb64PDF   string
	Error            error
}

// The APIs are a little weird for now since I started with a baseline
// of how PostNord does things, which is a semi-convoluted way to represent
// multiple parcels in a single shipment. Where processing the output requires
// tracing back through the input.

func FetchLabels(ctx context.Context, deliveryOption pulid.ID, collis []*ent.Colli) ([]FetchLabelOutput, error) {
	shipments, err := common.GroupPackagesBySenderReceiver(ctx, collis)
	if err != nil {
		return nil, err
	}

	shipmentConfigs, err := createShipmentConfig(ctx, deliveryOption, shipments)
	if err != nil {
		return nil, err
	}

	output := make([]FetchLabelOutput, 0)

	for _, s := range shipmentConfigs {
		// TODO: request labels outside this loop to re-use token (perf)
		allLabels, err := requestLabels(ctx, deliveryOption, s)
		// TODO: Potentially continue on err?
		if err != nil {
			return nil, err
		}

		output = append(output, allLabels...)
	}

	return output, nil
}

func SaveLabelData(ctx context.Context, resp FetchLabelOutput) (*common.CreateShipment, error) {

	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	createShip := tx.ShipmentUSPS.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetTenantID(view.TenantID()).
		SetTrackingNumber(resp.ResponseMetadata.TrackingNumber).
		SetPostage(resp.ResponseMetadata.Postage)

	expectedDelivery, err := time.Parse(time.DateOnly, resp.ResponseMetadata.Commitment.ScheduleDeliveryDate)
	if err == nil {
		createShip = createShip.SetScheduledDeliveryDate(expectedDelivery)
	}

	err = createShip.Exec(ctx)
	if err != nil {
		return nil, err
	}

	sp, err := tx.ShipmentParcel.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetStatus(shipmentparcel.StatusPending).
		SetItemID(resp.ResponseMetadata.TrackingNumber).
		SetColliID(resp.Package.DelivrioColliID).
		SetTenantID(view.TenantID()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	_, err = utils.CreateShipmentDocument(ctx, sp, &resp.Responseb64PDF, nil)
	if err != nil {
		return nil, err
	}

	return &common.CreateShipment{
		Shipment: resp.Package.DelivrioShipmentID,
		Labels:   []string{resp.Responseb64PDF},
	}, nil
}

func createShipmentConfig(ctx context.Context, deliveryOption pulid.ID, groupedSameDeliveryPacks [][]*ent.Colli) ([]ShipmentConfig, error) {

	shipments := make([]ShipmentConfig, 0)

	for _, s := range groupedSameDeliveryPacks {
		shipment, err := shipmentFromGroupedParcels(
			ctx,
			s,
			//additionalServices,
		)
		if err != nil {
			return nil, err
		}
		shipments = append(shipments, *shipment)
	}

	return shipments, nil
}

// Switch out with test vs production
const baseURLTest = "https://api-cat.usps.com"

// Set both to test in case of accidental toggle during testing
const baseURL = "https://api.usps.com"

type BasicAuthRoundTripper struct {
	Transport    http.RoundTripper
	ClientID     string
	ClientSecret string
}

// Add the required USPS basic auth header to the initial
// oauth2 request
func (rt *BasicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	authString := fmt.Sprintf("%s:%s", rt.ClientID, rt.ClientSecret)
	encodedAuthString := base64.StdEncoding.EncodeToString([]byte(authString))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", encodedAuthString))
	fmt.Println("Request Method:", req.Method)
	fmt.Println("Request URL:", req.URL.String())
	fmt.Println("Request Headers:", req.Header)
	if req.Body != nil {
		body, _ := io.ReadAll(req.Body)
		fmt.Println("Request Body:", string(body))
	}
	resp, err := rt.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Log the response data
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

	// Re-assign the response body to the original value for further reading
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return resp, nil
}

func requestLabels(ctx context.Context, deliveryOption pulid.ID, shipment ShipmentConfig) ([]FetchLabelOutput, error) {
	cli := ent.FromContext(ctx)
	do, err := cli.DeliveryOptionUSPS.Query().
		WithDeliveryOption().
		WithCarrierAdditionalServiceUSPS().
		Where(
			deliveryoptionusps.HasDeliveryOptionWith(
				deliveryoption.ID(deliveryOption),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	agreement, err := do.Edges.DeliveryOption.
		QueryCarrier().
		QueryCarrierUSPS().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	toggledURL := baseURLTest
	if !agreement.IsTestAPI {
		toggledURL = baseURL
	}

	labelURLPath := "/labels/v3/label"
	if shipment.IsReturn {
		labelURLPath = "/labels/v3/return-label"
	}

	labelURL, err := url.JoinPath(toggledURL, labelURLPath)
	if err != nil {
		return nil, err
	}

	tokenURL, err := url.JoinPath(toggledURL, "/oauth2/v3/token")
	if err != nil {
		return nil, err
	}

	config := &clientcredentials.Config{
		ClientID:     agreement.ConsumerKey,
		ClientSecret: agreement.ConsumerSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"payments", "labels"},
		// I think? Check here if not working.
		AuthStyle: oauth2.AuthStyleInParams,
	}

	cliWithBasicAuth := &http.Client{
		Transport: &BasicAuthRoundTripper{
			Transport:    http.DefaultTransport,
			ClientID:     agreement.ConsumerKey,
			ClientSecret: agreement.ConsumerSecret,
		},
	}
	cliWithLogger := &http.Client{
		Transport: &common.LoggerAuthRoundTripper{
			Transport: http.DefaultTransport,
		},
	}

	ctxWithBasic := context.WithValue(ctx, oauth2.HTTPClient, cliWithBasicAuth)
	ctxWithLogger := context.WithValue(ctx, oauth2.HTTPClient, cliWithLogger)
	ts := config.TokenSource(ctxWithBasic)
	uspsClient := oauth2.NewClient(ctxWithLogger, ts)

	token, err := paymentTokenFromTestCredentials(ctx, uspsClient, agreement, toggledURL)
	if err != nil {
		return nil, fmt.Errorf("request labels: %w", err)
	}

	selectedService, err := do.Edges.DeliveryOption.
		QueryCarrierService().
		QueryCarrierServiceUSPS().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("request labels: service: %w", err)
	}

	pickupDay, err := utils.PickupDate(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]FetchLabelOutput, 0)
	for _, pack := range shipment.Packages {

		uspsPacking, err := pack.Packaging.QueryPackagingUSPS().
			WithPackagingUSPSProcessingCategory().
			WithPackagingUSPSRateIndicator().
			Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("request labels: %w", err)
		}

		pendingLabelRequest, err := toUSPSLabelRequest(
			ctx,
			pickupDay,
			shipment,
			pack,
			selectedService.APIKey,
			uspsPacking.Edges.PackagingUSPSRateIndicator,
			uspsPacking.Edges.PackagingUSPSProcessingCategory.ProcessingCategory,
		)
		if err != nil {
			return nil, fmt.Errorf("request labels: to request: %w", err)
		}

		resp, b64PDF, err := labelRequest(uspsClient, token, labelURL, pendingLabelRequest)
		if err != nil {
			// Perhaps remove this check? so other labels can process?
			return nil, fmt.Errorf("request labels: %w", err)
		}
		output = append(output, FetchLabelOutput{
			Package:          pack,
			ResponseMetadata: resp,
			Responseb64PDF:   b64PDF,
			Error:            err,
		})
	}

	return output, nil

}

func toUSPSLabelRequest(ctx context.Context, pickup time.Time, shipment ShipmentConfig, parcel PackageConfig, mailClass carrierserviceusps.APIKey, rateIndicator *ent.PackagingUSPSRateIndicator, processingCategory packaginguspsprocessingcategory.ProcessingCategory) (*uspsrequest2.Label, error) {

	output := &uspsrequest2.Label{
		ImageInfo: uspsrequest2.ImageInfo{
			ImageType:     "PDF",
			LabelType:     "4X6LABEL",
			ShipInfo:      false,
			ReceiptOption: "NONE",
			// Don't print postage cost on label
			SuppressPostage:  true,
			SuppressMailDate: false,
			// Also print a return label
			ReturnLabel: false,
		},
		ToAddress:          uspsrequest2.Address{},
		FromAddress:        uspsrequest2.Address{},
		PackageDescription: uspsrequest2.PackageDescription{},
	}

	from := shipment.ConsignorAddress

	output.FromAddress = uspsrequest2.Address{
		StreetAddress:    from.AddressOne,
		SecondaryAddress: from.AddressTwo,
		City:             from.City,
		State:            strings.ToUpper(from.State),
		ZIPCode:          from.Zip,
		Urbanization:     "",
		FirstName:        shipment.ConsignorAddress.FirstName,
		LastName:         shipment.ConsignorAddress.LastName,
		Firm:             from.Company,
		Phone:            shipment.ConsignorAddress.PhoneNumber,
		IgnoreBadAddress: false,
		// This is a third party unique identifier for the sender used for correlation purposes.
		PlatformUserID: nil,
		Email:          &shipment.ConsignorAddress.Email,
		// Only for recipient; DELIVRIO doesn't support for now
		ParcelLockerDelivery: nil,
		FacilityID:           nil,
	}

	to := shipment.ConsigneeAddress

	output.ToAddress = uspsrequest2.Address{
		StreetAddress:    to.AddressOne,
		SecondaryAddress: to.AddressTwo,
		City:             to.City,
		State:            strings.ToUpper(to.State),
		ZIPCode:          to.Zip,
		FirstName:        shipment.ConsigneeAddress.FirstName,
		LastName:         shipment.ConsigneeAddress.LastName,
		Firm:             to.Company,
		Phone:            shipment.ConsigneeAddress.PhoneNumber,
		IgnoreBadAddress: false,
		// This is a third party unique identifier for the sender used for correlation purposes.
		PlatformUserID: nil,
		Email:          &shipment.ConsigneeAddress.Email,
		// Only for recipient; DELIVRIO doesn't support for now
		ParcelLockerDelivery: nil,
		FacilityID:           nil,
	}

	pounds, err := utils.PoundsFromOrderLines(ctx, parcel.Items)
	if err != nil {
		return nil, fmt.Errorf("request labels: %w", err)
	}

	output.PackageDescription = uspsrequest2.PackageDescription{
		WeightUOM:     "lb",
		Weight:        pounds,
		DimensionsUOM: "in",
		Length:        utils.CmToInches(parcel.Packaging.LengthCm),
		Height:        utils.CmToInches(parcel.Packaging.HeightCm),
		Width:         utils.CmToInches(parcel.Packaging.WidthCm),
		MailClass:     mailClass,
		// Own packaging, needs to be configurable to support flat rate boxes, etc
		RateIndicator: rateIndicator.Code,
		// https://pe.usps.com/text/dmm300/201.htm#ep1097220
		ProcessingCategory:           processingCategory,
		DestinationEntryFacilityType: "NONE",
		CustomerReference: []uspsrequest2.CustomerReference{
			{
				ReferenceNumber:      parcel.PublicOrderID,
				PrintReferenceNumber: true,
			},
		},
		ExtraServices:  []int64{},
		MailingDate:    pickup.Format(time.DateOnly),
		CarrierRelease: false,
	}

	return output, nil

}

func labelRequest(client *http.Client, token string, requestURL string, requestBody *uspsrequest2.Label) (*uspsresponse2.Label, string, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(requestBody)
	if err != nil {
		return nil, "", err
	}

	req, err := http.NewRequest(http.MethodPost, requestURL, &buf)
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Payment-Authorization-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		contentType := resp.Header.Get("Content-Type")
		_, params, _ := mime.ParseMediaType(contentType)
		// Create a multipart reader to read the response body
		mr := multipart.NewReader(resp.Body, params["boundary"])

		var respLabel uspsresponse2.Label
		b64PDF := ""
		count := 0
		// Assume there are only 2 parts..
		for {

			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, "", fmt.Errorf("error reading part: %w", err)
			}

			respBody, err := io.ReadAll(part)
			if err != nil {
				return nil, "", fmt.Errorf("error reading part body: %w", err)
			}

			if count == 0 {
				err := json.Unmarshal(respBody, &respLabel)
				if err != nil {
					return nil, "", err
				}
			} else {
				b64PDF = string(respBody)
			}
			count++
		}

		return &respLabel, b64PDF, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return nil, "", fmt.Errorf("fetching USPS label: response status: %v: %v", resp.StatusCode, string(body))

}

func paymentTokenFromTestCredentials(ctx context.Context, client *http.Client, agreement *ent.CarrierUSPS, selectedBaseURL string) (string, error) {
	payment := uspsrequest2.Payment{Roles: make([]uspsrequest2.Role, 0)}
	payment.Roles = append(payment.Roles, uspsrequest2.Role{
		RoleName:      "SHIPPER",
		Crid:          agreement.Crid,
		Mid:           agreement.Mid,
		ManifestMID:   agreement.ManifestMid,
		AccountType:   "EPS",
		AccountNumber: agreement.EpsAccountNumber,
	})
	payment.Roles = append(payment.Roles, uspsrequest2.Role{
		RoleName:      "LABEL_OWNER",
		Crid:          agreement.Crid,
		Mid:           agreement.Mid,
		ManifestMID:   agreement.ManifestMid,
		AccountType:   "EPS",
		AccountNumber: agreement.EpsAccountNumber,
	})
	payment.Roles = append(payment.Roles, uspsrequest2.Role{
		RoleName:      "PAYER",
		Crid:          agreement.Crid,
		Mid:           agreement.Mid,
		ManifestMID:   agreement.ManifestMid,
		AccountType:   "EPS",
		AccountNumber: agreement.EpsAccountNumber,
	})

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payment)
	if err != nil {
		return "", err
	}

	paymentURL, err := url.JoinPath(selectedBaseURL, "/payments/v3/payment-authorization")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, paymentURL, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusOK {
		var respToken uspsresponse2.Payment
		err := json.Unmarshal(body, &respToken)
		if err != nil {
			return "", err
		}

		return respToken.PaymentAuthorizationToken, nil
	}

	return "", fmt.Errorf("fetching USPS token: response status: %v: %v", resp.StatusCode, string(body))
}

func shipmentFromGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.Colli,
	// additionalServices []AdditionalService,
) (*ShipmentConfig, error) {
	cli := ent.FromContext(ctx)
	v := viewer.FromContext(ctx)

	firstParcel := groupedParcels[0]

	parcels := make([]PackageConfig, 0)
	for _, p := range groupedParcels {

		agreement, err := p.QueryDeliveryOption().
			QueryCarrier().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipment, err := cli.Shipment.Create().
			SetCarrier(agreement).
			SetShipmentPublicID(fmt.Sprintf("%v", time.Now())).
			SetTenantID(v.TenantID()).
			SetStatus(shipment2.StatusPending).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		packaging, err := common.ColliPackaging(ctx, p)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if ent.IsNotFound(err) {
			return nil, fmt.Errorf("usps requires colli packaging is added")
		}

		ord, err := p.QueryOrder().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		orderLines, err := p.QueryOrderLines().
			All(ctx)
		if err != nil {
			return nil, err
		}
		parcels = append(parcels, PackageConfig{
			PublicOrderID:      ord.OrderPublicID,
			DelivrioShipmentID: shipment.ID,
			DelivrioColliID:    p.ID,
			Packaging:          packaging,
			Items:              orderLines,
		})
	}

	return &ShipmentConfig{
		ConsignorAddress:  firstParcel.Edges.Sender,
		ConsigneeAddress:  firstParcel.Edges.Recipient,
		ParcelShopAddress: firstParcel.Edges.ParcelShop,
		//AdditionalServices: additionalServices,
		Packages: parcels,
	}, nil
}
