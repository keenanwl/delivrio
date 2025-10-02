package utils

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/currency"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"log"
	"strings"
)

type ProductVariantQuantity struct {
	OrderLineID pulid.ID `json:"orderLineID"`
	VariantID   pulid.ID `json:"variantID"`
	Units       int      `json:"units"`
	Price       float64  `json:"price"`
	Discount    float64  `json:"discount"`
	Currency    string   `json:"currency"`
}

func ToCurrency(ctx context.Context, cli *ent.Client, input string) (*ent.Currency, error) {

	cc := currency.DefaultCurrencyCode
	switch strings.ToLower(input) {
	case strings.ToLower(currency.CurrencyCodeDKK.String()):
		cc = currency.CurrencyCodeDKK
	case strings.ToLower(currency.CurrencyCodeUSD.String()):
		cc = currency.CurrencyCodeUSD
	case strings.ToLower(currency.CurrencyCodeEUR.String()):
		cc = currency.CurrencyCodeEUR
	default:
		log.Printf("unsupported currency code: %v", input)
	}

	return cli.Currency.Query().
		Where(currency.CurrencyCodeEQ(cc)).
		Only(ctx)
}

func CreateColli(
	ctx context.Context,
	orderID pulid.ID,
	input ent.CreateColliInput,
	deliveryOptionID *pulid.ID,
	deliveryPointID *pulid.ID,
	ccLocationID *pulid.ID,
	packagingID *pulid.ID,
	recipientAddress *ent.AddressCreate,
	senderAddress *ent.AddressCreate,
	products []*ProductVariantQuantity,
) (*ent.Colli, error) {
	v := viewer.FromContext(ctx)
	tx := ent.TxFromContext(ctx)

	newRecipientAddress, err := recipientAddress.
		SetTenantID(v.TenantID()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	newSenderAddress, err := senderAddress.
		SetTenantID(v.TenantID()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	create := tx.Colli.Create().
		SetInput(input).
		SetSender(newSenderAddress).
		SetRecipient(newRecipientAddress).
		SetOrderID(orderID).
		SetNillableParcelShopID(deliveryPointID).
		SetNillableDeliveryOptionID(deliveryOptionID).
		SetNillableClickCollectLocationID(ccLocationID).
		SetNillablePackagingID(packagingID).
		SetStatus(colli.StatusPending).
		SetTenantID(v.TenantID())

	newColli, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}

	for _, line := range products {

		cur, err := ToCurrency(ctx, tx.Client(), line.Currency)
		if err != nil {
			return nil, err
		}

		err = tx.OrderLine.Create().
			SetTenantID(v.TenantID()).
			SetColli(newColli).
			SetProductVariantID(line.VariantID).
			SetUnits(line.Units).
			SetUnitPrice(line.Price).
			SetCurrency(cur).
			Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return newColli, err
}

func DuplicateColli(ctx context.Context, fromColliID pulid.ID) (*ent.Colli, error) {
	tx := ent.TxFromContext(ctx)

	copyFrom, err := tx.Colli.Query().
		WithOrder().
		WithDeliveryOption().
		WithRecipient().
		WithSender().
		WithPackaging().
		Where(colli.ID(fromColliID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	recipientAddress, err := copyFrom.Edges.RecipientOrErr()
	if err != nil {
		return nil, fmt.Errorf("duplicate: %w", err)
	}
	senderAddress, err := copyFrom.Edges.SenderOrErr()
	if err != nil {
		return nil, fmt.Errorf("duplicate: %w", err)
	}

	var deliveryOptionID *pulid.ID
	if copyFrom.Edges.DeliveryOption != nil {
		deliveryOptionID = &copyFrom.Edges.DeliveryOption.ID
	}

	var deliveryPointID *pulid.ID
	if copyFrom.Edges.ParcelShop != nil {
		deliveryPointID = &copyFrom.Edges.ParcelShop.ID
	}

	var ccLocationID *pulid.ID
	if copyFrom.Edges.ClickCollectLocation != nil {
		ccLocationID = &copyFrom.Edges.ClickCollectLocation.ID
	}

	var packagingID *pulid.ID
	if copyFrom.Edges.Packaging != nil {
		packagingID = &copyFrom.Edges.Packaging.ID
	}

	recipientAddressCountry, err := recipientAddress.Country(ctx)
	if err != nil {
		return nil, err
	}

	senderAddressCountry, err := senderAddress.Country(ctx)
	if err != nil {
		return nil, err
	}

	return CreateColli(
		ctx,
		copyFrom.Edges.Order.ID,
		ent.CreateColliInput{
			DeliveryOptionID: deliveryOptionID,
		},
		deliveryOptionID,
		deliveryPointID,
		ccLocationID,
		packagingID,
		recipientAddress.CloneEntity(tx).
			SetCountry(recipientAddressCountry),
		senderAddress.CloneEntity(tx).
			SetCountry(senderAddressCountry),
		[]*ProductVariantQuantity{},
	)
}
