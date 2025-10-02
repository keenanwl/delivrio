package returncollihooks

import (
	"context"
	"delivrio.io/go/ent/returncolli"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"entgo.io/ent"
)

func UpdateReturnColli() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ReturnColliFunc(func(ctx context.Context, m *ent2.ReturnColliMutation) (ent.Value, error) {
			nextStatus, statusUpdated := m.Status()
			if statusUpdated && nextStatus != returncolli.StatusOpened && nextStatus != returncolli.StatusDeleted {
				returnColliIDs, err := m.IDs(ctx)
				if err != nil {
					return nil, err
				}

				returnCollisOrderLines, err := m.Client().ReturnColli.Query().
					Where(returncolli.IDIn(returnColliIDs...)).
					QueryReturnOrderLine().
					WithOrderLine().
					All(ctx)
				if err != nil {
					return nil, err
				}

				for _, rol := range returnCollisOrderLines {
					err = orderLineUnitsOK(ctx, rol.Edges.OrderLine, rol.Units, &rol.ID)
					if err != nil {
						return nil, err
					}
				}

			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}
