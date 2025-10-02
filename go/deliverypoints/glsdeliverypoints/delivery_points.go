package glsdeliverypoints

import (
	"bytes"
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/deliverypoints/glsdeliverypoints/glsresponse"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/addressglobal"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/parcelshop"
	"delivrio.io/go/ent/parcelshopgls"
	"delivrio.io/go/utils"
	"encoding/xml"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var conf *appconfig.DelivrioConfig
var confSet = false
var tracer = otel.Tracer("delivery-point-GLS")

const requestTimeout = time.Second * 10

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("labels: may not set config twice")
	}
	conf = c
	confSet = true
}

// Should only return locations where the customer can pick-up their package.
// Drop point count should be >= Service points anyways.
func DeliveryPoints(ctx context.Context, nearestToAddress deliverypoints.DropPointLookupAddress, maxCount int) ([]*ent.ParcelShop, error) {

	ctx, span := tracer.Start(ctx, "DeliveryPoints")
	defer span.End()

	span.SetAttributes(
		attribute.Int("maxResults", maxCount),
		attribute.String("street", nearestToAddress.Address1),
		attribute.String("ZipCode", nearestToAddress.Zip),
		attribute.String("country", nearestToAddress.Country.Alpha2),
	)

	url := "http://www.gls.dk/webservices_v4/wsShopFinder.asmx"

	// Define the data for the SOAP request
	data := struct {
		Street           string
		ZipCode          string
		CountryIso3166A2 string
		Amount           int
	}{
		Street:           nearestToAddress.Address1,
		ZipCode:          nearestToAddress.Zip,
		CountryIso3166A2: nearestToAddress.Country.Alpha2,
		Amount:           maxCount,
	}

	// Define the SOAP XML template
	soapXMLTemplate := strings.TrimSpace(`<?xml version="1.0" encoding="utf-8"?>
		<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		  <soap:Body>
		    <GetParcelShopDropPoint xmlns="http://gls.dk/webservices/">
		      <street>{{ .Street }}</street>
		      <zipcode>{{ .ZipCode }}</zipcode>
		      <countryIso3166A2>{{ .CountryIso3166A2 }}</countryIso3166A2>
		      <Amount>{{ .Amount }}</Amount>
		    </GetParcelShopDropPoint>
		  </soap:Body>
		</soap:Envelope>`)

	// Create a new template and parse the SOAP XML template
	tmpl, err := template.New("soapXML").Parse(soapXMLTemplate)
	if err != nil {
		return nil, err
	}

	// Execute the template with the data
	var soapXMLBuffer bytes.Buffer
	if err := tmpl.Execute(&soapXMLBuffer, data); err != nil {
		return nil, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, &soapXMLBuffer)
	if err != nil {
		return nil, err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "text/xml")

	rdump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(rdump))

	span.AddEvent("FireGLS")
	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	span.AddEvent("ResponseGLS")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rdump, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(rdump))

	if resp.StatusCode == 200 {

		var result glsresponse.Envelope
		if err := xml.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		return saveDeliveryPoints(ctx, result)

	}
	return nil, fmt.Errorf("GLS SOAP request failed with status: %s: %v", resp.Status, string(body))
}

func saveDeliveryPoints(ctx context.Context, resp glsresponse.Envelope) ([]*ent.ParcelShop, error) {
	db := ent.FromContext(ctx)
	output := make([]*ent.ParcelShop, 0)
	for _, sp := range resp.Body.GetParcelShopDropPointResponse.GetParcelShopDropPointResult.ParcelShops {

		tx, err := db.Tx(ctx)
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()

		txCtx := ent.NewTxContext(ctx, tx)
		coun, err := tx.Country.Query().
			Where(country.Alpha2(sp.CountryCodeISO3166A2)).
			Only(ctx)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}

		glsCB, err := tx.CarrierBrand.Query().
			Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDGLS)).
			Only(ctx)
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}

		ps, err := tx.ParcelShopGLS.Query().
			WithParcelShop().
			Where(parcelshopgls.GLSParcelShopID(sp.Number)).
			Only(txCtx)
		if ent.IsNotFound(err) {

			lat, err := strconv.ParseFloat(sp.Latitude, 64)
			if err != nil {
				log.Println("gls: can't convert string lat to float: ", err)
			}
			lon, err := strconv.ParseFloat(sp.Longitude, 64)
			if err != nil {
				log.Println("gls: can't convert string long to float: ", err)
			}

			delivery, err := tx.AddressGlobal.Create().
				SetCompany(sp.CompanyName).
				SetAddressOne(sp.Streetname).
				SetAddressTwo(sp.Streetname2).
				SetLatitude(lat).
				SetLatitude(lon).
				SetCity(sp.CityName).
				SetZip(sp.ZipCode).
				SetCountry(coun).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			ps, err := tx.ParcelShop.Create().
				SetCarrierBrand(glsCB).
				SetAddress(delivery).
				SetName(sp.CompanyName).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			err = tx.ParcelShopGLS.Create().
				SetParcelShop(ps).
				SetGLSParcelShopID(sp.Number).
				Exec(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}
			output = append(output, ps.Unwrap())
		} else if err != nil {
			return nil, utils.Rollback(tx, err)
		} else {
			if ps.Edges.ParcelShop.LastUpdated.Before(deliverypoints.UpdateInterval) {
				err = tx.AddressGlobal.Update().
					SetCompany(sp.CompanyName).
					SetAddressOne(sp.Streetname).
					SetAddressTwo(sp.Streetname2).
					SetCity(sp.CityName).
					SetZip(sp.ZipCode).
					SetCountry(coun).
					Where(
						addressglobal.HasParcelShopWith(parcelshop.ID(ps.ID)),
					).Exec(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				err = ps.Edges.ParcelShop.Update().SetLastUpdated(time.Now()).Exec(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}
			}

			output = append(output, ps.Edges.ParcelShop.Unwrap())
		}
		err = tx.Commit()
		if err != nil {
			return nil, utils.Rollback(tx, err)
		}
	}

	return output, nil

}
