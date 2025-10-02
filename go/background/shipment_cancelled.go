package background

import (
	"context"
	"delivrio.io/go/carrierapis/easypostapis"
	"delivrio.io/go/carrierapis/uspsapis"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"fmt"
	"log"
	"time"
)

func handleCancelledShipmentSync(ctx context.Context) {
	jobContext.Mu.Lock()
	defer jobContext.Mu.Unlock()
	cli := ent.FromContext(ctx)

	allTenants, err := getSystemTenants(ctx)
	if err != nil {
		log.Println("sync cancelled: get tenants: ", err)
		return
	}

	for _, t := range allTenants {

		nextCtx := viewer.MergeViewerTenantID(viewer.NewBackgroundContext(ctx), t.ID)
		systemEvent, err := newSystemEvent(nextCtx, systemevents.EventTypeSyncCancelledShipments)
		if err != nil {
			log.Printf("sync cancelled: new system event: %s", err)
			continue
		}

		ship, err := cli.ShipmentParcel.Query().
			Where(shipmentparcel.And(
				shipmentparcel.HasShipmentWith(shipment.StatusEQ(shipment.StatusDeleted)),
				shipmentparcel.CancelSyncedAtIsNil(),
			)).
			Limit(10).
			All(ctx)
		if err != nil {
			log.Printf("sync cancelled: %v", err)
			continue
		}

		total := 0
		allErrors := make([]error, 0)

		for _, s := range ship {
			car, err := s.QueryShipment().
				QueryCarrier().
				WithCarrierBrand().
				Only(ctx)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}

			switch car.Edges.CarrierBrand.InternalID {
			case carrierbrand.InternalIDUSPS:
				agreement, err := car.QueryCarrierUSPS().
					Only(ctx)
				if err != nil {
					allErrors = append(allErrors, err)
					continue
				}
				err = uspsapis.CancelByTrackingCode(ctx, agreement, s.ItemID)
				if err != nil {
					allErrors = append(allErrors, err)
					continue
				}
			case carrierbrand.InternalIDEasyPost:
				agreement, err := car.QueryCarrierEasyPost().
					Only(ctx)
				if err != nil {
					allErrors = append(allErrors, err)
					continue
				}
				shipEasyPost, err := s.QueryShipment().
					QueryShipmentEasyPost().
					Only(ctx)
				if err != nil {
					allErrors = append(allErrors, err)
					continue
				}
				err = easypostapis.CancelByEasyPostID(ctx, agreement, shipEasyPost.EpShipmentID)
				if err != nil {
					allErrors = append(allErrors, err)
					continue
				}
			default:
				// Not implemented; may or may not exist
			}

			err = s.Update().
				SetCancelSyncedAt(time.Now()).
				Exec(ctx)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}
			total++
		}

		status := systemevents.StatusSuccess
		if len(allErrors) > 0 {
			status = systemevents.StatusFail
		}

		err = cli.SystemEvents.Update().
			Where(systemevents.ID(systemEvent.ID)).
			SetStatus(status).
			SetDescription(fmt.Sprintf("Shipment cancelled sync success for %v collis", total)).
			SetData(utils.JoinErrors(allErrors, ", ")).
			Exec(nextCtx)
		if err != nil {
			log.Println("updating sync cancelled system event failed: ", err)
			continue
		}

	}

}
