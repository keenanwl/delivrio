package bringapis

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/bringapis/bringrequest"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"time"
)

type PackageConfig struct {
	DelivrioShipmentID pulid.ID
	DelivrioColliID    pulid.ID
	Packaging          *ent.Packaging
	OrderLines         []*ent.OrderLine
}

type ShipmentConfig struct {
	bringrequest.AuthenticationHeaders
	PublicOrderID       string
	ConsignorAddress    *ent.Address
	ConsigneeAddress    *ent.Address
	ParcelShopCountry   *ent.Country
	ParcelShopID        string // Bring ID
	Packages            []PackageConfig
	BringService        string
	BringCustomerNumber string
	ElectronicCustoms   bool
	ShippingDate        time.Time
	Test                bool
}

type FetchLabelOutput struct {
	Package           PackageConfig
	ConsignmentNumber string
	PackageNumber     string // "Barcode"
	ResponseB64PDF    string
	Error             error
}

func FetchLabels(ctx context.Context, collis []*ent.Colli) ([]*FetchLabelOutput, error) {
	shipmentGroups, err := common.GroupPackagesBySenderReceiver(ctx, collis)
	if err != nil {
		return nil, err
	}

	allShipmentConfigs := make([]*ShipmentConfig, 0)
	for _, sg := range shipmentGroups {
		shipmentConfig, err := shipmentFromGroupedParcels(ctx, sg)
		if err != nil {
			return nil, err
		}
		allShipmentConfigs = append(allShipmentConfigs, shipmentConfig)
	}

	output := make([]*FetchLabelOutput, 0)

	for _, s := range allShipmentConfigs {
		// TODO: request labels outside this loop to re-use token (perf)
		allLabels, err := CreateOrderFetchLabel(ctx, s)
		// TODO: Potentially continue on err?
		if err != nil {
			return nil, err
		}

		output = append(output, allLabels...)
	}

	return output, nil
}

func SaveLabelData(ctx context.Context, resp *FetchLabelOutput) (*common.CreateShipment, error) {
	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	ctxTX, _, err := cli.OpenTx(ctx)
	if err != nil {
		return nil, err
	}
	tx := ent.TxFromContext(ctxTX)
	defer tx.Rollback()

	err = tx.ShipmentBring.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetTenantID(view.TenantID()).
		SetConsignmentNumber(resp.PackageNumber).
		Exec(ctxTX)
	if err != nil {
		return nil, err
	}

	sp, err := tx.ShipmentParcel.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetStatus(shipmentparcel.StatusPending).
		SetItemID(resp.PackageNumber).
		SetColliID(resp.Package.DelivrioColliID).
		SetTenantID(view.TenantID()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	_, err = utils.CreateShipmentDocument(ctx, sp, &resp.ResponseB64PDF, nil)
	if err != nil {
		return nil, err
	}

	return &common.CreateShipment{
		Shipment: resp.Package.DelivrioShipmentID,
		Labels:   []string{resp.ResponseB64PDF},
	}, nil
}

func shipmentFromGroupedParcels(
	ctx context.Context,
	groupedParcels []*ent.Colli,
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

		packaging, err := common.ColliPackaging(ctx, p)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if ent.IsNotFound(err) {
			return nil, fmt.Errorf("bring requires colli packaging is added")
		}

		orderLines, err := p.QueryOrderLines().
			All(ctx)
		if err != nil {
			return nil, err
		}
		parcels = append(parcels, PackageConfig{
			DelivrioShipmentID: shipment.ID,
			DelivrioColliID:    p.ID,
			Packaging:          packaging,
			OrderLines:         orderLines,
		})
	}

	var parcelShopCountry *ent.Country
	var parcelShopID string

	parcelShop, err := firstParcel.QueryParcelShop().
		WithParcelShopDAO().
		WithAddress().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if !ent.IsNotFound(err) {
		psc, err := parcelShop.Edges.Address.Country(ctx)
		if err != nil {
			return nil, err
		}
		parcelShopCountry = psc
		parcelShopID = parcelShop.Edges.ParcelShopDAO.ShopID
	}

	ord, err := firstParcel.QueryOrder().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	dOpt, err := firstParcel.QueryDeliveryOption().
		WithDeliveryOptionBring().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	service, err := dOpt.QueryCarrierService().
		QueryCarrierServiceBring().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	agreement, err := dOpt.QueryCarrier().
		QueryCarrierBring().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	pickupDay, err := utils.PickupDate(ctx)
	if err != nil {
		return nil, err
	}

	return &ShipmentConfig{
		AuthenticationHeaders: NewAuthentication(conf),
		PublicOrderID:         ord.OrderPublicID,
		ConsignorAddress:      firstParcel.Edges.Sender,
		ConsigneeAddress:      firstParcel.Edges.Recipient,
		ParcelShopCountry:     parcelShopCountry,
		ParcelShopID:          parcelShopID,
		Packages:              parcels,
		BringService:          service.APIRequest,
		BringCustomerNumber:   agreement.CustomerNumber,
		ElectronicCustoms:     dOpt.Edges.DeliveryOptionBring.ElectronicCustoms,
		ShippingDate:          pickupDay,
		Test:                  agreement.Test,
	}, nil
}

func NewAuthentication(conf *appconfig.DelivrioConfig) bringrequest.AuthenticationHeaders {
	return bringrequest.AuthenticationHeaders{
		APIKey:    conf.Bring.APIKey,
		APIUID:    conf.Bring.APIUID,
		ClientURL: conf.BaseURL,
	}
}
