package uspsrequest

import (
	"delivrio.io/go/ent/carrierserviceusps"
	"delivrio.io/go/ent/packaginguspsprocessingcategory"
)

type Label struct {
	ImageInfo          ImageInfo          `json:"imageInfo"`
	ToAddress          Address            `json:"toAddress"`
	FromAddress        Address            `json:"fromAddress"`
	SenderAddress      *Address           `json:"senderAddress,omitempty"`
	ReturnAddress      *Address           `json:"returnAddress,omitempty"`
	PackageDescription PackageDescription `json:"packageDescription"`
	CustomsForm        *CustomsForm       `json:"customsForm,omitempty"`
}

type CustomsForm struct {
	ContentComments     string    `json:"contentComments"`
	RestrictionType     string    `json:"restrictionType"`
	RestrictionComments string    `json:"restrictionComments"`
	Aesitn              string    `json:"AESITN"`
	InvoiceNumber       string    `json:"invoiceNumber"`
	LicenseNumber       string    `json:"licenseNumber"`
	CertificateNumber   string    `json:"certificateNumber"`
	CustomsContentType  string    `json:"customsContentType"`
	ImportersReference  string    `json:"importersReference"`
	ImportersContact    string    `json:"importersContact"`
	ExportersReference  string    `json:"exportersReference"`
	ExportersContact    string    `json:"exportersContact"`
	Contents            []Content `json:"contents"`
}

type Content struct {
	ItemDescription string  `json:"itemDescription"`
	ItemQuantity    int64   `json:"itemQuantity"`
	ItemValue       int64   `json:"itemValue"`
	WeightUOM       string  `json:"weightUOM"`
	ItemWeight      float64 `json:"itemWeight"`
	HSTariffNumber  string  `json:"HSTariffNumber"`
	CountryofOrigin string  `json:"countryofOrigin"`
	ItemCategory    string  `json:"itemCategory"`
	ItemSubcategory string  `json:"itemSubcategory"`
}

type Address struct {
	StreetAddress        string  `json:"streetAddress"`
	SecondaryAddress     string  `json:"secondaryAddress"`
	City                 string  `json:"city"`
	State                string  `json:"state"`
	ZIPCode              string  `json:"ZIPCode"`
	ZIPPlus4             *string `json:"ZIPPlus4,omitempty"`
	Urbanization         string  `json:"urbanization"`
	FirstName            string  `json:"firstName,omitempty"`
	LastName             string  `json:"lastName,omitempty"`
	Firm                 string  `json:"firm,omitempty"`
	Phone                string  `json:"phone,omitempty"`
	IgnoreBadAddress     bool    `json:"ignoreBadAddress"`
	PlatformUserID       *string `json:"platformUserId,omitempty"`
	Email                *string `json:"email,omitempty"`
	ParcelLockerDelivery *bool   `json:"parcelLockerDelivery,omitempty"`
	FacilityID           *string `json:"facilityId,omitempty"`
}

type ImageInfo struct {
	ImageType        string `json:"imageType"`
	LabelType        string `json:"labelType"`
	ShipInfo         bool   `json:"shipInfo"`
	ReceiptOption    string `json:"receiptOption"`
	SuppressPostage  bool   `json:"suppressPostage"`
	SuppressMailDate bool   `json:"suppressMailDate"`
	ReturnLabel      bool   `json:"returnLabel"`
}

type PackageDescription struct {
	WeightUOM                       string                                             `json:"weightUOM"`
	Weight                          float64                                            `json:"weight"`
	DimensionsUOM                   string                                             `json:"dimensionsUOM"`
	Length                          int64                                              `json:"length"`
	Height                          int64                                              `json:"height"`
	Width                           int64                                              `json:"width"`
	Girth                           *int64                                             `json:"girth,omitempty"`
	DestinationEntryFacilityAddress *DestinationEntryFacilityAddress                   `json:"destinationEntryFacilityAddress,omitempty"`
	MailClass                       carrierserviceusps.APIKey                          `json:"mailClass"`
	RateIndicator                   string                                             `json:"rateIndicator"`
	ProcessingCategory              packaginguspsprocessingcategory.ProcessingCategory `json:"processingCategory"`
	DestinationEntryFacilityType    string                                             `json:"destinationEntryFacilityType"`
	PackageOptions                  *PackageOptions                                    `json:"packageOptions,omitempty"`
	CustomerReference               []CustomerReference                                `json:"customerReference"`
	ExtraServices                   []int64                                            `json:"extraServices"`
	MailingDate                     string                                             `json:"mailingDate"`
	CarrierRelease                  bool                                               `json:"carrierRelease"`
}

type CustomerReference struct {
	ReferenceNumber      string `json:"referenceNumber"`
	PrintReferenceNumber bool   `json:"printReferenceNumber"`
}

type DestinationEntryFacilityAddress struct {
	StreetAddress    string `json:"streetAddress"`
	SecondaryAddress string `json:"secondaryAddress"`
	City             string `json:"city"`
	State            string `json:"state"`
	ZIPCode          string `json:"ZIPCode"`
	ZIPPlus4         string `json:"ZIPPlus4"`
	Urbanization     string `json:"urbanization"`
}

type PackageOptions struct {
	PackageValue                 int64       `json:"packageValue"`
	NonDeliveryOption            string      `json:"nonDeliveryOption"`
	RedirectAddress              Address     `json:"redirectAddress"`
	ContentType                  string      `json:"contentType"`
	Containers                   []Container `json:"containers"`
	AncillaryServiceEndorsements string      `json:"ancillaryServiceEndorsements"`
}

type Container struct {
	ContainerID string `json:"containerID"`
	SortType    string `json:"sortType"`
}
