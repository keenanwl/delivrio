//go:build internal

package seed

import (
	"context"
	"delivrio.io/go/ent"
	"encoding/json"
	"os"
)

type InternalUser struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Password       string `json:"password"`
	TenantID       string `json:"tenantID"`
	LanguageID     string `json:"languageID"`
	IsAccountOwner bool   `json:"isAccountOwner"`
	IsGlobalAdmin  bool   `json:"isGlobalAdmin"`
	ID             string `json:"id,omitempty"`
}

// Internal builds have other default users configured outside this repo
func AdminUsers(ctx context.Context) {
	tx := ent.TxFromContext(ctx)

	usersFile, err := os.ReadFile("../../delivrio-internal/config/seed/users_internal.json")
	if err != nil {
		panic(err)
	}

	var input []InternalUser
	err = json.Unmarshal(usersFile, &input)
	if err != nil {
		panic(err)
	}

	for i, user := range input {
		u1 := tx.User.Create().
			SetName(user.Name).
			SetEmail(user.Email).
			SetPhoneNumber("").
			SetPassword(user.Password).
			SetTenantID(tenantID).
			SetLanguageID(languageID).
			SetIsAccountOwner(user.IsAccountOwner).
			SetIsGlobalAdmin(user.IsGlobalAdmin).
			SaveX(ctx)
		if i == 0 {
			adminUserID = u1.ID
		}
	}

}
