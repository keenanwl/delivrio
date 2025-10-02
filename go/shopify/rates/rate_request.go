package rates

type ShopifyRatesRequest struct {
	Rate ShopifyRateRequest `json:"rate"`
}

type ShopifyRateRequest struct {
	Origin      OriginDestinationAddressRequest `json:"origin"`
	Destination OriginDestinationAddressRequest `json:"destination"`
	Items       []ItemRequest                   `json:"items"`
	Currency    string                          `json:"currency"`
	Locale      string                          `json:"locale"`
}

type OriginDestinationAddressRequest struct {
	Country     *string `json:"country"`
	PostalCode  *string `json:"postal_code"`
	Province    *string `json:"province"`
	City        *string `json:"city"`
	Name        *string `json:"name"`
	Address1    *string `json:"address1"`
	Address2    *string `json:"address2"`
	Address3    *string `json:"address3"`
	Phone       *string `json:"phone"`
	Fax         *string `json:"fax"`
	Email       *string `json:"email"`
	AddressType *string `json:"address_type"`
	CompanyName *string `json:"company_name"`
}

type ItemRequest struct {
	Name               string      `json:"name"`
	Sku                string      `json:"sku"`
	Quantity           int         `json:"quantity"`
	Grams              int         `json:"grams"`
	Price              int         `json:"price"`
	Vendor             string      `json:"vendor"`
	RequiresShipping   bool        `json:"requires_shipping"`
	Taxable            bool        `json:"taxable"`
	FulfillmentService string      `json:"fulfillment_service"`
	Properties         interface{} `json:"properties"`
	ProductID          int64       `json:"product_id"`
	VariantID          int64       `json:"variant_id"`
}
