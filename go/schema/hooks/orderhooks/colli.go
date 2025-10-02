package orderhooks

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/documentfile"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/schema/hooks"
	"delivrio.io/go/utils"
	"entgo.io/ent"
)

func CreateColliBarcode() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ColliFunc(func(ctx context.Context, m *ent2.ColliMutation) (ent.Value, error) {
			// Add support for auto increment on DB level here
			m.SetInternalBarcode(utils.NextShipmentBarcodeID())
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func DeleteColli() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ColliFunc(func(ctx context.Context, m *ent2.ColliMutation) (ent.Value, error) {
			tx := ent2.FromContext(ctx)
			errs := hooks.NewValidationError()

			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			currentCollis, err := tx.Colli.Query().
				WithOrder().
				WithShipmentParcel().
				WithOrderLines().
				Where(colli.IDIn(ids...)).
				All(ctx)
			if err != nil {
				return nil, err
			}

			for _, c := range currentCollis {
				allColli, err := c.Edges.Order.Colli(ctx)
				if err != nil {
					return nil, err
				}

				if len(allColli) <= 1 {
					errs.SetError("order", "A order must have at least one package")
				}

				if c.Edges.ShipmentParcel != nil {
					errs.SetError("shipment", "A package with a shipment may not be deleted")
				}

				if len(c.Edges.OrderLines) > 0 {
					errs.SetError("order_lines", "A package with a order lines may not be deleted")
				}
			}

			// Order is probably always defined?
			err = updateOrderStatusIfRequired(ctx, currentCollis[0].Edges.Order)
			if err != nil {
				return nil, err
			}

			if len(errs.InvalidFields(ctx)) > 0 {
				return nil, errs
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpDelete|ent.OpDeleteOne)
}

// Docs (like packing slip) get cached for performance, but need to be re-fetched
// if the underlying data has been updated
func UpdateColliClearDocs() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ColliFunc(func(ctx context.Context, m *ent2.ColliMutation) (ent.Value, error) {
			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			_, err = m.Client().DocumentFile.Delete().
				Where(documentfile.HasColliWith(colli.IDIn(ids...))).
				Exec(ctx)
			if err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne)
}
