package uspsapis

import (
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/returncolli"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"time"
)

type AdditionalService string

func FetchSingleLabel(ctx context.Context, o common.ReturnOrderDeliveryOptionsColliIDs) (*FetchLabelOutput, error) {
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
		return nil, err
	}

	shipments, err := createReturnShipments(ctx, o.DeliveryOptionID, grouped)
	if err != nil {
		return nil, err
	}

	output := make([]FetchLabelOutput, 0)

	for _, shipment := range shipments {
		response, err := requestLabels(ctx, o.DeliveryOptionID, shipment)
		if err != nil {
			return nil, err
		}
		output = append(output, response...)
	}

	if len(output) == 0 {
		return nil, fmt.Errorf("usps: returns: unexpected empty results")
	}

	return &output[0], nil
}

func shipmentFromReturnGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.ReturnColli,
	additionalServices []AdditionalService,
) (*ShipmentConfig, error) {
	tx := ent.FromContext(ctx)
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

		shipment, err := tx.Shipment.Create().
			SetCarrier(agreement).
			SetShipmentPublicID(fmt.Sprintf("%v", time.Now())).
			SetTenantID(v.TenantID()).
			SetStatus(shipment2.StatusPending).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		packaging, err := p.QueryPackaging().
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if ent.IsNotFound(err) {
			// Default to outbound packaging if Return Colli does not have packaging
			outboundColli, err := p.QueryReturnOrderLine().
				QueryOrderLine().
				QueryColli().
				First(ctx)
			if err != nil {
				return nil, err
			}

			outboundPackaging, err := common.ColliPackaging(ctx, outboundColli)
			if err != nil && !ent.IsNotFound(err) {
				return nil, err
			} else if ent.IsNotFound(err) {
				return nil, fmt.Errorf("usps requires return colli packaging")
			}
			packaging = outboundPackaging
		}

		ord, err := firstParcel.QueryOrder().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		products, err := p.
			QueryReturnOrderLine().
			QueryOrderLine().
			All(ctx)
		if err != nil {
			return nil, err
		}
		parcels = append(parcels, PackageConfig{
			PublicOrderID:      ord.OrderPublicID,
			DelivrioShipmentID: shipment.ID,
			DelivrioColliID:    p.ID,
			Items:              products,
			Packaging:          packaging,
		})
	}

	return &ShipmentConfig{
		ConsignorAddress:   firstParcel.Edges.Sender,
		ConsigneeAddress:   firstParcel.Edges.Recipient,
		ParcelShopAddress:  nil,
		AdditionalServices: additionalServices,
		Packages:           parcels,
		IsReturn:           true,
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

	addServices, err := do.QueryDeliveryOptionUSPS().
		QueryCarrierAdditionalServiceUSPS().
		All(ctx)
	if err != nil {
		return nil, err
	}

	shipments := make([]ShipmentConfig, 0)
	for _, s := range groupedSameDeliveryPacks {
		additionalServices := make([]AdditionalService, 0)
		for _, as := range addServices {
			additionalServices = append(additionalServices, AdditionalService(as.APICode))
		}

		shipment, err := shipmentFromReturnGroupedParcels(
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
