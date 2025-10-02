package hookgls

import (
	"context"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"entgo.io/ent"
)

func UpdateDeliveryOptionGLS() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryOptionGLSFunc(func(ctx context.Context, m *ent2.DeliveryOptionGLSMutation) (ent.Value, error) {
			//errs := hooks.NewValidationError()

			/*name, nameSet := m.WorkstationName()
			if nameSet && len(name) > 10 {
				errs.SetError(deliveryoptiongls.FieldName, "name may not be longer than 10 characters")
				return nil, errs
			}*/

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}
