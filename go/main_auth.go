package main

import (
	"delivrio.io/go/appconfig"
	"delivrio.io/go/ent"
	entUser "delivrio.io/go/ent/user"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"time"
)

func authHandler(conf appconfig.DelivrioConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var loginVals login
		if err := httputils.UnmarshalRequestBody(r, &loginVals); err != nil {
			httputils.JSONResponse(w, http.StatusUnauthorized, httputils.Map{"message": "unrecognized credentials"})
			return
		}
		_, isValid := isValidRespondentCode(r, loginVals.Password, loginVals.Email)
		if !isValid {
			httputils.JSONResponse(w, http.StatusUnauthorized, httputils.Map{"message": "unrecognized credentials"})
			return
		}

		// Create the Claims
		claims := httputils.Map{
			"email": loginVals.Email,
			// We validate every request for now, so a long expiration is not as important
			// "exp" is defined as part of the standard (?)
			"exp": Clock().Add(time.Hour * 72).Unix(),
		}

		tokenAuth := jwtauth.New("HS256", []byte(conf.JWTKey), nil, jwt.WithClock(Clock))

		_, tokenString, err := tokenAuth.Encode(claims)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, httputils.Map{"message": "unrecognized credentials"})
			return
		}

		httputils.JSONResponse(w, http.StatusOK, httputils.Map{"token": tokenString})
		return
	}

}

type login struct {
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

func isValidRespondentCode(r *http.Request, password string, email string) (pulid.ID, bool) {
	// Temporary ctx override to get around privacy
	// since we need to be able to lookup the user to log them in
	v := viewer.UserViewer{Role: viewer.Background}
	ctx := viewer.NewContext(r.Context(), v)
	tx := ent.TxFromContext(r.Context())

	user, err := tx.User.
		Query().
		Where(entUser.Email(email)).
		Only(ctx)
	if err != nil {
		return "", false
	}

	return user.ID, utils.CheckPasswordHash(password, user.Hash)

}
