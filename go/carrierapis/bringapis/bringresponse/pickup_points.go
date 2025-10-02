package bringresponse

type ResponsePickupPoints struct {
	PickupPoint []PickupPoint `json:"pickupPoint"`
}
type Welcome struct {
	PickupPoint []PickupPoint `json:"pickupPoint"`
}

type PickupPoint struct {
	ID                    string        `json:"id"`
	UnitID                string        `json:"unitId"`
	Name                  string        `json:"name"`
	Address               string        `json:"address"`
	PostalCode            string        `json:"postalCode"`
	City                  string        `json:"city"`
	CountryCode           CountryCode   `json:"countryCode"`
	VisitingAddress       string        `json:"visitingAddress"`
	VisitingPostalCode    string        `json:"visitingPostalCode"`
	VisitingCity          string        `json:"visitingCity"`
	OpeningHoursNorwegian string        `json:"openingHoursNorwegian"`
	OpeningHoursEnglish   string        `json:"openingHoursEnglish"`
	OpeningHoursFinnish   string        `json:"openingHoursFinnish"`
	OpeningHoursDanish    string        `json:"openingHoursDanish"`
	OpeningHoursSwedish   string        `json:"openingHoursSwedish"`
	Latitude              float64       `json:"latitude"`
	Longitude             float64       `json:"longitude"`
	UtmX                  string        `json:"utmX"`
	UtmY                  string        `json:"utmY"`
	GoogleMapsLink        string        `json:"googleMapsLink"`
	DistanceInKM          string        `json:"distanceInKm"`
	DistanceType          DistanceType  `json:"distanceType"`
	DurationInMinutes     int64         `json:"durationInMinutes"`
	DurationType          DurationType  `json:"durationType"`
	Type                  string        `json:"type"`
	AdditionalServiceCode string        `json:"additionalServiceCode"`
	RouteMapsLink         string        `json:"routeMapsLink"`
	Status                Status        `json:"status"`
	Capabilities          []interface{} `json:"capabilities"`
	OpeningHours          []OpeningHour `json:"openingHours"`
	TemporaryOpeningHours []interface{} `json:"temporaryOpeningHours"`
	SpecialOpeningHours   []interface{} `json:"specialOpeningHours"`
	Photos                []interface{} `json:"photos"`
	Coordinate            Coordinate    `json:"coordinate"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OpeningHour struct {
	Day     Day    `json:"day"`
	Opening string `json:"opening"`
	Closing string `json:"closing"`
}

type CountryCode string
type DistanceType string
type DurationType string
type Day string

const (
	Friday    Day = "FRIDAY"
	Monday    Day = "MONDAY"
	Saturday  Day = "SATURDAY"
	Sunday    Day = "SUNDAY"
	Thursday  Day = "THURSDAY"
	Tuesday   Day = "TUESDAY"
	Wednesday Day = "WEDNESDAY"
)

type Status string

const (
	Active Status = "ACTIVE"
)
