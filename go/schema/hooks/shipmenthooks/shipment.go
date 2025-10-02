package shipmenthooks

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/consolidation"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/schema/hooks/history"
	"entgo.io/ent"
	"fmt"
)

// UpdateShipmentConnectedEntities updates the connected consolidations/collis/pallets,
// so the application code only has to be concerned with deleting the shipment itself.
func UpdateShipmentConnectedEntities() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ShipmentFunc(func(ctx context.Context, m *ent2.ShipmentMutation) (ent.Value, error) {
			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			nextStatus, exists := m.Status()
			if exists && nextStatus == shipment.StatusDeleted {
				for _, id := range ids {

					ship, err := m.Client().Shipment.Query().
						WithConsolidation().
						WithShipmentParcel().
						WithShipmentPallet().
						Where(shipment.ID(id)).
						Only(ctx)
					if err != nil {
						return nil, err
					}

					if ship.Edges.Consolidation != nil {
						err := ship.Edges.Consolidation.Update().
							ClearShipment().
							AddCancelledShipmentIDs(id).
							SetStatus(consolidation.StatusPending).
							Exec(ctx)
						if err != nil {
							return nil, err
						}
					}

					// Ignore ShipmentPallet status since the controlling
					// "Shipment" has the status deleted.
					for _, sp := range ship.Edges.ShipmentPallet {
						pa, err := sp.QueryPallet().
							Only(ctx)
						if err != nil && !ent2.IsNotFound(err) {
							return nil, err
						} else if ent2.IsNotFound(err) {
							// May have already been removed in the rest of the application
							continue
						}
						err = pa.Update().
							ClearShipmentPallet().
							AddCancelledShipmentPallet(sp).
							Exec(ctx)
						if err != nil {
							return nil, err
						}
					}

					// Ignore ShipmentParcel status since the controlling
					// "Shipment" has the status deleted.
					for _, p := range ship.Edges.ShipmentParcel {
						c, err := p.QueryColli().
							Only(ctx)
						if err != nil && !ent2.IsNotFound(err) {
							return nil, err
						} else if ent2.IsNotFound(err) {
							// May have already been removed in the rest of the application
							continue
						}
						err = c.Update().
							ClearShipmentParcel().
							AddCancelledShipmentParcel(p).
							SetStatus(colli.StatusPending).
							Exec(history.FromContext(ctx).
								SetDescription(fmt.Sprintf("Update colli status after shipment cancelled: %v", colli.StatusPending)).
								Ctx())
						if err != nil {
							return nil, err
						}
					}

				}
			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}
