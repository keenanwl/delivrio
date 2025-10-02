package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

func ProductionBase(ctx context.Context) {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Background, Context: pulid.MustNew("CH")})

	Languages(ctx)
	SeedPlans(ctx)
	SeedSignup(ctx)
	SeedTenant(ctx)
	CarrierBrands(ctx)
	Countries(ctx)
	AllCarrierServices(ctx)
	Currency(ctx)
	AccessRights(ctx)
	ConnectionBrands(ctx)
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Background,
		Context: pulid.MustNew("CH"),
		Tenant:  GetTenantID(),
	})
	LocationTags(ctx)
	AdminUsers(ctx)
}

func AllCarrierServices(ctx context.Context) {
	DFServices(ctx)
	PNServices(ctx)
	BringServices(ctx)
	DAOServices(ctx)
	EasyPostServices(ctx)
	GLSServices(ctx)
	USPSServices(ctx)
	USPSAdditionalServices(ctx)
	USPSRateIndicators(ctx)
}

func Base(ctx context.Context) {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Background,
		Context: pulid.MustNew("CH"),
	})
	ProductionBase(ctx)
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Background,
		Context: pulid.MustNew("CH"),
		Tenant:  GetTenantID(),
	})
	Location(ctx)
	SeedCarrierConnection(ctx)
	DeliveryOption(ctx)

}

func ProductionBaseTx(ctx context.Context) {
	client := ent.FromContext(ctx)
	tx, _ := client.Tx(ctx)
	defer tx.Rollback()
	ctx = ent.NewTxContext(ctx, tx)
	ProductionBase(ctx)

	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}
