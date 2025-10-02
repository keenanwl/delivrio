package background

import (
	"context"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/shopify/ordersync"
	"log"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/connectionshopify"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/viewer"
)

func handleShopifyOrderSync(ctx context.Context) {

	db := ent.FromContext(ctx)
	shops, err := db.ConnectionShopify.Query().
		WithTenant().
		Where(connectionshopify.HasConnectionWith(connection.SyncOrders(true))).
		All(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for _, s := range shops {
		lastSyncEvent, err := db.SystemEvents.Query().
			Where(systemevents.And(
				systemevents.EventTypeEQ(systemevents.EventTypeShopifyOrderSync),
				systemevents.EventTypeIDEQ(s.ID.String()),
				systemevents.StatusEQ(systemevents.StatusSuccess),
			)).
			Order(ent.Desc(systemevents.FieldUpdatedAt)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			log.Println(err)
			continue
		}

		timeout := time.Duration(-5) * time.Second
		cutoff := time.Now().Add(timeout)

		// prevent
		lastSyncTime := s.SyncFrom
		if lastSyncEvent != nil {
			lastSyncTime = lastSyncEvent.UpdatedAt
		}

		if lastSyncTime.Before(cutoff) {
			evt, err := db.SystemEvents.Create().
				SetStatus(systemevents.StatusRunning).
				SetEventTypeID(s.ID.String()).
				SetDescription("Running...").
				SetData("").
				SetUpdatedAt(time.Now()).
				SetEventType(systemevents.EventTypeShopifyOrderSync).
				SetTenantID(s.Edges.Tenant.ID).
				Save(ctx)

			if err != nil {
				log.Println(err)
				continue
			}

			// TODO should have some sort of abstraction
			currentView := viewer.FromContext(ctx)
			v := viewer.UserViewer{
				Role:    viewer.BackgroundForTenant,
				Tenant:  s.Edges.Tenant.ID,
				Context: currentView.ContextID(),
			}
			ctxCancel, cancel := context.WithCancel(viewer.NewContext(ctx, v))
			jobContext.RunningJobs[evt.ID] = cancel

			ordersync.ProcessShopifyOrderSync(ctxCancel, s, s.SyncFrom, lastSyncTime, evt.ID)
		}

	}
}

func handleShopifyOrderCancelledSync(ctx context.Context) {

	db := ent.FromContext(ctx)
	shops, err := db.ConnectionShopify.Query().
		WithTenant().
		Where(connectionshopify.HasConnectionWith(connection.SyncOrders(true))).
		All(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	for _, s := range shops {
		lastSyncEvent, err := db.SystemEvents.Query().
			Where(systemevents.And(
				systemevents.EventTypeEQ(systemevents.EventTypeShopifyOrderCancelledSync),
				systemevents.EventTypeIDEQ(s.ID.String()),
				systemevents.StatusEQ(systemevents.StatusSuccess),
			)).
			Order(ent.Desc(systemevents.FieldUpdatedAt)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			log.Println(err)
			continue
		}

		timeout := time.Duration(-5) * time.Second
		cutoff := time.Now().Add(timeout)

		lastSyncTime := time.Time{}
		if lastSyncEvent != nil {
			lastSyncTime = lastSyncEvent.UpdatedAt
		}

		if lastSyncTime.Before(cutoff) {
			evt, err := db.SystemEvents.Create().
				SetStatus(systemevents.StatusRunning).
				SetEventTypeID(s.ID.String()).
				SetDescription("Running...").
				SetData("").
				SetUpdatedAt(time.Now()).
				SetEventType(systemevents.EventTypeShopifyOrderCancelledSync).
				SetTenantID(s.Edges.Tenant.ID).
				Save(ctx)

			if err != nil {
				log.Println(err)
				continue
			}

			// TODO should have some sort of abstraction
			currentView := viewer.FromContext(ctx)
			v := viewer.UserViewer{
				Role:    viewer.BackgroundForTenant,
				Tenant:  s.Edges.Tenant.ID,
				Context: currentView.ContextID(),
			}
			ctxCancel, cancel := context.WithCancel(viewer.NewContext(ctx, v))
			jobContext.RunningJobs[evt.ID] = cancel

			ordersync.ProcessShopifyOrderCancelledSync(ctxCancel, s, lastSyncTime, evt.ID)
		}

	}
}
