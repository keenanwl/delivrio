package background

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/connectionlookup"
	"delivrio.io/go/ent/systemevents"
	"log"
	"time"
)

func handleLogCleanup(ctx context.Context) {
	jobContext.Mu.Lock()
	defer jobContext.Mu.Unlock()
	cli := ent.FromContext(ctx)

	oneDay := time.Hour * 24 * -1
	fiveDays := time.Hour * 24 * -5

	_, err := cli.SystemEvents.Delete().
		Where(systemevents.CreatedAtLTE(time.Now().Add(oneDay))).
		Exec(ctx)
	if err != nil {
		log.Printf("error: cleanup logs: %s", err)
	}

	_, err = cli.ConnectionLookup.Delete().
		Where(connectionlookup.CreatedAtLTE(time.Now().Add(fiveDays))).
		Exec(ctx)
	if err != nil {
		log.Printf("error: cleanup logs: %s", err)
	}
}
