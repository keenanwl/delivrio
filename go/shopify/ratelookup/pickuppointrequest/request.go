package pickuppointrequest

type FetchInput struct {
	Typename    string       `json:"__typename,omitempty"`
	Allocations []Allocation `json:"allocations"`
	Cart        Cart         `json:"cart"`
}

type Allocation struct {
	Typename        string         `json:"__typename,omitempty"`
	DeliveryAddress MailingAddress `json:"deliveryAddress"`
}

type Cart struct {
	Typename string     `json:"__typename,omitempty"`
	Lines    []CartLine `json:"lines"`
}

type CartLine struct {
	Typename    string       `json:"__typename,omitempty"`
	Quantity    int          `json:"quantity"`
	Merchandise Merchandise  `json:"merchandise"`
	Cost        CartLineCost `json:"cost"`
}

type Merchandise struct {
	Typename string  `json:"__typename"`
	ID       string  `json:"id"`
	Product  Product `json:"product,omitempty"`
}

type Product struct {
	Typename string `json:"__typename,omitempty"`
	ID       string `json:"id,omitempty"`
}

type CartLineCost struct {
	Typename    string  `json:"__typename,omitempty"`
	TotalAmount MoneyV2 `json:"totalAmount"`
}

type MoneyV2 struct {
	Typename string `json:"__typename"`
	// Serialized decimal string...
	Amount string `json:"amount"`
}

type MailingAddress struct {
	Typename     string  `json:"__typename,omitempty"`
	CountryCode  string  `json:"countryCode,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Zip          string  `json:"zip,omitempty"`
	ProvinceCode string  `json:"provinceCode,omitempty"`
	City         string  `json:"city,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
}
