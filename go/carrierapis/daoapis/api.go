package daoapis

import (
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/carrierapis/daoapis/daorequest"
	"delivrio.io/go/carrierapis/daoapis/daoresponse"
	"delivrio.io/go/seed"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type LabelRecipient struct {
	PostalCode string
	Address    string
	Name       string
	Mobile     string
	Email      string
}
type LabelSender struct {
	SenderID         string
	SenderName       string
	SenderAddress    string
	SenderPostalCode string
	SenderMobile     string
	SenderEmail      string
}

type LabelPackaging struct {
	Length string // CM
	Height string
	Width  string
	Weight string // Grams
}

// Create order and get label are
// 2 different requests
type LabelPDFRequest struct {
	daorequest.Authentication
	Barcode   string
	PaperSize daorequest.LabelPaperSize
	Format    daorequest.LabelFormat // default: JSON, only relevant on response failure; otherwise PDF
}

type HomeLabelRequest struct {
	daorequest.Authentication
	LabelRecipient
	LabelSender
	LabelPackaging
	Test      string                 // blank or 1
	Format    daorequest.LabelFormat // default: JSON
	Reference string
}

type ShopLabelRequest struct {
	daorequest.Authentication
	LabelRecipient
	LabelSender
	LabelPackaging
	ShopID    string
	Test      string                 // blank or 1
	Format    daorequest.LabelFormat // default: JSON
	Reference string
}

type ReturnLabelRequestQueryFields struct {
	daorequest.Authentication
	LabelRecipient
	LabelSender
	LabelType daorequest.LabelType
	Test      string                 // blank or 1
	Format    daorequest.LabelFormat // default: JSON
	Reference string
}

// CreateOrderFetchLabel DAO requires two separate requests
// to first create the order
// and secondarily to receive a PDF of the label
func CreateOrderFetchLabel(ctx context.Context, pc *PackageConfig) (*FetchLabelOutput, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	auth := daorequest.Authentication{
		CustomerID: pc.Agreement.CustomerID,
		Code:       pc.Agreement.APIKey,
	}

	createOrderReq, err := NewOrderRequest(ctx, pc, auth, pc.Agreement.Test)
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

	var body daoresponse.ShipmentCreateResponse
	err = json.Unmarshal(data, &body)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(body.ErrorCodeParsed(), daoresponse.StatusOK.String()) {
		return nil, fmt.Errorf("dao error: %s", body.ErrorMessage)
	}

	labelByte, err := FetchLabel(ctx, httpClient, body.Result.Barcode, auth)
	if err != nil {
		return nil, err
	}

	return &FetchLabelOutput{
		Package:        pc,
		Barcode:        body.Result.Barcode,
		ResponseB64PDF: b64.StdEncoding.EncodeToString(labelByte),
		Error:          nil,
	}, nil

}

func FetchLabel(ctx context.Context, client *http.Client, barcode string, auth daorequest.Authentication) ([]byte, error) {
	formValues := &url.Values{}
	formValues = authFormData(formValues, auth)
	formValues.Add("stregkode", barcode)
	formValues.Add("format", daorequest.LabelFormatJSON.String())
	// Consider making configurable in the future
	formValues.Add("papir", daorequest.LabelPaperSize100x150.String())

	u, err := url.Parse("https://api.dao.as/HentLabel.php")
	if err != nil {
		return nil, err
	}

	u.RawQuery = formValues.Encode()

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, fmt.Errorf("dao: fetch label: %w", err)
	}

	reqDump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(reqDump))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("dao: fetch label: %w", err)
	}
	defer resp.Body.Close()

	respDump, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(respDump))

	// Response is always 200 even when status is failed
	// So we check for PDF
	contentType := resp.Header.Get("Content-Type")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("dao: read body: %w", err)
	}

	// Success (we hope)
	if strings.EqualFold(contentType, "application/pdf") {
		return body, nil
	}

	var responseErrData daoresponse.BaseResponse
	err = json.Unmarshal(body, &responseErrData)
	if err != nil {
		return nil, fmt.Errorf("dao: unmarshall: %w", err)
	}

	return nil, fmt.Errorf("dao: fetch label: expected PDF: %s", responseErrData.ErrorMessage)

}

func testVal(test bool) string {
	output := ""
	if test {
		output = "1"
	}
	return output
}

func NewOrderRequest(ctx context.Context, pc *PackageConfig, auth daorequest.Authentication, test bool) (*http.Request, error) {

	sender := LabelSender{
		SenderID:         "", // Stored (@DAO) sender profile?
		SenderName:       standardizeSpaces(fmt.Sprintf("%s %s %s", pc.ConsigneeAddress.FirstName, pc.ConsigneeAddress.LastName, pc.ConsigneeAddress.Company)),
		SenderAddress:    standardizeSpaces(fmt.Sprintf("%s %s", pc.ConsigneeAddress.AddressOne, pc.ConsigneeAddress.AddressTwo)),
		SenderPostalCode: pc.ConsigneeAddress.Zip,
		SenderMobile:     pc.ConsigneeAddress.PhoneNumber,
		SenderEmail:      pc.ConsigneeAddress.Email,
	}

	recipient := LabelRecipient{
		PostalCode: pc.ConsignorAddress.Zip,
		Address:    standardizeSpaces(fmt.Sprintf("%s %s", pc.ConsignorAddress.AddressOne, pc.ConsignorAddress.AddressTwo)),
		Name:       standardizeSpaces(fmt.Sprintf("%s %s %s", pc.ConsignorAddress.FirstName, pc.ConsignorAddress.LastName, pc.ConsignorAddress.Company)),
		Mobile:     pc.ConsignorAddress.PhoneNumber,
		Email:      pc.ConsignorAddress.Email,
	}

	weightGrams, err := common.ColliWeightGram(ctx, pc.OrderLines)
	if err != nil {
		return nil, fmt.Errorf("dao: label request: %w", err)
	}

	packagingData := LabelPackaging{
		Length: fmt.Sprintf("%v", pc.Packaging.LengthCm),
		Height: fmt.Sprintf("%v", pc.Packaging.HeightCm),
		Width:  fmt.Sprintf("%v", pc.Packaging.WidthCm),
		Weight: fmt.Sprintf("%v", weightGrams),
	}

	switch pc.CarrierService.InternalID {
	case seed.DAOServiceHOME.String():
		return homeRequest(ctx, HomeLabelRequest{
			Authentication: auth,
			LabelSender:    sender,
			LabelRecipient: recipient,
			LabelPackaging: packagingData,
			Test:           testVal(test),
			Format:         daorequest.LabelFormatJSON,
			Reference:      pc.OrderPublicID,
		})
	case seed.DAOServiceSHOP.String():
		return shopRequest(ctx, ShopLabelRequest{
			Authentication: auth,
			LabelSender:    sender,
			LabelRecipient: recipient,
			LabelPackaging: packagingData,
			Test:           testVal(test),
			Format:         daorequest.LabelFormatJSON,
			Reference:      pc.OrderPublicID,
			ShopID:         pc.ParcelShopID,
		})
	case seed.DAOServiceSHOPReturn.String():
		return returnRequest(ctx, ReturnLabelRequestQueryFields{
			Authentication: auth,
			LabelSender:    sender,
			LabelRecipient: recipient,
			LabelType:      daorequest.LabelTypeWithLabel,
			Test:           testVal(test),
			Format:         daorequest.LabelFormatJSON,
			Reference:      pc.OrderPublicID,
		})
	}

	return nil, fmt.Errorf("dao: label request: service not found")
}

func homeRequest(ctx context.Context, data HomeLabelRequest) (*http.Request, error) {
	query := homeFormData(data)

	u, err := url.Parse("https://api.dao.as/DAODirekte/leveringsordre.php")
	if err != nil {
		return nil, err
	}

	u.RawQuery = query.Encode()

	output, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dao: label request: %w", err)
	}

	output.PostForm = *query

	return output, nil
}

func shopRequest(ctx context.Context, data ShopLabelRequest) (*http.Request, error) {
	query := shopFormData(data)

	u, err := url.Parse("https://api.dao.as/DAOPakkeshop/leveringsordre.php")
	if err != nil {
		return nil, err
	}

	u.RawQuery = query.Encode()

	output, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dao: label request: %w", err)
	}

	output.PostForm = *query

	return output, nil
}

func returnRequest(ctx context.Context, data ReturnLabelRequestQueryFields) (*http.Request, error) {
	query := returnFormData(data)

	u, err := url.Parse("https://api.dao.as/DAOPakkeshop/leveringsordre.php")
	if err != nil {
		return nil, err
	}

	u.RawQuery = query.Encode()

	output, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dao: label request: %w", err)
	}

	output.PostForm = *query

	return output, nil
}

func homeFormData(data HomeLabelRequest) *url.Values {

	formValues := &url.Values{}
	formValues.Add("test", data.Test)
	formValues.Add("format", data.Format.String())
	formValues.Add("faktura", data.Reference)

	return packagingFormData(
		authFormData(
			recipientFormData(
				senderFormData(
					formValues,
					data.LabelSender,
				),
				data.LabelRecipient,
			),
			data.Authentication,
		),
		data.LabelPackaging,
	)
}

func shopFormData(data ShopLabelRequest) *url.Values {

	formValues := &url.Values{}
	formValues.Add("test", data.Test)
	formValues.Add("format", data.Format.String())
	formValues.Add("faktura", data.Reference)
	formValues.Add("shopid", data.ShopID)

	return packagingFormData(
		authFormData(
			recipientFormData(
				senderFormData(
					formValues,
					data.LabelSender,
				),
				data.LabelRecipient,
			),
			data.Authentication,
		),
		data.LabelPackaging,
	)
}

func returnFormData(data ReturnLabelRequestQueryFields) *url.Values {

	formValues := &url.Values{}

	formValues.Add("type", data.LabelType.String())
	formValues.Add("test", data.Test)
	formValues.Add("format", data.Format.String())
	formValues.Add("faktura", data.Reference)

	return authFormData(
		recipientFormData(
			senderFormData(
				formValues,
				data.LabelSender,
			),
			data.LabelRecipient,
		),
		data.Authentication,
	)
}

func authFormData(formValues *url.Values, data daorequest.Authentication) *url.Values {
	formValues.Add("kundeid", data.CustomerID)
	formValues.Add("kode", data.Code)
	return formValues
}

func senderFormData(formValues *url.Values, data LabelSender) *url.Values {
	formValues.Add("afsender", data.SenderID)
	formValues.Add("afsender_adresse", data.SenderAddress)
	formValues.Add("afsender_postnr", data.SenderPostalCode)
	formValues.Add("afs_mobil", data.SenderMobile)
	formValues.Add("afs_email", data.SenderEmail)
	return formValues
}

func recipientFormData(formValues *url.Values, data LabelRecipient) *url.Values {
	formValues.Add("postnr", data.PostalCode)
	formValues.Add("adresse", data.Address)
	formValues.Add("navn", data.Name)
	formValues.Add("mobil", data.Mobile)
	formValues.Add("email", data.Email)
	return formValues
}

func packagingFormData(formValues *url.Values, data LabelPackaging) *url.Values {
	formValues.Add("l", data.Length)
	formValues.Add("h", data.Height)
	formValues.Add("b", data.Width)
	formValues.Add("vaegt", data.Weight)
	return formValues
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
