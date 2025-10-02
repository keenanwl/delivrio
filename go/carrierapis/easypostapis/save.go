package easypostapis

import (
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"strconv"
	"time"
)

func SaveLabelData(ctx context.Context, resp FetchLabelOutput) (*common.CreateShipment, error) {

	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	rate := 0.0
	if resp.ResponseShipment.SelectedRate != nil {
		r, err := strconv.ParseFloat(resp.ResponseShipment.SelectedRate.Rate, 10)
		if err != nil {
			return nil, err
		}
		rate = r
	}

	// Rough "days"
	expectedDelivery := time.Now().Add(time.Hour * 24 * time.Duration(resp.ResponseShipment.SelectedRate.EstDeliveryDays))

	createShip := tx.ShipmentEasyPost.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetTenantID(view.TenantID()).
		SetTrackingNumber(resp.ResponseShipment.TrackingCode).
		SetEpShipmentID(resp.ResponseShipment.ID).
		SetRate(rate).
		SetEstDeliveryDate(expectedDelivery)

	err := createShip.Exec(ctx)
	if err != nil {
		return nil, err
	}

	sp, err := tx.ShipmentParcel.Create().
		SetShipmentID(resp.Package.DelivrioShipmentID).
		SetStatus(shipmentparcel.StatusPending).
		SetItemID(resp.ResponseShipment.TrackingCode).
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
