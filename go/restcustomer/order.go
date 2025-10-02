package restcustomer

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

// swagger:model Order
type Order struct {
	// Internal DELIVRIO ID
	ID string `json:"id"`
	// Visible to end-customer. Must be unique.
	// example: SH123456789
	PublicID string `json:"public_id"`
	// ID from system external to DELIVRIO
	// example: 99999911111888888222222
	ExternalID *string `json:"external_id"`
	// Comment only visible within DELIVRIO
	CommentInternal *string `json:"comment_internal"`
	// Comment may be included in packaging or messing to end-customer
	CommentExternal *string `json:"comment_external"`
	// The name of the connection. Will fail if connection name duplicated in DELIVRIO.
	ConnectionName string `json:"connection_name"`
	// An order may contain >= 1 colli
	Colli []Colli `json:"colli"`
}

// swagger:model Colli
type Colli struct {
	// May contain duplicate product variants at different prices/currencies
	OrderLines      []OrderLine `json:"order_lines"`
	DeliveryAddress Address     `json:"delivery_address"`
	// When omitted, connection "sender" address will be used
	SenderAddress Address `json:"sender_address"`
	Status        string  `json:"status"`
}

// swagger:model OrderResponse
type OrderResponse struct {
	Order Order `json:"order"`
	// Existing shipments. A single order may be split into multiple shipments
	// so the same orderID may be found multiple times in the array
	Shipments []ShipmentResponse `json:"shipments"`
}

func checkOrderIDs(ctx context.Context, orderID string, externalOrderID string) (pulid.ID, *GeneralError) {

	if len(orderID) == 0 && len(externalOrderID) == 0 {
		return "", &GeneralError{
			Message: `either "id" or "external-id" are required query parameters`,
			Status:  http.StatusBadRequest,
		}
	}

	if len(orderID) > 0 && len(externalOrderID) > 0 {
		return "", &GeneralError{
			Message: `"id" and "external-id" are mutually exclusive query parameters`,
			Status:  http.StatusBadRequest,
		}
	}

	cli := ent.FromContext(ctx)

	orderQuery := cli.Order.Query()
	if len(orderID) > 0 {
		orderQuery = orderQuery.Where(order.ID(pulid.ID(orderID)))
	} else {
		orderQuery = orderQuery.Where(order.ExternalID(externalOrderID))
	}

	id, err := orderQuery.OnlyID(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", &GeneralError{
			Message: fmt.Sprintf("order: check ID: %v", err),
			Status:  http.StatusBadRequest,
		}
	} else if ent.IsNotFound(err) {
		return "", &GeneralError{
			Message: fmt.Sprintf("order: not found"),
			Status:  http.StatusNotFound,
		}
	}

	return id, nil
}

// HandleOrderGet godoc
//
//	@Summary		Returns the order information
//	@Description	Fetches order details by ID or external ID
//	@Tags			orders
//	@ID				get-order
//	@Accept			json
//	@Produce		json
//	@Param			id			query		string	false	"Order ID"		unique
//	@Param			external-id	query		string	false	"External ID"	unique
//	@Success		200			{object}	OrderResponse
//	@Failure		400			{object}	GeneralError
//	@Failure		404			{object}	GeneralError
//	@Failure		500			{object}	GeneralError
//	@Security		ApiKeyAuth
//	@Router			/order [get]
func HandleOrderGet(w http.ResponseWriter, r *http.Request) {

	ctx, span := tracer.Start(r.Context(), "HandleOrderGet")
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
		httputils.JSONResponse(w, http.StatusInternalServerError, errResponse)
		return
	}

	cli := ent.FromContext(ctx)

	ord, err := cli.Order.Query().
		WithConnection().
		WithColli().
		Where(order.ID(orderID)).
		Only(r.Context())
	if err != nil {
		span.SetStatus(codes.Error, "query order failed")
		span.RecordError(err)
		httputils.JSONResponse(w, http.StatusBadRequest, GeneralError{
			Message: fmt.Errorf("rest: get order: %w", err).Error(),
			Status:  http.StatusNotFound,
		})
		return
	}

	// Omitting TX isolation for now...
	outputColli := make([]Colli, 0)
	for _, col := range ord.Edges.Colli {
		sender, err := col.QuerySender().
			WithCountry().
			Only(r.Context())
		if err != nil {
			span.SetStatus(codes.Error, "query package sender failed")
			span.RecordError(err)
			httputils.JSONResponse(w, http.StatusBadRequest, GeneralError{
				Message: fmt.Errorf("rest: get order: %w", err).Error(),
				Status:  http.StatusNotFound,
			})
			return
		}

		recipient, err := col.QueryRecipient().
			WithCountry().
			Only(r.Context())
		if err != nil {
			span.SetStatus(codes.Error, "query recipient failed")
			span.RecordError(err)
			httputils.JSONResponse(w, http.StatusBadRequest, GeneralError{
				Message: fmt.Errorf("rest: get order: %w", err).Error(),
				Status:  http.StatusNotFound,
			})
			return
		}

		ol, err := col.QueryOrderLines().
			WithProductVariant().
			WithCurrency().
			All(r.Context())
		if err != nil {
			span.SetStatus(codes.Error, "query order lines failed")
			span.RecordError(err)
			httputils.JSONResponse(w, http.StatusBadRequest, GeneralError{
				Message: fmt.Errorf("rest: get order: %w", err).Error(),
				Status:  http.StatusNotFound,
			})
			return
		}

		outputOrderLines := make([]OrderLine, 0)
		for _, o := range ol {
			pvID := o.Edges.ProductVariant.ID.String()
			outputOrderLines = append(outputOrderLines, OrderLine{
				ExternalProductVariantID: &o.Edges.ProductVariant.ExternalID,
				ProductVariantID:         &pvID,
				Units:                    o.Units,
				Price:                    o.UnitPrice,
				Currency:                 o.Edges.Currency.Display,
			})
		}

		outputColli = append(outputColli, Colli{
			OrderLines: outputOrderLines,
			DeliveryAddress: Address{
				FirstName:     recipient.FirstName,
				LastName:      recipient.LastName,
				Company:       recipient.Company,
				VATNumber:     recipient.VatNumber,
				StreetOne:     recipient.AddressOne,
				StreetTwo:     recipient.AddressTwo,
				PostalCode:    recipient.Zip,
				City:          recipient.City,
				CountryAlpha2: recipient.Edges.Country.Alpha2,
				State:         recipient.State,
				Email:         recipient.Email,
				PhoneNumber:   recipient.PhoneNumber,
			},
			SenderAddress: Address{
				FirstName:     sender.FirstName,
				LastName:      sender.LastName,
				Company:       sender.Company,
				VATNumber:     sender.VatNumber,
				StreetOne:     sender.AddressOne,
				StreetTwo:     sender.AddressTwo,
				PostalCode:    sender.Zip,
				City:          sender.City,
				CountryAlpha2: sender.Edges.Country.Alpha2,
				State:         sender.State,
				Email:         sender.Email,
				PhoneNumber:   sender.PhoneNumber,
			},
		})
	}

	httputils.JSONResponse(w, http.StatusOK, OrderResponse{
		Order: Order{
			PublicID:        ord.OrderPublicID,
			ExternalID:      &ord.ExternalID,
			CommentInternal: &ord.CommentInternal,
			CommentExternal: &ord.CommentExternal,
			ConnectionName:  ord.Edges.Connection.Name,
			Colli:           outputColli,
		},
		Shipments: nil,
	})
	return

}
