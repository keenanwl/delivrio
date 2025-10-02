package daoapis

import (
	"context"
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

// PackageConfig is all single level  since packages aren't grouped
type PackageConfig struct {
	Agreement          *ent.CarrierDAO
	ConsignorAddress   *ent.Address
	ConsigneeAddress   *ent.Address
	ParcelShopID       string
	CarrierService     *ent.CarrierService
	OrderPublicID      string
	DelivrioShipmentID pulid.ID
	DelivrioColliID    pulid.ID
	Packaging          *ent.Packaging
	OrderLines         []*ent.OrderLine
}

type FetchLabelOutput struct {
	Package        *PackageConfig
	Barcode        string
	ResponseB64PDF string
	Error          error
}

func FetchLabels(ctx context.Context, collis []*ent.Colli) ([]*FetchLabelOutput, error) {
	output := make([]*FetchLabelOutput, 0)

	for _, c := range collis {

		config, err := shipmentFromColli(ctx, c)
		if err != nil {
			return nil, err
		}

		daoLabel, err := CreateOrderFetchLabel(ctx, config)
		// TODO: Potentially continue on err?
		if err != nil {
			return nil, err
		}

		output = append(output, daoLabel)
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

	err = tx.ShipmentDAO.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetTenantID(view.TenantID()).
		SetBarcodeID(resp.Barcode).
		Exec(ctxTX)
	if err != nil {
		return nil, err
	}

	sp, err := tx.ShipmentParcel.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetStatus(shipmentparcel.StatusPending).
		SetItemID(resp.Barcode).
		SetColliID(resp.Package.DelivrioColliID).
		SetTenantID(view.TenantID()).
		Save(ctxTX)
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

// Each colli gets it's own shipment since the API
// doesn't support batching
func shipmentFromColli(
	ctx context.Context,
	colli *ent.Colli,
) (*PackageConfig, error) {
	cli := ent.FromContext(ctx)
	v := viewer.FromContext(ctx)

	car, err := colli.QueryDeliveryOption().
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

	packaging, err := common.ColliPackaging(ctx, colli)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if ent.IsNotFound(err) {
		return nil, fmt.Errorf("dao requires colli packaging is added")
	}

	ord, err := colli.QueryOrder().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	orderLines, err := colli.QueryOrderLines().
		All(ctx)
	if err != nil {
		return nil, err
	}

	sender, err := colli.Sender(ctx)
	if err != nil {
		return nil, err
	}
	recipient, err := colli.Recipient(ctx)
	if err != nil {
		return nil, err
	}

	shopID := ""
	parcelShopDAO, err := colli.QueryParcelShop().
		QueryParcelShopDAO().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if !ent.IsNotFound(err) {
		shopID = parcelShopDAO.ShopID
	}

	do, err := colli.DeliveryOption(ctx)
	if err != nil {
		return nil, fmt.Errorf("dao: label config: %w", err)
	}

	service, err := do.QueryCarrierService().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("dao: label config: %w", err)
	}

	daoAgreement, err := do.QueryCarrier().
		QueryCarrierDAO().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("dao: label config: %w", err)
	}

	return &PackageConfig{
		Agreement:          daoAgreement,
		ConsignorAddress:   sender,
		ConsigneeAddress:   recipient,
		ParcelShopID:       shopID,
		CarrierService:     service,
		OrderPublicID:      ord.OrderPublicID,
		DelivrioShipmentID: shipment.ID,
		DelivrioColliID:    colli.ID,
		Packaging:          packaging,
		OrderLines:         orderLines,
	}, nil
}
