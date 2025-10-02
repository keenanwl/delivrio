package endpoints

import (
	"delivrio.io/go/ent/tenant"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/utils/httputils"
	"fmt"
	"net/http"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/language"
	plan2 "delivrio.io/go/ent/plan"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

type RegistrationResponse struct {
	Success   bool     `json:"success"`
	Message   string   `json:"message"`
	UserPULID pulid.ID `json:"user_pulid"`
}

type RegistrationUserTenantInput struct {
	UserInput   ent.CreateUserInput   `json:"user_input"`
	TenantInput ent.CreateTenantInput `json:"tenant_input"`
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	ctx := viewer.NewBackgroundContext(r.Context())

	tx := ent.TxFromContext(r.Context())

	var input RegistrationUserTenantInput

	err := httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	plan, err := tx.Plan.Query().Where(plan2.LabelEQ("Free")).Only(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	lang, err := tx.Language.Query().
		Where(language.InternalIDEQ(language.InternalIDEN)).
		Only(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	count, err := tx.Tenant.Query().
		Where(tenant.NameEqualFold(input.TenantInput.Name)).
		Count(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	} else if count > 0 {
		httputils.JSONResponse(w, http.StatusBadRequest, RegistrationResponse{
			Success: false,
			Message: fmt.Sprintf("Company name is not available"),
		})
		return
	}

	t, err := tx.Tenant.Create().
		SetInput(input.TenantInput).
		SetDefaultLanguage(lang).
		SetPlan(plan).
		Save(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	count, err = tx.User.Query().
		Where(user.EmailEqualFold(input.UserInput.Email)).
		Count(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	} else if count > 0 {
		httputils.JSONResponse(w, http.StatusBadRequest, RegistrationResponse{
			Success: false,
			Message: fmt.Sprintf("Email is not available"),
		})
		return
	}

	user, err := tx.User.Create().
		SetInput(input.UserInput).
		SetIsAccountOwner(true).
		SetTenant(t).
		Save(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, RegistrationResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, RegistrationResponse{
		Success:   true,
		UserPULID: user.ID,
	})
	return
}
