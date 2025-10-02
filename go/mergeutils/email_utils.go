package mergeutils

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/emailtemplate"
	"delivrio.io/go/ent/notification"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"github.com/mailgun/mailgun-go"
	"net/url"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("emailutils: may not set config twice")
	}
	conf = c
	confSet = true
}

func SendTransactionalEmail(templ, subject, toEmail string, data any) (bool, error) {
	mergedMsg, err := MergeTemplate(templ, data)
	if err != nil {
		return false, fmt.Errorf("send tx email: valid merge: %w", err)
	}

	mergedSubject, err := MergeTemplate(subject, data)
	if err != nil {
		return false, fmt.Errorf("send tx email: valid subject merge: %w", err)
	}

	mg := mailgun.NewMailgun(conf.Email.Mailgun.MGDomain, conf.Email.Mailgun.MGAPIKey)
	mg.SetAPIBase(conf.Email.Mailgun.MGURL)
	if err != nil {
		return false, fmt.Errorf("send tx email: mg: %w", err)
	}

	m := mg.NewMessage(
		fmt.Sprintf("%v <no-reply@%s>", conf.Email.DefaultFrom, conf.Email.Mailgun.MGDomain),
		mergedSubject.String(),
		"",
		toEmail,
	)
	m.SetHtml(mergedMsg.String())
	_, _, err = mg.Send(m)
	if err != nil {
		return false, fmt.Errorf("send tx email: %w", err)
	}

	return true, nil
}

func SendCCNotifyEmail(ctx context.Context, shipmentParcel *ent.ShipmentParcel) error {

	if shipmentParcel.Status == shipmentparcel.StatusAwaitingCcPickup {

		colli, err := shipmentParcel.QueryColli().
			WithRecipient().
			Only(ctx)
		if err != nil {
			return err
		}

		do, err := colli.QueryDeliveryOption().
			WithEmailClickCollectAtStore().
			Only(ctx)
		if err != nil {
			return err
		}
		if do.ClickCollect && do.Edges.EmailClickCollectAtStore != nil {
			email := do.Edges.EmailClickCollectAtStore

			emailData := CCAwaitingPickupNotice{
				CustomerFirstName: "Test name",
			}

			_, err = SendTransactionalEmail(
				email.HTMLTemplate,
				email.Subject,
				"klinsly@gmail.com", //colli.Edges.Recipient.Email,
				emailData,
			)
		}
	}

	return nil

}

func SendColliPackedConfirmation(ctx context.Context, col *ent.Colli) error {
	notif, err := col.QueryOrder().
		QueryConnection().
		QueryNotifications().
		Where(notification.HasEmailTemplateWith(emailtemplate.MergeTypeEQ(emailtemplate.MergeTypeOrderPicked))).
		WithEmailTemplate().
		All(ctx)
	if err != nil {
		return err
	}

	ord, err := col.Order(ctx)
	if err != nil {
		return err
	}

	merge, err := colliPackedToMerge(ctx, ord.OrderPublicID, col)
	if err != nil {
		return err
	}

	recip, err := col.Recipient(ctx)
	if err != nil {
		return err
	}

	for _, n := range notif {
		_, err = SendTransactionalEmail(
			n.Edges.EmailTemplate.HTMLTemplate,
			n.Edges.EmailTemplate.Subject,
			recip.Email,
			merge,
		)
		if err != nil {
			return err
		}
	}

	return nil

}

func SendOrderConfirmation(ctx context.Context, ord *ent.Order) error {
	notif, err := ord.QueryConnection().
		QueryNotifications().
		Where(notification.HasEmailTemplateWith(emailtemplate.MergeTypeEQ(emailtemplate.MergeTypeOrderConfirmation))).
		WithEmailTemplate().
		All(ctx)
	if err != nil {
		return err
	}

	collis, err := ord.QueryColli().
		WithOrderLines().
		WithRecipient().
		All(ctx)
	if err != nil {
		return err
	}

	// TODO: refactor to single summary email
	for _, c := range collis {
		merge, err := colliConfirmationToMerge(ctx, ord.OrderPublicID, c)
		if err != nil {
			return err
		}

		for _, n := range notif {
			_, err = SendTransactionalEmail(
				n.Edges.EmailTemplate.HTMLTemplate,
				n.Edges.EmailTemplate.Subject,
				c.Edges.Recipient.Email,
				merge,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func colliConfirmationToMerge(ctx context.Context, orderPublicID string, colli *ent.Colli) (*OrderConfirmation, error) {
	recip, err := colli.Recipient(ctx)
	if err != nil {
		return nil, err
	}

	dp, err := colli.QueryParcelShop().
		QueryAddress().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	dropPointAdr := &DropPointAddress{}
	if dp != nil {
		dropPointAdr, err = addressToDropPointAddress(ctx, dp)
		if err != nil {
			return nil, err
		}
	}

	orderConfirm, err := addressToCustomerAddress(ctx, recip)
	if err != nil {
		return nil, err
	}

	orderLines, err := orderLinesToMerge(ctx, colli.Edges.OrderLines)
	if err != nil {
		return nil, err
	}

	return &OrderConfirmation{
		CustomerAddress:  *orderConfirm,
		DropPointAddress: *dropPointAdr,
		OrderPublicID:    orderPublicID,
		OrderLines:       orderLines,
	}, nil
}

func colliPackedToMerge(ctx context.Context, orderPublicID string, colli *ent.Colli) (*ColliPacked, error) {
	recip, err := colli.Recipient(ctx)
	if err != nil {
		return nil, err
	}

	send, err := colli.Sender(ctx)
	if err != nil {
		return nil, err
	}

	dp, err := colli.QueryParcelShop().
		QueryAddress().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	recipAdr, err := addressToCustomerAddress(ctx, recip)
	if err != nil {
		return nil, err
	}

	senderAdr, err := addressToSenderAddress(ctx, send)
	if err != nil {
		return nil, err
	}

	dropPointAdr := &DropPointAddress{}
	if dp != nil {
		dropPointAdr, err = addressToDropPointAddress(ctx, dp)
		if err != nil {
			return nil, err
		}
	}

	tracking, err := colli.QueryShipmentParcel().
		Where(shipmentparcel.HasShipmentWith(shipment.StatusNotIn(shipment.StatusDeleted))).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	car, err := colli.QueryDeliveryOption().
		QueryCarrierService().
		QueryCarrierBrand().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	orderLines, err := orderLinesToMerge(ctx, colli.Edges.OrderLines)
	if err != nil {
		return nil, err
	}

	return &ColliPacked{
		CustomerAddress:      *recipAdr,
		DropPointAddress:     *dropPointAdr,
		SenderAddress:        *senderAdr,
		OrderPublicID:        orderPublicID,
		OrderLines:           orderLines,
		TrackingID:           tracking.ItemID,
		ExpectedDeliveryDate: tracking.ExpectedAt,
		CarrierName:          car.Label,
	}, nil
}

func returnColliBase(ctx context.Context, returnColli *ent.ReturnColli) (*ReturnBase, error) {
	recip, err := returnColli.Sender(ctx)
	if err != nil {
		return nil, err
	}

	recipMerge, err := addressToCustomerAddress(ctx, recip)
	if err != nil {
		return nil, err
	}

	returnDestination, err := returnColli.Recipient(ctx)
	if err != nil {
		return nil, err
	}

	returnDestinationMerge, err := addressToReturnAddress(ctx, returnDestination)
	if err != nil {
		return nil, err
	}

	returnOrderLines, err := returnColli.ReturnOrderLine(ctx)
	if err != nil {
		return nil, err
	}

	returnOrderLinesMerge, err := returnOrderLinesToMerge(ctx, returnOrderLines)
	if err != nil {
		return nil, err
	}

	deliveryOption, err := returnColli.DeliveryOption(ctx)
	if err != nil {
		return nil, err
	}

	ord, err := returnColli.Order(ctx)
	if err != nil {
		return nil, err
	}

	return &ReturnBase{
		CustomerAddress:  *recipMerge,
		ReturnAddress:    *returnDestinationMerge,
		OrderPublicID:    ord.OrderPublicID,
		ReturnMethodName: deliveryOption.Name,
		TrackingID:       "",
		OrderLines:       returnOrderLinesMerge,
	}, nil
}

func ReturnColliURL(baseURL string, path string, returnColliID pulid.ID, orderPublicID string) (*url.URL, error) {
	rawURL, err := url.Parse(fmt.Sprintf("%v%v%v", baseURL, delivrioroutes.API, path))
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add(delivrioroutes.QueryParamReturnColliID, returnColliID.String())
	params.Add(delivrioroutes.QueryParamOrderPublicID, orderPublicID)
	rawURL.RawQuery = params.Encode()

	return rawURL, nil
}

func ReturnColliConfirmationLabelMerge(ctx context.Context, returnColli *ent.ReturnColli) (*ReturnConfirmationLabel, error) {
	base, err := returnColliBase(ctx, returnColli)
	if err != nil {
		return nil, err
	}

	downloadURL, err := ReturnColliURL(
		conf.BaseURL,
		delivrioroutes.ReturnLabelDownload,
		returnColli.ID,
		base.OrderPublicID,
	)
	if err != nil {
		return nil, err
	}

	viewURL, err := ReturnColliURL(
		conf.BaseURL,
		delivrioroutes.ReturnLabel,
		returnColli.ID,
		base.OrderPublicID,
	)
	if err != nil {
		return nil, err
	}

	viewPNGURL, err := ReturnColliURL(
		conf.BaseURL,
		delivrioroutes.ReturnLabelPNG,
		returnColli.ID,
		base.OrderPublicID,
	)
	if err != nil {
		return nil, err
	}

	return &ReturnConfirmationLabel{
		ReturnBase:       *base,
		LabelDownloadURL: downloadURL.String(),
		LabelURL:         viewURL.String(),
		LabelPNGURL:      viewPNGURL.String(),
	}, nil
}

func ReturnColliConfirmationQRMerge(ctx context.Context, returnColli *ent.ReturnColli) (*ReturnConfirmationQRCode, error) {
	base, err := returnColliBase(ctx, returnColli)
	if err != nil {
		return nil, err
	}

	downloadURL, err := ReturnColliURL(
		conf.BaseURL,
		delivrioroutes.ReturnQRCodeDownload,
		returnColli.ID,
		base.OrderPublicID,
	)
	if err != nil {
		return nil, err
	}

	viewURL, err := ReturnColliURL(
		conf.BaseURL,
		delivrioroutes.ReturnQRCode,
		returnColli.ID,
		base.OrderPublicID,
	)
	if err != nil {
		return nil, err
	}

	return &ReturnConfirmationQRCode{
		ReturnBase:        *base,
		QRCodeDownloadURL: downloadURL.String(),
		QRCodeURL:         viewURL.String(),
	}, nil
}

func ReturnColliReceivedMerge(ctx context.Context, returnColli *ent.ReturnColli) (*ReturnReceived, error) {
	base, err := returnColliBase(ctx, returnColli)
	if err != nil {
		return nil, err
	}

	return &ReturnReceived{
		ReturnBase: *base,
	}, nil
}

func ReturnColliAcceptedMerge(ctx context.Context, returnColli *ent.ReturnColli) (*ReturnAccepted, error) {
	base, err := returnColliBase(ctx, returnColli)
	if err != nil {
		return nil, err
	}

	return &ReturnAccepted{
		ReturnBase: *base,
	}, nil
}

func addressToReturnAddress(ctx context.Context, recip *ent.Address) (*ReturnAddress, error) {
	c, err := recip.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &ReturnAddress{
		ReturnFirstName: recip.FirstName,
		ReturnLastName:  recip.LastName,
		ReturnCompany:   recip.Company,
		ReturnEmail:     recip.Email,
		ReturnAddress1:  recip.AddressOne,
		ReturnAddress2:  recip.AddressTwo,
		ReturnZip:       recip.Zip,
		ReturnCity:      recip.City,
		ReturnState:     recip.State,
		ReturnCountry:   c.Alpha2,
	}, nil
}

func addressToCustomerAddress(ctx context.Context, adr *ent.Address) (*CustomerAddress, error) {
	c, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &CustomerAddress{
		CustomerFirstName: adr.FirstName,
		CustomerLastName:  adr.LastName,
		CustomerCompany:   adr.Company,
		CustomerEmail:     adr.Email,
		CustomerAddress1:  adr.AddressOne,
		CustomerAddress2:  adr.AddressTwo,
		CustomerZip:       adr.Zip,
		CustomerCity:      adr.City,
		CustomerState:     adr.State,
		CustomerCountry:   c.Alpha2,
	}, nil
}

func addressToSenderAddress(ctx context.Context, adr *ent.Address) (*SenderAddress, error) {
	c, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &SenderAddress{
		SenderFirstName: adr.FirstName,
		SenderLastName:  adr.LastName,
		SenderCompany:   adr.Company,
		SenderEmail:     adr.Email,
		SenderAddress1:  adr.AddressOne,
		SenderAddress2:  adr.AddressTwo,
		SenderZip:       adr.Zip,
		SenderCity:      adr.City,
		SenderState:     adr.State,
		SenderCountry:   c.Alpha2,
	}, nil
}

func addressToDropPointAddress(ctx context.Context, adr *ent.AddressGlobal) (*DropPointAddress, error) {
	c, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &DropPointAddress{
		DropPointCompany:  adr.Company,
		DropPointAddress1: adr.AddressOne,
		DropPointAddress2: adr.AddressTwo,
		DropPointZip:      adr.Zip,
		DropPointCity:     adr.City,
		DropPointState:    adr.State,
		DropPointCountry:  c.Alpha2,
	}, nil
}

func returnOrderLinesToMerge(ctx context.Context, returnOrderLines ent.ReturnOrderLines) ([]ReturnOrderLine, error) {
	output := make([]ReturnOrderLine, 0)

	for _, rol := range returnOrderLines {
		ol, err := rol.OrderLine(ctx)
		if err != nil {
			return nil, err
		}

		productInfo, err := orderLineProductInfo(ctx, ol)
		if err != nil {
			return nil, err
		}

		lineClaim, err := rol.ReturnPortalClaim(ctx)
		if err != nil {
			return nil, err
		}

		output = append(output, ReturnOrderLine{
			OrderLine: OrderLine{
				Quantity:             fmt.Sprintf("%v", rol.Units),
				Price:                fmt.Sprintf("%v", ol.UnitPrice),
				Total:                fmt.Sprintf("%v", float64(rol.Units)*ol.UnitPrice),
				OrderLineProductInfo: *productInfo,
			},
			ReturnClaimName:        lineClaim.Name,
			ReturnClaimDescription: lineClaim.Description,
		})
	}

	return output, nil
}

func orderLinesToMerge(ctx context.Context, orderLines ent.OrderLines) ([]OrderLine, error) {
	output := make([]OrderLine, 0)

	for _, ol := range orderLines {
		productInfo, err := orderLineProductInfo(ctx, ol)
		if err != nil {
			return nil, err
		}

		output = append(output, OrderLine{
			Quantity:             fmt.Sprintf("%v", ol.Units),
			Price:                fmt.Sprintf("%v", ol.UnitPrice),
			Total:                fmt.Sprintf("%v", float64(ol.Units)*ol.UnitPrice),
			OrderLineProductInfo: *productInfo,
		})
	}

	return output, nil
}

func orderLineProductInfo(ctx context.Context, oderLine *ent.OrderLine) (*OrderLineProductInfo, error) {
	variant, err := oderLine.QueryProductVariant().
		WithProduct().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	firstImg := ""
	if len(variant.Edges.ProductImage) > 0 {
		firstImg = variant.Edges.ProductImage[0].URL
	} else {
		productImg, err := variant.Edges.Product.QueryProductImage().
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if productImg != nil {
			firstImg = productImg.URL
		}
	}

	return &OrderLineProductInfo{
		ProductName:          variant.Edges.Product.Title,
		ProductVariantName:   variant.Description,
		ProductFirstImageURL: firstImg,
	}, nil
}
