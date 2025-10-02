package restcustomer

import (
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/utils/httputils"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

type ReturnOrderLine struct {
	OrderLineExternalID string `json:"order_line_external_id"`
	ReturnReasonID      string `json:"return_reason_id"`
	ReturnReasonName    string `json:"return_reason_name"`
}

type ReturnColli struct {
	Status         string            `json:"status" validate:"oneof='Opened' 'Pending' 'Inbound' 'Received' 'Accepted' 'Declined' 'Deleted'"`
	CarrierLabelID string            `json:"carrier_label_id"`
	OrderLines     []ReturnOrderLine `json:"order_lines"`
}

type ReturnResponse struct {
	ReturnCollis []ReturnColli `json:"return_collis"`
}

// HandleReturnGet godoc
//
//	@Summary		Get return information
//	@Description	Responds with the return information for the OrderID provided.
//	@Tags			returns
//	@ID				get-return
//	@Accept			json
//	@Produce		json
//	@Param			id				query	string	false	"Order ID"	unique
//	@Param			external-id		query	string	false	"External ID"	unique
//	@Success		200				{object}	ReturnResponse
//	@Failure		400				{object}	GeneralError
//	@Failure		404				{object}	GeneralError
//	@Failure		500				{object}	GeneralError
//	@Security		ApiKeyAuth
//	@Router			/return [get]
func HandleReturnGet(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "HandleReturnGet")
	defer span.End()

	queryOrderID := r.URL.Query().Get(`id`)
	queryExternalOrderID := r.URL.Query().Get(`external-id`)

	span.SetAttributes(
		attribute.String("orderID", queryOrderID),
		attribute.String("externalOrderID", queryExternalOrderID),
	)

	orderID, errResponse := checkOrderIDs(ctx, queryOrderID, queryExternalOrderID)
	if errResponse != nil {
		span.SetStatus(codes.Error, "checkOrderIDs failed")
		span.RecordError(fmt.Errorf("%v", errResponse.Message))
		httputils.JSONResponse(w, errResponse.Status, errResponse)
		return
	}

	cli := ent.FromContext(ctx)

	output := ReturnResponse{
		ReturnCollis: make([]ReturnColli, 0),
	}

	returnCollis, err := cli.Order.Query().
		Where(order.ID(orderID)).
		QueryReturnColli().
		WithReturnOrderLine().
		Order(returncolli.ByCreatedAt()).
		All(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "query return collis failed")
		span.RecordError(err)
		httputils.JSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	for _, rc := range returnCollis {

		orderLines := make([]ReturnOrderLine, 0)
		for _, rol := range rc.Edges.ReturnOrderLine {
			ol, err := rol.QueryOrderLine().
				Only(ctx)
			if err != nil {
				span.RecordError(err)
				httputils.JSONResponse(w, http.StatusInternalServerError, err)
				return
			}

			orderLines = append(orderLines, ReturnOrderLine{
				OrderLineExternalID: ol.ExternalID,
				ReturnReasonID:      rol.Edges.ReturnPortalClaim.ID.String(),
				ReturnReasonName:    rol.Edges.ReturnPortalClaim.Name,
			})
		}

		output.ReturnCollis = append(output.ReturnCollis, ReturnColli{
			Status:         string(rc.Status),
			CarrierLabelID: "needs implementation",
			OrderLines:     orderLines,
		})
	}

	httputils.JSONResponse(w, http.StatusOK, output)
	return
}
