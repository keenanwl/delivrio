package daoapis

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/daoapis/daorequest"
	"delivrio.io/go/carrierapis/daoapis/daoresponse"
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/addressglobal"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/parcelshop"
	"delivrio.io/go/ent/parcelshopdao"
	"delivrio.io/go/utils"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("daoapis: may not set config twice")
	}
	conf = c
	confSet = true
}

type ctxKey struct{}

func DeliveryOptionIDFromContext(ctx context.Context) pulid.ID {
	v, _ := ctx.Value(ctxKey{}).(pulid.ID)
	return v
}

func WithDeliveryOptionID(parent context.Context, deliveryOptionID pulid.ID) context.Context {
	return context.WithValue(parent, ctxKey{}, deliveryOptionID)
}

func DeliveryPoints(ctx context.Context, nearestToAddress deliverypoints.DropPointLookupAddress, maxCount int) ([]*ent.ParcelShop, error) {
	cli := ent.FromContext(ctx)

	points, err := getPickupPoints(ctx, nearestToAddress.Zip)
	if err != nil {
		return nil, err
	}

	carrierBrand, err := cli.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDDAO)).
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

		ps, err := cli.ParcelShopDAO.Query().
			WithParcelShop().
			Where(parcelshopdao.ShopID(p.ShopId)).
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

			lat, err := strconv.ParseFloat(p.Latitude, 64)
			if err != nil {
				log.Println("dao: can't convert string lat to float: ", err)
			}
			lon, err := strconv.ParseFloat(p.Longitude, 64)
			if err != nil {
				log.Println("dao: can't convert string long to float: ", err)
			}

			visiting, err := tx.AddressGlobal.Create().
				SetAddressOne(p.Adresse).
				SetCity(p.Bynavn).
				SetZip(p.Postnr).
				SetLatitude(lat).
				SetLatitude(lon).
				SetCountryID(nearestToAddress.Country.ID).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			ps, err := tx.ParcelShop.Create().
				SetCarrierBrand(carrierBrand).
				SetAddress(visiting).
				SetName(p.Navn).
				Save(txCtx)
			if err != nil {
				return nil, utils.Rollback(tx, err)
			}

			err = tx.ParcelShopDAO.Create().
				SetParcelShop(ps).
				SetShopID(p.ShopId).
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
					SetAddressOne(p.Adresse).
					SetCity(p.Bynavn).
					SetZip(p.Postnr).
					SetCountryID(nearestToAddress.Country.ID).
					Where(
						addressglobal.HasParcelShopWith(parcelshop.ID(ps.ID)),
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

func getPickupPoints(ctx context.Context, zip string) ([]daoresponse.ParcelShops, error) {
	cli := ent.FromContext(ctx)
	doID := DeliveryOptionIDFromContext(ctx)

	// We need the agreement
	agreement, err := cli.DeliveryOption.Query().
		Where(deliveryoption.ID(doID)).
		QueryCarrier().
		QueryCarrierDAO().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(`https://api.dao.as/DAOPakkeshop/FindPakkeshop.php`)
	if err != nil {
		return nil, err
	}

	vals := &url.Values{}
	vals = authFormData(vals, daorequest.Authentication{
		CustomerID: agreement.CustomerID,
		Code:       agreement.APIKey,
	})
	vals.Set("postnr", zip)
	vals.Set("format", "json")
	vals.Set("antal", "5")

	u.RawQuery = vals.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	rdump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(rdump))

	httpCli := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rdump, _ = httputil.DumpResponse(resp, true)
	fmt.Println(string(rdump))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dao: fetch pickup points: %v", resp.StatusCode)
	}

	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respBody daoresponse.ResponseDeliveryPoints
	err = json.Unmarshal(bodyData, &respBody)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(respBody.ErrorCodeParsed(), daoresponse.StatusOK.String()) {
		return nil, fmt.Errorf("dao error: %s", respBody.ErrorMessage)
	}

	return respBody.Results.ParcelShops, nil
}
