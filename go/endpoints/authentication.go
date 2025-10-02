package endpoints

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/workstation"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/models/printer"
	"delivrio.io/shared-utils/pulid"
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func validWorkstationRequest(w http.ResponseWriter, r *http.Request) (*ent.Workstation, context.Context, error) {

	fullID := r.URL.Query().Get("id")
	computerID := r.URL.Query().Get("computer-id")
	deviceType := r.URL.Query().Get("device-type")
	if len(fullID) == 0 || len(computerID) == 0 || len(deviceType) == 0 {
		httputils.JSONResponse(w, http.StatusBadRequest, printer.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("missing authentication properties"),
		})
		return nil, nil, fmt.Errorf("missing authentication properties")
	}

	cli := ent.FromContext(r.Context())

	auth, err := b64.StdEncoding.DecodeString(fullID)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printer.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return nil, nil, err
	}
	log.Printf("received decoded string: %v", string(auth))
	tokenWorkstationTenant := strings.Split(string(auth), ":")

	if len(tokenWorkstationTenant) != 3 {
		httputils.JSONResponse(w, http.StatusBadRequest, printer.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("unexpected registration token format"),
		})
		return nil, nil, err
	}

	token := tokenWorkstationTenant[0]
	wsID := pulid.ID(tokenWorkstationTenant[1])
	tID := pulid.ID(tokenWorkstationTenant[2])
	ctxID := pulid.MustNew("CH")

	viewCtx := viewer.NewContext(
		r.Context(),
		viewer.UserViewer{
			Role:    viewer.Anonymous,
			Tenant:  tID,
			Context: ctxID,
		},
	)

	ws, err := cli.Workstation.Query().
		WithSelectedUser().
		Where(
			workstation.And(
				workstation.ID(wsID),
				workstation.TenantID(tID),
			),
		).Only(viewCtx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printer.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return nil, nil, err
	}

	if deviceType == "app" && ws.DeviceType != workstation.DeviceTypeApp {
		httputils.JSONResponse(w, http.StatusUnauthorized, printer.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("Expected device type 'app'"),
		})
		return nil, nil, err
	}

	if ws.Status == workstation.StatusDisabled || !ws.ArchivedAt.IsZero() {
		httputils.JSONResponse(w, http.StatusUnauthorized, printer.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("This workstation has been disabled by an administrator"),
		})
		return nil, nil, err
	}

	if !utils.CheckPasswordHash(token, ws.RegistrationCode) {
		httputils.JSONResponse(w, http.StatusUnauthorized, printer.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("your token was not recognized"),
		})
		return nil, nil, err
	}

	if len(ws.WorkstationID) != 0 && ws.WorkstationID != pulid.ID(computerID) {
		httputils.JSONResponse(w, http.StatusUnauthorized, printer.PrintClientPingResponse{
			Success: false,
			Message: fmt.Sprintf("Registration token may only be used once. Another computer has already utilized this registration token"),
		})
		return nil, nil, err
	}

	updateWS := ws.Update().
		SetWorkstationID(pulid.ID(computerID)).
		SetLastPing(time.Now())

	if ws.Status == workstation.StatusPending {
		updateWS.SetStatus(workstation.StatusActive)
	}

	err = updateWS.Exec(viewCtx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, printer.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return nil, nil, err
	}

	if ws.Edges.SelectedUser != nil {
		viewCtx = viewer.NewContext(
			viewCtx,
			viewer.UserViewer{
				Role:    viewer.Anonymous,
				Tenant:  tID,
				MyID:    ws.Edges.SelectedUser.ID,
				Context: ctxID,
			},
		)
	}

	// ViewCTX has tenant info
	return ws, viewCtx, nil
}
