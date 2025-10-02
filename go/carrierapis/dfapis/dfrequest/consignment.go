package dfrequest

import "delivrio.io/go/ent/carrierdf"

type APICredentials struct {
	ClientID  string `yaml:"client_id"`
	GrantType string `yaml:"grant_type"` // Always password?
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Resource  string `yaml:"resource"`
}

type APIAccessConfig struct {
	// Payload requests
	CustomerID      string
	AccessToken     string
	AgreementNumber string

	// Oauth2
	APICredentials
}

type ShippingType string

const (
	ShippingTypeUnknown                ShippingType = "Unknown"
	ShippingTypeStykgods               ShippingType = "Stykgods"
	ShippingTypeErhversbilpakker       ShippingType = "Erhversbilpakker"
	ShippingTypeBilpakke               ShippingType = "Bilpakke"
	ShippingTypePalleEnhedsForsendelse ShippingType = "PalleEnhedsForsendelse"
	ErhvervspakkeVolumen               ShippingType = "ErhvervspakkeVolumen"
)

func (p ShippingType) String() string {
	return string(p)
}

type WhoPays string

const (
	WhoPaysUnknown WhoPays = "Unknown"
	WhoPaysPrepaid WhoPays = "Prepaid"
	WhoPaysCollect WhoPays = "Collect"
)

func (p WhoPays) String() string {
	return string(p)
}

type ConsignmentNoteType string

const (
	ConsignmentNoteTypeUnknown ConsignmentNoteType = "Unknown"
	ConsignmentNoteTypePickup  ConsignmentNoteType = "Pickup"
	ConsignmentNoteTypeReturn  ConsignmentNoteType = "Return"
	// Missing relocation? Or maybe that is what
	// change of address means? Relocation?
	ConsignmentNoteTypeChangeOfAddress ConsignmentNoteType = "ChangeOfAddress"
)

func (p ConsignmentNoteType) String() string {
	return string(p)
}

type Consignment struct {
	ConsignmentNumber    *string               `json:"ConsignmentNumber,omitempty"`
	CustomID             *int64                `json:"CustomId,omitempty"`
	ConsignmentDate      *string               `json:"ConsignmentDate,omitempty"`
	ShippingType         ShippingType          `json:"ShippingType,omitempty"`
	ConsignmentNoteType  ConsignmentNoteType   `json:"ConsignmentNoteType,omitempty"`
	AgreementNumber      string                `json:"AgreementNumber,omitempty"`
	HubAgreement         string                `json:"HubAgreement,omitempty"`
	WhoPays              carrierdf.WhoPays     `json:"WhoPays,omitempty"`
	WhoPaysOriginal      *string               `json:"WhoPaysOriginal,omitempty"`
	Sender               Address               `json:"Sender,omitempty"`
	Receiver             Address               `json:"Receiver,omitempty"`
	Initiator            Address               `json:"Initiator,omitempty"`
	Pickup               *Address              `json:"Pickup,omitempty"`
	DeliveryNotification *DeliveryNotification `json:"DeliveryNotification,omitempty"`
	Goods                []Good                `json:"Goods,omitempty"`
	PickupTime           *PickupTime           `json:"PickupTime,omitempty"`
	DeliveryTime         *DeliveryTime         `json:"DeliveryTime,omitempty"`
	PreBooking           *bool                 `json:"PreBooking,omitempty"`
	FrankaturCode        *string               `json:"FrankaturCode,omitempty"`
	LimitedQuantityGoods *LimitedQuantityGoods `json:"LimitedQuantityGoods,omitempty"`
	Locked               *bool                 `json:"Locked,omitempty"`
	LockedTime           *string               `json:"LockedTime,omitempty"`
	LockedReason         *string               `json:"LockedReason,omitempty"`
	ProductCode          *string               `json:"ProductCode,omitempty"`
	ShopID               *string               `json:"ShopId,omitempty"`
	ExchangePallets      *ExchangePallets      `json:"ExchangePallets,omitempty"`
	Insurance            *Insurance            `json:"Insurance,omitempty"`
	ForceColliScanning   []string              `json:"ForceColliScanning,omitempty"`
	ServiceCodes         []string              `json:"ServiceCodes,omitempty"`
	SenderReference      *string               `json:"SenderReference,omitempty"`
	Reference1           *string               `json:"Reference1,omitempty"`
	Reference2           *string               `json:"Reference2,omitempty"`
	Reference3           *string               `json:"Reference3,omitempty"`
	Reference4           *string               `json:"Reference4,omitempty"`
	Reference5           *string               `json:"Reference5,omitempty"`
	Department           *string               `json:"Department,omitempty"`
	DeliveryRemark       *string               `json:"DeliveryRemark,omitempty"`
	DeliveryRemark2      *string               `json:"DeliveryRemark2,omitempty"`
	DeliveryRemark3      *string               `json:"DeliveryRemark3,omitempty"`
	DeliveryRemark4      *string               `json:"DeliveryRemark4,omitempty"`
	DeliveryRemark5      *string               `json:"DeliveryRemark5,omitempty"`
	PickupRemarks        *string               `json:"PickupRemarks,omitempty"`
	ID                   *int64                `json:"Id,omitempty"`
	Route                *Route                `json:"Route,omitempty"`
	Partner              *string               `json:"Partner,omitempty"`
	LastUpdated          *string               `json:"LastUpdated,omitempty"`
	Created              *string               `json:"Created,omitempty"`
	SourceSystem         *string               `json:"SourceSystem,omitempty"`
	Source               *string               `json:"Source,omitempty"`
	OriginalSource       *string               `json:"OriginalSource,omitempty"`
	HubBilling           *string               `json:"HubBilling,omitempty"`
	BillingInfo          *BillingInfo          `json:"BillingInfo,omitempty"`
	Size                 *Size                 `json:"Size,omitempty"`
	Deleted              *bool                 `json:"Deleted,omitempty"`
	TransportType        *string               `json:"TransportType,omitempty"`
	ConsignmentGUID      *string               `json:"ConsignmentGuid,omitempty"`
}

type BillingInfo struct {
	AlternativeCustomerNumber *string      `json:"AlternativeCustomerNumber,omitempty"`
	Price                     *int64       `json:"Price,omitempty"`
	NetPrice                  *int64       `json:"NetPrice,omitempty"`
	NetPriceCurrency          *string      `json:"NetPriceCurrency,omitempty"`
	GrossPrice                *int64       `json:"GrossPrice,omitempty"`
	GrossPriceCurrency        *string      `json:"GrossPriceCurrency,omitempty"`
	InvoiceNumber             *string      `json:"InvoiceNumber,omitempty"`
	InvoiceDate               *string      `json:"InvoiceDate,omitempty"`
	Fees                      []Fee        `json:"Fees,omitempty"`
	EnergyFee                 *int64       `json:"EnergyFee,omitempty"`
	PriceSplits               []PriceSplit `json:"PriceSplits,omitempty"`
	BilledConsignmentNumber   *string      `json:"BilledConsignmentNumber,omitempty"`
	NetPriceOriginal          *int64       `json:"NetPriceOriginal,omitempty"`
	FeesOrignal               []Fee        `json:"FeesOrignal,omitempty"`
	EnergyFeeOriginal         *int64       `json:"EnergyFeeOriginal,omitempty"`
	InvoiceNumberOriginal     *string      `json:"InvoiceNumberOriginal,omitempty"`
}

type Fee struct {
	ID                  *int64  `json:"Id,omitempty"`
	Fid                 *int64  `json:"Fid,omitempty"`
	FeeType             *string `json:"FeeType,omitempty"`
	Code                *string `json:"Code,omitempty"`
	Amount              *int64  `json:"Amount,omitempty"`
	AmountCurrency      *string `json:"AmountCurrency,omitempty"`
	CargoHub1Amount     *int64  `json:"CargoHub1Amount,omitempty"`
	CargoHub2Amount     *int64  `json:"CargoHub2Amount,omitempty"`
	CarrierRoute1Amount *int64  `json:"CarrierRoute1Amount,omitempty"`
	CarrierRoute2Amount *int64  `json:"CarrierRoute2Amount,omitempty"`
	CarrierRoute3Amount *int64  `json:"CarrierRoute3Amount,omitempty"`
	Description         *string `json:"Description,omitempty"`
}

type PriceSplit struct {
	Carrier         *string `json:"Carrier,omitempty"`
	Fee             *int64  `json:"Fee,omitempty"`
	RouteDeduction  *int64  `json:"RouteDeduction,omitempty"`
	Amount          *int64  `json:"Amount,omitempty"`
	GrossAmount     *int64  `json:"GrossAmount,omitempty"`
	BonusOffsetting *int64  `json:"BonusOffsetting,omitempty"`
	Budget          *string `json:"Budget,omitempty"`
	Route           *string `json:"Route,omitempty"`
}

type DeliveryNotification struct {
	Email []string `json:"Email,omitempty"`
	SMS   []string `json:"SMS,omitempty"`
}

type DeliveryTime struct {
	DotIntervalStart      *string `json:"DotIntervalStart,omitempty"`
	DotIntervalEnd        *string `json:"DotIntervalEnd,omitempty"`
	DotType               *string `json:"DotType,omitempty"`
	ExpectedDeliveryDate  *string `json:"ExpectedDeliveryDate,omitempty"`
	ActualDeliveryDate    *string `json:"ActualDeliveryDate,omitempty"`
	RequestedDeliveryDate *string `json:"RequestedDeliveryDate,omitempty"`
}

type ExchangePallets struct {
	FullPallets   *int64 `json:"FullPallets,omitempty"`
	HalfPallets   *int64 `json:"HalfPallets,omitempty"`
	QuaterPallets *int64 `json:"QuaterPallets,omitempty"`
}

type Good struct {
	DID                 *int64          `json:"DId,omitempty"`
	NumberOfItems       int64           `json:"NumberOfItems,omitempty"`
	Type                string          `json:"Type,omitempty"`
	Description         string          `json:"Description,omitempty"`
	Weight              float64         `json:"Weight,omitempty"`
	Volume              *int64          `json:"Volume,omitempty"`
	Amount              *int64          `json:"Amount,omitempty"`
	Length              *int            `json:"Length,omitempty"`
	Width               *int            `json:"Width,omitempty"`
	Height              *int            `json:"Height,omitempty"`
	Stackable           *bool           `json:"Stackable,omitempty"`
	LoadMeters          *int64          `json:"LoadMeters,omitempty"`
	SenderRef           *string         `json:"SenderRef,omitempty"`
	DangerousGoods      []DangerousGood `json:"DangerousGoods,omitempty"`
	Products            []Product       `json:"Products,omitempty"`
	ColliCodes          []ColliCode     `json:"ColliCodes,omitempty"`
	FragtpligtigVaegt   *int64          `json:"FragtpligtigVaegt,omitempty"`
	FragtpligtigRumfang *int64          `json:"FragtpligtigRumfang,omitempty"`
}

type ColliCode struct {
	ID          *int64  `json:"Id,omitempty"`
	FID         *int64  `json:"FId,omitempty"`
	GID         *int64  `json:"GId,omitempty"`
	Barcode     *string `json:"Barcode,omitempty"`
	IsPrinted   *bool   `json:"IsPrinted,omitempty"`
	ColliNumber *int64  `json:"ColliNumber,omitempty"`
}

type DangerousGood struct {
	ID           *int64  `json:"Id,omitempty"`
	HazardCode   *string `json:"HazardCode,omitempty"`
	UNDGnumber   *string `json:"UNDGnumber,omitempty"`
	Weight       *int64  `json:"Weight,omitempty"`
	Count        *int64  `json:"Count,omitempty"`
	PackingGroup *string `json:"PackingGroup,omitempty"`
	Packaging    *string `json:"Packaging,omitempty"`
	Unit         *string `json:"Unit,omitempty"`
	FgPoint      *int64  `json:"FgPoint,omitempty"`
	Content      *string `json:"Content,omitempty"`
	Content2     *string `json:"Content2,omitempty"`
	Content3     *string `json:"Content3,omitempty"`
	Content4     *string `json:"Content4,omitempty"`
	Content5     *string `json:"Content5,omitempty"`
	Lq           *bool   `json:"LQ,omitempty"`
	Neq          *int64  `json:"NEQ,omitempty"`
}

type Product struct {
	VID     *int64  `json:"VId,omitempty"`
	Content *string `json:"Content,omitempty"`
	Count   *int64  `json:"Count,omitempty"`
	Weight  *int64  `json:"Weight,omitempty"`
	Lwh     *string `json:"Lwh,omitempty"`
	Extra   *string `json:"Extra,omitempty"`
}

type Address struct {
	Name                  string                 `json:"Name,omitempty"`
	Name2                 string                 `json:"Name2,omitempty"`
	Name3                 string                 `json:"Name3,omitempty"`
	Name4                 *string                `json:"Name4,omitempty"`
	Street                string                 `json:"Street,omitempty"`
	Street2               string                 `json:"Street2,omitempty"`
	Town                  string                 `json:"Town,omitempty"`
	Zipcode               string                 `json:"Zipcode,omitempty"`
	CustomerID            *string                `json:"CustomerId,omitempty"`
	Country               string                 `json:"Country,omitempty"`
	Phone                 string                 `json:"Phone,omitempty"`
	Email                 string                 `json:"Email,omitempty"`
	ContactPerson         string                 `json:"ContactPerson,omitempty"`
	ContactPersonPhone    string                 `json:"ContactPersonPhone,omitempty"`
	ContactPersonEmail    string                 `json:"ContactPersonEmail,omitempty"`
	Position              *Position              `json:"Position,omitempty"`
	PositionForGPSBearing *PositionForGPSBearing `json:"PositionForGPSBearing,omitempty"`
}

type Position struct {
	Latitude  *int64 `json:"Latitude,omitempty"`
	Longitude *int64 `json:"Longitude,omitempty"`
}

type PositionForGPSBearing struct {
	Latitude  *string `json:"Latitude,omitempty"`
	Longitude *string `json:"Longitude,omitempty"`
}

type Insurance struct {
	Amount         *int64  `json:"Amount,omitempty"`
	AmountCurrency *string `json:"AmountCurrency,omitempty"`
	Type           *string `json:"Type,omitempty"`
	Premium        *int64  `json:"Premium,omitempty"`
}

type LimitedQuantityGoods struct {
	Weight          *int64           `json:"Weight,omitempty"`
	Identifications []Identification `json:"Identifications,omitempty"`
}

type Identification struct {
	UNNumber    *string `json:"UNNumber,omitempty"`
	HazardClass *string `json:"HazardClass,omitempty"`
	Weight      *int64  `json:"Weight,omitempty"`
}

type PickupTime struct {
	PickupIntervalStart *string `json:"PickupIntervalStart,omitempty"`
	PickupIntervalEnd   *string `json:"PickupIntervalEnd,omitempty"`
	ActualPickupDate    *string `json:"ActualPickupDate,omitempty"`
}

type Route struct {
	HubIncomming     *string `json:"HubIncomming,omitempty"`
	HubOutgoing      *string `json:"HubOutgoing,omitempty"`
	CarrierIncomming *int64  `json:"CarrierIncomming,omitempty"`
	CarrierTransit   *int64  `json:"CarrierTransit,omitempty"`
	CarrierOutgoing  *int64  `json:"CarrierOutgoing,omitempty"`
	SortingCode      *string `json:"SortingCode,omitempty"`
}

type Size struct {
	TotalWeight        *int64  `json:"TotalWeight,omitempty"`
	TotalLoadingMeters *int64  `json:"TotalLoadingMeters,omitempty"`
	TotalVolumn        *int64  `json:"TotalVolumn,omitempty"`
	NumberOfItems      *string `json:"NumberOfItems,omitempty"`
}
