package endpoints

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/apitoken"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	b64 "encoding/base64"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"net/http"
	"strings"
	"time"
)

const DelivrioApiHeaderKey = "X-DELIVRIO-Key"

func CheckRESTAPICredentials() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := viewer.NewBackgroundContext(r.Context())
			cli := ent.FromContext(r.Context())
			apiKey := r.Header.Get(DelivrioApiHeaderKey)

			decodedToken, err := b64.StdEncoding.DecodeString(apiKey)
			if err != nil {
				http.Error(w, "1Authentication failed", http.StatusUnauthorized)
				return
			}

			tokenAndID := strings.Split(string(decodedToken), ":")
			if len(tokenAndID) != 2 {
				http.Error(w, "2Authentication failed", http.StatusUnauthorized)
				return
			}

			token, err := cli.APIToken.Query().
				WithUser().
				Where(apitoken.ID(pulid.ID(tokenAndID[1]))).
				Only(ctx)
			if token == nil || err != nil {
				http.Error(w, "3Authentication failed", http.StatusUnauthorized)
				return
			}

			if !utils.CheckPasswordHash(tokenAndID[0], token.HashedToken) {
				http.Error(w, "4Authentication failed", http.StatusUnauthorized)
				return
			}

			err = token.Update().SetLastUsed(time.Now()).Exec(ctx)
			if err != nil {
				http.Error(w, "5Authentication failed", http.StatusUnauthorized)
				return
			}

			tenantCTX := viewer.UserViewer{
				Role:    viewer.CustomerTenant,
				Tenant:  token.TenantID,
				MyID:    token.Edges.User.ID,
				Context: pulid.MustNew("CH"),
			}

			r = r.WithContext(viewer.NewContext(r.Context(), tenantCTX))

			next.ServeHTTP(w, r)
			return
		})
	}

}

func AddSpan(serverID string, parentName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceCtx, span := otel.Tracer(serverID).
				Start(r.Context(), parentName)
			r = r.WithContext(traceCtx)
			next.ServeHTTP(w, r)
			span.End()
		})
	}

}

func AddTX(cli *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tx, err := cli.Tx(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer func() {
				if r := recover(); r != nil {
					_ = tx.Rollback()
					panic(r)
				}
			}()

			ctx := r.Context()
			ctx = ent.NewTxContext(ctx, tx)
			r = r.WithContext(ctx)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			if utils.StatusInList(ww.Status(), []int{http.StatusOK, http.StatusCreated, http.StatusAccepted}) {
				err := tx.Commit()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

			// Rollback in case of error
			utils.Rollback(tx, err)
			return
		})
	}
}

func AddClient(cli *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(ent.NewContext(r.Context(), cli))
			next.ServeHTTP(w, r)
		})
	}
}
