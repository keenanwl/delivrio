package easypostapis

import (
	"context"
	"delivrio.io/go/ent"
	"fmt"
	"github.com/EasyPost/easypost-go/v4"
)

func CancelByEasyPostID(ctx context.Context, carrier *ent.CarrierEasyPost, id string) error {
	cli := easypost.New(carrier.APIKey)
	ship, err := cli.RefundShipment(id)
	if err != nil {
		return err
	}

	if ship.RefundStatus != "submitted" && ship.RefundStatus != "refunded" {
		return fmt.Errorf("invalid shipment status: %s", ship.Status)
	}

	return nil
}
