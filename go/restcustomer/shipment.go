package restcustomer

import (
	"context"
	"delivrio.io/go/carrierapis/labels"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/orderline"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"net/http"
)

type ShipmentsRequest struct {
	ShipmentRequests []ShipmentCreate `json:"shipment_requests"`
}

type ColliRequest struct {
	// List of external order lines IDs to be included in this parcel.
	// For instance, allows Shopify order line IDs lookup.
	OrderLineExternalIDs []string `json:"order_line_external_ids"`
}

// ShipmentCreate splits orders into requested collis. Contacts carrier for labels and returns labels.
type ShipmentCreate struct {
	// A list of individual parcels for which a label needs to be generated.
	// The labels are grouped when the carrier supports this feature.
	Collis []ColliRequest `json:"collis"`
	// Order status
	// Settings order_status to Dispatched will split any remaining order lines into their own parcel, and then cancel that parcel.
	OrderStatus string `json:"order_status" validate:"oneof='Pending' 'Partially_dispatched' 'Dispatched'"`
}

type ShipmentsResponse struct {
	Shipments []ShipmentResponse `json:"shipments"`
}

type ShipmentErrorType string

const (
	NotFound          ShipmentErrorType = "NOT_FOUND"
	CarrierValidation ShipmentErrorType = "CARRIER_VALIDATION"
)

type ShipmentError struct {
	Message string            `json:"message"`
	Type    ShipmentErrorType `json:"type"`
	// HTTP Status of the error
	Status int `json:"status"`
}

type ShipmentResponse struct {
	OrderID     string        `json:"order_id"`
	OrderStatus string        `json:"order_status"`
	Parcels     []ParcelLabel `json:"parcels"`
	Error       ShipmentError `json:"error,omitempty"`
}

// HandleShipmentsCreate godoc
//
//	@Summary		Create or update shipments
//	@Description	Returns the shipment information based on the provided request. Be aware partial success may return status 207.
//	@Tags			shipments
//	@ID				post-shipments
//	@Accept			json
//	@Produce		json
//	@Param			ShipmentsRequest	body	ShipmentsRequest	true	"Shipments request body"
//	@Success		200		{object}	OrderResponse
//	@Success		207		{object}	OrderResponse
//	@Failure		400		{object}	GeneralError
//	@Failure		404		{object}	GeneralError
//	@Failure		500		{object}	GeneralError
//	@Security		ApiKeyAuth
//	@Router			/shipments [post]
func HandleShipmentsCreate(w http.ResponseWriter, r *http.Request) {

	ctx, span := tracer.Start(r.Context(), "HandleShipmentsCreate")
	defer span.End()

	var input ShipmentsRequest
	err := httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, GeneralError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	output := ShipmentsResponse{
		Shipments: make([]ShipmentResponse, 0),
	}

	multiStatus := false

	for _, ship := range input.ShipmentRequests {

		preparedColliIDs, _, errResponse := prepareCollis(ctx, ship.Collis)
		if errResponse != nil {
			multiStatus = true
			output.Shipments = append(output.Shipments, ShipmentResponse{
				Error: ShipmentError{
					Message: errResponse.Message,
					Status:  errResponse.Status,
				},
			})
			continue
		}

		sortedCollis, err := labels.SortCollis(ctx, preparedColliIDs)
		if err != nil {
			multiStatus = true
			output.Shipments = append(output.Shipments, ShipmentResponse{
				Error: ShipmentError{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				},
			})
			continue
		}

		createdShipments, err := labels.RequestAndSave(ctx, sortedCollis)
		if err != nil {
			multiStatus = true
			output.Shipments = append(output.Shipments, ShipmentResponse{
				Error: ShipmentError{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				},
			})
			continue
		}

		outputParcels := make([]ParcelLabel, 0)
		for _, s := range createdShipments {
			outputParcels = append(outputParcels, ParcelLabel{
				Link:     "1234",
				LabelPDF: s.AllLabels,
			})
		}

		output.Shipments = append(output.Shipments, ShipmentResponse{
			OrderID:     "",
			OrderStatus: "",
			Parcels:     outputParcels,
		})

	}

	if multiStatus {
		httputils.JSONResponse(w, http.StatusMultiStatus, output)
		return
	}

	httputils.JSONResponse(w, http.StatusOK, output)
	return

}

func rollbackPrepareColli(tx *ent.Tx, code int, err error) ([]pulid.ID, []string, *GeneralError) {
	tx.Rollback()
	return nil, nil, &GeneralError{
		Message: fmt.Errorf("preparecolli: %w", err).Error(),
		Status:  code,
	}
}

func prepareCollis(ctx context.Context, colliRequests []ColliRequest) ([]pulid.ID, []string, *GeneralError) {
	cli := ent.FromContext(ctx)

	tx, err := cli.Tx(ctx)
	if err != nil {
		return nil, nil, &GeneralError{
			Message: fmt.Errorf("preparecolli: %w", err).Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	defer tx.Rollback()

	ctx = ent.NewTxContext(ctx, tx)

	colliToShip := make([]pulid.ID, 0)
	allOrderLineExternalIDs := make([]string, 0)
	for _, orderLineGroupToShip := range colliRequests {

		if len(orderLineGroupToShip.OrderLineExternalIDs) == 0 {
			return rollbackPrepareColli(tx, http.StatusBadRequest, fmt.Errorf("expected external order IDs"))
		}

		allOrderLineExternalIDs = append(allOrderLineExternalIDs, orderLineGroupToShip.OrderLineExternalIDs...)

		ol, err := tx.OrderLine.Query().
			WithColli().
			Where(orderline.ExternalIDIn(orderLineGroupToShip.OrderLineExternalIDs...)).
			All(ctx)
		if err != nil {
			return rollbackPrepareColli(tx, http.StatusInternalServerError, err)
		}

		if len(ol) != len(orderLineGroupToShip.OrderLineExternalIDs) {
			return rollbackPrepareColli(tx, http.StatusBadRequest, fmt.Errorf(
				"expected found order lines count (%v) to match requested count (%v)",
				len(ol),
				len(orderLineGroupToShip.OrderLineExternalIDs),
			))
		}

		firstColli := ol[0].Edges.Colli
		colliToShip = append(colliToShip, firstColli.ID)
		for _, o := range ol {
			// Ensure all subsequent order lines are in the requested colli
			// The first available colli is just an arbitrary grouping
			// Suffers from, potentially not matching addresses on existing collis
			if o.Edges.Colli.ID != firstColli.ID {
				err := o.Update().
					SetColli(firstColli).
					Exec(ctx)
				if err != nil {
					return rollbackPrepareColli(tx, http.StatusBadRequest, err)
				}
			}
		}

		orderLinesToMove, err := tx.OrderLine.Query().Where(
			orderline.And(
				orderline.Not(orderline.ExternalIDIn(orderLineGroupToShip.OrderLineExternalIDs...)),
				orderline.HasColliWith(colli.ID(firstColli.ID)),
			),
		).All(ctx)
		if err != nil {
			return rollbackPrepareColli(tx, http.StatusBadRequest, err)
		}

		if len(orderLinesToMove) > 0 {
			newColli, err := utils.DuplicateColli(ctx, firstColli.ID)
			if err != nil {
				return rollbackPrepareColli(tx, http.StatusBadRequest, err)
			}

			for _, ol := range orderLinesToMove {
				err = ol.Update().
					SetColli(newColli).
					Exec(ctx)
				if err != nil {
					return rollbackPrepareColli(tx, http.StatusBadRequest, err)
				}
			}
		}

	}

	err = tx.Commit()
	if err != nil {
		return rollbackPrepareColli(tx, http.StatusBadRequest, err)
	}
	return colliToShip, allOrderLineExternalIDs, nil
}
