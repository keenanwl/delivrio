package postnorddeliverypoints

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/deliverypoints/postnorddeliverypoints/postnordresponse"
	"delivrio.io/go/ent/businesshoursperiod"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/addressglobal"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/parcelshop"
	"delivrio.io/go/ent/parcelshoppostnord"
	"delivrio.io/go/utils"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("postnordservicepoint: may not set config twice")
	}
	conf = c
	confSet = true
}

func DeliveryPoints(ctx context.Context, nearestToAddress deliverypoints.DropPointLookupAddress, maxCount int) ([]*ent.ParcelShop, error) {
	db := ent.FromContext(ctx)

	u, err := url.Parse("https://atapi2.postnord.com")
	if err != nil {
		return nil, err
	}

	request, err := fetchServicePoints(u, conf.PostNord.APIKey, nearestToAddress.Country.Alpha2, nearestToAddress.Zip)
	if err != nil {
		return nil, err
	}

	rdump, _ := httputil.DumpRequest(request, true)
	fmt.Println(string(rdump))

	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pnCarrierBrand, err := db.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDPostNord)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*ent.ParcelShop, 0)

	rdump, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(rdump))

	if resp.StatusCode == 200 {
		var body postnordresponse.ServicePointsResponse
		err = utils.ReadBody(resp.Body, &body)
		if err != nil {
			return nil, err
		}

		count := 0
		for _, sp := range body.ServicePointInformationResponse.ServicePoints {
			if count >= maxCount {
				break
			}
			count++

			tx, err := db.Tx(ctx)
			if err != nil {
				return nil, err
			}
			defer tx.Rollback()

			txCtx := ent.NewTxContext(ctx, tx)

			pspn, err := tx.ParcelShopPostNord.Query().
				WithParcelShop().
				Where(parcelshoppostnord.Pudoid(sp.PudoID)).
				Only(txCtx)
			if ent.IsNotFound(err) {

				var lat float64 = 0
				var lon float64 = 0
				if len(sp.Coordinates) > 0 && sp.Coordinates[0].Northing != nil && sp.Coordinates[0].Easting != nil {
					lat = *sp.Coordinates[0].Northing
					lon = *sp.Coordinates[0].Easting
				}

				visiting, err := tx.AddressGlobal.Create().
					SetAddressOne(fmt.Sprintf("%v %v", sp.VisitingAddress.StreetName, sp.VisitingAddress.StreetNumber)).
					SetCity(sp.VisitingAddress.City).
					SetZip(sp.VisitingAddress.PostalCode).
					SetCountryID(nearestToAddress.Country.ID).
					SetLatitude(lat).
					SetLongitude(lon).
					Save(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				delivery, err := tx.AddressGlobal.Create().
					SetAddressOne(fmt.Sprintf("%v %v", sp.DeliveryAddress.StreetName, sp.DeliveryAddress.StreetNumber)).
					SetCity(sp.DeliveryAddress.City).
					SetZip(sp.DeliveryAddress.PostalCode).
					SetCountryID(nearestToAddress.Country.ID).
					SetLatitude(lat).
					SetLongitude(lon).
					Save(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				ps, err := tx.ParcelShop.Create().
					SetCarrierBrand(pnCarrierBrand).
					SetAddress(visiting).
					SetName(sp.Name).
					Save(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				err = tx.ParcelShopPostNord.Create().
					SetParcelShop(ps).
					SetAddressDelivery(delivery).
					SetPudoid(sp.PudoID).
					SetServicePointID(sp.ServicePointID).
					SetTypeID(strconv.FormatInt(sp.Type.TypeID, 10)).
					Exec(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				openHoursCreate := make([]*ent.BusinessHoursPeriodCreate, 0)
				for _, h := range sp.OpeningHours.PostalServices {
					days, err := openRangeToBusinessHours(txCtx, ps, h)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}
					openHoursCreate = append(openHoursCreate, days...)
				}

				err = tx.BusinessHoursPeriod.CreateBulk(openHoursCreate...).
					Exec(ctx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				output = append(output, ps.Unwrap())

			} else if err != nil {
				return nil, utils.Rollback(tx, err)
			} else {

				if pspn.Edges.ParcelShop.LastUpdated.Before(deliverypoints.UpdateInterval) {

					visitingAddressOne := fmt.Sprintf("%v %v", sp.VisitingAddress.StreetName, sp.VisitingAddress.StreetNumber)

					err = tx.AddressGlobal.Update().
						SetAddressOne(visitingAddressOne).
						SetCity(sp.VisitingAddress.City).
						SetZip(sp.VisitingAddress.PostalCode).
						SetCountryID(nearestToAddress.Country.ID).
						Where(
							addressglobal.HasParcelShopWith(parcelshop.ID(pspn.ID)),
						).Exec(txCtx)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}

					deliveryAddressOne := fmt.Sprintf("%v %v", sp.DeliveryAddress.StreetName, sp.DeliveryAddress.StreetNumber)

					err = tx.AddressGlobal.Update().
						SetAddressOne(deliveryAddressOne).
						SetCity(sp.DeliveryAddress.City).
						SetZip(sp.DeliveryAddress.PostalCode).
						SetCountryID(nearestToAddress.Country.ID).
						Where(
							addressglobal.HasParcelShopPostNordDeliveryWith(parcelshoppostnord.ID(pspn.ID)),
						).Exec(txCtx)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}

					_, err = tx.BusinessHoursPeriod.Delete().
						Where(businesshoursperiod.HasParcelShopWith(parcelshop.ID(pspn.Edges.ParcelShop.ID))).
						Exec(ctx)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}

					openHoursCreate := make([]*ent.BusinessHoursPeriodCreate, 0)
					for _, h := range sp.OpeningHours.PostalServices {
						days, err := openRangeToBusinessHours(txCtx, pspn.Edges.ParcelShop, h)
						if err != nil {
							return nil, err
						}
						openHoursCreate = append(openHoursCreate, days...)
					}

					err = tx.BusinessHoursPeriod.CreateBulk(openHoursCreate...).
						Exec(ctx)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}

					err = pspn.Edges.ParcelShop.Update().
						SetLastUpdated(time.Now()).
						Exec(txCtx)
					if err != nil {
						return nil, utils.Rollback(tx, err)
					}

				}

				// Unwrapping means we can't use it any longer in this transaction
				output = append(output, pspn.Edges.ParcelShop.Unwrap())
			}

			err = tx.Commit()
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

		}

	} else {
		return nil, fmt.Errorf("post nord request failed with status: %s", resp.Status)
	}

	return output, nil

}

const openHourFormat = "15:04"

func openRangeToBusinessHours(ctx context.Context, ps *ent.ParcelShop, input postnordresponse.PostalService) ([]*ent.BusinessHoursPeriodCreate, error) {
	tx := ent.TxFromContext(ctx)

	openTime, err := time.Parse(openHourFormat, input.OpenTime)
	if err != nil {
		return nil, err
	}

	closeTime, err := time.Parse(openHourFormat, input.CloseTime)
	if err != nil {
		return nil, err
	}
	output := make([]*ent.BusinessHoursPeriodCreate, 0)

	if input.OpenDay != input.CloseDay {
		output = append(output, tx.BusinessHoursPeriod.Create().
			SetParcelShop(ps).
			SetOpening(openTime).
			SetClosing(time.Date(0, 0, 0, 23, 59, 0, 0, time.Local)).
			SetDayOfWeek(businesshoursperiod.DayOfWeek(strings.ToUpper(input.OpenDay.String()))))

		output = append(output, tx.BusinessHoursPeriod.Create().
			SetParcelShop(ps).
			SetOpening(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)). // Zero values = midnight
			SetClosing(closeTime).
			SetDayOfWeek(businesshoursperiod.DayOfWeek(strings.ToUpper(input.CloseDay.String()))))

		for _, otherDay := range dayDistance(input.OpenDay, input.CloseDay) {
			output = append(output, tx.BusinessHoursPeriod.Create().
				SetParcelShop(ps).
				SetOpening(time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)). // Zero values = midnight
				SetClosing(time.Date(0, 0, 0, 23, 59, 0, 0, time.Local)).
				SetDayOfWeek(businesshoursperiod.DayOfWeek(strings.ToUpper(otherDay.String()))))
		}

	} else if input.OpenDay == input.CloseDay {
		output = append(output,
			tx.BusinessHoursPeriod.Create().
				SetParcelShop(ps).
				SetOpening(openTime).
				SetClosing(closeTime).
				SetDayOfWeek(businesshoursperiod.DayOfWeek(strings.ToUpper(input.OpenDay.String()))))
	}

	return output, nil

}

func dayDistance(dayOpen, dayClose postnordresponse.Day) []postnordresponse.Day {
	days := make(map[postnordresponse.Day]int)
	days[postnordresponse.Monday] = 0
	days[postnordresponse.Tuesday] = 1
	days[postnordresponse.Wednesday] = 2
	days[postnordresponse.Thursday] = 3
	days[postnordresponse.Friday] = 4
	days[postnordresponse.Saturday] = 5
	days[postnordresponse.Sunday] = 6

	lookup := make(map[int]postnordresponse.Day)
	lookup[0] = postnordresponse.Monday
	lookup[1] = postnordresponse.Tuesday
	lookup[2] = postnordresponse.Wednesday
	lookup[3] = postnordresponse.Thursday
	lookup[4] = postnordresponse.Friday
	lookup[5] = postnordresponse.Saturday
	lookup[6] = postnordresponse.Sunday

	distance := days[dayClose] - days[dayOpen]

	if distance < 0 {
		distance = 7 + distance
	}

	allDaysInBetween := make([]postnordresponse.Day, 0)
	for i := 1; i < distance; i++ {
		day := (days[dayOpen] + i) % 7
		allDaysInBetween = append(allDaysInBetween, lookup[day])
	}

	return allDaysInBetween
}

func fetchServicePoints(baseURL *url.URL, apiKey string, countryCode string, zip string) (*http.Request, error) {

	baseURL.Path = "/rest/businesslocation/v5/servicepoints/bypostalcode"

	query := baseURL.Query()
	query.Set("apikey", apiKey)
	query.Set("returnType", "json")
	query.Set("countryCode", countryCode)
	query.Set("postalCode", zip)
	query.Set("context", "optionalservicepoint")
	baseURL.RawQuery = query.Encode()
	fullUrl := baseURL.String()

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil

}
