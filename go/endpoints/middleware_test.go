package endpoints

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/seed"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	b64 "encoding/base64"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCheckRESTAPICredentials(t *testing.T) {

	cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := ent.NewContext(viewer.NewContext(context.Background(), viewer.UserViewer{Role: viewer.Background, Context: pulid.MustNew("CH")}), cli)

	tx, err := cli.Tx(ctx)
	require.NoError(t, err)
	ctx = ent.NewTxContext(ctx, tx)
	seed.Base(ctx)

	validToken := "000000000-000000000-000000000"
	token := utils.HashPasswordX(validToken)
	savedToken := tx.APIToken.Create().
		SetTenantID(seed.GetTenantID()).
		SetName("Working key").
		SetHashedToken(token).
		SetUserID(seed.GetAdminUser()).
		SaveX(ctx)
	err = tx.Commit()
	require.NoError(t, err)

	r := chi.NewRouter()
	r.Use(AddClient(cli))
	r.Use(CheckRESTAPICredentials())
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		httputils.JSONResponse(w, http.StatusOK, map[string]string{"success": "yes"})
	})

	// Define a test case with different scenarios.
	testCases := []struct {
		Name           string
		Header         map[string]string
		ExpectedStatus int
	}{
		{
			Name: "ValidAPIKeyValidID",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte(fmt.Sprintf("%v:%v", validToken, savedToken.ID)),
				),
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Name: "InvalidBase64",
			Header: map[string]string{
				DelivrioApiHeaderKey: fmt.Sprintf("%v:%v", validToken, savedToken.ID),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name: "ValidAPIKeyInvalidID",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte(fmt.Sprintf("%v:%v", validToken, pulid.MustNew("AT"))),
				),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name: "InvalidAPIKeyValidID",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte(fmt.Sprintf("%v:%v", "000000000-000000000-000000001", savedToken.ID)),
				),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name: "InvalidAPIKey",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte("invalid-api-key:invalid-id"),
				),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name: "MissingID",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte(fmt.Sprintf("%v:", validToken)),
				),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name: "MissingKey",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte(fmt.Sprintf(":%v", savedToken.ID)),
				),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name:           "NoAPIKey",
			Header:         map[string]string{},
			ExpectedStatus: http.StatusUnauthorized,
		},
		{
			Name: "Empty",
			Header: map[string]string{
				DelivrioApiHeaderKey: b64.StdEncoding.EncodeToString(
					[]byte(""),
				),
			},
			ExpectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new HTTP request with the specified headers.
			req, err := http.NewRequest(http.MethodGet, "/test", nil)
			require.NoError(t, err)
			for key, value := range tc.Header {
				req.Header.Set(key, value)
			}

			w := httptest.NewRecorder()

			// Perform the request using the router.
			r.ServeHTTP(w, req)

			// Check if the response status code matches the expected status.
			require.Equal(t, tc.ExpectedStatus, w.Code, tc.Name)
		})
	}
}
