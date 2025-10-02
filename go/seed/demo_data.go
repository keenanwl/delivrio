package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

func DemoData(ctx context.Context) {

	c := ent.FromContext(ctx)

	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Background,
		Context: pulid.MustNew("CH"),
	})

	ctx, tx, err := c.OpenTx(ctx)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	Base(ctx)
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Background,
		Context: pulid.MustNew("CH"),
		Tenant:  GetTenantID(),
	})
	DemoLocation(ctx)
	Products(ctx, 2)

	SeedOrders(ctx, 3)

	tx.Commit()

}
