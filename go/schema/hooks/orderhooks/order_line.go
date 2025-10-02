package orderhooks

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/documentfile"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/orderline"
	"entgo.io/ent"
	"fmt"
)

func DeleteOrderLine() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.OrderLineFunc(func(ctx context.Context, m *ent2.OrderLineMutation) (ent.Value, error) {
			/*			tx := ent2.FromContext(ctx)
						errs := hooks.NewValidationError()

						ids, err := m.IDs(ctx)
						if err != nil {
							return nil, err
						}
						currentLines, err := tx.OrderLine.Query().WithShipmentOrderLines().Where(orderline.IDIn(ids...)).All(ctx)
						if err != nil {
							return nil, err
						}

						for _, line := range currentLines {
							if len(line.Edges.ShipmentOrderLines) > 0 {
								errs.SetError(orderline.EdgeShipmentOrderLines, "order lines that have been added to shipments may not be deleted")
								return nil, errs
							}
						}*/

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpDelete)
}

// Docs (like packing slip) get cached for performance, but need to be re-fetched
// if the underlying data has been updated
func UpdateOrderLineClearDocs() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.OrderLineFunc(func(ctx context.Context, m *ent2.OrderLineMutation) (ent.Value, error) {
			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			_, err = m.Client().DocumentFile.Delete().
				Where(documentfile.HasColliWith(colli.HasOrderLinesWith(orderline.IDIn(ids...)))).
				Exec(ctx)
			if err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}

func CreateOrderLineClearDocs() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.OrderLineFunc(func(ctx context.Context, m *ent2.OrderLineMutation) (ent.Value, error) {

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			ol, ok := v.(*ent2.OrderLine)
			if !ok {
				return nil, fmt.Errorf("clear docs: could not determine val type")
			}

			_, err = m.Client().DocumentFile.Delete().
				Where(documentfile.HasColliWith(colli.HasOrderLinesWith(orderline.ID(ol.ID)))).
				Exec(ctx)
			if err != nil {
				return nil, err
			}

			return v, nil
		})
	}
	return hook.On(hk, ent.OpCreate)
}
