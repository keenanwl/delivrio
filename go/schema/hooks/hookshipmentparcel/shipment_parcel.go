package hookshipmentparcel

import (
	"context"
	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/hook"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/schema/hooks"
	"delivrio.io/go/schema/hooks/history"
	"entgo.io/ent"
	"fmt"
	"time"
)

func CreateShipmentParcelDeliveryEstimate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ShipmentParcelFunc(func(ctx context.Context, m *ent2.ShipmentParcelMutation) (ent.Value, error) {
			tx := ent2.TxFromContext(ctx)
			colliID, _ := m.ColliID()

			do, err := tx.Colli.Query().
				Where(
					colli.ID(colliID),
				).
				QueryDeliveryOption().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			// TODO: more prevision
			estimate := time.Duration(do.DeliveryEstimateFrom) * time.Hour * 24
			m.SetExpectedAt(time.Now().Add(estimate))

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent2.OpCreate)

}

func CreateShipmentChangeColliStatus() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ShipmentParcelFunc(func(ctx context.Context, m *ent2.ShipmentParcelMutation) (ent.Value, error) {
			errs := hooks.NewValidationError()

			colliID, exists := m.ColliID()
			if !exists {
				return nil, fmt.Errorf("colli is required when creating a new ShipmentParcel")
			}

			mutateVal, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			err = m.Client().Colli.Update().
				Where(colli.ID(colliID)).
				SetStatus(colli.StatusDispatched).
				Exec(history.NewConfig(ctx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Shipment created; colli marked as Dispatched").
					Ctx())
			if err != nil {
				return nil, err
			}

			if len(errs.InvalidFields(ctx)) > 0 {
				return nil, errs
			}

			return mutateVal, nil
		})
	}
	// When a shipment is cancelled/deleted we should update the status as well
	return hook.On(hk, ent2.OpCreate)
}

func UpdateShipmentStatusOnShipmentParcelMutation() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ShipmentParcelFunc(func(ctx context.Context, m *ent2.ShipmentParcelMutation) (ent.Value, error) {
			mutateVal, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			updateIDs, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			shipmentParcels, err := m.Client().ShipmentParcel.Query().
				WithShipment().
				Where(shipmentparcel.IDIn(updateIDs...)).
				All(ctx)
			if err != nil {
				return nil, err
			}

			for _, sp := range shipmentParcels {
				pendingParcelCount, err := sp.QueryShipment().
					QueryShipmentParcel().
					Where(shipmentparcel.StatusEQ(shipmentparcel.StatusPending)).
					Count(ctx)
				if err != nil {
					return nil, err
				}

				printedParcelCount, err := sp.QueryShipment().
					QueryShipmentParcel().
					Where(shipmentparcel.StatusEQ(shipmentparcel.StatusPrinted)).
					Count(ctx)
				if err != nil {
					return nil, err
				}

				if pendingParcelCount == 0 {
					err = m.Client().Shipment.Update().
						SetStatus(shipment.StatusDispatched).
						Where(shipment.ID(sp.Edges.Shipment.ID)).
						Exec(ctx)
					if err != nil {
						return nil, err
					}
				} else if printedParcelCount > 0 {
					err = m.Client().Shipment.Update().
						SetStatus(shipment.StatusPartially_dispatched).
						Where(shipment.ID(sp.Edges.Shipment.ID)).
						Exec(ctx)
					if err != nil {
						return nil, err
					}
				}

			}

			return mutateVal, nil
		})
	}
	// When a shipment is cancelled/deleted we should update the status as well
	return hook.On(hk, ent2.OpUpdate|ent2.OpUpdateOne)
}
