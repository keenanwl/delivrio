package postnordrequest

type Label struct {
	MessageDate     string      `json:"messageDate"`
	Application     Application `json:"application"`
	UpdateIndicator string      `json:"updateIndicator"`
	Shipment        []Shipment  `json:"shipment"`
}

type Application struct {
	ApplicationID int64  `json:"applicationId"`
	Name          string `json:"name"`
	Version       string `json:"version"`
}

type Shipment struct {
	ShipmentIdentification ShipmentIdentification `json:"shipmentIdentification"`
	DateAndTimes           DateAndTimes           `json:"dateAndTimes"`
	Service                Service                `json:"service"`
	// Assumed string
	FreeText         []string         `json:"freeText"`
	NumberOfPackages NumberOfPackages `json:"numberOfPackages"`
	TotalGrossWeight GrossWeight      `json:"totalGrossWeight"`
	Parties          Parties          `json:"parties"`
	GoodsItem        []*GoodsItem     `json:"goodsItem"`
}

type DateAndTimes struct {
	LoadingDate string `json:"loadingDate"`
}

type GoodsItem struct {
	PackageTypeCode string `json:"packageTypeCode"`
	Items           []Item `json:"items"`
}

type Item struct {
	ItemIdentification ItemIdentification `json:"itemIdentification"`
	GrossWeight        GrossWeight        `json:"grossWeight"`
}

type GrossWeight struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type ItemIdentification struct {
	ItemID     string `json:"itemId"`
	ItemIDType string `json:"itemIdType"`
}

type NumberOfPackages struct {
	Value int64 `json:"value"`
}

type Parties struct {
	Consignor     Consignor      `json:"consignor"`
	Consignee     Consignee      `json:"consignee"`
	DeliveryParty *DeliveryParty `json:"deliveryParty,omitempty"`
}

type Consignee struct {
	Party ConsigneeParty `json:"party"`
}

type ConsigneeParty struct {
	NameIdentification NameIdentification `json:"nameIdentification"`
	Address            Address            `json:"address"`
	Contact            Contact            `json:"contact"`
}

type Address struct {
	Streets     []string `json:"streets"`
	PostalCode  string   `json:"postalCode"`
	City        string   `json:"city"`
	CountryCode string   `json:"countryCode"`
}

type Contact struct {
	ContactName  *string `json:"contactName"`
	EmailAddress *string `json:"emailAddress"`
	SMSNo        *string `json:"smsNo"`
}

type NameIdentification struct {
	Name        *string `json:"name,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
}

type Consignor struct {
	IssuerCode          string              `json:"issuerCode"`
	PartyIdentification PartyIdentification `json:"partyIdentification"`
	Party               ConsignorParty      `json:"party"`
}

type ConsignorParty struct {
	NameIdentification NameIdentification `json:"nameIdentification"`
	Address            Address            `json:"address"`
}

type PartyIdentification struct {
	PartyID     string `json:"partyId"`
	PartyIDType string `json:"partyIdType"`
}

type DeliveryParty struct {
	PartyIdentification PartyIdentification `json:"partyIdentification"`
	Party               ConsignorParty      `json:"party"`
}

type Service struct {
	BasicServiceCode      string   `json:"basicServiceCode"`
	AdditionalServiceCode []string `json:"additionalServiceCode"`
}

type ShipmentIdentification struct {
	ShipmentID string `json:"shipmentId"`
}
