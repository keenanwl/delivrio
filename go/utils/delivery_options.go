package utils

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/shared-utils/pulid"
)

func UpdateDeliveryOption(ctx context.Context, id pulid.ID, inputDeliveryOption ent.UpdateDeliveryOptionInput) error {
	tx := ent.TxFromContext(ctx)

	err := tx.DeliveryOption.Update().
		SetInput(inputDeliveryOption).
		Where(deliveryoption.ID(id)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
