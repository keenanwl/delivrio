package mergeutils

import (
	"delivrio.io/go/schema/fieldjson"
	"html/template"
	"time"
)

type BasicDocVariables struct {
	DocCreatedDate string
	DocCreatedTime string
}

type OrdersList struct {
	BasicDocVariables
	CarrierName   string
	RangeFromDate string
	RangeToDate   string
	ShipmentCount string
	Orders        []OrderRow
}

type OrderRow struct {
	RowCount            string
	OrderID             string
	ShipmentCreatedDate string
	ShipmentID          string
	ShipmentTrackingID  string
	// For sorting
	created time.Time
}

type PackingSlip struct {
	BasicDocVariables
	CustomerAddress
	SenderAddress
	DropPointAddress
	DeliveryOptionName    string
	DeliveryOptionCarrier string
	OrderPublicID         string
	OrderCommentExternal  string
	OrderCommentInternal  string
	OrderNoteAttributes   fieldjson.NoteAttributes
	OrderLines            []OrderLine
	DELIVRIOBarcodeImgTag template.HTML
	DELIVRIOBarcodeImgSrc string
	DELIVRIOBarcode       string
}

func testOrderRow() []OrderRow {
	return []OrderRow{
		{
			RowCount:            "1",
			OrderID:             "#222222222",
			ShipmentCreatedDate: "2022-01-01",
			ShipmentID:          "09:22",
			ShipmentTrackingID:  "00000000011111111111111",
			created:             time.Now(),
		},
		{
			RowCount:            "2",
			OrderID:             "#333333333",
			ShipmentCreatedDate: "2022-01-02",
			ShipmentID:          "11:22",
			ShipmentTrackingID:  "0000000001111111111222",
			created:             time.Now(),
		},
	}
}
