package restcustomer

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/utils/httputils"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

// ReturnReason Each reason is attached to a specific return portal
type ReturnReason struct {
	// Internal DELIVRIO ID
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Restockable bool   `json:"restockable"`
	//Refundable       bool   `json:"refundable"`
	ReturnPortalID   string `json:"return_portal_id"`
	ReturnPortalName string `json:"return_portal_name"`
}

type ReturnReasonsResponse struct {
	Reasons []ReturnReason `json:"reasons"`
}

// HandleReturnReasonsGet godoc
//
//	@Summary		Get return reasons
//	@Description	Responds with the return information for the OrderID provided.
//	@Tags			returns
//	@ID				get-return-reasons
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	ReturnReasonsResponse
//	@Failure		400		{object}	GeneralError
//	@Failure		404		{object}	GeneralError
//	@Failure		500		{object}	GeneralError
//	@Security		ApiKeyAuth
//	@Router			/return/reasons [get]
func HandleReturnReasonsGet(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "HandleReturnReasonsGet")
	defer span.End()

	cli := ent.FromContext(ctx)

	claims, err := cli.ReturnPortalClaim.Query().
		WithReturnPortal().
		All(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "query return claims failed")
		span.RecordError(err)
		httputils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	output := ReturnReasonsResponse{
		Reasons: make([]ReturnReason, 0),
	}

	for _, cl := range claims {

		output.Reasons = append(output.Reasons, ReturnReason{
			ID:          cl.ID.String(),
			Name:        cl.Name,
			Description: cl.Description,
			Restockable: cl.Restockable,
			//Refundable:       cl.,
			ReturnPortalID:   cl.Edges.ReturnPortal.ID.String(),
			ReturnPortalName: cl.Edges.ReturnPortal.Name,
		})
	}

	httputils.JSONResponse(w, http.StatusOK, output)
	return
}
