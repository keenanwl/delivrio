package returncollihooks

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/orderline"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/ent/returnorderline"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent"
)

type HookReturnOrderLineErr struct {
	Message string
}

func (e HookReturnOrderLineErr) Error() string {
	return e.Message
}

func CreateReturnOrderLine() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ReturnOrderLineFunc(func(ctx context.Context, m *ent2.ReturnOrderLineMutation) (ent.Value, error) {
			orderLineID, exists := m.OrderLineID()
			if !exists {
				return nil, HookReturnOrderLineErr{"order line is required"}
			}

			orderLine, err := m.Client().OrderLine.Query().
				Where(
					orderline.ID(orderLineID),
				).Only(ctx)
			if err != nil {
				return nil, err
			}

			unitsToCreate, exists := m.Units()
			if !exists {
				return nil, HookReturnOrderLineErr{"return units must be defined"}
			}

			err = orderLineUnitsOK(ctx, orderLine, unitsToCreate, nil)
			if err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func orderLineUnitsOK(ctx context.Context, orderLine *ent2.OrderLine, unitsToCreate int, excludeReturnColliOrderLine *pulid.ID) error {

	pred := returnorderline.HasReturnColliWith(
		returncolli.StatusNotIn(
			// Users can "Open" multiple
			returncolli.StatusDeleted, returncolli.StatusOpened,
		),
	)

	if excludeReturnColliOrderLine != nil {
		pred = returnorderline.And(
			pred,
			returnorderline.IDNEQ(*excludeReturnColliOrderLine),
		)
	}

	existingReturnOrderLines, err := orderLine.
		QueryReturnOrderLine().
		Where(pred).
		All(ctx)
	if err != nil {
		return err
	}

	maxTotalUnits := orderLine.Units

	currentUnits := 0
	for _, rol := range existingReturnOrderLines {
		currentUnits += rol.Units
	}

	if currentUnits+unitsToCreate > maxTotalUnits {
		return HookReturnOrderLineErr{"Total returned units may not exceed total delivered units"}
	}

	return nil
}
