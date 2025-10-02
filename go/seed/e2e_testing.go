package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

func E2E(ctx context.Context) {
	cli := ent.FromContext(ctx)
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Background,
		Context: pulid.MustNew("CH"),
	})
	ctx, tx, _ := cli.OpenTx(ctx)
	E2EUsers(ctx)
	tx.Commit()
}

func E2EUsers(ctx context.Context) {
	tx := ent.TxFromContext(ctx)
	tx.User.Create().
		SetName("Blinky").
		SetSurname("Bill").
		SetEmail("bb@example.com").
		SetPhoneNumber("+45 99 11 11 00").
		SetPassword("000111kkddmmmaasasÆÆ!`'~^~¨").
		SetTenantID(tenantID).
		SetLanguageID(languageID).
		SetIsAccountOwner(false).
		SetIsGlobalAdmin(false).
		SaveX(ctx)
}
