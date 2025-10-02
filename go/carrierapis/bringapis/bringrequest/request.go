package bringrequest

import "delivrio.io/go/ent/currency"

type AuthenticationHeaders struct {
	APIKey    string // X-Mybring-API-Key
	APIUID    string // X-Mybring-API-Uid: Seems like it is just the account email?
	ClientURL string // X-Bring-Client-URL: Undocumented? But API complains if it is not there
}

type Request struct {
	Consignments  []Consignment `json:"consignments"`
	SchemaVersion int64         `json:"schemaVersion"`
	TestIndicator bool          `json:"testIndicator"`
}

type Consignment struct {
	CorrelationID    string    `json:"correlationId"`
	Packages         []Package `json:"packages"`
	Parties          Parties   `json:"parties"`
	Product          Product   `json:"product"`
	ShippingDateTime string    `json:"shippingDateTime"`
}

type Package struct {
	CorrelationID    string      `json:"correlationId"`
	Dimensions       Dimensions  `json:"dimensions"`
	GoodsDescription string      `json:"goodsDescription"`
	PackageType      interface{} `json:"packageType"`
	WeightInKg       float64     `json:"weightInKg"`
}

type Dimensions struct {
	HeightInCM int64 `json:"heightInCm"`
	LengthInCM int64 `json:"lengthInCm"`
	WidthInCM  int64 `json:"widthInCm"`
}

type Parties struct {
	Recipient   Recipient    `json:"recipient"`
	Sender      Sender       `json:"sender"`
	PickupPoint *PickupPoint `json:"pickupPoint,omitempty"`
}

type PickupPoint struct {
	CountryCode string `json:"countryCode"`
	BringID     string `json:"id"`
}

type Recipient struct {
	AdditionalAddressInfo string           `json:"additionalAddressInfo"`
	AddressLine           string           `json:"addressLine"`
	AddressLine2          string           `json:"addressLine2"`
	City                  string           `json:"city"`
	Contact               RecipientContact `json:"contact"`
	CountryCode           string           `json:"countryCode"`
	Name                  string           `json:"name"`
	PostalCode            string           `json:"postalCode"`
	Reference             string           `json:"reference"`
}

type RecipientContact struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type Sender struct {
	AdditionalAddressInfo string        `json:"additionalAddressInfo"`
	AddressLine           string        `json:"addressLine"`
	AddressLine2          string        `json:"addressLine2"`
	City                  string        `json:"city"`
	Contact               SenderContact `json:"contact"`
	CountryCode           string        `json:"countryCode"`
	Name                  string        `json:"name"`
	PostalCode            string        `json:"postalCode"`
	Reference             string        `json:"reference"`
}

type SenderContact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type Product struct {
	AdditionalServices     []AdditionalService     `json:"additionalServices"`
	CustomerNumber         string                  `json:"customerNumber"`
	EDICustomsDeclarations *EDICustomsDeclarations `json:"ediCustomsDeclarations,omitempty"`
	ID                     string                  `json:"id"`
}

type EDICustomsDeclarations struct {
	EDICustomsDeclaration []EDICustomsDeclaration `json:"ediCustomsDeclaration"`
	NatureOfTransaction   string                  `json:"natureOfTransaction"`
}

type EDICustomsDeclaration struct {
	Quantity             int                   `json:"quantity,omitempty"`
	GoodsDescription     string                `json:"goodsDescription,omitempty"`
	CustomsArticleNumber string                `json:"customsArticleNumber,omitempty"`
	ItemNetWeightInKg    float64               `json:"itemNetWeightInKg,omitempty"`
	TarriffLineAmount    float64               `json:"tarriffLineAmount,omitempty"` // Unknown if supports floats?
	Currency             currency.CurrencyCode `json:"currency,omitempty"`
	CountryOfOrigin      string                `json:"countryOfOrigin,omitempty"`
}

type AdditionalService struct {
	Email  string `json:"email"`
	ID     string `json:"id"`
	Mobile string `json:"mobile"`
}
