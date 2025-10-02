package bringapis

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/bringapis/bringrequest"
	"delivrio.io/go/carrierapis/bringapis/bringresponse"
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/addressglobal"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/parcelshop"
	"delivrio.io/go/ent/parcelshopbring"
	"delivrio.io/go/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("bringapis: may not set config twice")
	}
	conf = c
	confSet = true
}

func DeliveryPoints(ctx context.Context, nearestToAddress deliverypoints.DropPointLookupAddress, maxCount int) ([]*ent.ParcelShop, error) {
	cli := ent.FromContext(ctx)

	points, err := getPickupPoints(ctx, nearestToAddress.Country, nearestToAddress.Zip)
	if err != nil {
		return nil, err
	}

	carrierBrand, err := cli.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDBring)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*ent.ParcelShop, 0)

	count := 0
	for _, p := range points {
		if count >= maxCount {
			break
		}
		count++

		ps, err := cli.ParcelShopBring.Query().
			WithParcelShop().
			Where(parcelshopbring.BringID(p.ID)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if ent.IsNotFound(err) {

			tx, err := cli.Tx(ctx)
			if err != nil {
				return nil, err
			}
			defer tx.Rollback()

			txCtx := ent.NewTxContext(ctx, tx)

			visiting, err := tx.AddressGlobal.Create().
				SetAddressOne(p.VisitingAddress).
				SetCity(p.VisitingCity).
				SetZip(p.VisitingPostalCode).
				SetLatitude(p.Latitude).
				SetLatitude(p.Longitude).
				SetCountryID(nearestToAddress.Country.ID).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			delivery, err := tx.AddressGlobal.Create().
				SetAddressOne(p.Address).
				SetCity(p.City).
				SetZip(p.PostalCode).
				SetLatitude(p.Latitude).
				SetLatitude(p.Longitude).
				SetCountryID(nearestToAddress.Country.ID).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			ps, err := tx.ParcelShop.Create().
				SetCarrierBrand(carrierBrand).
				SetAddress(visiting).
				SetName(p.Name).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			err = tx.ParcelShopBring.Create().
				SetParcelShop(ps).
				SetAddressDelivery(delivery).
				SetBringID(p.ID).
				SetPointType(bringPointType(p.Type)).
				Exec(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			err = tx.Commit()
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			output = append(output, ps.Unwrap())
		} else {
			if ps.Edges.ParcelShop.LastUpdated.Before(deliverypoints.UpdateInterval) {

				tx, err := cli.Tx(ctx)
				if err != nil {
					return nil, err
				}
				defer tx.Rollback()

				txCtx := ent.NewTxContext(ctx, tx)

				err = tx.AddressGlobal.Update().
					SetAddressOne(p.VisitingAddress).
					SetCity(p.VisitingCity).
					SetZip(p.VisitingPostalCode).
					SetCountryID(nearestToAddress.Country.ID).
					Where(
						addressglobal.HasParcelShopWith(parcelshop.ID(ps.ID)),
					).Exec(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				err = tx.AddressGlobal.Update().
					SetAddressOne(p.Address).
					SetCity(p.City).
					SetZip(p.PostalCode).
					SetCountryID(nearestToAddress.Country.ID).
					Where(
						addressglobal.HasParcelShopBringDeliveryWith(parcelshopbring.ID(ps.ID)),
					).Exec(txCtx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}

				err = tx.ParcelShop.Update().
					SetLastUpdated(time.Now()).
					Where(parcelshop.ID(ps.ID)).
					Exec(ctx)
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}
				err = tx.Commit()
				if err != nil {
					return nil, utils.Rollback(tx, err)
				}
			}

			output = append(output, ps.Edges.ParcelShop)
		}
	}

	return output, nil

}

func bringPointType(val string) parcelshopbring.PointType {
	switch val {
	case "one":
		return parcelshopbring.PointTypeOne
	case "four":
		return parcelshopbring.PointTypeFour
	case "nineteen":
		return parcelshopbring.PointTypeNineteen
	case "twenty_one":
		return parcelshopbring.PointTypeTwentyOne
	case "thirty_two":
		return parcelshopbring.PointTypeThirtyTwo
	case "thirty_four":
		return parcelshopbring.PointTypeThirtyFour
	case "thirty_seven":
		return parcelshopbring.PointTypeThirtySeven
	case "thirty_eight":
		return parcelshopbring.PointTypeThirtyEight
	case "thirty_nine":
		return parcelshopbring.PointTypeThirtyNine
	case "eighty_five":
		return parcelshopbring.PointTypeEightyFive
	case "eighty_six":
		return parcelshopbring.PointTypeEightySix
	case "SmartPOST":
		return parcelshopbring.PointTypeSmartPOST
	case "Posti":
		return parcelshopbring.PointTypePosti
	case "Noutopiste":
		return parcelshopbring.PointTypeNoutopiste
	case "LOCKER":
		return parcelshopbring.PointTypeLOCKER
	default:
		return parcelshopbring.PointTypeUnknown
	}
}

func getPickupPoints(ctx context.Context, c *ent.Country, zip string) ([]bringresponse.PickupPoint, error) {
	u, err := url.Parse(fmt.Sprintf(
		`https://api.bring.com/pickuppoint/api/pickuppoint/%v/postalCode/%v`,
		c.Alpha2,
		zip,
	))
	if err != nil {
		return nil, err
	}

	vals := url.Values{}
	// We only support DK senders for now
	// Assume this is sender value?
	vals.Set("requestCountryCode", "DK")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = addAuthentication(req, bringrequest.AuthenticationHeaders{
		APIKey:    conf.Bring.APIKey,
		APIUID:    conf.Bring.APIUID,
		ClientURL: conf.BaseURL,
	})

	rdump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(rdump))

	cli := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rdump, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(rdump))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bring: fetch pickup points: %v", resp.StatusCode)
	}

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respBody bringresponse.ResponsePickupPoints
	err = json.Unmarshal(bodyData, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody.PickupPoint, nil
}
