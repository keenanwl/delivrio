package deliveryoptionhooks

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/hook"
	"entgo.io/ent"
	"fmt"
)

// Prevent Webshipper & Shipmondo from both being enabled
func PreventConflictingIntegrations() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.DeliveryOptionFunc(func(ctx context.Context, m *ent2.DeliveryOptionMutation) (ent.Value, error) {
			var err error

			webshipperEnabled, webshipperExists := m.WebshipperIntegration()
			shipmondoEnabled, shipmondoExists := m.ShipmondoIntegration()

			// OldX is not allowed for multi-update
			if (!shipmondoExists && webshipperExists) || (!webshipperExists && shipmondoExists) {
				webshipperEnabled, err = m.OldShipmondoIntegration(ctx)
				if err != nil {
					return nil, err
				}
				shipmondoEnabled, err = m.OldWebshipperIntegration(ctx)
				if err != nil {
					return nil, err
				}
			}

			if webshipperEnabled && shipmondoEnabled {
				return nil, fmt.Errorf("webshipper and shipmondo may not be simultaneously enabled")
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
