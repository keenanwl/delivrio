package systemeventhooks

import (
	"context"
	"errors"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/viewer"
	"entgo.io/ent"
)

func CreateProduct() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ProductFunc(func(ctx context.Context, m *ent2.ProductMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)

			if !view.Background() && view.CurrentRole() != viewer.BackgroundForTenant {
				return next.Mutate(ctx, m)
			}

			allErr := make([]error, 0)

			id, exists := m.ID()
			if !exists {
				allErr = append(allErr, errors.New("missing Product ID"))
			}

			create := m.Client().SystemEvents.Create().
				SetTenantID(view.TenantID()).
				SetEventType(systemevents.EventTypeBackgroundProductMutate).
				SetEventTypeID(id.String()).
				SetDescription("Background product create from connection")

			val, err := next.Mutate(ctx, m)
			if err != nil {
				allErr = append(allErr, err)
			}

			if len(allErr) == 0 {
				create = create.SetStatus(systemevents.StatusSuccess)
			} else {
				create = create.
					SetStatus(systemevents.StatusFail).
					SetData(err.Error())
			}
			// Add some observability here to process the error
			_, _ = create.Save(ctx)

			return val, err
		})
	}

	return hook.On(hk, ent.OpCreate)
}
func UpdateProduct() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ProductFunc(func(ctx context.Context, m *ent2.ProductMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)

			if !view.Background() && view.CurrentRole() != viewer.BackgroundForTenant {
				return next.Mutate(ctx, m)
			}

			allErr := make([]error, 0)

			ids, err := m.IDs(ctx)
			if err != nil {
				allErr = append(allErr, err)
			}

			if len(ids) == 0 {
				allErr = append(allErr, errors.New("missing Product ID(s)"))
			}

			create := m.Client().SystemEvents.Create().
				SetTenantID(view.TenantID()).
				SetEventType(systemevents.EventTypeBackgroundProductMutate).
				SetEventTypeID(ids[0].String()).
				SetDescription("Background product update from connection")

			val, err := next.Mutate(ctx, m)
			if err != nil {
				allErr = append(allErr, err)
			}

			if len(allErr) == 0 {
				create = create.SetStatus(systemevents.StatusSuccess)
			} else {
				create = create.
					SetStatus(systemevents.StatusFail).
					SetData(err.Error())
			}
			create.SaveX(ctx)

			return val, err
		})
	}

	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}
