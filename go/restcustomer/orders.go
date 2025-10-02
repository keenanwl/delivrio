package restcustomer

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/productvariant"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type OrdersCreateRequest struct {
	Orders []OrderCreate `json:"orders"`
}

type OrdersCreateResponse struct {
	// 1:1 output corresponding to posted Orders
	Orders []OrdersSaved `json:"orders"`
}

type OrdersSaved struct {
	// Whether the order was saved.
	Success bool `json:"success"`
	// Save error. Present when Success is false.
	Error string `json:"error"`
	// Internal DELIVRIO ID. Present when save is successful.
	OrderID string `json:"order_id"`
}

// OrderCreate represents the payload to create an order.
type OrderCreate struct {
	// PublicID is visible to the end-customer and must be unique.
	// Example: SH123456789
	PublicID string `json:"public_id" validate:"required"`

	// ExternalID is the ID from a system external to DELIVRIO.
	// Example: 99999911111888888222222
	ExternalID *string `json:"external_id"`

	// CommentInternal is a comment visible only within DELIVRIO.
	// Example: Internal note about the order.
	CommentInternal *string `json:"comment_internal"`

	// CommentExternal is a comment that may be included in packaging or messaging to the end-customer.
	// Example: Thank you for your order!
	CommentExternal *string `json:"comment_external"`

	// ConnectionName is the name of the connection.
	// Example: MainWarehouseConnection
	// Note: This will fail if the connection name is duplicated in DELIVRIO.
	ConnectionName string `json:"connection_name" validate:"required"`

	// DeliveryOptionID is the ID of the delivery option.
	// Example: DEL12345
	// Note: Defaults to the optional connection default if not provided.
	DeliveryOptionID *string `json:"delivery_option_id" validate:"required"`

	// OrderLines contains the items in the order. It may contain duplicate product variants at different prices or in different currencies.
	OrderLines []OrderLine `json:"order_lines" validate:"required"`

	// DeliveryAddress is the address where the order should be delivered.
	DeliveryAddress Address `json:"delivery_address" validate:"required"`

	// SenderAddress is the address of the sender. If omitted, the connection's "sender" address will be used.
	SenderAddress *Address `json:"sender_address" validate:"required"`
}

type Address struct {
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Company    string `json:"company"`
	VATNumber  string `json:"vat_number"`
	StreetOne  string `json:"street_one" validate:"required"`
	StreetTwo  string `json:"street_two" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	City       string `json:"city" validate:"required"`
	// Alpha2
	// example: "DK" | "DE" | "FR"
	CountryAlpha2 string `json:"country" validate:"required"`
	State         string `json:"state"`
	// Some carriers require depending on service.
	Email string `json:"email" validate:"required"`
	// Include country code.
	// Some carriers require depending on service.
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type OrderLine struct {
	// Mutually exclusive with product_variant_id
	ExternalProductVariantID *string `json:"external_product_id"`
	// Mutually exclusive with external_product_id
	ProductVariantID *string `json:"product_variant_id"`

	Units         int     `json:"units"`
	Price         float64 `json:"price"`
	TotalDiscount float64 `json:"total_discount"`
	Currency      string  `json:"currency" validate:"oneof='DKK' 'EUR' 'USD'"`
}

// HandleOrderCreate godoc
//
//	@Summary		Create orders
//	@Description	Responds with the save info for each POSTed order. Be aware partial success will return status 207. Max orders per request: 5
//	@Tags			orders
//	@ID				create-orders
//	@Accept			json
//	@Produce		json
//	@Param			body	body	OrdersCreateRequest	true	"Order creation request"
//	@Success		200		{object}	OrdersCreateResponse
//	@Success		207		{object}	OrdersCreateResponse
//	@Failure		400		{object}	GeneralError
//	@Failure		404		{object}	GeneralError
//	@Failure		500		{object}	GeneralError
//	@Security		ApiKeyAuth
//	@Router			/orders [post]
func HandleOrderCreate(w http.ResponseWriter, r *http.Request) {

	ctx, span := tracer.Start(r.Context(), "HandleOrderCreate")
	defer span.End()

	var input OrdersCreateRequest
	err := httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest,
			GeneralError{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		return
	}

	span.SetAttributes(
		attribute.Int("orderCount", len(input.Orders)),
	)

	if len(input.Orders) > 5 {
		errMaxOrders := fmt.Errorf("max create orders exceed; got %v", len(input.Orders))
		span.SetStatus(codes.Error, "max orders hit")
		span.RecordError(errMaxOrders)
		httputils.JSONResponse(w, http.StatusBadRequest, GeneralError{
			Message: errMaxOrders.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	output := OrdersCreateResponse{
		Orders: make([]OrdersSaved, 0),
	}

	cli := ent.FromContext(ctx)
	multiStatus := 0

	for _, o := range input.Orders {

		ctx, tx, err := cli.OpenTx(ctx)
		if err != nil {
			span.SetStatus(codes.Error, "tx failed")
			span.RecordError(err)
			httputils.JSONResponse(w, http.StatusInternalServerError, GeneralError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		defer tx.Rollback()

		orderID, err := createOrder(ctx, o)
		if err != nil {
			span.SetStatus(codes.Error, "createOrder failed")
			span.RecordError(err)
			tx.Rollback()
			multiStatus++
			output.Orders = append(output.Orders, OrdersSaved{
				Success: false,
				Error:   err.Error(),
			})
		} else {
			err := tx.Commit()
			if err != nil {
				span.SetStatus(codes.Error, "commit failed")
				span.RecordError(err)
				output.Orders = append(output.Orders, OrdersSaved{
					Success: false,
					Error:   err.Error(),
				})
			} else {
				output.Orders = append(output.Orders, OrdersSaved{
					Success: true,
					OrderID: orderID.String(),
				})
			}
		}
	}

	if multiStatus > 0 {
		span.AddEvent("multistatus", trace.WithAttributes(attribute.Int("failCreate", multiStatus)))
		httputils.JSONResponse(w, http.StatusMultiStatus, output)
		return
	}

	span.AddEvent("createdAllOrders")
	httputils.JSONResponse(w, http.StatusOK, output)
	return

}

func createOrder(ctx context.Context, o OrderCreate) (pulid.ID, error) {

	view := viewer.FromContext(ctx)
	tx := ent.TxFromContext(ctx)

	conn, err := tx.Connection.Query().
		WithDefaultDeliveryOption().
		Where(connection.NameEqualFold(o.ConnectionName)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) && !ent.IsNotSingular(err) {
		return "", err
	} else if ent.IsNotSingular(err) {
		return "", fmt.Errorf("rest: found multiple connections with name: %v", o.ConnectionName)
	} else if ent.IsNotFound(err) {
		return "", fmt.Errorf("rest: connection with name not found: %v", o.ConnectionName)
	}

	ord, err := tx.Order.Create().
		SetTenantID(view.TenantID()).
		SetConnection(conn).
		SetNillableExternalID(o.ExternalID).
		SetOrderPublicID(o.PublicID).
		SetNillableCommentExternal(o.CommentExternal).
		SetNillableCommentInternal(o.CommentInternal).
		SetStatus(order.StatusPending).
		Save(history.NewConfig(ctx).
			SetDescription("Add order").
			SetOrigin(changehistory.OriginRestAPI).
			Ctx())
	if err != nil {
		return "", err
	}

	sender, err := getSender(ctx, o.SenderAddress, conn)
	if err != nil {
		return "", err
	}

	recipient, err := getRecipient(ctx, o.DeliveryAddress)
	if err != nil {
		return "", err
	}

	historyDescription := "Add colli"

	var deliveryOptionID *pulid.ID
	if deliveryOptionID == nil && conn.Edges.DefaultDeliveryOption != nil {
		deliveryOptionID = &conn.Edges.DefaultDeliveryOption.ID
		historyDescription = "Add colli with connection default delivery option"
	} else if o.DeliveryOptionID != nil {
		pID := pulid.ID(*o.DeliveryOptionID)
		deliveryOptionID = &pID
	}

	col, err := tx.Colli.Create().
		SetRecipient(recipient).
		SetOrder(ord).
		SetSender(sender).
		SetTenantID(view.TenantID()).
		SetStatus(colli.StatusPending).
		SetNillableDeliveryOptionID(deliveryOptionID).
		Save(history.NewConfig(ctx).
			SetDescription(historyDescription).
			SetOrigin(changehistory.OriginRestAPI).
			Ctx())
	if err != nil {
		return "", err
	}

	err = getOrderLines(ctx, o.OrderLines, col)
	if err != nil {
		return "", err
	}

	return ord.ID, nil

}

func getOrderLines(ctx context.Context, o []OrderLine, colli *ent.Colli) error {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	for i, ol := range o {

		if ol.ExternalProductVariantID != nil && ol.ProductVariantID != nil {
			return fmt.Errorf("both external_product_variant_id and product_variant_id set (offset %v)", i)
		}

		cur, err := utils.ToCurrency(ctx, tx.Client(), ol.Currency)
		if err != nil {
			return fmt.Errorf("currency lookup: %w", err)
		}

		create := tx.OrderLine.Create().
			SetTenantID(view.TenantID()).
			SetColli(colli).
			SetUnits(ol.Units).
			SetUnitPrice(ol.Price).
			SetDiscountAllocationAmount(ol.TotalDiscount).
			SetCurrency(cur)
		if ol.ExternalProductVariantID != nil {
			pv, err := tx.ProductVariant.Query().
				Where(productvariant.ExternalIDEQ(*ol.ExternalProductVariantID)).
				Only(ctx)
			if err != nil {
				return fmt.Errorf("save order line (offset %v): %w", i, err)
			}
			create = create.SetProductVariant(pv)
		} else if ol.ProductVariantID != nil {
			create = create.SetProductVariantID(pulid.ID(*ol.ProductVariantID))
		} else {
			return fmt.Errorf("either external_product_variant_id or product_variant_id is required (offset %v)", i)
		}

		err = create.Exec(ctx)
		if err != nil {
			return fmt.Errorf("save order line (offset %v): %w", i, err)
		}

	}

	return nil

}

func getRecipient(ctx context.Context, o Address) (*ent.Address, error) {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	deliveryCountry, err := tx.Country.Query().
		Where(country.Alpha2EqualFold(o.CountryAlpha2)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return tx.Address.Create().
		SetTenantID(view.TenantID()).
		SetCompany(o.Company).
		SetAddressOne(o.StreetOne).
		SetAddressTwo(o.StreetTwo).
		SetCity(o.City).
		SetZip(o.PostalCode).
		SetState(o.State).
		SetEmail(o.Email).
		SetFirstName(o.FirstName).
		SetLastName(o.LastName).
		SetPhoneNumber(o.PhoneNumber).
		SetCountry(deliveryCountry).
		SetVatNumber(o.VATNumber).
		Save(ctx)
}

func getSender(ctx context.Context, o *Address, conn *ent.Connection) (*ent.Address, error) {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	if o != nil {
		senderCountry, err := tx.Country.Query().
			Where(country.Alpha2EqualFold(o.CountryAlpha2)).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		return tx.Address.Create().
			SetTenantID(view.TenantID()).
			SetCompany(o.Company).
			SetAddressOne(o.StreetOne).
			SetAddressTwo(o.StreetTwo).
			SetCity(o.City).
			SetZip(o.PostalCode).
			SetState(o.State).
			SetCountry(senderCountry).
			SetVatNumber(o.VATNumber).
			SetEmail(o.Email).
			SetFirstName(o.FirstName).
			SetLastName(o.LastName).
			SetPhoneNumber(o.PhoneNumber).
			Save(ctx)
	}

	connSenderLocation, err := conn.QuerySenderLocation().
		WithAddress().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	senderCountry, err := connSenderLocation.Edges.Address.Country(ctx)
	if err != nil {
		return nil, err
	}

	return connSenderLocation.Edges.Address.CloneEntity(tx).
		SetTenantID(view.TenantID()).
		SetCountry(senderCountry).
		Save(ctx)
}
