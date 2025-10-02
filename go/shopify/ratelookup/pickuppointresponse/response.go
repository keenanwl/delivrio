package pickuppointresponse

type PickupPointDeliveryOption struct {
	Cost        float64     `json:"cost,omitempty"`
	PickupPoint PickupPoint `json:"pickupPoint"`
}

type PickupPoint struct {
	Address       PickupAddress   `json:"address"`
	BusinessHours []BusinessHours `json:"businessHours"`
	ExternalID    string          `json:"externalId"`
	Name          string          `json:"name"`
	Provider      Provider        `json:"provider"`
}

type BusinessHours struct {
	Day     Weekday               `json:"day"`
	Periods []BusinessHoursPeriod `json:"periods"`
}

type BusinessHoursPeriod struct {
	ClosingTime string `json:"closingTime"`
	OpeningTime string `json:"openingTime"`
}

type Weekday string

const (
	WeekdayFriday    Weekday = "FRIDAY"
	WeekdayMonday    Weekday = "MONDAY"
	WeekdaySaturday  Weekday = "SATURDAY"
	WeekdaySunday    Weekday = "SUNDAY"
	WeekdayThursday  Weekday = "THURSDAY"
	WeekdayTuesday   Weekday = "TUESDAY"
	WeekdayWednesday Weekday = "WEDNESDAY"
)

type Provider struct {
	LogoURL string `json:"logoUrl"`
	Name    string `json:"name"`
}

type PickupAddress struct {
	Address1     string  `json:"address1"`
	Address2     *string `json:"address2,omitempty"`
	City         string  `json:"city"`
	Country      *string `json:"country,omitempty"`
	CountryCode  string  `json:"countryCode"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Phone        *string `json:"phone,omitempty"`
	Province     *string `json:"province,omitempty"`
	ProvinceCode *string `json:"provinceCode,omitempty"`
	Zip          *string `json:"zip,omitempty"`
}
