package endpoints

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/workspacerecentscan"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"net/http"
)

type StatusUpdateInput struct {
	ScanID         pulid.ID `json:"scan_id"`
	IsNotify       bool     `json:"is_notify"`
	IsMarkPickedUp bool     `json:"is_mark_picked_up"`
}

type StatusUpdateResponse struct {
	Success       bool                  `json:"success"`
	Message       string                `json:"message"`
	CurrentStatus shipmentparcel.Status `json:"current_status"`
}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {
	ws, _, err := validWorkstationRequest(w, r)
	if err != nil {
		// Writing already happens
		return
	}

	var input StatusUpdateInput
	err = httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, StatusUpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx := viewer.NewContext(r.Context(), viewer.UserViewer{
		Role:   viewer.Anonymous,
		Tenant: ws.TenantID,
	})

	cli := ent.FromContext(r.Context())

	record, err := cli.WorkspaceRecentScan.Query().
		WithShipmentParcel().
		Where(workspacerecentscan.ID(input.ScanID)).
		Only(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, StatusUpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	status := shipmentparcel.StatusPickedUp
	if input.IsNotify {
		status = shipmentparcel.StatusAwaitingCcPickup
		err = mergeutils.SendCCNotifyEmail(ctx, record.Edges.ShipmentParcel)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, StatusUpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	err = record.Edges.ShipmentParcel.Update().
		SetStatus(status).
		Exec(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, StatusUpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, StatusUpdateResponse{
		Success:       true,
		Message:       "",
		CurrentStatus: status,
	})
	return

}
