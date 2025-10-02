package mergeutils

import (
	"context"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/orderline"
	"delivrio.io/go/ent/product"
	"delivrio.io/go/ent/productvariant"
	"time"
)

type CCAwaitingPickupNotice struct {
	CustomerFirstName string
}

type DropPointAddress struct {
	DropPointCompany  string
	DropPointAddress1 string
	DropPointAddress2 string
	DropPointZip      string
	DropPointCity     string
	DropPointState    string
	DropPointCountry  string
}

type CustomerAddress struct {
	CustomerFirstName string
	CustomerLastName  string
	CustomerCompany   string
	CustomerEmail     string
	CustomerAddress1  string
	CustomerAddress2  string
	CustomerZip       string
	CustomerCity      string
	CustomerState     string
	CustomerCountry   string
}

type SenderAddress struct {
	SenderFirstName string
	SenderLastName  string
	SenderCompany   string
	SenderEmail     string
	SenderAddress1  string
	SenderAddress2  string
	SenderZip       string
	SenderCity      string
	SenderState     string
	SenderCountry   string
}

type ColliPacked struct {
	CustomerAddress
	DropPointAddress
	SenderAddress

	OrderPublicID        string
	OrderLines           []OrderLine
	TrackingID           string
	CarrierName          string
	ExpectedDeliveryDate time.Time
}

func NewTestColliPacked(ctx context.Context) *ColliPacked {
	return &ColliPacked{
		CustomerAddress:      testCustomerAddress(),
		DropPointAddress:     testDropPointAddress(),
		SenderAddress:        testSenderAddress(),
		OrderPublicID:        "SH008899",
		OrderLines:           testOrderLines(ctx),
		TrackingID:           "0000009999999988888888877777777766666",
		ExpectedDeliveryDate: time.Now(),
		CarrierName:          "PostNord",
	}
}

func testCustomerAddress() CustomerAddress {
	return CustomerAddress{
		CustomerFirstName: "Sven",
		CustomerLastName:  "Svendsen",
		CustomerCompany:   "",
		CustomerEmail:     "sven@example.com",
		CustomerAddress1:  "888 Local St.",
		CustomerAddress2:  "APT 1",
		CustomerZip:       "90210",
		CustomerCity:      "Beverly Hills",
		CustomerState:     "CA",
		CustomerCountry:   "USA",
	}
}

func testSenderAddress() SenderAddress {
	return SenderAddress{
		SenderFirstName: "Test",
		SenderLastName:  "Sender",
		SenderCompany:   "A Company Inc",
		SenderEmail:     "test@example.com",
		SenderAddress1:  "999 Main st.",
		SenderAddress2:  "Apt 44",
		SenderZip:       "8000",
		SenderCity:      "Aarhus",
		SenderState:     "",
		SenderCountry:   "Denmark",
	}
}

func testDropPointAddress() DropPointAddress {
	return DropPointAddress{
		DropPointCompany:  "A Company Inc",
		DropPointAddress1: "111 Main st.",
		DropPointAddress2: "Apt 1",
		DropPointZip:      "8000",
		DropPointCity:     "Aarhus",
		DropPointState:    "",
		DropPointCountry:  "Denmark",
	}
}

func testOrderLines(ctx context.Context) []OrderLine {
	cli := ent.FromContext(ctx)
	orderLines, err := cli.OrderLine.Query().
		Where(orderline.HasProductVariantWith(productvariant.HasProductWith(product.HasProductImage()))).
		Limit(3).
		All(ctx)

	// Use test data if available
	if err == nil && len(orderLines) > 0 {
		out, _ := orderLinesToMerge(ctx, orderLines)
		return out
	}

	return []OrderLine{
		{
			Quantity: "5",
			Price:    "10",
			Total:    "50.00",
			OrderLineProductInfo: OrderLineProductInfo{
				ProductName:          "Peanuts",
				ProductVariantName:   "The really big bag",
				ProductFirstImageURL: "demo-img-url",
			},
		},
		{
			Quantity: "50",
			Price:    "100",
			Total:    "5,000.00",
			OrderLineProductInfo: OrderLineProductInfo{
				ProductName:          "Shoes",
				ProductVariantName:   "Size UK-15 Black",
				ProductFirstImageURL: "demo-img-url2",
			},
		},
	}

}

func NewTestOrderConfirmation(ctx context.Context) *OrderConfirmation {
	return &OrderConfirmation{
		CustomerAddress:  testCustomerAddress(),
		DropPointAddress: testDropPointAddress(),
		OrderPublicID:    "#10000010001000",
		OrderLines:       testOrderLines(ctx),
	}
}

type OrderConfirmation struct {
	CustomerAddress
	DropPointAddress

	OrderPublicID string
	OrderLines    []OrderLine
}

type ReturnBase struct {
	CustomerAddress
	ReturnAddress

	OrderPublicID    string
	ReturnMethodName string
	TrackingID       string

	OrderLines []ReturnOrderLine
}

type ReturnConfirmationLabel struct {
	ReturnBase

	LabelDownloadURL string
	// For viewing in browser
	LabelURL string
	// For displaying in emails
	LabelPNGURL string
}
type ReturnConfirmationQRCode struct {
	ReturnBase

	QRCodeDownloadURL string
	// For viewing in browser
	QRCodeURL string
}
type ReturnReceived struct {
	ReturnBase
}
type ReturnAccepted struct {
	ReturnBase
}

type ReturnAddress struct {
	ReturnFirstName string
	ReturnLastName  string
	ReturnCompany   string
	ReturnEmail     string
	ReturnAddress1  string
	ReturnAddress2  string
	ReturnZip       string
	ReturnCity      string
	ReturnState     string
	ReturnCountry   string
}

type ReturnOrderLine struct {
	OrderLine

	ReturnClaimName        string
	ReturnClaimDescription string
}

type OrderLine struct {
	Quantity string
	Price    string
	Total    string
	OrderLineProductInfo
}

type OrderLineProductInfo struct {
	ProductName          string
	ProductVariantName   string
	ProductFirstImageURL string
}

func NewTestReturnBase(ctx context.Context) ReturnBase {

	cli := ent.FromContext(ctx)

	productImg := "http://localhost:888/no-product-images-found-upload-for-better-test-data"

	arbitraryProduct, err := cli.Product.Query().
		Where(product.HasProductImage()).
		WithProductImage().
		First(ctx)
	if err == nil {
		productImg = arbitraryProduct.Edges.ProductImage[0].URL
	}

	return ReturnBase{
		CustomerAddress: CustomerAddress{
			CustomerFirstName: "John",
			CustomerLastName:  "Doe",
			CustomerCompany:   "ABC Corporation",
			CustomerEmail:     "johndoe@example.com",
			CustomerAddress1:  "123 Main Street",
			CustomerAddress2:  "Apt 4B",
			CustomerZip:       "12345",
			CustomerCity:      "New York",
			CustomerState:     "NY",
			CustomerCountry:   "United States",
		},
		ReturnAddress: ReturnAddress{
			ReturnFirstName: "Jane",
			ReturnLastName:  "Smith",
			ReturnCompany:   "XYZ Inc.",
			ReturnEmail:     "janesmith@example.com",
			ReturnAddress1:  "456 Elm Avenue",
			ReturnAddress2:  "Suite 7",
			ReturnCity:      "Los Angeles",
			ReturnState:     "CA",
			ReturnCountry:   "United States",
		},
		OrderPublicID:    "ORD123456",
		ReturnMethodName: "Standard Shipping",
		TrackingID:       "TRK78901234",
		OrderLines: []ReturnOrderLine{
			{

				ReturnClaimName: "Sizing: too small",
				// Smiley included to hopefully catch UTF8 issues
				ReturnClaimDescription: "Did not fit as expected: too small 😊",
				OrderLine: OrderLine{
					Quantity: "2",
					Price:    "100",
					Total:    "200",
					OrderLineProductInfo: OrderLineProductInfo{
						ProductName:          "Adidas sweatshirt",
						ProductVariantName:   "Grey",
						ProductFirstImageURL: productImg,
					},
				},
			},
		},
	}
}

func NewTestReturnConfirmationLabel(ctx context.Context) ReturnConfirmationLabel {
	downloadURL, _ := ReturnColliURL("https://localhost:4401", delivrioroutes.ReturnLabelDownload, "1234", "#0003")
	viewerURL, _ := ReturnColliURL("https://localhost:4401", delivrioroutes.ReturnLabel, "1234", "#0003")

	return ReturnConfirmationLabel{
		ReturnBase:       NewTestReturnBase(ctx),
		LabelDownloadURL: downloadURL.String(),
		LabelURL:         viewerURL.String(),
		LabelPNGURL:      downloadURL.String(),
	}
}

func NewTestReturnConfirmationQRCode(ctx context.Context) ReturnConfirmationQRCode {
	downloadURL, _ := ReturnColliURL("https://localhost:4401", delivrioroutes.ReturnLabelDownload, "1234", "#0003")
	viewerURL, _ := ReturnColliURL("https://localhost:4401", delivrioroutes.ReturnLabel, "1234", "#0003")

	return ReturnConfirmationQRCode{
		ReturnBase:        NewTestReturnBase(ctx),
		QRCodeDownloadURL: downloadURL.String(),
		QRCodeURL:         viewerURL.String(),
	}
}

func NewTestReturnReceived(ctx context.Context) ReturnReceived {
	return ReturnReceived{
		ReturnBase: NewTestReturnBase(ctx),
	}
}

func NewTestReturnAccepted(ctx context.Context) ReturnAccepted {
	return ReturnAccepted{
		ReturnBase: NewTestReturnBase(ctx),
	}
}

func NewTestFilteredShipmentsEmail() FilteredShipmentsEmail {
	return FilteredShipmentsEmail{
		ArrivingToday:    2,
		ArrivingTomorrow: 3,
		Orders: []FilteredShipmentsEmailRow{
			{
				PublicID:        "ABC123",
				BarcodeLabel:    "XYZ789",
				DeliveryOption:  "Express",
				ExpectedArrival: "2020-06-19",
				CustomerAddress: CustomerAddress{
					CustomerFirstName: "John",
					CustomerLastName:  "Doe",
					CustomerCompany:   "Acme Inc",
					CustomerEmail:     "john.doe@example.com",
					CustomerAddress1:  "123 Main Street",
					CustomerCity:      "New York",
					CustomerState:     "NY",
					CustomerZip:       "10001",
					CustomerCountry:   "USA",
				},
			},
		},
	}
}

type FilteredShipmentsEmail struct {
	ArrivingToday    int
	ArrivingTomorrow int
	Orders           []FilteredShipmentsEmailRow
}

type FilteredShipmentsEmailRow struct {
	PublicID        string
	BarcodeLabel    string
	ExpectedArrival string

	DeliveryOption string

	CustomerAddress
}
