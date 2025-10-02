package viewer

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/shared-utils/pulid"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

const IdentityKey = "email"

func AddAnonymousViewer() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(NewContext(r.Context(), UserViewer{
				Role:    Anonymous,
				Context: pulid.MustNew("CH"),
			})))
		})
	}
}

func AddJWTViewer() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				httputils.JSONResponse(w, http.StatusUnauthorized,
					httputils.Map{"error": "JWT claims missing or invalid"})
				return
			}

			email, ok := claims[IdentityKey].(string)
			if ok {
				next.ServeHTTP(w, AddViewerByEmail(r, email))
				return
			}

			httputils.JSONResponse(w, http.StatusUnauthorized, httputils.Map{"error": "JWT missing identity claim"})
			return
		})
	}
}

func AddViewerByEmail(r *http.Request, email string) *http.Request {

	ctxID := pulid.MustNew("CH")
	client := ent.FromContext(r.Context())
	v := UserViewer{
		Role:    Anonymous,
		Context: ctxID,
	}
	if len(email) > 0 {

		u := client.User.Query().
			WithTenant().
			Where(user.EmailEqualFold(email)).
			OnlyX(NewContext(
				r.Context(),
				UserViewer{
					Role:    Background,
					Context: ctxID,
				},
			))

		v = UserViewer{
			Role:    CustomerTenant,
			Tenant:  u.Edges.Tenant.ID,
			MyID:    u.ID,
			Context: ctxID,
		}
		if u.IsGlobalAdmin {
			v = UserViewer{
				Role:    Admin,
				Tenant:  u.Edges.Tenant.ID,
				MyID:    u.ID,
				Context: ctxID,
			}
		}

	}

	return r.WithContext(NewContext(r.Context(), v))

}
