package endpoints

import (
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"fmt"
	"net/http"
)

/*const lookupUrl = `https://dawa.aws.dk/adresser/autocomplete?q=%s*`

type AddressAutocompleteRequest struct {
	Address string `json:"streetName"`
}

type AddressAutocompleteResponse struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	StreetName   string `json:"streetName"`
	StreetNumber string `json:"street_number"`
	FloorDoor    string `json:"floor_door"`
	PostCode     string `json:"post_code"`
	CityName     string `json:"city_name"`
	FullResult   string `json:"full_result"`
}
type DawaResult struct {
	Text    string      `json:"tekst"`
	Address DawaAddress `json:"adresse"`
}
type DawaAddress struct {
	ID                     string      `json:"id"`
	Status                 int64       `json:"status"`
	Darstatus              int64       `json:"darstatus"`
	Vejkode                string      `json:"vejkode"`
	Vejnavn                string      `json:"vejnavn"`
	Adresseringsvejnavn    string      `json:"adresseringsvejnavn"`
	Husnr                  string      `json:"husnr"`
	Floor                  *string     `json:"etage"`
	Door                   interface{} `json:"dør"`
	Supplerendebynavn      interface{} `json:"supplerendebynavn"`
	Postnr                 string      `json:"postnr"`
	Postnrnavn             string      `json:"postnrnavn"`
	Stormodtagerpostnr     interface{} `json:"stormodtagerpostnr"`
	Stormodtagerpostnrnavn interface{} `json:"stormodtagerpostnrnavn"`
	Kommunekode            string      `json:"kommunekode"`
	Adgangsadresseid       string      `json:"adgangsadresseid"`
	X                      float64     `json:"x"`
	Y                      float64     `json:"y"`
	Href                   string      `json:"href"`
}*/

func AddressAutocompleteHandler(w http.ResponseWriter, r *http.Request) {

	v := viewer.FromContext(r.Context())
	fmt.Println(v.TenantID())
	httputils.JSONResponse(w, http.StatusOK, httputils.Map{"SOME STRING": 1})
	return

	/*var request AddressAutocompleteRequest

	err := c.Bind(&request)
	check(err)

	resp, err := http.Get(fmt.Sprintf(lookupUrl, url.QueryEscape(request.Address)))
	check(err)

	if resp.StatusCode == 200 {

		var results []DawaResult
		err = server.readBody(resp.Body, &results)
		check(err)

		byrdieResults := make([]Address, 0)
		for _, adr := range results {

			floor := ""
			if adr.Address.Floor != nil {
				floor = *adr.Address.Floor
			}
			door, ok := adr.Address.Door.(string)
			if !ok {
				door = ""
			}

			byrdieResults = append(byrdieResults, Address{
				StreetName:   adr.Address.Adresseringsvejnavn,
				StreetNumber: adr.Address.Husnr,
				FloorDoor:    floor + ` ` + door,
				PostCode:     adr.Address.Postnr,
				CityName:     adr.Address.Postnrnavn,
				FullResult:   adr.Text,
			})
		}

		c.JSON(http.StatusOK, AddressAutocompleteResponse{
			Success:   true,
			Addresses: byrdieResults,
		})
		return
	}

	c.JSON(http.StatusOK, AddressAutocompleteResponse{
		Success:   false,
		Message:   "Lookup failed",
		Addresses: make([]Address, 0),
	})
	*/
}
