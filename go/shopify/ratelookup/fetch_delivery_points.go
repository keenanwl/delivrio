package ratelookup

import (
	"context"
	"delivrio.io/go/deliveryoptions"
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionlookup"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionrequest"
	"delivrio.io/go/shopify/rates"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func deliveryWindow(from, to int) (*string, *string) {

	if from == 0 && to == from {
		return nil, nil
	}

	fromOutput := bumpPastSunday(time.Now().AddDate(0, 0, from)).
		Format(ShopifyLookupDateFormat)
	toOutput := bumpPastSunday(time.Now().AddDate(0, 0, to)).
		Format(ShopifyLookupDateFormat)

	return &fromOutput, &toOutput
}

func fetchDeliveryPoints(
	ctx context.Context,
	trackingID *pulid.ID,
	do *ent.DeliveryOption,
	price *deliveryoptions.PricePair,
	nearestToAddress deliverypoints.DropPointLookupAddress,
	carrierLookupFn func(ctx context.Context, nearestToAddress deliverypoints.DropPointLookupAddress, maxCount int) ([]*ent.ParcelShop, error),
	carrierID carrierbrand.InternalID,
) ([]rates.RateResponse, error) {
	db := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)
	output := make([]rates.RateResponse, 0)

	servicePoints, err := carrierLookupFn(ctx, nearestToAddress, do.ClickOptionDisplayCount)
	if err != nil {
		return nil, err
	}

	minDeliveryWindow, maxDeliveryWindow := deliveryWindow(do.DeliveryEstimateFrom, do.DeliveryEstimateTo)

	for _, sp := range servicePoints {

		shopID, err := uniqueShopID(ctx, sp, carrierID)
		if err != nil {
			return nil, err
		}

		serviceCode := fmt.Sprintf("%v-%v", do.ID.String(), shopID)
		if trackingID != nil && *trackingID != "" {
			err := db.HypothesisTestDeliveryOptionLookup.Create().
				SetDeliveryOption(do).
				SetHypothesisTestDeliveryOptionRequestID(*trackingID).
				SetTenantID(view.TenantID()).
				Exec(ctx)
			if err != nil && !ent.IsConstraintError(err) {
				log.Println("Could not save hypothesis test lookup", err)
			}

			existingID, err := db.HypothesisTestDeliveryOptionLookup.Query().
				Where(
					hypothesistestdeliveryoptionlookup.And(
						hypothesistestdeliveryoptionlookup.HasDeliveryOptionWith(deliveryoption.ID(do.ID)),
						hypothesistestdeliveryoptionlookup.HasHypothesisTestDeliveryOptionRequestWith(
							hypothesistestdeliveryoptionrequest.ID(*trackingID),
						),
					),
				).OnlyID(ctx)
			if err != nil {
				log.Println("Could fetch save hypothesis test lookup", err)
			}

			serviceCode = fmt.Sprintf("%v-%v", existingID, shopID)
		}

		externalServiceCode, err := externalIntegrationCodeDropPoint(ctx, do, sp)
		if err != nil {

			log.Println("Could not calculate external service code", err)
		}
		if externalServiceCode != nil {
			serviceCode = *externalServiceCode
		}

		output = append(output, rates.RateResponse{
			ServiceName: fmt.Sprintf(
				"%v%v %v",
				strings.Repeat(invisibleOrderingCharacter, int(do.SortOrder)+1),
				do.Name,
				sp.Name,
			),
			// Should include parcel shop
			ServiceCode: serviceCode,
			// Shopify wants floats represented as ints
			TotalPrice:        fmt.Sprintf("%v", int(price.Price*100)),
			Description:       fmt.Sprintf("%v", do.Description),
			Currency:          price.Currency.CurrencyCode.String(),
			MinDeliveryDate:   minDeliveryWindow,
			MaxDeliveryDate:   maxDeliveryWindow,
			TotalPriceDecimal: price.Price,
			ParcelShop:        sp,
			DeliveryOption:    do,
		})
	}
	return output, nil
}

func externalIntegrationCode(do *ent.DeliveryOption, wsDropPoint WebshipperDropPoint, shipmondoDropPointID string) (*string, error) {
	var externalIntegrationCode *string
	if do.WebshipperIntegration {
		j, err := json.Marshal(WebshipperDeliveryPoint{
			ShippingRateID: do.WebshipperID,
			DropPoint:      wsDropPoint,
		})
		if err != nil {
			return nil, err
		}
		jString := string(j)
		externalIntegrationCode = &jString
	} else if do.ShipmondoIntegration {
		val := strings.Replace(do.ShipmondoDeliveryOption, "{{.DropPointID}}", shipmondoDropPointID, -1)
		externalIntegrationCode = &val
	}

	return externalIntegrationCode, nil
}

func externalIntegrationCodeDropPoint(ctx context.Context, do *ent.DeliveryOption, sp *ent.ParcelShop) (*string, error) {

	if !do.WebshipperIntegration && !do.ShipmondoIntegration {
		return nil, nil
	}

	cb, err := sp.CarrierBrand(ctx)
	if err != nil {
		return nil, err
	}

	wsDropPoint := WebshipperDropPoint{}
	shipmondoDropPointID := ""

	switch cb.InternalID {
	case carrierbrand.InternalIDGLS:
		ps, err := sp.QueryParcelShopGLS().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		spAddress, err := sp.QueryAddress().
			WithCountry().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipmondoDropPointID = ps.GLSParcelShopID
		wsDropPoint = WebshipperDropPoint{
			DropPointID: ps.GLSParcelShopID,
			Name:        sp.Name,
			Address1:    spAddress.AddressOne,
			Zip:         spAddress.Zip,
			City:        spAddress.City,
			CountryCode: spAddress.Edges.Country.Alpha2,
			Distance:    0,
		}
		break
	case carrierbrand.InternalIDDAO:
		ps, err := sp.QueryParcelShopDAO().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		spAddress, err := sp.QueryAddress().
			WithCountry().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipmondoDropPointID = ps.ShopID
		wsDropPoint = WebshipperDropPoint{
			DropPointID: ps.ShopID,
			Name:        sp.Name,
			Address1:    spAddress.AddressOne,
			Zip:         spAddress.Zip,
			City:        spAddress.City,
			CountryCode: spAddress.Edges.Country.Alpha2,
			Distance:    0,
		}
		break
	case carrierbrand.InternalIDPostNord:
		ps, err := sp.QueryParcelShopPostNord().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		spAddress, err := sp.QueryAddress().
			WithCountry().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipmondoDropPointID = ps.ServicePointID
		wsDropPoint = WebshipperDropPoint{
			DropPointID: ps.ServicePointID,
			Name:        sp.Name,
			Address1:    spAddress.AddressOne,
			Zip:         spAddress.Zip,
			City:        spAddress.City,
			CountryCode: spAddress.Edges.Country.Alpha2,
			Distance:    0,
		}
		break
	case carrierbrand.InternalIDBring:
		ps, err := sp.QueryParcelShopBring().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		spAddress, err := sp.QueryAddress().
			WithCountry().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		shipmondoDropPointID = ps.BringID
		wsDropPoint = WebshipperDropPoint{
			DropPointID: ps.BringID,
			Name:        sp.Name,
			Address1:    spAddress.AddressOne,
			Zip:         spAddress.Zip,
			City:        spAddress.City,
			CountryCode: spAddress.Edges.Country.Alpha2,
			Distance:    0,
		}
		break
	}

	return externalIntegrationCode(do, wsDropPoint, shipmondoDropPointID)
}

func uniqueShopID(ctx context.Context, ps *ent.ParcelShop, carrierID carrierbrand.InternalID) (string, error) {
	switch carrierID {
	case carrierbrand.InternalIDPostNord:
		pspn, err := ps.ParcelShopPostNord(ctx)
		if err != nil {
			return "", err
		}
		return pspn.Pudoid, nil
	case carrierbrand.InternalIDGLS:
		psgls, err := ps.ParcelShopGLS(ctx)
		if err != nil {
			return "", err
		}
		return psgls.GLSParcelShopID, nil
	}
	return "", fmt.Errorf("carrier brand not recognized: %v", carrierID)
}
