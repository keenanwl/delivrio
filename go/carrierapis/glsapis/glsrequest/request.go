package glsrequest

type Shipment struct {
	UserName     string    `json:"UserName"`
	Password     string    `json:"Password"`
	Customerid   string    `json:"Customerid"`
	Contactid    string    `json:"Contactid"`
	ShipmentDate string    `json:"ShipmentDate"`
	Reference    string    `json:"Reference"`
	Addresses    Addresses `json:"Addresses"`
	Parcels      []Parcel  `json:"Parcels"`
	Services     Services  `json:"Services"`
}

type Addresses struct {
	Delivery           Address  `json:"Delivery"`
	AlternativeShipper *Address `json:"AlternativeShipper,omitempty"`
	Pickup             *Address `json:"Pickup,omitempty"`
}

type Address struct {
	Name1      string `json:"Name1"`
	Name2      string `json:"Name2"`
	Name3      string `json:"Name3"`
	Street1    string `json:"Street1"`
	CountryNum string `json:"CountryNum"`
	ZipCode    string `json:"ZipCode"`
	City       string `json:"City"`
	Contact    string `json:"Contact"`
	Email      string `json:"Email"`
	Phone      string `json:"Phone"`
	Mobile     string `json:"Mobile"`
}

type Parcel struct {
	Weight                float64 `json:"Weight"`
	Reference             string  `json:"Reference,omitempty"`
	Comment               string  `json:"Comment,omitempty"`
	AddOnLiabilityService int64   `json:"AddOnLiabilityService,omitempty"`
}

type Services struct {
	PrivateDelivery       *string `json:"PrivateDelivery,omitempty"`
	FlexDelivery          *string `json:"FlexDelivery,omitempty"`
	NotificationEmail     *string `json:"NotificationEmail,omitempty"`
	ShopReturn            *string `json:"ShopReturn,omitempty"`
	DirectShop            *string `json:"DirectShop,omitempty"`
	ShopDelivery          *string `json:"ShopDelivery,omitempty"`
	Express12             *string `json:"Express12,omitempty"`
	Express10             *string `json:"Express10,omitempty"`
	Deposit               *string `json:"Deposit,omitempty"`
	AddOnLiabilityService *int    `json:"AddOnLiabilityService,omitempty"`
}
