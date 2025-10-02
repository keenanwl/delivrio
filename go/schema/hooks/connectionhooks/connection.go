package connectionhooks

import (
	"context"
	"delivrio.io/go/appconfig"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/deliveryrule"
	"delivrio.io/go/ent/hook"
	"entgo.io/ent"
	"fmt"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("connectionhooks: may not set config twice")
	}
	conf = c
	confSet = true
}

func UpdateCurrencyDependencies() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ConnectionFunc(func(ctx context.Context, m *ent2.ConnectionMutation) (ent.Value, error) {

			nextOrderSyncStatus, exists := m.SyncOrders()
			if exists && nextOrderSyncStatus && conf.LimitedSystem {
				return nil, fmt.Errorf("syncing orders is not supported on this system, please disable")
			}

			nextCurrency, exists := m.CurrencyID()
			if exists {
				updateIDs, err := m.IDs(ctx)
				if err != nil {
					return nil, err
				}
				for _, id := range updateIDs {
					conn, err := m.Client().Connection.Query().
						Where(connection.ID(id)).
						WithCurrency().
						Only(ctx)
					if err != nil {
						return nil, err
					}

					if conn.Edges.Currency.ID != nextCurrency {
						allAffectedDeliveryRules, err := conn.QueryDeliveryOption().
							QueryDeliveryRule().
							IDs(ctx)
						if err != nil {
							return nil, err
						}

						err = m.Client().DeliveryRule.Update().
							SetCurrencyID(nextCurrency).
							Where(deliveryrule.IDIn(allAffectedDeliveryRules...)).
							Exec(ctx)
						if err != nil {
							return nil, fmt.Errorf("update connection currency: %w", err)
						}
					}
				}
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}
