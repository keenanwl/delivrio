package orderhooks

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
	"fmt"
)

func UpdateOrderStatusOnColliMutate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ColliFunc(func(ctx context.Context, m *ent2.ColliMutation) (ent.Value, error) {

			colliIDs := make([]pulid.ID, 0)
			var err error
			if m.Op() != ent.OpCreate {
				colliIDs, err = m.IDs(ctx)
				if err != nil {
					return nil, fmt.Errorf("hook: colli: order status: %w", err)
				}
			}

			mutateVal, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, fmt.Errorf("hook: colli: order status: %w", err)
			}

			// We need IDs at different times thus the
			if m.Op() == ent.OpCreate {
				createID, exists := m.ID()
				if !exists {
					return nil, fmt.Errorf("hook: colli: order status: ID should exist after create")
				}
				colliIDs = []pulid.ID{createID}
			}

			ordersToCheck, err := m.Client().Colli.Query().
				Where(colli.IDIn(colliIDs...)).
				QueryOrder().
				Unique(true).
				All(ctx)
			if err != nil {
				return nil, fmt.Errorf("hook: colli: order status: %w", err)
			}

			if len(colliIDs) == 0 {
				return mutateVal, nil
			}

			for _, o := range ordersToCheck {
				err = updateOrderStatusIfRequired(ctx, o)
				if err != nil {
					return nil, fmt.Errorf("hook: colli: order status: %w", err)
				}
			}

			return mutateVal, nil
		})
	}
	return hook.On(hk, ent2.OpCreate|ent2.OpUpdateOne|ent2.OpUpdate|ent2.OpDeleteOne|ent2.OpDelete)
}

func updateOrderStatusIfRequired(ctx context.Context, ord *ent2.Order) error {

	collis, err := ord.Colli(ctx)
	if err != nil {
		return err
	}

	allHaveShipments := true
	partiallyDispatched := false
	allCollisCancelled := true
	for _, c := range collis {

		if c.Status != colli.StatusCancelled {
			allCollisCancelled = false
		}

		count, err := c.QueryShipmentParcel().
			QueryShipment().
			Where(shipment.StatusNEQ(shipment.StatusDeleted)).
			Count(ctx)
		if count == 0 {
			allHaveShipments = false
			continue
		} else if err != nil {
			return err
		}
		partiallyDispatched = true

	}

	status := order.StatusPending

	if allCollisCancelled {
		status = order.StatusCancelled
	} else if allHaveShipments {
		status = order.StatusDispatched
	} else if partiallyDispatched {
		status = order.StatusDispatched
	}

	// Bail if status already set
	if status == ord.Status {
		return nil
	}

	err = ord.Update().
		SetStatus(status).
		Exec(history.FromContext(ctx).SetDescription("Update order status").Ctx())
	if err != nil {
		return err
	}

	return nil

}
