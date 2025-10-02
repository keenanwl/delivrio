package glsapis

import (
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierservicegls"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/returncolli"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"time"
)

func FetchSingleLabelGLS(ctx context.Context, o common.ReturnOrderDeliveryOptionsColliIDs) (string, error) {
	db := ent.FromContext(ctx)

	allReturnCollis, err := db.ReturnColli.Query().
		WithSender().
		WithRecipient().
		Where(returncolli.ID(o.ReturnColliID)).
		All(ctx)
	if err != nil {
		return "", err
	}

	grouped, err := common.GroupReturnPackages(ctx, allReturnCollis)
	if err != nil {
		return "", err
	}

	shipments, err := createReturnShipments(ctx, o.DeliveryOptionID, grouped)
	if err != nil {
		return "", err
	}

	for _, shipment := range shipments {
		response, err := requestLabels(ctx, o.DeliveryOptionID, shipment)
		if err != nil {
			return "", err
		}
		return response.PDF, nil
	}
	return "", fmt.Errorf("fetchlabels: empty return")
}

func glsShipmentFromReturnGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.ReturnColli,
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

		products, err := p.
			QueryReturnOrderLine().
			QueryOrderLine().
			QueryProductVariant().
			All(ctx)
		if err != nil {
			return nil, err
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
		ParcelShopAddress:  nil,
		AdditionalServices: additionalServices,
		Packages:           parcels,
	}, nil
}

func createReturnShipments(ctx context.Context, deliveryOption pulid.ID, groupedSameDeliveryPacks [][]*ent.ReturnColli) ([]ShipmentConfig, error) {
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
		additionalServices := make([]AdditionalService, 0)
		for _, ad := range addServices {
			switch ad.InternalID {
			case seed.GLSDeliveryPointServiceInternalID:
				// Not relevant for returns
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
		shipment, err := glsShipmentFromReturnGroupedParcels(
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
