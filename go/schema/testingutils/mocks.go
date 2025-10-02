package testingutils

import (
	"context"

	"delivrio.io/go/ent"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

const TenantID = pulid.ID("T9999999")
const UserID = pulid.ID("U0000000")

func DefaultTenant(client *ent.Client, ctx context.Context) *ent.TenantCreate {

	plan := client.Plan.Create().
		SetRank(0).
		SetPriceDkk(1000).
		SetLabel("Plan 1").
		SaveX(ctx)

	lang := client.Language.Create().SetLabel("DK").SaveX(ctx)

	return client.Tenant.Create().
		SetPlan(plan).
		SetDefaultLanguage(lang).
		SetName("Tenant 1").
		SetID(TenantID)
}
func DefaultUserViewer(client *ent.Client) context.Context {

	ctx := viewer.NewContext(context.Background(), viewer.UserViewer{
		Role:   viewer.CustomerTenant,
		Tenant: TenantID,
		MyID:   UserID,
	})

	return ent.NewContext(ctx, client)
}
func DefaultBackgroundViewer() context.Context {
	return viewer.NewContext(context.Background(), viewer.UserViewer{
		Role: viewer.Background,
	})
}
func DefaultUser(client *ent.Client) *ent.UserCreate {
	return client.User.Create().
		SetID(UserID).
		SetTenantID(TenantID).
		SetName("J{ø¤hn").
		SetHash("1919191").
		SetSurname("Powell").
		SetEmail("jp@example.com").
		SetPhoneNumber("+45 71 20 11 18")
}
