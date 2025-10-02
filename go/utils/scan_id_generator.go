package utils

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/viewer"
	"entgo.io/ent/dialect/sql"
	"sync/atomic"
)

var BarcodeCounter atomic.Int64

// For tests, otherwise should be overwritten
func init() {
	BarcodeCounter.Store(1_000_000)
}

func InitBarcodeCounter(ctx context.Context) error {
	viewCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Background})
	cli := ent.FromContext(viewCtx)
	sp, err := cli.Colli.Query().
		Order(colli.ByID(sql.OrderDesc())).
		First(viewCtx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	} else if ent.IsNotFound(err) {
		BarcodeCounter.Store(1_000_001)
		return nil
	}
	BarcodeCounter.Store(sp.InternalBarcode)
	return nil
}

// Since we can't rely on auto increment from a single database type
// Global across all tenants. However unique index only enforced
// on the teant level.
func NextShipmentBarcodeID() int64 {
	return BarcodeCounter.Add(1)
}
