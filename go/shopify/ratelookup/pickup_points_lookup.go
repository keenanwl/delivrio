package ratelookup

import (
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/ent"
	"delivrio.io/go/shopify/ratelookup/pickuppointrequest"
	"delivrio.io/go/shopify/ratelookup/pickuppointresponse"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// For the Pickup Points app
const openHourFormat = "15:04"

func PickupPointsHandler(w http.ResponseWriter, r *http.Request) {
	db := ent.FromContext(r.Context())

	ctx := viewer.NewContext(r.Context(), viewer.UserViewer{
		Role: viewer.Background,
	})

	conn, err := db.Connection.Query().
		WithConnectionShopify().
		First(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "Unknown token"})
		return
	}

	var req pickuppointrequest.FetchInput
	err = httputils.UnmarshalRequestBody(r, &req)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("INPUT REQ", req)

	var address pickuppointrequest.MailingAddress
	if len(req.Allocations) > 0 {
		address = req.Allocations[0].DeliveryAddress
	}

	shopifyItems := make([]shopifyItem, 0)
	for _, i := range req.Cart.Lines {

		costToFloat, err := strconv.ParseFloat(i.Cost.TotalAmount.Amount, 10)
		if err != nil {
			httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		shopifyItems = append(shopifyItems, shopifyItem{
			VariantID: i.Merchandise.ID,
			ProductID: i.Merchandise.Product.ID,
			Quantity:  i.Quantity,
			Price:     costToFloat,
		})
	}

	allVariants, totalProductsWeight, err := sortFetchOrderItems(ctx, conn.Edges.ConnectionShopify, shopifyItems)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	lookupCountry, err := rawCountryToDelivrio(ctx, fmt.Sprintf("%s", address.CountryCode))
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Default to true since we don't have the address here?
	deliveryOptions, trackingID, err := filterByHypothesisTestAndCompanyField(ctx, time.Now(), conn, req.Cart, address, true)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	nearestToAddress := deliverypoints.DropPointLookupAddress{
		Address1: address.Address1,
		Zip:      address.Zip,
		Country:  lookupCountry,
	}

	output, err := waitForRates(
		ctx,
		trackingID,
		deliveryOptions,
		allVariants,
		nearestToAddress,
		totalProductsWeight,
	)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(output)

	allPickups := make([]pickuppointresponse.PickupPointDeliveryOption, 0)
	for _, rate := range output.Rates {

		if rate.ParcelShop == nil {
			// We only want opens that are pickup points
			// for this endpoint
			continue
		}

		adr, err := rate.ParcelShop.QueryAddress().
			WithCountry().
			Only(ctx)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		hours, err := rate.ParcelShop.QueryBusinessHoursPeriod().
			All(ctx)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		allBusinessHours := make([]pickuppointresponse.BusinessHours, 0)
		for _, h := range hours {
			allBusinessHours = append(allBusinessHours, pickuppointresponse.BusinessHours{
				Day: pickuppointresponse.Weekday(h.DayOfWeek.String()),
				Periods: []pickuppointresponse.BusinessHoursPeriod{
					{
						ClosingTime: h.Closing.Format(time.TimeOnly),
						OpeningTime: h.Opening.Format(time.TimeOnly),
					},
				},
			})
		}

		cb, err := rate.ParcelShop.QueryCarrierBrand().
			Only(ctx)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		allPickups = append(allPickups, pickuppointresponse.PickupPointDeliveryOption{
			Cost: rate.TotalPriceDecimal,
			PickupPoint: pickuppointresponse.PickupPoint{
				Address: pickuppointresponse.PickupAddress{
					Address1:     adr.AddressOne,
					Address2:     &adr.AddressTwo,
					City:         adr.City,
					Country:      &adr.Edges.Country.Label,
					CountryCode:  adr.Edges.Country.Alpha2,
					Latitude:     adr.Latitude,
					Longitude:    adr.Longitude,
					Phone:        nil,
					Province:     &adr.State,
					ProvinceCode: nil,
					Zip:          &adr.Zip,
				},
				BusinessHours: allBusinessHours,
				ExternalID:    rate.ServiceCode,
				Name:          rate.ServiceName,
				Provider: pickuppointresponse.Provider{
					LogoURL: "https://cdn.shopify.com/s/files/1/0577/7778/2844/files/PostNord_Logo_DELIVRIO_White_Circle.webp?v=1712829323",
					Name:    cb.Label,
				},
			},
		})
	}

	httputils.JSONResponse(w, http.StatusOK, allPickups)
	return

}
