package postnordresponse

type ServicePointsResponse struct {
	ServicePointInformationResponse *ServicePointInformationResponse `json:"servicePointInformationResponse,omitempty"`
}

type ServicePointInformationResponse struct {
	CustomerSupports []interface{}  `json:"customerSupports,omitempty"`
	ServicePoints    []ServicePoint `json:"servicePoints,omitempty"`
}

type ServicePoint struct {
	Name                  string            `json:"name,omitempty"`
	ServicePointID        string            `json:"servicePointId,omitempty"`
	PudoID                string            `json:"pudoId,omitempty"`
	PhoneNoToCashRegister interface{}       `json:"phoneNoToCashRegister"`
	RoutingCode           interface{}       `json:"routingCode"`
	HandlingOffice        interface{}       `json:"handlingOffice"`
	LocationDetail        *string           `json:"locationDetail"`
	RouteDistance         interface{}       `json:"routeDistance"`
	Pickup                *Pickup           `json:"pickup,omitempty"`
	VisitingAddress       *Address          `json:"visitingAddress,omitempty"`
	DeliveryAddress       *Address          `json:"deliveryAddress,omitempty"`
	NotificationArea      *NotificationArea `json:"notificationArea,omitempty"`
	Coordinates           []Coordinate      `json:"coordinates,omitempty"`
	OpeningHours          *OpeningHours     `json:"openingHours,omitempty"`
	Type                  *Type             `json:"type,omitempty"`
}

type Coordinate struct {
	CountryCode *CountryCode `json:"countryCode,omitempty"`
	Northing    *float64     `json:"northing,omitempty"`
	Easting     *float64     `json:"easting,omitempty"`
	SrID        *SrID        `json:"srId,omitempty"`
}

type Address struct {
	CountryCode           CountryCode `json:"countryCode,omitempty"`
	City                  string      `json:"city,omitempty"`
	StreetName            string      `json:"streetName,omitempty"`
	StreetNumber          string      `json:"streetNumber,omitempty"`
	PostalCode            string      `json:"postalCode,omitempty"`
	AdditionalDescription interface{} `json:"additionalDescription"`
}

type NotificationArea struct {
	PostalCodes []string `json:"postalCodes,omitempty"`
}

type OpeningHours struct {
	PostalServices []PostalService `json:"postalServices,omitempty"`
	SpecialDates   []SpecialDate   `json:"specialDates,omitempty"`
}

type PostalService struct {
	CloseDay  Day    `json:"closeDay,omitempty"`
	CloseTime string `json:"closeTime,omitempty"`
	OpenDay   Day    `json:"openDay,omitempty"`
	OpenTime  string `json:"openTime,omitempty"`
}

type SpecialDate struct {
	CloseTime *SpecialDateCloseTime `json:"closeTime,omitempty"`
	EndDate   interface{}           `json:"endDate"`
	IsClosed  *bool                 `json:"isClosed,omitempty"`
	OpenTime  *SpecialDateOpenTime  `json:"openTime,omitempty"`
	Reason    *string               `json:"reason,omitempty"`
	StartDate interface{}           `json:"startDate"`
}

type Pickup struct {
	CashOnDelivery     interface{}   `json:"cashOnDelivery"`
	HeavyGoodsProducts []interface{} `json:"heavyGoodsProducts,omitempty"`
	Products           []Product     `json:"products,omitempty"`
}

type Product struct {
	Name      *string    `json:"name,omitempty"`
	TimeSlots *TimeSlots `json:"timeSlots,omitempty"`
}

type TimeSlots struct {
	AvailableForPickupEarlyCollect []interface{} `json:"availableForPickupEarlyCollect,omitempty"`
	AvailableForPickupStandard     []interface{} `json:"availableForPickupStandard,omitempty"`
}

type Type struct {
	GroupTypeID   int64         `json:"groupTypeId,omitempty"`
	GroupTypeName GroupTypeName `json:"groupTypeName,omitempty"`
	TypeID        int64         `json:"typeId,omitempty"`
	TypeName      TypeName      `json:"typeName,omitempty"`
	BoxType       interface{}   `json:"boxType"`
}

type CountryCode string

const (
	Dk CountryCode = "DK"
)

type SrID string

const (
	Epsg4326 SrID = "EPSG:4326"
)

type City string

const (
	AarhusC City = "AARHUS C"
)

type Day string

func (d Day) String() string {
	return string(d)
}

const (
	Friday    Day = "Friday"
	Monday    Day = "Monday"
	Saturday  Day = "Saturday"
	Sunday    Day = "Sunday"
	Thursday  Day = "Thursday"
	Tuesday   Day = "Tuesday"
	Wednesday Day = "Wednesday"
)

type SpecialDateCloseTime string

const (
	Empty SpecialDateCloseTime = " "
)

type SpecialDateOpenTime string

const (
	TjekForÅbningstidPåDetEnkeltePosthus SpecialDateOpenTime = "Tjek for åbningstid på det enkelte posthus"
)

type GroupTypeName string

const (
	ServiceAgent GroupTypeName = "Service agent"
	ServiceBox   GroupTypeName = "Service box"
)

type TypeName string

const (
	ParcelBoxLocation              TypeName = "Parcel Box Location"
	PostbutikMedBegrænsetSortiment TypeName = "Postbutik med begrænset sortiment"
	Posthus                        TypeName = "Posthus"
)
