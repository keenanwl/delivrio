package endpoints

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/workspacerecentscan"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"net/http"
)

type SignatureUploadInput struct {
	Base64PNG string   `json:"base_64_png"`
	ScanID    pulid.ID `json:"scan_id"`
}

type SignatureUploadResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SignatureUpload(w http.ResponseWriter, r *http.Request) {
	ws, _, err := validWorkstationRequest(w, r)
	if err != nil {
		// Writing already happens
		return
	}

	var input SignatureUploadInput
	err = httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, SignatureUploadResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx := viewer.NewContext(r.Context(), viewer.UserViewer{
		Role:   viewer.Anonymous,
		Tenant: ws.TenantID,
	})

	cli := ent.FromContext(ctx)

	scan, err := cli.WorkspaceRecentScan.Query().
		WithShipmentParcel().
		WithUser().
		Where(workspacerecentscan.ID(input.ScanID)).
		Only(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, SignatureUploadResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	imgPath, err := utils.SaveImage(fmt.Sprintf("data:image/png;base64,%s", input.Base64PNG))
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, SignatureUploadResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = scan.Edges.ShipmentParcel.Update().
		AppendCcPickupSignatureUrls([]string{imgPath}).
		Exec(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, SignatureUploadResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, SignatureUploadResponse{
		Success: true,
		Message: "",
	})
	return

}
