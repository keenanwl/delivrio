package endpoints

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/otkrequests"
	"delivrio.io/go/viewer"
)

func generateToken() string {

	token, err := GenerateRandomString(32)
	if err != nil {
		panic(err)
	}

	return token

}

func replaceToken(ctx context.Context, user *ent.User, newToken string) error {
	tx := ent.TxFromContext(ctx)
	_, err := tx.OTKRequests.Create().
		SetOtk(newToken).
		SetUsers(user).
		SetTenantID(user.TenantID).
		Save(viewer.NewBackgroundContext(ctx))
	return err
}

// Refactor to return err instead of bool
func ValidOtk(ctx context.Context, otk string) (bool, ent.User) {
	tx := ent.TxFromContext(ctx)
	// Security
	time.Sleep(1 * time.Second)

	request, err := tx.OTKRequests.Query().
		Where(otkrequests.Otk(otk)).
		Only(viewer.NewBackgroundContext(ctx))

	if err != nil {
		return false, ent.User{}
	}

	if u, err := request.QueryUsers().Only(viewer.NewBackgroundContext(ctx)); err == nil {
		return true, *u
	}

	return false, ent.User{}

}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an Err if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
