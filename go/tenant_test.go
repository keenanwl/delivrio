package main

import (
	"context"
	"testing"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/viewer"
)

func TestTenantPrivacy(t *testing.T) {
	ctx := context.Background()
	cli := open(ctx, t)
	defer cli.Close()

	viewAdminctx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Background})

	l := cli.Language.Create().
		SetLabel("EN").
		SaveX(viewAdminctx)

	p1 := cli.Plan.Create().
		SetLabel("Free").
		SetPriceDkk(0).
		SetRank(0).
		SaveX(viewAdminctx)

	tenant := cli.Tenant.Create().
		SetPlan(p1).
		SetDefaultLanguage(l).
		SetName("ÆÅø ` 111 ").
		SaveX(viewAdminctx)
	queryUser := cli.User.Create().
		SetName("Nerø").
		SetTenantID(tenant.ID).
		SetPhoneNumber("").
		SetPassword("").
		SetEmail("address.email@example.com").
		SaveX(viewAdminctx)

	tenant2 := cli.Tenant.Create().
		SetPlan(p1).
		SetDefaultLanguage(l).
		SetName("Tenant 2").
		SaveX(viewAdminctx)
	cli.User.Create().
		SetName("2").
		SetTenantID(tenant2.ID).
		SetPhoneNumber("").
		SetPassword("").
		SetEmail("address.email2@example.com").
		SaveX(viewAdminctx)

	cli.User.Create().
		SetName("3").
		SetTenantID(tenant2.ID).
		SetPhoneNumber("").
		SetPassword("").
		SetEmail("address.email3@example.com").
		SaveX(viewAdminctx)

	view1ctx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.CustomerTenant, Tenant: tenant.ID})

	users, err := cli.User.Query().
		Where(user.IDNotIn(queryUser.ID)).
		All(view1ctx)
	if len(users) != 0 || err != nil {
		t.Fatalf("expected all users to be filtered from view: %v", err)
	}

}

func open(ctx context.Context, t *testing.T) *ent.Client {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(
		//
		),
	}
	cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)

	// Run the auto migration tool.
	if err := cli.Schema.Create(ctx); err != nil {
		t.Fatalf("failed creating schema resources: %v", err)
	}

	return cli

}
