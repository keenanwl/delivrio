package endpoints

import (
	"delivrio.io/go/utils/httputils"
	"net/http"
)

func GetLabels(w http.ResponseWriter, r *http.Request) {
	// Write the JSON response using the helper function
	response := map[string]bool{"ok": true}
	httputils.JSONResponse(w, http.StatusOK, response)
}
