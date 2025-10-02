package postnordapis

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/carrierapis/postnordapis/postnordresponse"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httputil"
	"strings"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/deliveryoptionpostnord"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("labels: may not set config twice")
	}
	conf = c
	confSet = true
}

type CreateShipment struct {
	Shipment pulid.ID
	Labels   []string
}

func FetchLabels(ctx context.Context, deliveryOption pulid.ID, collis []*ent.Colli) ([]PostNordLabelResponseData, error) {
	shipments, err := common.GroupPackagesBySenderReceiver(ctx, collis)
	if err != nil {
		return nil, err
	}

	formattedShipments, err := createShipmentConfigs(ctx, deliveryOption, shipments)
	if err != nil {
		return nil, err
	}

	response, err := requestLabels(ctx, deliveryOption, formattedShipments, false)
	if err != nil {
		return nil, err
	}

	return checkResponse(formattedShipments, *response)

}

func FetchReturnLabels(ctx context.Context, deliveryOption pulid.ID, collis []*ent.ReturnColli) ([]PostNordLabelResponseData, error) {
	shipments, err := common.GroupReturnPackages(ctx, collis)
	if err != nil {
		return nil, err
	}

	formattedShipments, err := createReturnShipmentConfigs(ctx, deliveryOption, shipments)
	if err != nil {
		return nil, err
	}

	response, err := requestLabels(ctx, deliveryOption, formattedShipments, true)
	if err != nil {
		return nil, err
	}

	return checkResponse(formattedShipments, *response)
}

func createShipmentConfigs(ctx context.Context, deliveryOption pulid.ID, groupedSameDeliveryPacks [][]*ent.Colli) ([]ShipmentConfig, error) {
	tx := ent.FromContext(ctx)

	// Query is duplicated
	do, err := tx.DeliveryOptionPostNord.Query().
		WithDeliveryOption().
		WithCarrierAddServPostNord().
		Where(
			deliveryoptionpostnord.HasDeliveryOptionWith(
				deliveryoption.ID(deliveryOption),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	additionalServiceIDs := make([]string, 0)
	for _, ad := range do.Edges.CarrierAddServPostNord {
		additionalServiceIDs = append(additionalServiceIDs, ad.APICode)
	}

	carrierServicePostNord, err := do.Edges.DeliveryOption.QueryCarrierService().
		QueryCarrierServicePostNord().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	shipments := make([]ShipmentConfig, 0)
	for _, s := range groupedSameDeliveryPacks {
		shipment, err := ShipmentConfigFromGroupedParcels(
			ctx,
			s,
			carrierServicePostNord.APICode,
			additionalServiceIDs,
		)
		if err != nil {
			return nil, err
		}
		shipments = append(shipments, *shipment)
	}

	return shipments, nil
}

func createReturnShipmentConfigs(ctx context.Context, deliveryOption pulid.ID, groupedSameDeliveryPacks [][]*ent.ReturnColli) ([]ShipmentConfig, error) {
	cli := ent.FromContext(ctx)

	// Query is duplicated
	do, err := cli.DeliveryOptionPostNord.Query().
		WithDeliveryOption().
		WithCarrierAddServPostNord().
		Where(
			deliveryoptionpostnord.HasDeliveryOptionWith(
				deliveryoption.ID(deliveryOption),
			),
		).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("query DO PN: %w", err)
	}

	additionalServiceIDs := make([]string, 0)
	for _, ad := range do.Edges.CarrierAddServPostNord {
		additionalServiceIDs = append(additionalServiceIDs, ad.APICode)
	}

	carrierServicePostNord, err := do.Edges.DeliveryOption.
		QueryCarrierService().
		QueryCarrierServicePostNord().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("carrier service: %w", err)
	}

	shipments := make([]ShipmentConfig, 0)
	for _, s := range groupedSameDeliveryPacks {
		shipment, err := createShipmentConfigFromReturnGroupedParcels(
			ctx,
			s,
			carrierServicePostNord.APICode,
			additionalServiceIDs,
		)
		if err != nil {
			return nil, fmt.Errorf("shipment config from grouped: %w", err)
		}
		shipments = append(shipments, *shipment)
	}

	return shipments, nil
}

func requestLabels(ctx context.Context, deliveryOption pulid.ID, shipments []ShipmentConfig, requestQR bool) (*postnordresponse.Label, error) {
	tx := ent.FromContext(ctx)

	if len(shipments) == 0 {
		return nil, fmt.Errorf("found 0 shipments to request labels for")
	}

	do, err := tx.DeliveryOptionPostNord.Query().
		WithDeliveryOption().
		WithCarrierAddServPostNord().
		Where(
			deliveryoptionpostnord.HasDeliveryOptionWith(
				deliveryoption.ID(deliveryOption),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	agreement, err := do.Edges.DeliveryOption.QueryCarrier().
		QueryCarrierPostNord().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	labelConfig := RequestConfig{
		PartyID:        agreement.CustomerNumber,
		Shipments:      shipments,
		PostNordAPIKey: conf.PostNord.APIKey,
		RequestQRCode:  requestQR,
	}

	urlWithParams, payload, err := generateV3LabelPayload(ctx, labelConfig)
	if err != nil {
		return nil, err
	}

	request, err := generateRequest(ctx, urlWithParams, payload)
	if err != nil {
		return nil, err
	}

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	response, err := fireRequest(request)
	if err != nil {
		return nil, err
	}

	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(responseDump))

	body, err := io.ReadAll(response.Body)
	if err != nil {
		response.Body.Close()
		return nil, err
	}
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		var responseData postnordresponse.Label
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			return nil, err
		}

		return &responseData, nil
	}

	var responseData postnordresponse.ErrorComposite
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}
	details := make([]string, 0)
	for _, f := range responseData.CompositeFault.Faults {
		details = append(details, f.ExplanationText)
	}
	return nil, fmt.Errorf("PostNord request returned %v: %v: %v", response.StatusCode, responseData.Message, strings.Join(details, "; "))

}

func SaveFlattenedResponse(ctx context.Context, data []PostNordLabelResponseData) ([]*common.CreateShipment, error) {
	cli := ent.FromContext(ctx)
	v := viewer.FromContext(ctx)

	allLabels := make([]string, 0)
	output := make(map[pulid.ID]*common.CreateShipment, 0)

	for _, r := range data {

		ctxTX, _, err := cli.OpenTx(ctx)
		if err != nil {
			return nil, err
		}
		tx := ent.TxFromContext(ctxTX)
		defer tx.Rollback()

		// Should only be 1:1 with shipment; likely refactor?
		// or use the on-conflict handler
		err = tx.ShipmentPostNord.Create().
			SetBookingID(r.BookingID).
			SetItemID("<deprecated>").                     // Deprecate??
			SetShipmentReferenceNo(r.ShippingReferenceNo). // Not provided on PN test API any longer?; Duplicate of BookingID anyways?
			SetShipmentID(r.DelivrioShipmentID).
			SetTenantID(v.TenantID()).
			OnConflict().
			UpdateNewValues().
			Exec(ctxTX)
		if err != nil {
			return nil, err
		}

		createParcel, err := tx.ShipmentParcel.Create().
			SetStatus(shipmentparcel.StatusPending).
			SetShipmentID(r.DelivrioShipmentID).
			SetItemID(r.ItemID).
			SetTenantID(v.TenantID()).
			SetColliID(r.DelivrioColliID).
			Save(ctxTX)
		if err != nil {
			return nil, err
		}

		if len(r.LabelPDF) > 0 {
			allLabels = append(allLabels, r.LabelPDF)
			if _, ok := output[r.DelivrioShipmentID]; ok {
				output[r.DelivrioShipmentID].Labels = append(output[r.DelivrioShipmentID].Labels, r.LabelPDF)
			} else {
				output[r.DelivrioShipmentID] = &common.CreateShipment{
					Shipment: r.DelivrioShipmentID,
					Labels:   []string{r.LabelPDF},
				}
			}
		}

		err = tx.Commit()
		if err != nil {
			return nil, err
		}

		_, err = utils.CreateShipmentDocument(ctx, createParcel, &r.LabelPDF, &r.LabelZPL)
		if err != nil {
			return nil, err
		}

	}

	outputSlice := make([]*common.CreateShipment, 0)
	for _, s := range output {
		outputSlice = append(outputSlice, s)
	}
	return outputSlice, nil
}

func ShipmentConfigFromGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.Colli,
	basicServiceCode string,
	additionalServices []string,
) (*ShipmentConfig, error) {
	cli := ent.FromContext(ctx)
	v := viewer.FromContext(ctx)

	firstParcel := groupedParcels[0]

	parcels := make([]PackageConfig, 0)
	for _, p := range groupedParcels {

		car, err := p.QueryDeliveryOption().
			QueryCarrier().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipment, err := cli.Shipment.Create().
			SetCarrier(car).
			SetShipmentPublicID(fmt.Sprintf("%v", time.Now())).
			SetTenantID(v.TenantID()).
			SetStatus(shipment2.StatusPending).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		orderLines, err := p.QueryOrderLines().
			All(ctx)
		if err != nil {
			return nil, err
		}
		parcels = append(parcels, PackageConfig{
			DelivrioShipmentID: shipment.ID,
			DelivrioColliID:    p.ID,
			Items:              orderLines,
		})
	}

	return &ShipmentConfig{
		ConsignorAddress:   firstParcel.Edges.Sender,
		ConsigneeAddress:   firstParcel.Edges.Recipient,
		ParcelShopAddress:  firstParcel.Edges.ParcelShop,
		BasicServices:      basicServiceCode,
		AdditionalServices: additionalServices,
		Packages:           parcels,
	}, nil
}

func createShipmentConfigFromReturnGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.ReturnColli,
	basicServiceCode string,
	additionalServices []string,
) (*ShipmentConfig, error) {
	tx := ent.FromContext(ctx)
	v := viewer.FromContext(ctx)

	firstParcel := groupedParcels[0]

	parcels := make([]PackageConfig, 0)
	for _, p := range groupedParcels {

		car, err := p.QueryDeliveryOption().
			QueryCarrier().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipment, err := tx.Shipment.Create().
			SetCarrier(car).
			SetShipmentPublicID(fmt.Sprintf("%v", time.Now())).
			SetTenantID(v.TenantID()).
			SetStatus(shipment2.StatusPending).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		orderLines, err := p.
			QueryReturnOrderLine().
			QueryOrderLine().
			All(ctx)
		if err != nil {
			return nil, err
		}
		parcels = append(parcels, PackageConfig{
			DelivrioShipmentID: shipment.ID,
			DelivrioColliID:    p.ID,
			Items:              orderLines,
		})
	}

	return &ShipmentConfig{
		ConsignorAddress:   firstParcel.Edges.Sender,
		ConsigneeAddress:   firstParcel.Edges.Recipient,
		ParcelShopAddress:  nil,
		BasicServices:      basicServiceCode,
		AdditionalServices: additionalServices,
		Packages:           parcels,
	}, nil
}

type PostNordLabelResponseData struct {
	Success             bool
	DelivrioShipmentID  pulid.ID
	DelivrioColliID     pulid.ID
	BookingID           string
	ItemID              string
	ShippingReferenceNo string
	LabelZPL            string
	LabelPDF            string
	QRCodePNG           string
	LabelOffset         int
}

func checkResponse(config []ShipmentConfig, resp postnordresponse.Label) ([]PostNordLabelResponseData, error) {
	output := make([]PostNordLabelResponseData, 0)

	labelCount := 0
	for si, s := range config {

		if len(resp.BookingResponse.IDInformation)-1 < si {
			return nil, fmt.Errorf("expected IDInformation length to match input length %v != %v",
				len(resp.BookingResponse.IDInformation)-1, si,
			)
		}

		for pi, p := range s.Packages {

			if resp.BookingResponse.IDInformation[si].Status != "OK" {
				output = append(output, PostNordLabelResponseData{
					Success:            false,
					DelivrioShipmentID: p.DelivrioShipmentID,
				})
				continue
			}

			// Shipment is reported as OK at this point, so we can
			// be a little more aggressive returning errors, since we shouldn't fail here.
			if len(resp.BookingResponse.IDInformation[si].IDS)-1 < pi {
				return nil, fmt.Errorf("expected IDInformation.IDs length to match input length %v != %v",
					len(resp.BookingResponse.IDInformation[si].IDS)-1, pi)
			}

			output = append(output, PostNordLabelResponseData{
				Success:             true,
				DelivrioShipmentID:  p.DelivrioShipmentID,
				DelivrioColliID:     p.DelivrioColliID,
				BookingID:           resp.BookingResponse.BookingID,
				ItemID:              resp.BookingResponse.IDInformation[si].IDS[pi].Value,
				ShippingReferenceNo: "",
				LabelZPL:            "",
				LabelPDF:            "",
				QRCodePNG:           "",
				LabelOffset:         labelCount,
			})

			labelCount++

		}
	}

	printoutsByItemID := make(map[string][]postnordresponse.Printout)
	for li, l := range resp.LabelPrintout {
		if len(l.ItemIDS) != 1 {
			return nil, fmt.Errorf("expected LabelPrintout[%v].ItemIDS length to be 1; got: %v",
				li,
				len(l.ItemIDS))
		}

		if printoutsByItemID[l.ItemIDS[0].ItemIDS] == nil {
			printoutsByItemID[l.ItemIDS[0].ItemIDS] = make([]postnordresponse.Printout, 0)
		}
		printoutsByItemID[l.ItemIDS[0].ItemIDS] = append(printoutsByItemID[l.ItemIDS[0].ItemIDS], l.Printout)
	}

	for li, l := range output {
		if allPrintouts, ok := printoutsByItemID[l.ItemID]; ok {
			for _, val := range allPrintouts {
				if val.LabelFormat == "ZPL" {
					output[li].LabelZPL = val.Data
				} else if val.LabelFormat == "PDF" {
					output[li].LabelPDF = val.Data
				} else if val.LabelFormat == "PNG" {
					output[li].QRCodePNG = val.Data
				} else {
					return nil, fmt.Errorf(
						"expected LabelFormat: ZPL | PDF | PNG; got: %v",
						val.LabelFormat)
				}
			}
		}
	}

	return output, nil
}
