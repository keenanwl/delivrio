package endpoints

import (
	"delivrio.io/go/utils/httputils"
	"net/http"

	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
)

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password"`
	Otk         string `json:"otk"`
}

type ResetPasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	var request ResetPasswordRequest

	err := httputils.UnmarshalRequestBody(r, &request)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, ResetPasswordResponse{
			Success: false,
			Message: "An error occurred",
		})
		return
	}

	validOtk, user := ValidOtk(r.Context(), request.Otk)

	if !validOtk {
		httputils.JSONResponse(w, http.StatusOK, ResetPasswordResponse{
			Success: false,
			Message: "Your reset password session has expired",
		})
		return
	}

	// Check password valid first
	_, err = user.Update().
		SetHash(utils.HashPasswordX(request.NewPassword)).
		Save(viewer.NewBackgroundContext(r.Context()))

	httputils.JSONResponse(w, http.StatusOK, ResetPasswordResponse{
		Success: err == nil,
		Message: "",
	})
	return

}
