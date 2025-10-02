package glsdeliverypoints

import (
	"context"
	"delivrio.io/go/deliverypoints/glsdeliverypoints/glsresponse"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_saveDeliveryPoints(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := ent.NewContext(context.Background(), client)

	tx, _ := client.Tx(ctx)
	ctx = ent.NewTxContext(ctx, tx)
	seed.Base(ctx)

	tx.Commit()

	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Admin,
		Context: pulid.MustNew("CH"),
		Tenant:  seed.GetTenantID(),
	})

	ctx = ent.NewContext(ctx, client)

	shops, err := saveDeliveryPoints(ctx, sampleData())
	require.NoError(t, err)
	require.Len(t, shops, 5, "expect count to match sample data input")
	shops, err = saveDeliveryPoints(ctx, sampleData())
	require.Len(t, shops, 5, "expect unchanged count also when updating")
	require.Equal(t, "Bruuns Galleri informationen 1 sal", shops[0].Name)
	require.Equal(t, "Cbdhelten", shops[1].Name)
	require.Equal(t, "Silvan Lille Torv", shops[2].Name)
	require.Equal(t, "SuperBrugsen Vesterbro Torv", shops[3].Name)
	require.Equal(t, "Vingummi Nørregade", shops[4].Name)

}

func sampleData() glsresponse.Envelope {

	return glsresponse.Envelope{
		Body: glsresponse.Body{
			GetParcelShopDropPointResponse: glsresponse.GetParcelShopDropPointResponse{
				GetParcelShopDropPointResult: glsresponse.GetParcelShopDropPointResult{
					ParcelShops: []glsresponse.PakkeshopData{
						{
							Number:                       "99558",
							CompanyName:                  "Bruuns Galleri informationen 1 sal",
							Streetname:                   "M.P. Bruuns Gade 25",
							Streetname2:                  "Pakkeshop: 99558",
							ZipCode:                      "8000",
							CityName:                     "Aarhus C",
							CountryCode:                  "008",
							CountryCodeISO3166A2:         "DK",
							Telephone:                    "-",
							Longitude:                    "10.2048",
							Latitude:                     "56.1491",
							DistanceMetersAsTheCrowFlies: "430",
							OpeningHours: []glsresponse.Weekday{
								{
									Day: "Monday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "20:00",
									},
									Breaks: "",
								},
								{
									Day: "Tuesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "20:00",
									},
									Breaks: "",
								},
								{
									Day: "Wednesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "20:00",
									},
									Breaks: "",
								},
								{
									Day: "Thursday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "20:00",
									},
									Breaks: "",
								},
								{
									Day: "Friday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "20:00",
									},
									Breaks: "",
								},
								{
									Day: "Saturday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Sunday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "18:00",
									},
									Breaks: "",
								},
							},
						},
						{
							Number:                       "96936",
							CompanyName:                  "Cbdhelten",
							Streetname:                   "Thorvaldsensgade 3",
							Streetname2:                  "Pakkeshop: 96936",
							ZipCode:                      "8000",
							CityName:                     "Aarhus C",
							CountryCode:                  "008",
							CountryCodeISO3166A2:         "DK",
							Telephone:                    "-",
							Longitude:                    "10.1984",
							Latitude:                     "56.1554",
							DistanceMetersAsTheCrowFlies: "461",
							OpeningHours: []glsresponse.Weekday{
								{
									Day: "Monday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Tuesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Wednesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Thursday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Friday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Saturday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "18:00",
									},
									Breaks: "",
								},
								{
									Day: "Sunday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "12:00",
										To:   "16:00",
									},
									Breaks: "",
								},
							},
						},
						{
							Number:                       "96315",
							CompanyName:                  "Silvan Lille Torv",
							Streetname:                   "Lille Torv 6",
							Streetname2:                  "Pakkeshop: 96315",
							ZipCode:                      "8000",
							CityName:                     "Aarhus C",
							CountryCode:                  "008",
							CountryCodeISO3166A2:         "DK",
							Telephone:                    "-",
							Longitude:                    "10.2069",
							Latitude:                     "56.1578",
							DistanceMetersAsTheCrowFlies: "554",
							OpeningHours: []glsresponse.Weekday{
								{
									Day: "Monday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
								{
									Day: "Tuesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
								{
									Day: "Wednesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
								{
									Day: "Thursday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
								{
									Day: "Friday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
								{
									Day: "Saturday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
								{
									Day: "Sunday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "09:00",
										To:   "19:00",
									},
									Breaks: "",
								},
							},
						},
						{
							Number:                       "95940",
							CompanyName:                  "SuperBrugsen Vesterbro Torv",
							Streetname:                   "Vesterbro Torv 1",
							Streetname2:                  "Pakkeshop: 95940",
							ZipCode:                      "8000",
							CityName:                     "Aarhus C",
							CountryCode:                  "008",
							CountryCodeISO3166A2:         "DK",
							Telephone:                    "-",
							Longitude:                    "10.1992",
							Latitude:                     "56.158",
							DistanceMetersAsTheCrowFlies: "641",
							OpeningHours: []glsresponse.Weekday{
								{
									Day: "Monday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
								{
									Day: "Tuesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
								{
									Day: "Wednesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
								{
									Day: "Thursday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
								{
									Day: "Friday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
								{
									Day: "Saturday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
								{
									Day: "Sunday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "07:00",
										To:   "22:00",
									},
									Breaks: "",
								},
							},
						},
						{
							Number:                       "95105",
							CompanyName:                  "Vingummi Nørregade",
							Streetname:                   "Nørregade 49",
							Streetname2:                  "Pakkeshop: 95105",
							ZipCode:                      "8000",
							CityName:                     "Aarhus C",
							CountryCode:                  "008",
							CountryCodeISO3166A2:         "DK",
							Telephone:                    "-",
							Longitude:                    "10.2096",
							Latitude:                     "56.1607",
							DistanceMetersAsTheCrowFlies: "908",
							OpeningHours: []glsresponse.Weekday{
								{
									Day: "Monday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "20:30",
									},
									Breaks: "",
								},
								{
									Day: "Tuesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "20:30",
									},
									Breaks: "",
								},
								{
									Day: "Wednesday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "20:30",
									},
									Breaks: "",
								},
								{
									Day: "Thursday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "20:30",
									},
									Breaks: "",
								},
								{
									Day: "Friday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "21:00",
									},
									Breaks: "",
								},
								{
									Day: "Saturday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "10:00",
										To:   "21:00",
									},
									Breaks: "",
								},
								{
									Day: "Sunday",
									OpenAt: struct {
										From string `xml:"From"`
										To   string `xml:"To"`
									}{
										From: "11:00",
										To:   "20:30",
									},
									Breaks: "",
								},
							},
						},
					},
				},
			},
		},
	}

}
