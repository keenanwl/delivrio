package dfapis

import (
	"bytes"
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/carrierapis/dfapis/dfrequest"
	"delivrio.io/go/carrierapis/dfapis/dfresponse"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierdf"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/consolidation"
	"delivrio.io/go/ent/documentfile"
	"delivrio.io/go/ent/pallet"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentpallet"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("dfapis: may not set config twice")
	}
	conf = c
	confSet = true
}

type ConsignmentConfig struct {
	dfrequest.APIAccessConfig
	ConsolidationID   pulid.ID
	Test              bool
	WhoPays           carrierdf.WhoPays
	DFID              *string // Empty on first pre-book
	Prebook           bool
	DFConsignmentID   *string
	ShipmentID        pulid.ID
	DeliveryOption    *ent.DeliveryOption
	ConsignorAddress  *ent.Address
	ConsigneeAddress  *ent.Address
	ParcelShopAddress *ent.ParcelShop
	Collis            []*ColliConfig
}

// Maps to Order or Pallet&Orders
type ColliConfig struct {
	DelivrioShipmentID pulid.ID
	// Mutually exclusive
	DelivrioColliID  pulid.ID
	DelivrioPalletID pulid.ID

	OrderLines []*ent.OrderLine
}

func (cc *ColliConfig) IsPallet() bool {
	return len(cc.DelivrioPalletID) > 0
}

type ConsignmentOutput struct {
	// DF returns different schema if there are no actual updates
	// When this occurs, we can just output the existing labels
	IsNoop      bool
	DFID        string
	ColliOutput []ColliOutput
	Config      *ConsignmentConfig
}

type ColliOutput struct {
	Package        *ColliConfig
	DFColliNumber  int
	DfColliID      int
	Barcode        string
	ResponseB64PDF string
	Error          error
}

func PalletToColliConfig(ctx context.Context, shipmentID pulid.ID, p *ent.Pallet) (*ColliConfig, error) {
	orderLines, err := p.QueryOrders().
		QueryColli().
		QueryOrderLines().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return &ColliConfig{
		DelivrioShipmentID: shipmentID,
		DelivrioPalletID:   p.ID,
		OrderLines:         orderLines,
	}, nil
}

func OrderToColliConfig(ctx context.Context, shipmentID pulid.ID, o *ent.Order) ([]*ColliConfig, error) {
	collis, err := o.QueryColli().
		WithOrderLines().
		All(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*ColliConfig, 0)
	for _, c := range collis {
		output = append(output, &ColliConfig{
			DelivrioShipmentID: shipmentID,
			DelivrioColliID:    c.ID,
			OrderLines:         c.Edges.OrderLines,
		})
	}

	return output, nil
}

// Summarize each duplicate colli (WIP)
// Just contents/description, since they are already part of
// the consolidation. ie sender/receiver is overwritten.
func toDFGoods(ctx context.Context, configs []*ColliConfig) ([]dfrequest.Good, error) {

	// 1:1 with input config
	flatGoods := make([]dfrequest.Good, 0)
	for _, cc := range configs {
		g, err := configToGood(ctx, cc)
		if err != nil {
			return nil, err
		}
		flatGoods = append(flatGoods, *g)
	}

	return flatGoods, nil

	// input config may now be > goods
	// but we finally process goods response
	// as individual labels, we just collapse so
	// the DF docs show 5 pallets on 1 line when
	// the pallets are identical
	// Needs implementing

}

func configToGood(ctx context.Context, config *ColliConfig) (*dfrequest.Good, error) {
	cli := ent.FromContext(ctx)

	totalPalletWeight, err := common.ColliWeightKG(ctx, config.OrderLines)
	if err != nil {
		return nil, fmt.Errorf("df: calculate pallet weight: %w", err)
	}

	// DF API docs have Weight listed as double,
	// but seems to require an +Int
	totalPalletWeight = math.Ceil(totalPalletWeight)

	var pack *ent.Packaging
	description := ""
	ref := ""
	if config.IsPallet() {
		p, err := cli.Pallet.Query().
			Where(pallet.ID(config.DelivrioPalletID)).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		description = p.Description
		ref = p.PublicID

		pack, err = p.QueryPackaging().
			WithPackagingDF().
			Only(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		c, err := cli.Colli.Query().
			Where(colli.ID(config.DelivrioColliID)).
			WithOrder().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		description = c.Edges.Order.CommentExternal
		ref = c.Edges.Order.OrderPublicID

		pack, err = common.ColliPackaging(ctx, c)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if ent.IsNotFound(err) {
			return nil, fmt.Errorf("df requires colli packaging is added")
		}
	}

	if pdf, err := pack.PackagingDF(ctx); err != nil || pdf == nil {
		return nil, fmt.Errorf("df: expected DF packaging type")
	}

	return &dfrequest.Good{
		DID:                 nil,
		NumberOfItems:       1,
		Type:                pack.Edges.PackagingDF.APIType.String(),
		Description:         description,
		Weight:              totalPalletWeight,
		Volume:              nil,
		Amount:              nil,
		Length:              &pack.LengthCm,
		Width:               &pack.WidthCm,
		Height:              &pack.HeightCm,
		Stackable:           &pack.Edges.PackagingDF.Stackable,
		LoadMeters:          nil,
		SenderRef:           &ref,
		DangerousGoods:      nil,
		Products:            nil,
		ColliCodes:          nil,
		FragtpligtigVaegt:   nil,
		FragtpligtigRumfang: nil,
	}, nil

}

func FetchLabels(ctx context.Context, prebook bool, c *ent.Consolidation) (*ConsignmentOutput, error) {
	co, err := configFromConsolidation(ctx, prebook, c)
	if err != nil {
		return nil, err
	}

	return createDFConsignmentFetchLabel(ctx, co)
}

func SaveLabel(ctx context.Context, consign *ConsignmentOutput) (*common.CreateShipment, error) {
	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	ctxTX, _, err := cli.OpenTx(ctx)
	if err != nil {
		return nil, err
	}
	tx := ent.TxFromContext(ctx)
	defer tx.Rollback()

	nextShipmentStatus := consolidation.StatusBooked
	if consign.Config.Prebook {
		nextShipmentStatus = consolidation.StatusPrebooked
	}

	st := shipment2.StatusBooked
	if consign.Config.Prebook {
		st = shipment2.StatusPrebooked
	}

	err = tx.Shipment.Update().
		// Saving DF ID happens right after initial POST request is successful
		SetStatus(st).
		Where(shipment2.ID(consign.Config.ShipmentID)).
		Exec(ctxTX)
	if err != nil {
		return nil, err
	}

	err = tx.Consolidation.Update().
		SetStatus(nextShipmentStatus).
		Where(consolidation.ID(consign.Config.ConsolidationID)).
		Exec(ctxTX)
	if err != nil {
		return nil, err
	}

	// Only create if this is a new shipment
	if consign.Config.DFConsignmentID == nil {
		// No extra fields yet.
		// Include ShipmentDF for consistency with other shipments
		err = tx.ShipmentDF.Create().
			SetShipmentID(consign.Config.ShipmentID).
			SetTenantID(view.TenantID()).
			Exec(ctxTX)
		if err != nil {
			return nil, err
		}
	}

	// Stop the TX to save to S3 possibly
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	savedLabels := make(map[int]*ent.DocumentFile)
	for _, c := range consign.ColliOutput {
		df, err := utils.CreateShipmentDocument(ctx, nil, &c.ResponseB64PDF, nil)
		if err != nil {
			return nil, err
		}

		savedLabels[c.DfColliID] = df
	}

	ctxTX, _, err = cli.OpenTx(ctx)
	if err != nil {
		return nil, err
	}
	tx = ent.TxFromContext(ctx)
	defer tx.Rollback()

	// Consider status change instead?
	_, err = tx.ShipmentPallet.Delete().
		Where(shipmentpallet.HasShipmentWith(shipment2.ID(consign.Config.ShipmentID))).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.ShipmentParcel.Delete().
		Where(shipmentparcel.HasShipmentWith(shipment2.ID(consign.Config.ShipmentID))).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	outputLabels := make([]string, 0)
	for _, c := range consign.ColliOutput {
		outputLabels = append(outputLabels, c.ResponseB64PDF)

		if c.Package.IsPallet() {
			err = tx.ShipmentPallet.Create().
				SetTenantID(view.TenantID()).
				SetLabelPdf(c.ResponseB64PDF).
				SetBarcode(c.Barcode).
				SetStatus(shipmentpallet.StatusPending).
				SetColliNumber(strconv.Itoa(c.DFColliNumber)).
				SetCarrierID(strconv.Itoa(c.DfColliID)).
				SetPalletID(c.Package.DelivrioPalletID).
				SetShipmentID(c.Package.DelivrioShipmentID).
				Exec(ctx)
			if err != nil {
				return nil, err
			}
		} else {
			sp, err := tx.ShipmentParcel.Create().
				SetTenantID(view.TenantID()).
				SetItemID(c.Barcode).
				SetColliID(c.Package.DelivrioColliID).
				SetStatus(shipmentparcel.StatusPending).
				SetShipmentID(c.Package.DelivrioShipmentID).
				Save(ctx)
			if err != nil {
				return nil, err
			}

			if savedLabels[c.DfColliID] != nil {
				err = tx.DocumentFile.Update().
					SetShipmentParcel(sp).
					Where(documentfile.ID(savedLabels[c.DfColliID].ID)).
					Exec(ctxTX)
				if err != nil {
					return nil, err
				}
			}

		}

	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &common.CreateShipment{
		Shipment: consign.Config.ShipmentID,
		Labels:   outputLabels,
	}, nil

}

func configFromConsolidation(ctx context.Context, prebook bool, c *ent.Consolidation) (*ConsignmentConfig, error) {
	do, err := c.QueryDeliveryOption().
		WithDeliveryOptionDF().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	agreement, err := do.QueryCarrier().
		QueryCarrierDF().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	sender, err := c.Sender(ctx)
	if err != nil {
		return nil, err
	}
	receiver, err := c.Recipient(ctx)
	if err != nil {
		return nil, err
	}

	if sender == nil || receiver == nil {
		return nil, fmt.Errorf("sender & recipient address are required")
	}

	pal, err := c.Pallets(ctx)
	if err != nil {
		return nil, err
	}
	ord, err := c.Orders(ctx)
	if err != nil {
		return nil, err
	}

	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	var currentConsignmentID *string
	ship, err := c.QueryShipment().
		// This means the previous shipment request did not go through
		// and we don't have a consignment ID to update
		Where(shipment2.StatusNotIn(shipment2.StatusPending)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if ent.IsNotFound(err) {
		// Archives a pending shipment
		err = archiveShipmentConsolidation(ctx, c)
		if err != nil {
			return nil, err
		}

		car, err := c.QueryDeliveryOption().
			QueryCarrier().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		ship, err = cli.Shipment.Create().
			SetCarrier(car).
			SetShipmentPublicID(fmt.Sprintf("%v", time.Now())).
			SetConsolidation(c).
			SetTenantID(view.TenantID()).
			SetStatus(shipment2.StatusPending).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {

		switch ship.Status {
		case shipment2.StatusPrebooked:
			currentConsignmentID = &ship.ShipmentPublicID
			break
		case shipment2.StatusBooked:
		case shipment2.StatusDispatched:
			return nil, fmt.Errorf("df: consignment already has booked shipment: remove existing before creating new")
		}

	}

	requestCollis := make([]*ColliConfig, 0)
	for _, p := range pal {
		c, err := PalletToColliConfig(ctx, ship.ID, p)
		if err != nil {
			return nil, err
		}
		requestCollis = append(requestCollis, c)
	}

	for _, o := range ord {
		c, err := OrderToColliConfig(ctx, ship.ID, o)
		if err != nil {
			return nil, err
		}
		requestCollis = append(requestCollis, c...)
	}

	return &ConsignmentConfig{
		APIAccessConfig: dfrequest.APIAccessConfig{
			CustomerID:      agreement.CustomerID,
			AgreementNumber: agreement.AgreementNumber,
			APICredentials: dfrequest.APICredentials{
				ClientID:  conf.DanskeFragtmaend.ClientID,
				GrantType: conf.DanskeFragtmaend.GrantType,
				Username:  conf.DanskeFragtmaend.Username,
				Password:  conf.DanskeFragtmaend.Password,
				Resource:  conf.DanskeFragtmaend.Resource,
			},
		},
		ConsolidationID:   c.ID,
		ShipmentID:        ship.ID,
		Test:              agreement.Test,
		WhoPays:           agreement.WhoPays,
		DFID:              nil,
		Prebook:           prebook,
		DFConsignmentID:   currentConsignmentID,
		DeliveryOption:    do,
		ConsignorAddress:  sender,
		ConsigneeAddress:  receiver,
		ParcelShopAddress: nil, // Not supported yet
		Collis:            requestCollis,
	}, nil
}

// consider ensuring this happens in a tx
// not huge now since a shipment should only be
// archived if it failed or was manually cancelled
func archiveShipmentConsolidation(ctx context.Context, c *ent.Consolidation) error {
	ship, err := c.QueryShipment().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	} else if ent.IsNotFound(err) {
		return nil
	}

	return c.Update().
		ClearShipment().
		AddCancelledShipments(ship).
		Exec(ctx)
}

func createDFConsignmentFetchLabel(ctx context.Context, cr *ConsignmentConfig) (*ConsignmentOutput, error) {
	httpClient := oauth2Client(ctx, cr.APICredentials, cr.Test)

	createOrderReq, err := NewConsignmentRequest(ctx, cr)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(createOrderReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		var errBody dfresponse.ErrResponse
		err = json.Unmarshal(data, &errBody)
		if err != nil {
			return nil, fmt.Errorf("df: cosnignment request failed: %v: %v", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("expected DF response code %v, got %v: %v", http.StatusCreated, resp.StatusCode, errBody)
	}

	isNoop, err := isNoopConsignmentUpdate(data)
	if err != nil {
		return nil, err
	}

	if isNoop {
		return &ConsignmentOutput{
			IsNoop: true,
		}, nil
	}

	consignmentNumber, goods, err := handleCreateConsignment(data)
	if err != nil {
		return nil, err
	}

	cli := ent.FromContext(ctx)
	// Not ideal, but we need to save the consignment ID
	// quickly in case the label request fails
	cli.Shipment.Update().
		SetShipmentPublicID(consignmentNumber).
		Where(shipment2.ID(cr.ShipmentID)).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	requestLabels, err := labelRequest(ctx, []string{consignmentNumber}, cr.Test)
	if err != nil {
		return nil, err
	}

	resp, err = httpClient.Do(requestLabels)
	if err != nil {
		return nil, fmt.Errorf("df: request labels: %w", err)
	}
	defer resp.Body.Close()

	labelsPDF, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	splitLabels, err := utils.SplitPDFPagesToB64(labelsPDF)
	if err != nil {
		return nil, fmt.Errorf("df: splitting PDF: %w", err)
	}

	colliOutput, err := iterateConsignmentResponse(
		ctx,
		cr.Collis,
		createColliOutputList(ctx, goods),
		splitLabels,
	)

	return &ConsignmentOutput{
		Config:      cr,
		DFID:        consignmentNumber,
		ColliOutput: colliOutput,
	}, nil
}

func handleCreateConsignment(data []byte) (string, []dfresponse.Good, error) {
	var body dfresponse.Consignment
	err := json.Unmarshal(data, &body)
	if err != nil {
		return "", nil, err
	}
	return body.ConsignmentNumber, body.Goods, nil
}

func isNoopConsignmentUpdate(data []byte) (bool, error) {
	var body dfresponse.NoUpdateConsignment
	err := json.Unmarshal(data, &body)
	if err != nil {
		return false, err
	}

	if body.UpdateNote != nil {
		return true, nil
	}

	return false, nil
}

// Must maintain list order (no maps!)
// Creates a long list of collis that we can then map back to the list of
// DELIVRIO collis.
func createColliOutputList(ctx context.Context, goods []dfresponse.Good) []dfresponse.ColliCode {
	output := make([]dfresponse.ColliCode, 0)
	for _, g := range goods {
		output = append(output, g.ColliCodes...)
	}
	return output
}

func iterateConsignmentResponse(
	ctx context.Context,
	inputConfig []*ColliConfig,
	dfColliCodes []dfresponse.ColliCode,
	splitLabels []string,
) ([]ColliOutput, error) {

	if len(inputConfig) != len(dfColliCodes) || len(inputConfig) != len(splitLabels) {
		return nil, fmt.Errorf(
			"df: expected response input colli count == df colli count == label count: %v, %v, %v",
			len(inputConfig),
			len(dfColliCodes),
			len(splitLabels),
		)
	}

	output := make([]ColliOutput, 0)
	for i, cc := range dfColliCodes {
		output = append(output, ColliOutput{
			Package:        inputConfig[i],
			DFColliNumber:  cc.ColliNumber,
			DfColliID:      cc.ID,
			Barcode:        cc.Barcode,
			ResponseB64PDF: splitLabels[i],
			Error:          nil,
		})
	}

	return output, nil
}

func labelRequest(ctx context.Context, consignmentIDs []string, test bool) (*http.Request, error) {

	if len(consignmentIDs) == 0 {
		return nil, fmt.Errorf("df: request labels: consignment IDs must be > 0")
	}

	fullURL, err := url.JoinPath(baseURL(test), `/v1/Report/GetLabelForPrint`)
	if err != nil {
		return nil, fmt.Errorf("df: fetch labels: %w", err)
	}
	u, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("df: fetch labels: %w", err)
	}

	vals := url.Values{}
	vals.Set("consignmentNumbers", strings.Join(consignmentIDs, ","))
	vals.Set("labelType", "Label10x19ForLabelPrinter")
	u.RawQuery = vals.Encode()

	return http.NewRequest("GET", u.String(), nil)
}

func baseURL(test bool) string {
	resourceURL := `https://apistaging.fragt.dk`
	if !test {
		resourceURL = "https://api.fragt.dk"
	}
	return resourceURL
}

func oauth2Client(ctx context.Context, config dfrequest.APICredentials, test bool) *http.Client {

	customParams := url.Values{}

	customParams.Set("Resource", baseURL(test))
	customParams.Set("Username", config.Username)
	customParams.Set("Password", config.Password)
	customParams.Set("grant_type", config.GrantType)

	reqConfig := &clientcredentials.Config{
		ClientID:       config.ClientID,
		ClientSecret:   "",
		TokenURL:       "https://sts.fragt.dk/adfs/oauth2/token",
		Scopes:         []string{},
		AuthStyle:      oauth2.AuthStyleInParams,
		EndpointParams: customParams,
	}

	cliWithLogger := &http.Client{
		Transport: &common.LoggerAuthRoundTripper{
			Transport: http.DefaultTransport,
		},
	}

	ctxWithLogger := context.WithValue(ctx, oauth2.HTTPClient, cliWithLogger)
	ts := reqConfig.TokenSource(ctxWithLogger)
	return oauth2.NewClient(ctxWithLogger, ts)
}

func NewConsignmentRequest(ctx context.Context, cr *ConsignmentConfig) (*http.Request, error) {
	u, err := url.Parse(baseURL(cr.Test))
	if err != nil {
		return nil, err
	}
	if !cr.Test {
		panic("Danske Fragtmænd: only test supported")
	}

	httpVerb := http.MethodPost
	if cr.DFConsignmentID != nil {
		u = u.JoinPath(fmt.Sprintf(`/v1/consignments/%s`, *cr.DFConsignmentID))
		httpVerb = http.MethodPut
	} else {
		u = u.JoinPath(`/v1/consignments`)
	}

	shippingType, noteType, err := shipmentTypeFromDeliveryOption(ctx, cr.DeliveryOption)
	if err != nil {
		return nil, err
	}

	sender, err := addressToDFAddress(ctx, cr.ConsignorAddress)
	if err != nil {
		return nil, err
	}

	receiver, err := addressToDFAddress(ctx, cr.ConsigneeAddress)
	if err != nil {
		return nil, err
	}

	goods, err := toDFGoods(ctx, cr.Collis)
	if err != nil {
		return nil, err
	}

	cons := &dfrequest.Consignment{
		WhoPays:             cr.WhoPays,
		ConsignmentNumber:   cr.DFConsignmentID,
		ShippingType:        shippingType,
		ConsignmentNoteType: noteType,
		HubAgreement:        "AR", // Constant? Not sure what this means
		AgreementNumber:     cr.AgreementNumber,
		Sender:              *sender,
		Initiator:           *sender,
		Receiver:            *receiver,
		PreBooking:          &cr.Prebook,
		ServiceCodes:        []string{"TEST"},
		Goods:               goods,
	}

	jsonBody, err := json.Marshal(cons)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(httpVerb, u.String(), bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func addressToDFAddress(ctx context.Context, adr *ent.Address) (*dfrequest.Address, error) {

	c, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &dfrequest.Address{
		Name:               adr.FirstName,
		Name2:              adr.LastName,
		Name3:              adr.Company,
		Name4:              nil,
		Street:             adr.AddressOne,
		Street2:            adr.AddressTwo,
		Town:               adr.City,
		Zipcode:            adr.Zip,
		CustomerID:         nil,
		Country:            c.Alpha2,
		Phone:              adr.PhoneNumber,
		Email:              adr.Email,
		ContactPerson:      fmt.Sprintf("%v %v", adr.FirstName, adr.LastName),
		ContactPersonPhone: adr.PhoneNumber,
		ContactPersonEmail: adr.Email,
	}, nil

}

func shipmentTypeFromDeliveryOption(ctx context.Context, do *ent.DeliveryOption) (dfrequest.ShippingType, dfrequest.ConsignmentNoteType, error) {

	service, err := do.CarrierService(ctx)
	if err != nil {
		return "", "", err
	}

	switch service.InternalID {
	case "DF_PICK_UP_CARGO":
		return dfrequest.ShippingTypeStykgods, dfrequest.ConsignmentNoteTypePickup, nil
	case "DF_PICK_UP_PALLET":
		return dfrequest.ShippingTypePalleEnhedsForsendelse, dfrequest.ConsignmentNoteTypePickup, nil
	case "DF_PICK_UP_VEHICLE_PARCEL":
		return dfrequest.ShippingTypeBilpakke, dfrequest.ConsignmentNoteTypePickup, nil
	case "DF_RETURN_CARGO":
		return dfrequest.ShippingTypeStykgods, dfrequest.ConsignmentNoteTypeReturn, nil
	case "DF_RETURN_PALLET":
		return dfrequest.ShippingTypePalleEnhedsForsendelse, dfrequest.ConsignmentNoteTypeReturn, nil
	case "DF_RETURN_VEHICLE_PARCEL":
		return dfrequest.ShippingTypeBilpakke, dfrequest.ConsignmentNoteTypeReturn, nil
	case "DF_RELOCATION_CARGO":
		return dfrequest.ShippingTypeStykgods, dfrequest.ConsignmentNoteTypeChangeOfAddress, nil
	case "DF_RELOCATION_PALLET":
		return dfrequest.ShippingTypePalleEnhedsForsendelse, dfrequest.ConsignmentNoteTypeChangeOfAddress, nil
	case "DF_RELOCATION_VEHICLE_PARCEL":
		return dfrequest.ShippingTypeBilpakke, dfrequest.ConsignmentNoteTypeChangeOfAddress, nil
	default:
		return "", "", fmt.Errorf("df shipment type: unknow service: %s", service.InternalID)
	}

}
