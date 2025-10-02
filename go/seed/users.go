//go:build !internal

package seed

import (
	"context"
	"delivrio.io/go/ent"
)

// Internal builds have other default users
func AdminUsers(ctx context.Context) {
	tx := ent.TxFromContext(ctx)
	u1 := tx.User.Create().
		SetName("Admin").
		SetEmail("admin@example.com").
		SetPhoneNumber("").
		SetPassword("delivrio++++").
		SetTenantID(tenantID).
		SetLanguageID(languageID).
		SetIsAccountOwner(true).
		SetIsGlobalAdmin(true).
		SaveX(ctx)
	adminUserID = u1.ID
}
