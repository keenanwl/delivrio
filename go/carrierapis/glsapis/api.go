package glsapis

import (
	"bytes"
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/carrierapis/glsapis/glsrequest"
	"delivrio.io/go/carrierapis/glsapis/glsresponse"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentgls"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var timeout = 15 * time.Second

const glsTimestampFormat = "20060102"

type GLSAPIAuth struct {
	UserName   string
	Password   string
	ContactID  string
	CustomerID string
}

type RequestConfig struct {
	GLSAPIAuth
	// GLS only supports a single shipment per HTTP request
	Shipment ShipmentConfig
}

type ShipmentConfig struct {
	GLSAPIAuth
	ConsignorAddress   *ent.Address
	ConsigneeAddress   *ent.Address
	ParcelShopAddress  *ent.ParcelShop
	AdditionalServices []AdditionalService
	Packages           []PackageConfig
}

// Includes both regular "services" and "add-ons" since GLS doesn't differentiate
// Seems they default to business parcels
type AdditionalService struct {
	Key   string
	Value string
}

type PackageConfig struct {
	DelivrioShipmentID pulid.ID
	DelivrioColliID    pulid.ID
	Items              []*ent.ProductVariant
}

type FetchLabelOutput struct {
	ShipmentConfig ShipmentConfig
	Response       *glsresponse.SuccessLabel
	Error          error
}

func FetchLabels(ctx context.Context, deliveryOption pulid.ID, collis []*ent.Colli) ([]FetchLabelOutput, error) {
	shipments, err := common.GroupPackagesBySenderReceiver(ctx, collis)
	if err != nil {
		return nil, err
	}

	shipmentConfigs, err := createShipmentConfig(ctx, deliveryOption, shipments)
	if err != nil {
		return nil, err
	}

	output := make([]FetchLabelOutput, 0)

	for _, s := range shipmentConfigs {
		resp, err := requestLabels(ctx, deliveryOption, s)
		// TODO: Potentially continue on err?
		if err != nil {
			return nil, err
		}

		output = append(output, FetchLabelOutput{
			ShipmentConfig: s,
			Response:       resp,
			Error:          err,
		})
	}

	return output, nil
}

func fireRequest(req *http.Request) (*http.Response, error) {
	// TODO: move up call chain to enable keep-alive if required
	client := http.Client{
		Timeout: timeout,
	}
	return client.Do(req)
}

func generateRequest(ctx context.Context, u *url.URL, payload interface{}) (*http.Request, error) {
	requestJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(requestJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

var ErrServiceNotSupported = fmt.Errorf("Deposit/AddOnLiabilityService service not supported")

func generateV1CreateShipment(ctx context.Context, config RequestConfig) (*url.URL, *glsrequest.Shipment, error) {
	u, err := url.Parse("http://api.gls.dk/ws/DK/V1/CreateShipment")
	if err != nil {
		return nil, nil, err
	}

	parcels := make([]glsrequest.Parcel, 0)
	for _, p := range config.Shipment.Packages {
		totalWeight := 0
		for _, i := range p.Items {
			if i.WeightG != nil {
				totalWeight += *i.WeightG
			}
		}
		parcels = append(parcels, glsrequest.Parcel{
			Weight: float64(totalWeight) / 1000,
		})
	}

	recip, err := delivrioRecipientToGLSAddress(ctx, config.Shipment.ConsigneeAddress)
	if err != nil {
		return nil, nil, err
	}

	sender, err := delivrioPackageSenderToGLSAddress(ctx, config.Shipment.ConsignorAddress)
	if err != nil {
		return nil, nil, err
	}

	// Default to standard outbound since biz deliveries don't include AdditionalServices
	var delivery = recip
	var pickup *glsrequest.Address
	var alternativeShipper = sender
	addServices := glsrequest.Services{}
	for _, a := range config.Shipment.AdditionalServices {

		switch a.Key {
		case "PrivateDelivery":
			addServices.PrivateDelivery = &a.Value
			break
		case "FlexDelivery":
			addServices.FlexDelivery = &a.Value
			break
		case "NotificationEmail":
			addServices.NotificationEmail = &a.Value
			break
		case "ShopReturn":
			delivery = sender
			pickup = recip
			alternativeShipper = nil
			addServices.ShopReturn = &a.Value
			break
		case "DirectShop":
			addServices.DirectShop = &a.Value
			break
		case "ShopDelivery":
			addServices.ShopDelivery = &a.Value
			break
		case "Express12":
			addServices.Express12 = &a.Value
			break
		case "Express10":
			addServices.Express10 = &a.Value
			break
		case "Deposit":
		case "AddOnLiabilityService":
		default:
			return nil, nil, ErrServiceNotSupported
		}

	}

	ship := &glsrequest.Shipment{
		UserName:     config.UserName,
		Password:     config.Password,
		Customerid:   config.CustomerID,
		Contactid:    config.ContactID,
		ShipmentDate: time.Now().Format(glsTimestampFormat),
		Reference:    "",
		Addresses: glsrequest.Addresses{
			Delivery:           *delivery,
			Pickup:             pickup,
			AlternativeShipper: alternativeShipper,
		},
		Parcels:  parcels,
		Services: addServices,
	}

	return u, ship, nil

}

func delivrioRecipientToGLSAddress(ctx context.Context, recipient *ent.Address) (*glsrequest.Address, error) {
	c, err := recipient.Country(ctx)
	if err != nil {
		return nil, err
	}
	return &glsrequest.Address{
		Name1: recipient.FirstName,
		Name2: recipient.LastName,
		// GLS caps at 40 characters
		Street1:    limitString(fmt.Sprintf("%v %v", recipient.AddressOne, recipient.AddressTwo), 40),
		CountryNum: c.Code,
		ZipCode:    recipient.Zip,
		City:       recipient.City,
		Contact:    fmt.Sprintf("%v %v", recipient.FirstName, recipient.LastName),
		Email:      recipient.Email,
		Phone:      recipient.PhoneNumber,
		Mobile:     recipient.PhoneNumber,
	}, nil
}

func limitString(input string, maxLength int) string {
	if len(input) <= maxLength {
		return input
	}
	return input[:maxLength]
}

func delivrioPackageSenderToGLSAddress(ctx context.Context, sender *ent.Address) (*glsrequest.Address, error) {
	c, err := sender.Country(ctx)
	if err != nil {
		return nil, err
	}
	return &glsrequest.Address{
		Name1:      sender.FirstName,
		Name2:      sender.LastName,
		Street1:    limitString(fmt.Sprintf("%v %v", sender.AddressOne, sender.AddressTwo), 40),
		CountryNum: c.Code,
		ZipCode:    sender.Zip,
		City:       sender.City,
		Contact:    fmt.Sprintf("%v %v", sender.FirstName, sender.LastName),
		Email:      sender.Email,
		Phone:      sender.PhoneNumber,
		Mobile:     sender.PhoneNumber,
	}, nil
}

func SaveLabelData(ctx context.Context, shipConfig ShipmentConfig, resp FetchLabelOutput) ([]common.CreateShipment, error) {

	if len(shipConfig.Packages) != len(resp.Response.Parcels) {
		return nil, fmt.Errorf("expected %v packages got %v", shipConfig.Packages, resp.Response.Parcels)
	}

	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	decodePDF, err := base64.StdEncoding.DecodeString(resp.Response.PDF)
	if err != nil {
		return nil, err
	}

	pdfPages, err := utils.SplitPDFPagesToB64(decodePDF)
	if err != nil {
		return nil, err
	}

	if len(pdfPages) != len(shipConfig.Packages) {
		return nil, fmt.Errorf("expected %v PDF pages got %v", shipConfig.Packages, len(pdfPages))
	}

	allCreateShipments := make([]common.CreateShipment, 0)

	for i, p := range shipConfig.Packages {
		sp, err := tx.ShipmentParcel.Create().
			SetShipmentID(p.DelivrioShipmentID).
			SetStatus(shipmentparcel.StatusPending).
			SetItemID(resp.Response.Parcels[i].ParcelNumber).
			SetColliID(p.DelivrioColliID).
			SetTenantID(view.TenantID()).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		_, err = utils.CreateShipmentDocument(ctx, sp, &pdfPages[i], nil)
		if err != nil {
			return nil, err
		}

		err = tx.ShipmentGLS.Update().
			SetConsignmentID(resp.Response.ConsignmentID).
			Where(shipmentgls.HasShipmentWith(shipment.ID(p.DelivrioShipmentID))).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		allCreateShipments = append(allCreateShipments, common.CreateShipment{
			Shipment: p.DelivrioShipmentID,
			Labels:   []string{pdfPages[i]},
		})
	}

	return allCreateShipments, nil

}
