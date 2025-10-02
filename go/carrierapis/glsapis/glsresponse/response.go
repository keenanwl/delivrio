package glsresponse

type SuccessLabel struct {
	ConsignmentID string   `json:"ConsignmentId"`
	PDF           string   `json:"PDF"`
	Parcels       []Parcel `json:"Parcels"`
}

type Parcel struct {
	UniqueNumber string  `json:"UniqueNumber"`
	ParcelNumber string  `json:"ParcelNumber"`
	NdiNumber    string  `json:"NdiNumber"`
	Routing      Routing `json:"Routing"`
}

type Routing struct {
	Primary2D   string `json:"Primary2D"`
	Secondary2D string `json:"Secondary2D"`
	NationalRef string `json:"NationalRef"`
}

type ErrorLabel struct {
	Message    string     `json:"Message"`
	ModelState ModelState `json:"ModelState"`
}

type ModelState struct {
	ShipmentPassword []string `json:"shipment.Password"`
}
