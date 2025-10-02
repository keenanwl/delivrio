package easypostapis

import (
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/deliveryoptioneasypost"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/EasyPost/easypost-go/v4"
)

type PackageConfig struct {
	PublicOrderID      string
	DelivrioShipmentID pulid.ID
	DelivrioColliID    pulid.ID
	Packaging          *ent.Packaging
	OrderLines         []*ent.OrderLine
	// Can potentially be in the shipment level?
	ExternalComment string
}

type FetchLabelOutput struct {
	Package          PackageConfig
	ResponseShipment *easypost.Shipment
	Responseb64PDF   string
	Error            error
}

type EasyPostAPIAuth struct {
	APIKey string
	Test   bool
}

type ShipmentConfig struct {
	EasyPostAPIAuth
	ConsignorAddress   *ent.Address
	ConsigneeAddress   *ent.Address
	ParcelShopAddress  *ent.ParcelShop
	Packages           []PackageConfig
	AdditionalServices []*ent.CarrierAdditionalServiceEasyPost
	IsReturn           bool
}

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

func createShipmentConfig(ctx context.Context, deliveryOption pulid.ID, groupedSameDeliveryPacks [][]*ent.Colli) ([]ShipmentConfig, error) {

	cli := ent.FromContext(ctx)
	shipments := make([]ShipmentConfig, 0)

	additionalServices, err := cli.DeliveryOption.Query().
		Where(deliveryoption.ID(deliveryOption)).
		QueryDeliveryOptionEasyPost().
		QueryCarrierAddServEasyPost().
		All(ctx)
	if err != nil {
		return nil, err
	}

	for _, s := range groupedSameDeliveryPacks {
		shipment, err := shipmentFromGroupedParcels(
			ctx,
			s,
			additionalServices,
		)
		if err != nil {
			return nil, err
		}
		shipments = append(shipments, *shipment)
	}

	return shipments, nil
}

func shipmentFromGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.Colli,
	additionalServices []*ent.CarrierAdditionalServiceEasyPost,
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
			return nil, fmt.Errorf("easy post requires colli packaging is added")
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
			OrderLines:         orderLines,
			ExternalComment:    ord.CommentExternal,
		})
	}

	return &ShipmentConfig{
		ConsignorAddress:   firstParcel.Edges.Sender,
		ConsigneeAddress:   firstParcel.Edges.Recipient,
		ParcelShopAddress:  firstParcel.Edges.ParcelShop,
		AdditionalServices: additionalServices,
		Packages:           parcels,
	}, nil
}

func requestLabels(ctx context.Context, deliveryOption pulid.ID, shipment ShipmentConfig) ([]FetchLabelOutput, error) {
	cli := ent.FromContext(ctx)
	do, err := cli.DeliveryOptionEasyPost.Query().
		WithDeliveryOption().
		WithCarrierAddServEasyPost().
		Where(
			deliveryoptioneasypost.HasDeliveryOptionWith(
				deliveryoption.ID(deliveryOption),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	agreement, err := do.Edges.DeliveryOption.
		QueryCarrier().
		QueryCarrierEasyPost().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	service, err := do.Edges.DeliveryOption.QueryCarrierService().
		QueryCarrierServEasyPost().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	pickupDay, err := utils.PickupDate(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]FetchLabelOutput, 0)
	for _, pack := range shipment.Packages {

		carrierAccount := ""
		if len(agreement.CarrierAccounts) == 1 {
			carrierAccount = agreement.CarrierAccounts[0]
		}

		pendingLabelRequest, err := toLabelRequest(
			ctx,
			pickupDay,
			shipment,
			pack,
			do.Edges.DeliveryOption.CustomsEnabled,
			do.Edges.DeliveryOption.CustomsSigner,
			do.Edges.CarrierAddServEasyPost,
			carrierAccount,
			service.APIKey.String(),
		)
		if err != nil {
			return nil, fmt.Errorf("request labels: to request: %w", err)
		}

		requestSubscriber := easypost.RequestHookEventSubscriber{
			Callback: myRequestHookEventSubscriberCallback,
			HookEventSubscriber: easypost.HookEventSubscriber{
				ID: "my-request-hook",
			},
		}

		easyPostClient := easypost.New(agreement.APIKey)
		easyPostClient.Hooks.AddRequestEventSubscriber(requestSubscriber)
		resp, err := easyPostClient.CreateShipment(pendingLabelRequest)
		if err != nil {
			// Perhaps remove this check? so other labels can process?
			return nil, fmt.Errorf("request labels: %w", err)
		}

		pdfLabel, err := fetchPDFURL(resp.PostageLabel.LabelPDFURL)
		if err != nil {
			// Perhaps remove this check? so other labels can process?
			return nil, fmt.Errorf("request labels: %w", err)
		}

		output = append(output, FetchLabelOutput{
			Package:          pack,
			Responseb64PDF:   utils.String2Base64(pdfLabel),
			ResponseShipment: resp,
			Error:            err,
		})
	}

	return output, nil

}

func myRequestHookEventSubscriberCallback(ctx context.Context, event easypost.RequestHookEvent) error {
	bod, _ := io.ReadAll(event.RequestBody)
	fmt.Printf("Making HTTP call to %s\n", event.Url)
	fmt.Printf("Making HTTP call to %s\n", string(bod))
	return nil
}

func fetchPDFURL(url string) ([]byte, error) {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := cli.Get(url)
	if err != nil {
		return nil, fmt.Errorf("easypost: fetch PDF from URL: %w", err)
	}
	defer resp.Body.Close()

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("easypost: read PDF body: %w", err)
	}

	return bod, nil
}

func toEasyPostOptions(pickup time.Time, machinable bool, customTxt string, additionalServices []*ent.CarrierAdditionalServiceEasyPost) *easypost.ShipmentOptions {
	dt := easypost.DateTimeFromTime(pickup)
	output := &easypost.ShipmentOptions{
		LabelDate: &dt,
		// Default to true until packaging updated
		Machinable:   true,
		LabelFormat:  "PDF",
		LabelSize:    "4x6",
		PrintCustom1: customTxt,
	}
	for _, additionalService := range additionalServices {
		switch additionalService.APIKey {
		case "delivery_confirmation":
			output.DeliveryConfirmation = additionalService.APIValue
			break
		case "address_validation_level":
			output.AddressValidationLevel = additionalService.APIValue
			break
		case "endorsement":
			output.Endorsement = additionalService.APIValue
			break
		}
	}
	return output
}

func toLabelRequest(
	ctx context.Context,
	pickup time.Time,
	shipment ShipmentConfig,
	parcel PackageConfig,
	includeCustoms bool,
	customsSigner string,
	additionalServices []*ent.CarrierAdditionalServiceEasyPost,
	carrierAccountID string,
	carrierService string,
) (*easypost.Shipment, error) {

	output := &easypost.Shipment{
		// Rate and buy in one go
		CarrierAccountIDs: []string{carrierAccountID},
		Service:           carrierService,
		ToAddress:         &easypost.Address{},
		FromAddress:       &easypost.Address{},
		Parcel:            &easypost.Parcel{},
		Reference:         parcel.PublicOrderID,
		Options:           toEasyPostOptions(pickup, true, parcel.ExternalComment, additionalServices),
	}

	from := shipment.ConsignorAddress
	fromCountry, err := from.QueryCountry().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	output.FromAddress = &easypost.Address{
		Street1: from.AddressOne,
		Street2: from.AddressTwo,
		City:    from.City,
		State:   strings.ToUpper(from.State),
		Zip:     from.Zip,
		Name:    strings.TrimSpace(fmt.Sprintf("%s %s", from.FirstName, from.LastName)),
		Company: from.Company,
		Phone:   from.PhoneNumber,
		Email:   from.Email,
		Country: fromCountry.Alpha2,
	}

	to := shipment.ConsigneeAddress
	toCountry, err := from.QueryCountry().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	output.ToAddress = &easypost.Address{
		Street1: to.AddressOne,
		Street2: to.AddressTwo,
		City:    to.City,
		State:   strings.ToUpper(to.State),
		Zip:     to.Zip,
		Name:    fmt.Sprintf("%s %s", to.FirstName, to.LastName),
		Company: to.Company,
		Phone:   to.PhoneNumber,
		Email:   to.Email,
		Country: toCountry.Alpha2,
	}

	if includeCustoms {
		items := make([]*easypost.CustomsItem, 0)

		for _, item := range parcel.OrderLines {
			product, err := item.QueryProductVariant().
				WithInventoryItem().
				WithProduct().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			hsCode, originCountry, err := utils.HsCode(ctx, product.Edges.InventoryItem, toCountry.ID)
			if err != nil {
				return nil, err
			}

			cur, err := item.QueryCurrency().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			orderLineOunces, err := utils.OuncesFromOrderLines(ctx, []*ent.OrderLine{item})
			if err != nil {
				return nil, err
			}

			items = append(items, &easypost.CustomsItem{
				Description:    fmt.Sprintf("%s: %s", product.Edges.Product.Title, product.Description),
				Quantity:       float64(item.Units),
				Value:          item.UnitPrice * float64(item.Units),
				Weight:         orderLineOunces,
				HSTariffNumber: hsCode,
				Code:           "",
				OriginCountry:  originCountry,
				Currency:       cur.CurrencyCode.String(),
			})
		}
		output.CustomsInfo = &easypost.CustomsInfo{
			// If the value of the goods is less than $2,500, then you pass the following EEL code: "NOEEI 30.37(a)"
			EELPFC:       "NOEEI 30.37(a)",
			ContentsType: "merchandise",
			// Only used with type = "Other"
			ContentsExplanation: "",
			CustomsCertify:      true,
			CustomsSigner:       customsSigner,
			NonDeliveryOption:   "return",
			RestrictionType:     "none",
			CustomsItems:        items,
		}
	}

	colliOunces, err := utils.OuncesFromOrderLines(ctx, parcel.OrderLines)
	if err != nil {
		return nil, err
	}

	output.Parcel = &easypost.Parcel{
		Weight: colliOunces,
		Length: utils.CmToInchesFloat(parcel.Packaging.LengthCm),
		Height: utils.CmToInchesFloat(parcel.Packaging.HeightCm),
		Width:  utils.CmToInchesFloat(parcel.Packaging.WidthCm),
	}

	return output, nil

}
