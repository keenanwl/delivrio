package rates

import "delivrio.io/go/ent"

type ShopifyRateResponses struct {
	Rates []RateResponse `json:"rates"`
}

type RateResponse struct {
	ServiceName     string  `json:"service_name"`
	ServiceCode     string  `json:"service_code"`
	TotalPrice      string  `json:"total_price"`
	Description     string  `json:"description"`
	Currency        string  `json:"currency"`
	PhoneRequired   *string `json:"phone_required,omitempty"`
	MinDeliveryDate *string `json:"min_delivery_date,omitempty"`
	MaxDeliveryDate *string `json:"max_delivery_date,omitempty"`

	// Don't serialize
	TotalPriceDecimal float64
	ParcelShop        *ent.ParcelShop
	DeliveryOption    *ent.DeliveryOption
}
