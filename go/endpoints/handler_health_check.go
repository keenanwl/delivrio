package endpoints

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/utils/httputils"
	printerUtils "delivrio.io/shared-utils/models/printer"
	"fmt"
	"log"
	"net/http"
)

func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {

	cli := ent.FromContext(r.Context())

	count, err := cli.Tenant.Query().
		Count(r.Context())
	if err != nil || count < 1 {
		log.Printf("health check: %v; count > 0: %v", err, count > 0)
		httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("health check: %v; count > 0: %v", err, count > 0),
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, printerUtils.PrintClientPingResponse{
		Success: true,
	})
	return

}
