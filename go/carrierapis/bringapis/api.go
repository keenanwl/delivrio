package bringapis

import (
	"bytes"
	"context"
	"delivrio.io/go/carrierapis/bringapis/bringrequest"
	"delivrio.io/go/carrierapis/bringapis/bringresponse"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/countryharmonizedcode"
	"delivrio.io/go/utils"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

// CreateOrderFetchLabel Bring requires two separate requests
// to first create the order and secondarily to download a PDF of the label
func CreateOrderFetchLabel(ctx context.Context, sc *ShipmentConfig) ([]*FetchLabelOutput, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	booking, err := newBaseBooking(ctx, sc)
	if err != nil {
		return nil, err
	}

	req, err := newBookingRequest(ctx, sc.AuthenticationHeaders, booking)
	if err != nil {
		return nil, err
	}

	reqDump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(reqDump))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respDump, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(respDump))

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var body bringresponse.ResponseBooking
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bodyStr, _ := json.Marshal(body)
		return nil, fmt.Errorf("bring error: %s", bodyStr)
	}

	if len(body.Consignments) != 1 {
		return nil, fmt.Errorf("expected 1 consignment in response, got %v", len(body.Consignments))
	}

	labelByte, err := fetchLabelData(ctx, httpClient, body.Consignments[0].Confirmation.Links.Labels)
	if err != nil {
		return nil, err
	}

	allLabels, err := utils.SplitPDFPagesToB64(labelByte)
	if err != nil {
		return nil, err
	}

	output := make([]*FetchLabelOutput, 0)

	for pi, rp := range body.Consignments[0].Confirmation.Packages {

		pack, err := packageConfigByIndex(sc.Packages, pi)
		if err != nil {
			return nil, err
		}

		output = append(output, &FetchLabelOutput{
			Package:           *pack,
			ConsignmentNumber: body.Consignments[0].Confirmation.ConsignmentNumber,
			PackageNumber:     rp.PackageNumber,
			ResponseB64PDF:    allLabels[pi],
			Error:             nil,
		})
	}

	return output, nil

}

func packageConfigByIndex(packages []PackageConfig, index int) (*PackageConfig, error) {
	if len(packages)-1 < index {
		return nil, fmt.Errorf("package slice less than index %v", index)
	}

	return &packages[index], nil
}

func fetchLabelData(ctx context.Context, client *http.Client, pdfURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", pdfURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	return nil, fmt.Errorf("bring: fetch label: got status: %v", resp.StatusCode)
}

func newBookingRequest(ctx context.Context, auth bringrequest.AuthenticationHeaders, requestBody *bringrequest.Request) (*http.Request, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.bring.com/booking-api/api/booking", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req = addAuthentication(req, auth)

	return req, nil
}

func newSender(ctx context.Context, orderReference string, adr *ent.Address) (*bringrequest.Sender, error) {
	coun, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &bringrequest.Sender{
		AdditionalAddressInfo: "",
		AddressLine:           adr.AddressOne,
		AddressLine2:          adr.AddressTwo,
		City:                  adr.City,
		Contact: bringrequest.SenderContact{
			Name:        fmt.Sprintf("%s %s", adr.FirstName, adr.LastName),
			PhoneNumber: adr.PhoneNumber,
		},
		CountryCode: coun.Alpha2,
		Name:        fmt.Sprintf("%s %s", adr.FirstName, adr.LastName),
		PostalCode:  adr.Zip,
		Reference:   orderReference,
	}, nil
}

func newRecipient(ctx context.Context, orderReference string, adr *ent.Address) (*bringrequest.Recipient, error) {
	coun, err := adr.Country(ctx)
	if err != nil {
		return nil, err
	}

	return &bringrequest.Recipient{
		AdditionalAddressInfo: "",
		AddressLine:           adr.AddressOne,
		AddressLine2:          adr.AddressTwo,
		City:                  adr.City,
		Contact: bringrequest.RecipientContact{
			Email:       adr.Email,
			Name:        fmt.Sprintf("%s %s", adr.FirstName, adr.LastName),
			PhoneNumber: adr.PhoneNumber,
		},
		CountryCode: coun.Alpha2,
		Name:        fmt.Sprintf("%s %s", adr.FirstName, adr.LastName),
		PostalCode:  adr.Zip,
		Reference:   orderReference,
	}, nil
}

func newPickupPoint(ctx context.Context, config *ShipmentConfig) *bringrequest.PickupPoint {
	if config.ParcelShopCountry == nil {
		return nil
	}

	return &bringrequest.PickupPoint{
		CountryCode: config.ParcelShopCountry.Alpha2,
		BringID:     config.ParcelShopID,
	}
}

func newPackages(ctx context.Context, packages []PackageConfig) ([]bringrequest.Package, error) {
	output := make([]bringrequest.Package, 0)

	for _, p := range packages {
		totalWeightKG, err := common.ColliWeightKG(ctx, p.OrderLines)
		if err != nil {
			return nil, err
		}

		output = append(output, bringrequest.Package{
			CorrelationID: "",
			Dimensions: bringrequest.Dimensions{
				HeightInCM: int64(p.Packaging.HeightCm),
				LengthInCM: int64(p.Packaging.LengthCm),
				WidthInCM:  int64(p.Packaging.WidthCm),
			},
			GoodsDescription: "",
			PackageType:      nil,
			WeightInKg:       totalWeightKG,
		})
	}

	return output, nil
}

func newConsignment(ctx context.Context, shipmentConfig *ShipmentConfig) (*bringrequest.Consignment, error) {
	sender, err := newSender(ctx, shipmentConfig.PublicOrderID, shipmentConfig.ConsignorAddress)
	if err != nil {
		return nil, err
	}

	recipient, err := newRecipient(ctx, shipmentConfig.PublicOrderID, shipmentConfig.ConsigneeAddress)
	if err != nil {
		return nil, err
	}

	packages, err := newPackages(ctx, shipmentConfig.Packages)
	if err != nil {
		return nil, err
	}

	var customs bringrequest.EDICustomsDeclarations
	if shipmentConfig.ElectronicCustoms {
		customs.NatureOfTransaction = "SALE_OF_GOODS"
		tariffLines, err := customsLines(ctx, shipmentConfig)
		if err != nil {
			return nil, err
		}

		customs.EDICustomsDeclaration = tariffLines
	}
	return &bringrequest.Consignment{
		CorrelationID: "",
		Packages:      packages,
		Parties: bringrequest.Parties{
			Recipient:   *recipient,
			Sender:      *sender,
			PickupPoint: newPickupPoint(ctx, shipmentConfig),
		},
		Product: bringrequest.Product{
			AdditionalServices:     make([]bringrequest.AdditionalService, 0),
			CustomerNumber:         shipmentConfig.BringCustomerNumber,
			EDICustomsDeclarations: &customs,
			ID:                     shipmentConfig.BringService,
		},
		ShippingDateTime: shipmentConfig.ShippingDate.Format(time.RFC3339),
	}, nil
}

// TODO: check that grouping by ProductVariant makes sense
func customsLines(ctx context.Context, shipmentConfig *ShipmentConfig) ([]bringrequest.EDICustomsDeclaration, error) {
	dest, err := shipmentConfig.ConsigneeAddress.Country(ctx)
	if err != nil {
		return nil, err
	}
	orderLineConsolidation := make(map[pulid.ID]bringrequest.EDICustomsDeclaration)
	for _, p := range shipmentConfig.Packages {
		for _, ol := range p.OrderLines {
			pv, err := ol.QueryProductVariant().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			val, err := toBringCustomsLine(ctx, ol, dest)
			if err != nil {
				return nil, err
			}

			if item, ok := orderLineConsolidation[pv.ID]; ok {
				orderLineConsolidation[pv.ID] = combineCustomsLines(item, *val)
			} else {
				orderLineConsolidation[pv.ID] = *val
			}
		}
	}

	output := make([]bringrequest.EDICustomsDeclaration, 0)
	for _, cd := range orderLineConsolidation {
		output = append(output, cd)
	}

	return output, nil
}

func combineCustomsLines(base bringrequest.EDICustomsDeclaration, adl ...bringrequest.EDICustomsDeclaration) bringrequest.EDICustomsDeclaration {
	for _, a := range adl {
		base = bringrequest.EDICustomsDeclaration{
			Quantity:             a.Quantity + base.Quantity,
			GoodsDescription:     base.GoodsDescription,
			CustomsArticleNumber: base.CustomsArticleNumber,
			ItemNetWeightInKg:    a.ItemNetWeightInKg + base.ItemNetWeightInKg,
			TarriffLineAmount:    a.TarriffLineAmount + base.TarriffLineAmount,
			Currency:             base.Currency,
			CountryOfOrigin:      base.CountryOfOrigin,
		}
	}

	return base
}

func toBringCustomsLine(ctx context.Context, ol *ent.OrderLine, dest *ent.Country) (*bringrequest.EDICustomsDeclaration, error) {
	currentLineWeightKG, err := common.ColliWeightKG(ctx, []*ent.OrderLine{ol})
	if err != nil {
		return nil, err
	}

	v, err := ol.QueryProductVariant().
		WithInventoryItem().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if v.Edges.InventoryItem == nil {
		return nil, fmt.Errorf("expected product to have inventory item for calculating customs")
	}

	tariffCode := ""
	if v.Edges.InventoryItem.Code != nil {
		tariffCode = *v.Edges.InventoryItem.Code
	}

	// Follows same structure as Shopify
	// https://shopify.dev/docs/api/admin-rest/2024-01/resources/inventoryitem
	countrySpecificCode, err := v.Edges.InventoryItem.QueryCountryHarmonizedCode().
		Where(countryharmonizedcode.HasCountryWith(country.ID(dest.ID))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if err == nil {
		tariffCode = countrySpecificCode.Code
	}

	cur, err := ol.QueryCurrency().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	coo, err := v.Edges.InventoryItem.QueryCountryOfOrigin().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &bringrequest.EDICustomsDeclaration{
		Quantity:             ol.Units,
		GoodsDescription:     v.Description,
		CustomsArticleNumber: tariffCode,
		ItemNetWeightInKg:    currentLineWeightKG,
		TarriffLineAmount:    common.NetOrderLinePrice(ol),
		Currency:             cur.CurrencyCode,
		CountryOfOrigin:      coo.Alpha2,
	}, nil
}

func newBaseBooking(ctx context.Context, shipmentConfig *ShipmentConfig) (*bringrequest.Request, error) {
	consignment, err := newConsignment(ctx, shipmentConfig)
	if err != nil {
		return nil, err
	}

	return &bringrequest.Request{
		Consignments: []bringrequest.Consignment{
			*consignment,
		},
		SchemaVersion: 1,
		TestIndicator: shipmentConfig.Test,
	}, nil
}

func addAuthentication(req *http.Request, auth bringrequest.AuthenticationHeaders) *http.Request {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("X-Mybring-API-Key", auth.APIKey)
	req.Header.Set("X-Mybring-API-Uid", auth.APIUID)
	req.Header.Set("X-Bring-Client-URL", auth.ClientURL)
	return req
}
