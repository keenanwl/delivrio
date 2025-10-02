package glsresponse

import "encoding/xml"

// Uses the old DK-WSDL setup since the new one only supports
// Lat/Long lookup, which we don't have

// Envelope represents the SOAP envelope.
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"Body"`
}

// Body represents the SOAP body.
type Body struct {
	XMLName                        xml.Name                       `xml:"Body"`
	GetParcelShopDropPointResponse GetParcelShopDropPointResponse `xml:"http://gls.dk/webservices/ GetParcelShopDropPointResponse"`
}

// GetParcelShopDropPointResponse represents the GetParcelShopDropPointResponse element.
type GetParcelShopDropPointResponse struct {
	XMLName                      xml.Name                     `xml:"GetParcelShopDropPointResponse"`
	GetParcelShopDropPointResult GetParcelShopDropPointResult `xml:"GetParcelShopDropPointResult"`
}

// GetParcelShopDropPointResult represents the GetParcelShopDropPointResult element.
type GetParcelShopDropPointResult struct {
	XMLName       xml.Name        `xml:"GetParcelShopDropPointResult"`
	AccuracyLevel string          `xml:"accuracylevel"`
	ParcelShops   []PakkeshopData `xml:"parcelshops>PakkeshopData"`
}

// PakkeshopData represents the PakkeshopData element.
type PakkeshopData struct {
	XMLName                      xml.Name  `xml:"PakkeshopData"`
	Number                       string    `xml:"Number"`
	CompanyName                  string    `xml:"CompanyName"`
	Streetname                   string    `xml:"Streetname"`
	Streetname2                  string    `xml:"Streetname2"`
	ZipCode                      string    `xml:"ZipCode"`
	CityName                     string    `xml:"CityName"`
	CountryCode                  string    `xml:"CountryCode"`
	CountryCodeISO3166A2         string    `xml:"CountryCodeISO3166A2"`
	Telephone                    string    `xml:"Telephone"`
	Longitude                    string    `xml:"Longitude"`
	Latitude                     string    `xml:"Latitude"`
	DistanceMetersAsTheCrowFlies string    `xml:"DistanceMetersAsTheCrowFlies"`
	OpeningHours                 []Weekday `xml:"OpeningHours>Weekday"`
}

// Weekday represents the Weekday element.
type Weekday struct {
	XMLName xml.Name `xml:"Weekday"`
	Day     string   `xml:"day"`
	OpenAt  struct {
		From string `xml:"From"`
		To   string `xml:"To"`
	} `xml:"openAt"`
	Breaks string `xml:"breaks"`
}
