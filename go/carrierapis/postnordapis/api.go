package postnordapis

import (
	"bytes"
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/carrierapis/postnordapis/postnordrequest"
	"delivrio.io/go/ent/returncolli"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/utils"
	"delivrio.io/shared-utils/pulid"
)

const timeout = 30 * time.Second
const postNordTimestampFormat = time.RFC3339Nano
const delivrioPostNordAppID = 1438
const qrCodeKey = "C2"

type RequestConfig struct {
	RequestQRCode bool
	PartyID       string
	// Max 100 labels, needs to be handled
	Shipments      []ShipmentConfig
	PostNordAPIKey string
}

type ShipmentConfig struct {
	PartyID            string
	PostNordAPIKey     string
	ConsignorAddress   *ent.Address
	ConsigneeAddress   *ent.Address
	ParcelShopAddress  *ent.ParcelShop
	BasicServices      string
	AdditionalServices []string
	Packages           []PackageConfig
}

type PackageConfig struct {
	DelivrioShipmentID pulid.ID
	DelivrioColliID    pulid.ID
	Items              []*ent.OrderLine
}

func FetchLabelPostNord(ctx context.Context, o common.ReturnOrderDeliveryOptionsColliIDs) (*PostNordLabelResponseData, error) {
	db := ent.FromContext(ctx)

	allReturnCollis, err := db.ReturnColli.Query().
		WithSender().
		WithRecipient().
		Where(returncolli.ID(o.ReturnColliID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	grouped, err := common.GroupReturnPackages(ctx, allReturnCollis)
	if err != nil {
		return nil, fmt.Errorf("group return packages: %w", err)
	}

	shipment, err := createReturnShipmentConfigs(ctx, o.DeliveryOptionID, grouped)
	if err != nil {
		return nil, fmt.Errorf("create return shipment config: %w", err)
	}

	response, err := requestLabels(ctx, o.DeliveryOptionID, shipment, false)
	if err != nil {
		return nil, fmt.Errorf("request label: %w", err)
	}

	responseCheck, err := checkResponse(shipment, *response)
	if err != nil {
		return nil, fmt.Errorf("check response: %w", err)
	}

	if len(responseCheck) == 0 {
		return nil, fmt.Errorf("return update: expected label response")
	}

	return &responseCheck[0], nil
}

func fireRequest(req *http.Request) (*http.Response, error) {
	// TODO: move up call chain to enable keep-alive if required
	client := http.Client{
		Timeout: timeout,
	}
	return client.Do(req)
}

func generateRequest(ctx context.Context, u *url.URL, payload interface{}) (*http.Request, error) {
	requestJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestJSON))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func generateV3LabelPayload(ctx context.Context, config RequestConfig) (*url.URL, *postnordrequest.Label, error) {

	// TODO: production URL? Pass as config
	baseURL, err := url.Parse("https://atapi2.postnord.com/rest/shipment/v3/edi/labels/pdf")
	if err != nil {
		return nil, nil, err
	}

	queryParams := baseURL.Query()

	queryParams.Set("apikey", config.PostNordAPIKey)
	queryParams.Set("multiZPL", "false")
	queryParams.Set("cutter", "false")
	queryParams.Set("dpmm", "8dpmm")
	queryParams.Set("fonts", "E:ARI000.TTF, E:ARIAL.TTF, E:T0003M_")
	queryParams.Set("labelType", "standard")
	queryParams.Set("labelLength", "190")
	queryParams.Set("pnInfoText", "false")
	queryParams.Set("labelsPerPage", "100")
	queryParams.Set("page", "1")
	queryParams.Set("processOffline", "false")
	queryParams.Set("storeLabel", "false")
	queryParams.Set("paperSize", "LABEL")
	if config.RequestQRCode {
		queryParams.Set("generateQrcodeImage", "true")
	}

	baseURL.RawQuery = queryParams.Encode()

	shipments := make([]postnordrequest.Shipment, 0)
	for _, s := range config.Shipments {
		shipment, err := shipmentFromDelivrioPackage(ctx, config.PartyID, s)
		if err != nil {
			return nil, nil, err
		}
		shipments = append(shipments, *shipment)
	}

	payload := &postnordrequest.Label{
		MessageDate: time.Now().Format(postNordTimestampFormat),
		Application: postnordrequest.Application{
			ApplicationID: delivrioPostNordAppID,
			Name:          appconfig.AppName,
			Version:       appconfig.AppVersion,
		},
		// Constant?
		UpdateIndicator: "Original",
		Shipment:        shipments,
	}

	return baseURL, payload, nil

}

func postNordItemsFromDelivrioProducts(ctx context.Context, lines []*ent.OrderLine) (float64, *postnordrequest.GoodsItem, error) {
	totalWeightKG, err := common.ColliWeightKG(ctx, lines)
	if err != nil {
		return 0, nil, err
	}

	goodItem := &postnordrequest.GoodsItem{
		PackageTypeCode: "PC",
		// Only send 1 package per shipment, so OrderLines always = 1?
		Items: []postnordrequest.Item{
			{
				// For pallets? Is required, but static for now
				ItemIdentification: postnordrequest.ItemIdentification{
					ItemID:     "0",
					ItemIDType: "SCSS",
				},
				GrossWeight: postnordrequest.GrossWeight{
					Value: totalWeightKG,
					Unit:  "KGM",
				},
			},
		},
	}

	return totalWeightKG, goodItem, nil

}

func postNordDeliveryPartyFromDelivrioParcelShop(ctx context.Context, parcelShop *ent.ParcelShop) (*postnordrequest.DeliveryParty, error) {

	pnParcelShop, err := parcelShop.ParcelShopPostNord(ctx)
	if err != nil {
		return nil, err
	}

	if pnParcelShop == nil {
		return nil, fmt.Errorf("expected ParcelShop entity to have exactly 1 ParcelShopPostNord")
	}

	adr, err := pnParcelShop.AddressDelivery(ctx)
	if err != nil {
		return nil, err
	}

	// Refactor Delivrio to match PN slice?
	streets := make([]string, 0)
	if len(adr.AddressOne) > 0 {
		streets = append(streets, adr.AddressOne)
	}

	if len(adr.AddressTwo) > 0 {
		streets = append(streets, adr.AddressTwo)
	}

	country, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &postnordrequest.DeliveryParty{
		PartyIdentification: postnordrequest.PartyIdentification{
			PartyID:     pnParcelShop.ServicePointID,
			PartyIDType: "156",
		},
		Party: postnordrequest.ConsignorParty{
			NameIdentification: postnordrequest.NameIdentification{
				Name: &parcelShop.Name,
			},
			Address: postnordrequest.Address{
				Streets:     streets,
				PostalCode:  adr.Zip,
				City:        adr.City,
				CountryCode: country.Alpha2,
			},
		},
	}, nil

}

func postNordConsignorFromDelivrioPackageSender(ctx context.Context, partyID string, sender *ent.Address) (*postnordrequest.Consignor, error) {
	streets := make([]string, 0)
	if len(sender.AddressOne) > 0 {
		streets = append(streets, sender.AddressOne)
	}

	if len(sender.AddressTwo) > 0 {
		streets = append(streets, sender.AddressTwo)
	}

	country, err := sender.Country(ctx)
	if err != nil {
		return nil, err
	}

	// Max 60 and not empty string
	var name *string
	if len(sender.FirstName) > 0 || len(sender.LastName) > 0 {
		n := strings.TrimSpace(utils.TruncateString(
			fmt.Sprintf("%s %s", sender.FirstName, sender.LastName),
			60,
		))
		name = &n
	}

	// Max 60 and not empty string (assumed from name)
	var companyName *string
	if len(sender.Company) > 0 {
		n := strings.TrimSpace(utils.TruncateString(fmt.Sprintf("%s", sender.Company), 60))
		companyName = &n
	}

	return &postnordrequest.Consignor{
		// Only DK issuers for now
		IssuerCode: "Z11",
		PartyIdentification: postnordrequest.PartyIdentification{
			PartyID:     partyID,
			PartyIDType: "160",
		},
		Party: postnordrequest.ConsignorParty{
			NameIdentification: postnordrequest.NameIdentification{
				Name:        name,
				CompanyName: companyName,
			},
			Address: postnordrequest.Address{
				Streets:     streets,
				PostalCode:  sender.Zip,
				City:        sender.City,
				CountryCode: country.Alpha2,
			},
		},
	}, nil
}

func postNordConsigneeFromDelivrioRecipient(ctx context.Context, recipient *ent.Address) (*postnordrequest.Consignee, error) {

	// Refactor Delivrio to match PN slice?
	// TODO: test max lines: result 15 and still working..But only 2 shown on label
	streets := make([]string, 0)
	if len(recipient.AddressOne) > 0 {
		streets = append(streets, recipient.AddressOne)
	}

	if len(recipient.AddressTwo) > 0 {
		streets = append(streets, recipient.AddressTwo)
	}

	country, err := recipient.Country(ctx)
	if err != nil {
		return nil, err
	}

	var name *string
	if len(recipient.FirstName) > 0 || len(recipient.LastName) > 0 {
		n := strings.TrimSpace(utils.TruncateString(
			fmt.Sprintf("%s %s", recipient.FirstName, recipient.LastName),
			60,
		))
		name = &n
	}

	var companyName *string
	if len(recipient.Company) > 0 {
		n := strings.TrimSpace(utils.TruncateString(fmt.Sprintf("%s", recipient.Company), 60))
		companyName = &n
	}

	return &postnordrequest.Consignee{
		Party: postnordrequest.ConsigneeParty{
			NameIdentification: postnordrequest.NameIdentification{
				Name:        name,
				CompanyName: companyName,
			},
			Address: postnordrequest.Address{
				Streets:     streets,
				PostalCode:  recipient.Zip,
				City:        recipient.City,
				CountryCode: country.Alpha2,
			},
			Contact: postnordrequest.Contact{
				ContactName:  name,
				EmailAddress: &recipient.Email,
				SMSNo:        &recipient.PhoneNumber,
			},
		},
	}, nil
}

func packagesFromDelivrioPackages(ctx context.Context, config ShipmentConfig) (float64, []*postnordrequest.GoodsItem, error) {
	var grossWeight float64 = 0
	shipmentPackages := make([]*postnordrequest.GoodsItem, 0)
	for _, p := range config.Packages {
		totalWeightKG, goodsItems, err := postNordItemsFromDelivrioProducts(ctx, p.Items)
		if err != nil {
			return 0, nil, err
		}
		shipmentPackages = append(shipmentPackages, goodsItems)
		grossWeight += totalWeightKG
	}
	return grossWeight, shipmentPackages, nil
}

// TODO: add unit tests
func shipmentFromDelivrioPackage(ctx context.Context, partyID string, config ShipmentConfig) (*postnordrequest.Shipment, error) {
	pickupDate, err := utils.PickupDate(ctx)
	if err != nil {
		return nil, err
	}

	totalWeightKG, goodsItems, err := packagesFromDelivrioPackages(ctx, config)
	if err != nil {
		return nil, err
	}
	var parcelShop *postnordrequest.DeliveryParty
	if config.ParcelShopAddress != nil {
		parcelShop, err = postNordDeliveryPartyFromDelivrioParcelShop(ctx, config.ParcelShopAddress)
		if err != nil {
			return nil, err
		}
	}

	consignor, err := postNordConsignorFromDelivrioPackageSender(ctx, partyID, config.ConsignorAddress)
	if err != nil {
		return nil, err
	}

	consignee, err := postNordConsigneeFromDelivrioRecipient(ctx, config.ConsigneeAddress)
	if err != nil {
		return nil, err
	}

	return &postnordrequest.Shipment{
		ShipmentIdentification: postnordrequest.ShipmentIdentification{
			// TODO give this an ID
			ShipmentID: "0",
		},
		DateAndTimes: postnordrequest.DateAndTimes{
			LoadingDate: pickupDate.Format(postNordTimestampFormat),
		},
		Service: postnordrequest.Service{
			BasicServiceCode:      config.BasicServices,
			AdditionalServiceCode: config.AdditionalServices,
		},
		FreeText: make([]string, 0),
		NumberOfPackages: postnordrequest.NumberOfPackages{
			// Needs to be more? Or do we just create multiple shipments?
			Value: int64(len(goodsItems)),
		},
		TotalGrossWeight: postnordrequest.GrossWeight{
			Value: totalWeightKG,
			Unit:  "KGM",
		},
		Parties: postnordrequest.Parties{
			Consignor:     *consignor,
			Consignee:     *consignee,
			DeliveryParty: parcelShop,
		},
		GoodsItem: goodsItems,
	}, nil

}
