package endpoints

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"errors"
	"fmt"
	"github.com/mailgun/mailgun-go"
	"log"
	"net/http"
	"net/url"
	"time"
)

var BaseURL *url.URL
var AppConfig *appconfig.DelivrioConfig

type RequestEmailRequest struct {
	Email string `json:"email"`
}

type RequestEmailResponse struct {
	Success bool `json:"success"`
}

func RequestEmailHandler(w http.ResponseWriter, r *http.Request) {

	var email RequestEmailRequest

	err := httputils.UnmarshalRequestBody(r, &email)
	if err != nil {
		log.Println("request login email: ", err)
		httputils.JSONResponse(w, http.StatusInternalServerError, RequestEmailResponse{
			Success: false,
		})
		return
	}

	err = sendEmail(r.Context(), email.Email)
	if err != nil {
		log.Println("request login email: ", err)
		httputils.JSONResponse(w, http.StatusInternalServerError, RequestEmailResponse{
			Success: false,
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, RequestEmailResponse{
		Success: true,
	})
	return

}

const emailNotFound = `email not found`

func sendEmail(ctx context.Context, email string) error {

	u, err := userExists(ctx, email)
	if err != nil {
		return err
	}

	if u.ID.String() != "" {

		token := generateToken()

		err := replaceToken(ctx, u, token)
		if err != nil {
			return err
		}

		msg := fmt.Sprintf(
			`
Hi %s,
Please find your link to reset your password:
%s/password-reset?otk=%s
		`,
			u.Email,
			BaseURL,
			token,
		)

		_, err = SendSimpleMessage(*u, msg)
		if err != nil {
			return err
		}

	} else {
		return errors.New(emailNotFound)
	}

	return err

}

func userExists(ctx context.Context, email string) (*ent.User, error) {
	tx := ent.TxFromContext(ctx)
	return tx.User.Query().
		Where(user.Email(email)).
		Only(viewer.NewBackgroundContext(ctx))
}

// Refactors to utils version?
func SendSimpleMessage(respondent ent.User, msg string) (string, error) {

	mg := mailgun.NewMailgun(AppConfig.Email.Mailgun.MGDomain, AppConfig.Email.Mailgun.MGAPIKey)
	mg.SetAPIBase(AppConfig.Email.Mailgun.MGURL)
	mg.Client().Timeout = time.Second * 5

	m := mg.NewMessage(
		"DELIVRIO <no-reply@mail.delivrio.io>",
		"Reset your DELIVRIO password",
		msg,
		respondent.Email,
	)
	_, id, err := mg.Send(m)

	return id, err
}
