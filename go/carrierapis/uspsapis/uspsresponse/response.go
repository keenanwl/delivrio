package uspsresponse

type Label struct {
	LabelAddress       LabelAddress   `json:"labelAddress"`
	RoutingInformation string         `json:"routingInformation"`
	TrackingNumber     string         `json:"trackingNumber"`
	Postage            float64        `json:"postage"`
	ExtraServices      []ExtraService `json:"extraServices"`
	Zone               string         `json:"zone"`
	Commitment         Commitment     `json:"commitment"`
	WeightUOM          string         `json:"weightUOM"`
	Weight             float64        `json:"weight"`
	Fees               []interface{}  `json:"fees"`
	Sku                string         `json:"SKU"`
}

type Commitment struct {
	Name                 string `json:"name"`
	ScheduleDeliveryDate string `json:"scheduleDeliveryDate"`
}

type ExtraService struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Sku   string  `json:"SKU"`
}

type LabelAddress struct {
	StreetAddress    string `json:"streetAddress"`
	SecondaryAddress string `json:"secondaryAddress"`
	City             string `json:"city"`
	State            string `json:"state"`
	Urbanization     string `json:"urbanization"`
	ZIPCode          string `json:"ZIPCode"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Firm             string `json:"firm"`
	IgnoreBadAddress bool   `json:"ignoreBadAddress"`
}
