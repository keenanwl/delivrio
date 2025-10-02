package history

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/orderhistory"
	"delivrio.io/go/ent/returncollihistory"
	"delivrio.io/go/viewer"
	"entgo.io/ent"
	"errors"
	"fmt"
)

type ctxHistoryKey struct{}

type ConfigCtx struct {
	ctx         context.Context
	description string
	origin      changehistory.Origin
}

func (c *ConfigCtx) SetDescription(d string) *ConfigCtx {
	c.description = d
	return c
}
func (c *ConfigCtx) SetOrigin(o changehistory.Origin) *ConfigCtx {
	c.origin = o
	return c
}
func (c *ConfigCtx) Ctx() context.Context {
	return context.WithValue(c.ctx, ctxHistoryKey{}, *c)
}

func NewConfig(parent context.Context) *ConfigCtx {
	return &ConfigCtx{
		ctx:         parent,
		description: "",
		origin:      changehistory.OriginUnknown,
	}
}

func FromContext(ctx context.Context) *ConfigCtx {
	h, found := ctx.Value(ctxHistoryKey{}).(ConfigCtx)
	if !found {
		return &ConfigCtx{
			ctx:    ctx,
			origin: changehistory.DefaultOrigin,
		}
	}

	if h.origin == "" {
		h.origin = changehistory.DefaultOrigin
	}

	return &h
}

func OrderCreate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.OrderFunc(func(ctx context.Context, m *ent2.OrderMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			o := v.(*ent2.Order)

			first := m.Client().OrderHistory.Query().
				Where(orderhistory.HasOrderWith(order.ID(o.ID))).
				FirstX(ctx)
			if first != nil {
				// TODO: upsert detected, don't double log
				return v, nil
			}

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			_, err = m.Client().OrderHistory.Create().
				SetOrder(o).
				SetDescription(hist.description).
				SetType(orderhistory.TypeCreate).
				SetTenantID(view.TenantID()).
				SetChangeHistoryID(view.ContextID()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating order history record: %w", err)
			}

			return v, nil
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func OrderUpdate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.OrderFunc(func(ctx context.Context, m *ent2.OrderMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			if len(ids) == 0 {
				return nil, errors.New("could not determine which order(s) is(are) being updated")
			}

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			op := orderhistory.TypeUpdate
			if m.Op() == ent.OpDelete || m.Op() == ent.OpDeleteOne {
				op = orderhistory.TypeDelete
			}

			for _, id := range ids {

				_, err = m.Client().OrderHistory.Create().
					SetOrderID(id).
					SetDescription(hist.description).
					SetType(op).
					SetTenantID(view.TenantID()).
					SetChangeHistoryID(view.ContextID()).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("creating order history record: %w", err)
				}
			}

			return v, nil
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne|ent.OpDeleteOne|ent.OpDelete)
}

func ColliCreate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ColliFunc(func(ctx context.Context, m *ent2.ColliMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			col := v.(*ent2.Colli)

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			ord, err := col.Order(ctx)
			if err != nil {
				return nil, err
			}

			_, err = m.Client().OrderHistory.Create().
				SetOrder(ord).
				SetDescription(hist.description).
				SetType(orderhistory.TypeCreate).
				SetTenantID(view.TenantID()).
				SetChangeHistoryID(view.ContextID()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating colli-order history record: %w", err)
			}

			return v, nil
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func ColliUpdate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ColliFunc(func(ctx context.Context, m *ent2.ColliMutation) (ent.Value, error) {

			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			if len(ids) == 0 {
				return next.Mutate(ctx, m)
			}

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			op := orderhistory.TypeUpdate
			if m.Op() == ent.OpDelete || m.Op() == ent.OpDeleteOne {
				op = orderhistory.TypeDelete
			}

			for _, id := range ids {

				ord, err := m.Client().Colli.Query().
					Where(colli.ID(id)).
					QueryOrder().
					Only(ctx)
				if err != nil {
					return nil, err
				}

				_, err = m.Client().OrderHistory.Create().
					SetOrderID(ord.ID).
					SetDescription(hist.description).
					SetType(op).
					SetTenantID(view.TenantID()).
					SetChangeHistoryID(view.ContextID()).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("creating colli-order history record: %w", err)
				}
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne|ent.OpDeleteOne|ent.OpDelete)
}

func ReturnColliCreate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ReturnColliFunc(func(ctx context.Context, m *ent2.ReturnColliMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			rCol := v.(*ent2.ReturnColli)

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			_, err = m.Client().ReturnColliHistory.Create().
				SetReturnColli(rCol).
				SetDescription(hist.description).
				SetType(returncollihistory.TypeCreate).
				SetTenantID(view.TenantID()).
				SetChangeHistoryID(view.ContextID()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating return colli history record: %w", err)
			}

			return v, nil
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func ReturnColliUpdate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ReturnColliFunc(func(ctx context.Context, m *ent2.ReturnColliMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			// Some updates are just to clear out any previously unfinished returns
			if len(ids) == 0 {
				return next.Mutate(ctx, m)
			}

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			op := returncollihistory.TypeUpdate
			if m.Op() == ent.OpDelete || m.Op() == ent.OpDeleteOne {
				op = returncollihistory.TypeDelete
			}

			for _, id := range ids {
				_, err = m.Client().ReturnColliHistory.Create().
					SetReturnColliID(id).
					SetDescription(hist.description).
					SetType(op).
					SetTenantID(view.TenantID()).
					SetChangeHistoryID(view.ContextID()).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("creating return colli history record: %w", err)
				}
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne|ent.OpDeleteOne|ent.OpDelete)
}

func PrintJobCreate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.PrintJobFunc(func(ctx context.Context, m *ent2.PrintJobMutation) (ent.Value, error) {
			view := viewer.FromContext(ctx)
			hist := FromContext(ctx)

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			pj := v.(*ent2.PrintJob)

			err = newChangeHistoryOrNoop(ctx, m.Client())
			if err != nil {
				return nil, err
			}

			col, err := pj.QueryColli().
				WithOrder().
				Only(ctx)
			if err != nil && !ent2.IsNotFound(err) {
				return nil, err
			} else if ent2.IsNotFound(err) {
				return nil, nil
			}

			_, err = m.Client().OrderHistory.Create().
				SetOrder(col.Edges.Order).
				SetDescription(hist.description).
				SetType(orderhistory.TypeCreate).
				SetTenantID(view.TenantID()).
				SetChangeHistoryID(view.ContextID()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("creating print job colli history record: %w", err)
			}

			return v, nil
		})
	}
	return hook.On(hk, ent.OpCreate)
}

func newChangeHistoryOrNoop(ctx context.Context, cli *ent2.Client) error {
	view := viewer.FromContext(ctx)
	hist := FromContext(ctx)

	createCH := cli.ChangeHistory.Create()
	if len(view.MyId()) > 0 {
		createCH.SetUserID(view.MyId())
	}

	err := createCH.
		SetTenantID(view.TenantID()).
		SetID(view.ContextID()).
		SetOrigin(hist.origin).
		OnConflictColumns(changehistory.FieldID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
