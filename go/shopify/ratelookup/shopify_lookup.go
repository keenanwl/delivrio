package ratelookup

import (
	"context"
	"delivrio.io/go/carrierapis/bringapis"
	"delivrio.io/go/carrierapis/daoapis"
	"delivrio.io/go/deliverypoints"
	"delivrio.io/go/deliverypoints/glsdeliverypoints"
	"delivrio.io/go/deliverypoints/postnorddeliverypoints"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionlookup"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionrequest"
	"delivrio.io/go/shopify/productsync"
	"delivrio.io/go/utils/httputils"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/exp/slices"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"delivrio.io/go/deliveryoptions"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/connectionshopify"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/productvariant"
	"delivrio.io/go/shopify/rates"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

var tracer = otel.Tracer("shopify-ratelookup")

const ShopifyLookupDateFormat = "2006-01-02 15:04:05 Z07:00"

// const invisibleOrderingCharacter = " "
// Disable the character until we decide if we need it
const invisibleOrderingCharacter = ""

type WebshipperDeliveryPoint struct {
	ShippingRateID int                 `json:"shipping_rate_id"`
	DropPoint      WebshipperDropPoint `json:"drop_point"`
}

type WebshipperDropPoint struct {
	DropPointID string  `json:"drop_point_id"`
	Name        string  `json:"name"`
	Address1    string  `json:"address_1"`
	Zip         string  `json:"zip"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	Distance    float64 `json:"distance"`
}

type shopifyItem struct {
	ProductID string
	VariantID string
	Quantity  int
	Price     float64
}

func sortFetchOrderItems(ctx context.Context, shop *ent.ConnectionShopify, shopifyItems []shopifyItem) ([]*deliveryoptions.ConstraintProductWeight, int, error) {
	allVariants := make([]*deliveryoptions.ConstraintProductWeight, 0)
	productsNotFound := make([]string, 0)
	// For second pass
	itemsNotFound := make([]shopifyItem, 0)
	totalProductsWeight := 0

	for _, item := range shopifyItems {
		variant, err := ProductVariantFromShopifyID(ctx, item.VariantID)
		if ent.IsNotFound(err) {
			productsNotFound = append(productsNotFound, item.ProductID)
			itemsNotFound = append(itemsNotFound, item)
		} else if err != nil {
			return nil, 0, err
		} else {
			lookup, err := VariantToProductLookup(ctx, variant, item.Quantity, item.Price)
			if err != nil {
				return nil, 0, err
			}
			allVariants = append(allVariants, lookup)
			totalProductsWeight += lookup.WeightG * lookup.Units
		}
	}

	err := productsync.RealTimeShopifyProductSync(ctx, shop.StoreURL, shop.APIKey, productsNotFound)
	if err != nil {
		log.Printf("error: real time product sync failed: %v\n", err)
	}

	for _, item := range itemsNotFound {
		variant, err := ProductVariantFromShopifyID(ctx, item.VariantID)
		if ent.IsNotFound(err) {
			log.Println("ID", item)
			log.Println("error: real time product sync: this shouldn't happen")
		} else if err != nil {
			return nil, 0, err
		} else {
			lookup, err := VariantToProductLookup(ctx, variant, item.Quantity, item.Price)
			if err != nil {
				return nil, 0, err
			}
			allVariants = append(allVariants, lookup)
			totalProductsWeight += lookup.WeightG * lookup.Units
		}
	}

	return allVariants, totalProductsWeight, nil
}

func rawCountryToDelivrio(ctx context.Context, lookup string) (*ent.Country, error) {
	db := ent.FromContext(ctx)

	lookup = strings.TrimSpace(lookup)
	if len(lookup) == 0 {
		return nil, fmt.Errorf("shopify: country not found: %s", lookup)
	}

	lookupCountry, err := db.Country.Query().
		Where(country.Alpha2EqualFold(lookup)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return lookupCountry, nil
}

func ShopifyLookupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := viewer.NewContext(r.Context(), viewer.UserViewer{
		Role: viewer.Background,
	})
	cli := ent.FromContext(ctx)

	lookupID := r.URL.Query().Get(delivrioroutes.QueryParamLookupID)
	shop, tenantID, found := validLookupID(ctx, lookupID)
	if !found {
		httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "Unknown token"})
		return
	}

	connect, err := shop.Connection(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var req rates.ShopifyRatesRequest
	err = httputils.UnmarshalRequestBody(r, &req)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		go logLookupRequest(cli, shop.TenantID, req, connect, 0, err)
		return
	}

	ctx = viewer.NewContext(r.Context(), viewer.UserViewer{
		Role:   viewer.Anonymous,
		Tenant: tenantID,
	})

	shopifyItems := make([]shopifyItem, 0)
	for _, i := range req.Rate.Items {
		shopifyItems = append(shopifyItems, shopifyItem{
			ProductID: strconv.FormatInt(i.ProductID, 10),
			VariantID: strconv.FormatInt(i.VariantID, 10),
			Quantity:  i.Quantity,
			Price:     float64(i.Price) / 100,
		})
	}

	allVariants, totalProductsWeight, err := sortFetchOrderItems(ctx, shop, shopifyItems)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		go logLookupRequest(cli, shop.TenantID, req, connect, 0, err)
		return
	}

	countryRaw := req.Rate.Destination.Country
	lookupCountry, err := rawCountryToDelivrio(ctx, fmt.Sprintf("%s", *countryRaw))
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		go logLookupRequest(cli, shop.TenantID, req, connect, 0, err)
		return
	}

	hasCompanyField := false
	if req.Rate.Destination.CompanyName != nil && len(*req.Rate.Destination.CompanyName) > 0 {
		hasCompanyField = true
	}

	deliveryOptions, trackingID, err := filterByHypothesisTestAndCompanyField(ctx, time.Now(), connect, req.Rate.Items, req.Rate.Destination, hasCompanyField)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, err.Error())
		go logLookupRequest(cli, shop.TenantID, req, connect, 0, err)
		return
	}

	nearestToAddress := deliverypoints.DropPointLookupAddress{
		Address1: formatAddress(req.Rate.Destination),
		Zip:      formatPostalCode(req.Rate.Destination),
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
		go logLookupRequest(cli, shop.TenantID, req, connect, 0, err)
		return
	}

	httputils.JSONResponse(w, http.StatusOK, output)
	go logLookupRequest(cli, shop.TenantID, req, connect, len(output.Rates), nil)
	return

}

func formatPostalCode(req rates.OriginDestinationAddressRequest) string {
	if req.PostalCode == nil {
		return ""
	}
	return fmt.Sprintf("%s", *req.PostalCode)
}

func formatAddress(req rates.OriginDestinationAddressRequest) string {
	if req.Address1 == nil {
		return ""
	}
	return fmt.Sprintf("%s", *req.Address1)
}

func logLookupRequest(
	cli *ent.Client,
	tenantID pulid.ID,
	payload rates.ShopifyRatesRequest,
	connection *ent.Connection,
	deliveryOptionCount int,
	err error,
) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("LookupRequest logger: ", err)
		}
	}()

	ctx := viewer.NewBackgroundContext(context.Background())
	payloadInput := ""
	pay, err := json.Marshal(payload)
	if err != nil {
		payloadInput = fmt.Sprintf("could not marshal payload: %v", err)
	} else {
		payloadInput = string(pay)
	}
	create := cli.ConnectionLookup.Create().
		SetTenantID(tenantID).
		SetPayload(payloadInput).
		SetOptionsOutputCount(deliveryOptionCount)

	if connection != nil {
		create = create.SetConnections(connection)
	}
	if err != nil {
		create = create.SetError(err.Error())
	}

	err = create.Exec(ctx)
	if err != nil {
		log.Println("Lookup request logger: save: ", err)
	}

}

func waitForRates(ctx context.Context,
	trackingID *pulid.ID,
	deliveryOptions []*ent.DeliveryOption,
	allVariants []*deliveryoptions.ConstraintProductWeight,
	nearestToAddress deliverypoints.DropPointLookupAddress,
	totalWeight int,
) (*rates.ShopifyRateResponses, error) {
	outputRates := make(map[int][]rates.RateResponse)
	var m = sync.Mutex{}
	var wg = sync.WaitGroup{}
	var errors []error

	for _, do := range deliveryOptions {
		wg.Add(1) // Increment the wait group counter for each iteration

		fmt.Printf("Wait for: %v", do)
		go func(do *ent.DeliveryOption) {
			defer wg.Done() // Decrement the wait group counter when the goroutine completes
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered from panic in waitForRates: %v", r)
				}
			}()

			r, err := checkDeliveryOption(ctx, trackingID, do, allVariants, nearestToAddress, totalWeight)
			if err != nil {
				// TODO: Do something with these?
				errors = append(errors, err)
				log.Printf("async rate request failed: %v", err)
				return
			}
			m.Lock()
			defer m.Unlock()
			outputRates[do.SortOrder] = r
			fmt.Println(do.Name, do.SortOrder, r)

		}(do)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	output := &rates.ShopifyRateResponses{
		Rates: make([]rates.RateResponse, 0),
	}

	sortTracker := make([]int, 0)
	for ord, _ := range outputRates {
		sortTracker = append(sortTracker, ord)
	}
	slices.Sort(sortTracker)

	for _, ord := range sortTracker {

		output.Rates = append(output.Rates, outputRates[ord]...)
		// Shopify fixed issue and seems to allow unlimited**
		// **tested to 50, and beyond that there are UX issues
		if len(output.Rates) > 50 {
			output.Rates = output.Rates[0:49]
			break
		}
	}

	fmt.Println(output.Rates)

	return output, nil
}

func checkDeliveryOption(
	ctx context.Context,
	trackingID *pulid.ID,
	do *ent.DeliveryOption,
	allVariants []*deliveryoptions.ConstraintProductWeight,
	nearestToAddress deliverypoints.DropPointLookupAddress,
	totalWeight int,
) ([]rates.RateResponse, error) {

	ctx, span := tracer.Start(ctx, "checkDeliveryOption")
	defer span.End()

	span.SetAttributes(
		attribute.String("street", nearestToAddress.Address1),
		attribute.String("ZipCode", nearestToAddress.Zip),
		attribute.String("country", nearestToAddress.Country.Alpha2),
	)

	db := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)
	matches, price, err := deliveryoptions.DeliveryOptionMatches(ctx, do, nearestToAddress.Zip, allVariants, nearestToAddress.Country.ID)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	output := make([]rates.RateResponse, 0)

	if !matches {
		return output, nil
	}

	deliveryPointOptional, deliveryPointRequired, err := deliveryoptions.ServicePointConfig(ctx, do)
	if err != nil {
		return nil, err
	}

	log.Printf("%v: %v: %v\n", do.Name, deliveryPointOptional, deliveryPointRequired)

	minDeliveryWindow, maxDeliveryWindow := deliveryWindow(do.DeliveryEstimateFrom, do.DeliveryEstimateTo)

	if deliveryPointRequired || deliveryPointOptional {
		cb, err := do.QueryCarrierService().
			QueryCarrierBrand().
			Only(ctx)
		if err != nil {
			return nil, err
		}
		deliveryPointRates := make([]rates.RateResponse, 0)
		switch cb.InternalID {
		case carrierbrand.InternalIDBring:
			deliveryPointRates, err = fetchDeliveryPoints(
				ctx,
				trackingID,
				do,
				price,
				nearestToAddress,
				bringapis.DeliveryPoints,
				carrierbrand.InternalIDBring,
			)
			if err != nil {
				span.RecordError(err)
				return nil, err
			}
			break
		case carrierbrand.InternalIDDAO:
			deliveryPointRates, err = fetchDeliveryPoints(
				ctx,
				trackingID,
				do,
				price,
				nearestToAddress,
				daoapis.DeliveryPoints,
				carrierbrand.InternalIDDAO,
			)
			if err != nil {
				span.RecordError(err)
				return nil, err
			}
			break
		case carrierbrand.InternalIDPostNord:
			deliveryPointRates, err = fetchDeliveryPoints(
				ctx,
				trackingID,
				do,
				price,
				nearestToAddress,
				postnorddeliverypoints.DeliveryPoints,
				carrierbrand.InternalIDPostNord,
			)
			if err != nil {
				span.RecordError(err)
				return nil, err
			}
			break
		case carrierbrand.InternalIDGLS:
			deliveryPointRates, err = fetchDeliveryPoints(
				ctx,
				trackingID,
				do,
				price,
				nearestToAddress,
				glsdeliverypoints.DeliveryPoints,
				carrierbrand.InternalIDGLS,
			)
			if err != nil {
				span.RecordError(err)
				return nil, err
			}
			break
		}

		output = append(output, deliveryPointRates...)

	} else if do.ClickCollect {
		span.AddEvent("Click&Collect")
		ccLocations, err := do.QueryClickCollectLocation().
			WithAddress().
			All(ctx)
		if err != nil {
			return nil, err
		}

		closestZips := CCLocationsToClosest(ccLocations, nearestToAddress.Zip, do.ClickOptionDisplayCount)

		for _, l := range closestZips {
			serviceCode := fmt.Sprintf("%v-%v", do.ID.String(), l.ID)

			output = append(output, rates.RateResponse{
				ServiceName: fmt.Sprintf(
					"%v%v %v",
					// Start at 1 so all lines have same start
					// Only within price ordering
					strings.Repeat(invisibleOrderingCharacter, int(do.SortOrder)+1),
					do.Name,
					l.Name,
				),
				// WS & Shipmondo integrations are disabled when C&C enabled
				// Check hooks.
				ServiceCode: serviceCode,
				// Shopify wants floats represented as ints
				TotalPrice:      fmt.Sprintf("%v", int(price.Price*100)),
				Description:     fmt.Sprintf("%v %v", do.Description, l.AddressFormatted),
				Currency:        price.Currency.CurrencyCode.String(),
				MinDeliveryDate: minDeliveryWindow,
				MaxDeliveryDate: maxDeliveryWindow,
			})
		}

	} else {

		serviceCode := fmt.Sprintf("%v", do.ID.String())

		if trackingID != nil && *trackingID != "" {
			existing, err := db.HypothesisTestDeliveryOptionLookup.Create().
				SetDeliveryOption(do).
				SetHypothesisTestDeliveryOptionRequestID(*trackingID).
				SetTenantID(view.TenantID()).
				Save(ctx)
			if err != nil && !ent.IsConstraintError(err) {
				log.Println("Could not save hypothesis test lookup", err)
			} else if ent.IsConstraintError(err) {
				existing, err = db.HypothesisTestDeliveryOptionLookup.Query().
					Where(hypothesistestdeliveryoptionlookup.HasHypothesisTestDeliveryOptionRequestWith(hypothesistestdeliveryoptionrequest.ID(*trackingID))).
					Only(ctx)
				if err != nil {
					log.Println("Could not fetch hypothesis test lookup", err)
				} else {
					serviceCode = fmt.Sprintf("%v", existing.ID)
				}
			} else {
				serviceCode = fmt.Sprintf("%v", existing.ID)
			}
		}

		// No drop point
		externalCode, err := externalIntegrationCode(do, WebshipperDropPoint{}, "")
		if err != nil {
			log.Println("Could not construct external drop point", err)
		}
		if externalCode != nil {
			serviceCode = *externalCode
		}

		output = append(output, rates.RateResponse{
			ServiceName: fmt.Sprintf(
				"%v%v",
				// Start at 1 so all lines have same start
				// Only within price ordering
				strings.Repeat(invisibleOrderingCharacter, int(do.SortOrder)+1),
				do.Name,
			),
			// Should include parcel shop
			ServiceCode: serviceCode,
			// Shopify wants floats represented as ints
			TotalPrice:      fmt.Sprintf("%v", int(price.Price*100)),
			Description:     fmt.Sprintf("%v", do.Description),
			Currency:        price.Currency.CurrencyCode.String(),
			MinDeliveryDate: minDeliveryWindow,
			MaxDeliveryDate: maxDeliveryWindow,
		})
	}

	return output, nil
}

// Assumes no carrier delivers on Sundays
// Just a quick heuristic for now
func bumpPastSunday(initial time.Time) time.Time {
	if initial.Weekday() == time.Sunday {
		return initial.AddDate(0, 0, 1)
	}

	return initial
}

func CCLocationsToClosest(ccLocations []*ent.Location, zip string, count int) []LocationZip {
	allZips := make([]LocationZip, 0)
	for _, l := range ccLocations {
		allZips = append(allZips, LocationZip{
			ID:   l.ID,
			Zip:  l.Edges.Address.Zip,
			Name: l.Name,
			AddressFormatted: fmt.Sprintf(
				"%v %v %v %v",
				l.Edges.Address.AddressOne,
				l.Edges.Address.AddressTwo,
				l.Edges.Address.City,
				l.Edges.Address.Zip,
			),
		})
	}
	return closestXLocations(zip, allZips, count)
}

func ProductVariantFromShopifyID(ctx context.Context, shopifyVariantID string) (*ent.ProductVariant, error) {
	db := ent.FromContext(ctx)

	return db.ProductVariant.Query().
		Where(productvariant.ExternalIDEqualFold(shopifyVariantID)).
		Only(ctx)
}

func VariantToProductLookup(ctx context.Context, variant *ent.ProductVariant, units int, price float64) (*deliveryoptions.ConstraintProductWeight, error) {
	product, err := variant.QueryProduct().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query product")
	}

	tags, err := product.QueryProductTags().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query tags")
	}

	tagIDs := make([]pulid.ID, 0)
	for _, t := range tags {
		tagIDs = append(tagIDs, t.ID)
	}
	return &deliveryoptions.ConstraintProductWeight{
		WeightG:       *variant.WeightG,
		ProductTagIDs: tagIDs,
		// TODO: fix this when we have SKU
		SKU:       variant.EanNumber,
		UnitPrice: price,
		Units:     units,
	}, nil
}

func validLookupID(ctx context.Context, lookupID string) (*ent.ConnectionShopify, pulid.ID, bool) {
	db := ent.FromContext(ctx)
	shop, err := db.ConnectionShopify.Query().
		WithTenant().
		Where(connectionshopify.LookupKeyEqualFold(lookupID)).
		Only(ctx)
	if err != nil || shop == nil {
		return nil, pulid.MustNew(""), false
	}
	return shop, shop.Edges.Tenant.ID, true
}
