package endpoints

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/user"
	"delivrio.io/go/ent/workspacerecentscan"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/models/printer"
	"delivrio.io/shared-utils/pulid"
	"entgo.io/ent/dialect/sql"
	"net/http"
	"time"
)

type LabelListScan struct {
	ScanID        pulid.ID              `json:"scan_id"`
	Barcode       string                `json:"barcode"`
	CurrentStatus shipmentparcel.Status `json:"current_status"`
	CreatedAt     time.Time             `json:"created_at"`
}

type LabelListResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Scans   []LabelListScan `json:"scans"`
}

type LabelRegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Found   bool   `json:"found"`
}

func ScannedLabelList(w http.ResponseWriter, r *http.Request) {

	output := make([]LabelListScan, 0)

	ws, _, err := validWorkstationRequest(w, r)
	if err != nil {
		// Writing already happens
		return
	}

	cli := ent.FromContext(r.Context())
	viewCtx := viewer.NewContext(r.Context(), viewer.UserViewer{
		Role:   viewer.Anonymous,
		Tenant: ws.TenantID,
	})

	list, err := cli.WorkspaceRecentScan.Query().
		WithShipmentParcel().
		Where(workspacerecentscan.HasUserWith(user.ID(ws.Edges.User.ID))).
		Order(workspacerecentscan.ByCreatedAt(sql.OrderDesc())).
		Limit(50).
		All(viewCtx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, LabelListResponse{
			Success: false,
			Message: err.Error(),
			Scans:   output,
		})
		return
	}

	for _, i := range list {
		output = append(
			output,
			LabelListScan{
				ScanID:        i.ID,
				CurrentStatus: i.Edges.ShipmentParcel.Status,
				Barcode:       i.Edges.ShipmentParcel.ItemID,
				CreatedAt:     i.CreatedAt,
			},
		)
	}

	httputils.JSONResponse(w, http.StatusOK, LabelListResponse{
		Success: true,
		Message: "",
		Scans:   output,
	})
	return

}

func ScannedLabelRegister(w http.ResponseWriter, r *http.Request) {

	ws, _, err := validWorkstationRequest(w, r)
	if err != nil {
		// Writing already happens
		return
	}

	var input printer.PrintClientPing
	err = httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printer.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	cli := ent.FromContext(r.Context())
	viewCtx := viewer.NewContext(r.Context(), viewer.UserViewer{
		Role:   viewer.Anonymous,
		Tenant: ws.TenantID,
	})

	ship, err := cli.ShipmentParcel.Query().
		Where(shipmentparcel.ItemIDContainsFold(input.LabelID)).
		// Should be Only in the future?
		First(viewCtx)
	if err != nil && !ent.IsNotFound(err) {
		httputils.JSONResponse(w, http.StatusBadRequest, LabelRegisterResponse{
			Success: false,
			Message: err.Error(),
			Found:   false,
		})
		return
	} else if ent.IsNotFound(err) {
		httputils.JSONResponse(w, http.StatusNotFound, LabelRegisterResponse{
			Success: true,
			Message: err.Error(),
			Found:   false,
		})
		return
	}

	lastScan, err := cli.WorkspaceRecentScan.Query().
		WithShipmentParcel().
		Where(workspacerecentscan.HasUserWith(user.ID(ws.Edges.User.ID))).
		Order(workspacerecentscan.ByCreatedAt(sql.OrderDesc())).
		First(viewCtx)
	if err != nil && !ent.IsNotFound(err) {
		httputils.JSONResponse(w, http.StatusBadRequest, LabelRegisterResponse{
			Success: false,
			Message: err.Error(),
			Found:   false,
		})
		return
	} else if err == nil && lastScan.Edges.ShipmentParcel.ID == ship.ID {
		// Skip creation if just scanned
		// maybe refactor to upsert..
		httputils.JSONResponse(w, http.StatusOK, LabelRegisterResponse{
			Success: true,
			Message: "",
			Found:   true,
		})
		return
	}

	err = cli.WorkspaceRecentScan.Create().
		SetTenantID(ws.TenantID).
		SetUserID(ws.Edges.User.ID).
		SetShipmentParcel(ship).
		Exec(viewCtx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, LabelRegisterResponse{
			Success: false,
			Message: err.Error(),
			Found:   false,
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, LabelRegisterResponse{
		Success: true,
		Message: "",
		Found:   true,
	})
	return

}
