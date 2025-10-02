package background

import (
	"context"
	dbsql "database/sql"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/connectionshopify"
	"delivrio.io/go/shopify/productsync"
	"delivrio.io/go/viewer"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/shared-utils/pulid"
)

type backgroundJobSync struct {
	Mu          sync.Mutex
	RunningJobs map[pulid.ID]context.CancelFunc
}

func (b *backgroundJobSync) MayAdd() bool {
	return len(b.RunningJobs) <= 10
}

var jobContext = backgroundJobSync{
	Mu:          sync.Mutex{},
	RunningJobs: make(map[pulid.ID]context.CancelFunc),
}

func HandleDBVacuum(ctx context.Context, db *dbsql.DB, lastVacuum time.Time) time.Time {

	if time.Since(lastVacuum) < time.Hour*24 {
		return lastVacuum
	}

	_, err := db.ExecContext(ctx, `SELECT pg_repack.repack('public.system_events')`)
	if err != nil {
		log.Printf("background job vacuum failed: %v", err)
		return lastVacuum
	}

	return time.Now()
}

func HandleBackgroundJobs(ctx context.Context, once bool) {

	for {

		handleFailed(ctx)
		handleShopifyProductSync(ctx)
		handleShopifyOrderSync(ctx)
		handleShopifyOrderCancelledSync(ctx)
		handleNotifications(ctx)
		handleCancelledShipmentSync(ctx)
		handleLogCleanup(ctx)

		if once {
			break
		}

		time.Sleep(5 * time.Second)
		log.Println("Running background task")
	}

}

func handleFailed(ctx context.Context) {

	db := ent.FromContext(ctx)
	jobContext.Mu.Lock()
	defer jobContext.Mu.Unlock()

	running, err := db.SystemEvents.Query().
		Where(
			systemevents.StatusEQ(systemevents.StatusRunning),
		).All(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	timeout := time.Duration(-5) * time.Minute
	cutoff := time.Now().Add(timeout)

	for _, r := range running {

		if r.UpdatedAt.Before(cutoff) {
			if jobContext.RunningJobs[r.ID] != nil {
				// Cancel job
				jobContext.RunningJobs[r.ID]()
			}
			err := db.SystemEvents.Update().
				SetStatus(systemevents.StatusFail).
				SetData(fmt.Sprintf("cancelled: job did not finish within %v minute timeout", timeout)).
				SetUpdatedAt(time.Now()).
				Where(systemevents.ID(r.ID)).
				Exec(ctx)
			if err != nil {
				log.Printf("Failed to cancel job: %v", err)
				return
			}
		}

	}

}

func handleShopifyProductSync(ctx context.Context) {

	db := ent.FromContext(ctx)
	jobContext.Mu.Lock()
	defer jobContext.Mu.Unlock()

	shops, err := db.ConnectionShopify.Query().
		WithTenant().
		Where(connectionshopify.HasConnectionWith(connection.SyncProducts(true))).
		All(ctx)
	if err != nil {
		log.Printf("background: query connections: %v\n", err)
		return
	}

	// Prevents 1 shop from always being synced first
	rand.Shuffle(len(shops), func(i, j int) {
		shops[i], shops[j] = shops[j], shops[i]
	})

	for _, s := range shops {
		lastSyncEvent, err := db.SystemEvents.Query().
			Where(systemevents.And(
				systemevents.EventTypeEQ(systemevents.EventTypeShopifyProductSync),
				systemevents.EventTypeIDEQ(s.ID.String()),
				systemevents.StatusEQ(systemevents.StatusSuccess),
			)).
			Order(ent.Desc(systemevents.FieldUpdatedAt)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			log.Printf("background: product sync: query system events: %v\n", err)
			continue
		}

		syncInterval := time.Duration(-5) * time.Minute
		cutoff := time.Now().Add(syncInterval)

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
				SetEventType(systemevents.EventTypeShopifyProductSync).
				SetTenantID(s.Edges.Tenant.ID).
				Save(ctx)
			if err != nil {
				log.Println(err)
			}

			currentView := viewer.FromContext(ctx)
			v := viewer.UserViewer{
				Role:   viewer.BackgroundForTenant,
				Tenant: s.Edges.Tenant.ID,
				// TODO: this needs to be refreshed
				Context: currentView.ContextID(),
			}
			ctxCancel, cancel := context.WithCancel(viewer.NewContext(ctx, v))
			jobContext.RunningJobs[evt.ID] = cancel

			productsync.ProcessShopifyProductSync(ctxCancel, *s, lastSyncTime, evt.ID)
		}

	}

}
