package postnordresponse

type Label struct {
	BookingResponse BookingResponse `json:"bookingResponse"`
	LabelPrintout   []LabelPrintout `json:"labelPrintout"`
}

type BookingResponse struct {
	BookingID     string          `json:"bookingId"`
	IDInformation []IDInformation `json:"idInformation"`
}

type IDInformation struct {
	Status     string     `json:"status"`
	References References `json:"references"`
	IDS        []ID       `json:"ids"`
	Urls       []URL      `json:"urls"`
}

type ID struct {
	IDType string `json:"idType"`
	Value  string `json:"value"`
}

type References struct {
	Item     []interface{} `json:"item"`
	Shipment []Shipment    `json:"shipment"`
}

type Shipment struct {
	ReferenceType string `json:"referenceType"`
	ReferenceNo   string `json:"referenceNo"`
}

type URL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type LabelPrintout struct {
	ItemIDS             []ItemID         `json:"itemIds"`
	Printout            Printout         `json:"printout"`
	PrintoutComposition map[string]int64 `json:"printoutComposition"`
}

type ItemID struct {
	ItemIDS          string `json:"itemIds"`
	BasicServiceCode string `json:"basicServiceCode"`
	Status           string `json:"status"`
}

type Printout struct {
	LabelFormat string `json:"labelFormat"`
	Encoding    string `json:"encoding"`
	Data        string `json:"data"`
	Type        string `json:"type"`
}
