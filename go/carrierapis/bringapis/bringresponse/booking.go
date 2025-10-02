package bringresponse

type ResponseBooking struct {
	Consignments []Consignment `json:"consignments"`
}

type Consignment struct {
	CorrelationID string       `json:"correlationId"`
	Confirmation  Confirmation `json:"confirmation"`
	Errors        []Error      `json:"errors"`
}

type Confirmation struct {
	ConsignmentNumber   string       `json:"consignmentNumber"`
	ProductSpecificData interface{}  `json:"productSpecificData"`
	Links               Links        `json:"links"`
	DateAndTimes        DateAndTimes `json:"dateAndTimes"`
	Packages            []Package    `json:"packages"`
}

type DateAndTimes struct {
	EarliestPickup   interface{} `json:"earliestPickup"`
	ExpectedDelivery interface{} `json:"expectedDelivery"`
}

type Links struct {
	Labels   string      `json:"labels"`
	Waybill  interface{} `json:"waybill"`
	Tracking string      `json:"tracking"`
}

type Package struct {
	PackageNumber string `json:"packageNumber"`
}

type Error struct {
	UniqueID                 string      `json:"uniqueId"`
	Code                     string      `json:"code"`
	Messages                 []Message   `json:"messages"`
	ConsignmentCorrelationID interface{} `json:"consignmentCorrelationId"`
	PackageCorrelationID     interface{} `json:"packageCorrelationId"`
}

type Message struct {
	Lang    string `json:"lang"`
	Message string `json:"message"`
}
