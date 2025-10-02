package glsapis

import (
	"context"
	"delivrio.io/go/carrierapis/glsapis/glsresponse"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierservicegls"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/deliveryoptiongls"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"io"
	"net/http/httputil"
	"time"
)

var (
	tracer = otel.Tracer("labels-GLS")
)

func requestLabels(ctx context.Context, deliveryOption pulid.ID, shipment ShipmentConfig) (*glsresponse.SuccessLabel, error) {

	ctx, span := tracer.Start(ctx, "requestLabels")
	defer span.End()

	cli := ent.FromContext(ctx)

	do, err := cli.DeliveryOptionGLS.Query().
		WithDeliveryOption().
		WithCarrierAdditionalServiceGLS().
		Where(
			deliveryoptiongls.HasDeliveryOptionWith(
				deliveryoption.ID(deliveryOption),
			),
		).Only(ctx)
	if err != nil {
		return nil, err
	}

	agreement, err := do.Edges.DeliveryOption.
		QueryCarrier().
		QueryCarrierGLS().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	labelConfig := RequestConfig{
		GLSAPIAuth: GLSAPIAuth{
			UserName:   agreement.GLSUsername,
			Password:   agreement.GLSPassword,
			ContactID:  agreement.ContactID,
			CustomerID: agreement.CustomerID,
		},
		Shipment: shipment,
	}

	urlWithParams, payload, err := generateV1CreateShipment(ctx, labelConfig)
	if err != nil {
		return nil, err
	}

	span.SetAttributes(
		attribute.String("requestURL", urlWithParams.String()),
	)

	request, err := generateRequest(ctx, urlWithParams, payload)
	if err != nil {
		return nil, err
	}

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(requestDump))

	span.AddEvent("FireGLS")
	response, err := fireRequest(request)
	span.AddEvent("ResponseGLS")
	if err != nil {
		return nil, err
	}

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
		var responseData glsresponse.SuccessLabel
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			return nil, err
		}

		return &responseData, nil
	}

	var responseData glsresponse.ErrorLabel
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("GLS request returned %v: %v: %v", response.StatusCode, responseData.Message, string(body))

}

func createShipmentConfig(ctx context.Context, deliveryOption pulid.ID, groupedSameDeliveryPacks [][]*ent.Colli) ([]ShipmentConfig, error) {
	tx := ent.FromContext(ctx)

	// Query is duplicated
	do, err := tx.DeliveryOption.Query().
		WithCarrier().
		WithCarrierService().
		Where(deliveryoption.ID(deliveryOption)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	addServices, err := do.QueryDeliveryOptionGLS().
		QueryCarrierAdditionalServiceGLS().
		All(ctx)
	if err != nil {
		return nil, err
	}

	carrierServiceGLS, err := do.Edges.CarrierService.
		QueryCarrierServiceGLS().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	shipments := make([]ShipmentConfig, 0)
	for _, s := range groupedSameDeliveryPacks {
		parcelShop, err := s[0].QueryParcelShop().
			QueryParcelShopGLS().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		additionalServices := make([]AdditionalService, 0)
		for _, ad := range addServices {
			switch ad.InternalID {
			case seed.GLSDeliveryPointServiceInternalID:
				additionalServices = append(additionalServices, AdditionalService{
					Key:   "ShopDelivery",
					Value: parcelShop.GLSParcelShopID,
				})
				break
			}
		}

		// All GLS services & add-ons end up in the same KV object
		if carrierServiceGLS.APIValue != carrierservicegls.APIValueNone && carrierServiceGLS.APIValue != carrierservicegls.APIValueNumericString {
			additionalServices = append(additionalServices, AdditionalService{
				Key:   *carrierServiceGLS.APIKey,
				Value: carrierServiceGLS.APIValue.String(),
			})
		}
		shipment, err := glsShipmentFromGroupedParcels(
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

func glsShipmentFromGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.Colli,
	additionalServices []AdditionalService,
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

		err = tx.ShipmentGLS.Create().
			SetShipment(shipment).
			SetTenantID(v.TenantID()).
			SetConsignmentID("").
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		products, err := p.QueryOrderLines().
			QueryProductVariant().
			All(ctx)
		if err != nil {
			return &ShipmentConfig{}, err
		}
		parcels = append(parcels, PackageConfig{
			DelivrioShipmentID: shipment.ID,
			DelivrioColliID:    p.ID,
			Items:              products,
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
