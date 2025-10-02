package ordersync

import (
	"bytes"
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrieradditionalservicegls"
	"delivrio.io/go/ent/carrieradditionalservicepostnord"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/hypothesistestdeliveryoptionlookup"
	"delivrio.io/go/ent/location"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/parcelshop"
	"delivrio.io/go/ent/parcelshopgls"
	"delivrio.io/go/ent/parcelshoppostnord"
	"delivrio.io/go/ent/product"
	"delivrio.io/go/ent/producttag"
	"delivrio.io/go/ent/productvariant"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/ent/workstation"
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/go/schema/hooks/history"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/go/seed"
	"delivrio.io/go/shopify"
	"delivrio.io/go/shopify/ordermodels"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/printerutils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
)

const fulfillmentOrderURL = `%v/admin/api/2024-04/orders/%v/fulfillment_orders.json`
const createFulfillmentURL = `%v/admin/api/2024-04/fulfillments.json`

func SyncOrderFulfilled(
	ctx context.Context,
	connection *ent.ConnectionShopify,
	orderExternalID string,
	shipment *ent.ShipmentParcel,
) []error {
	allErrors := make([]error, 0)

	updateURL := fmt.Sprintf(fulfillmentOrderURL, connection.StoreURL, orderExternalID)

	req, err := http.NewRequest(http.MethodGet, updateURL, nil)
	if err != nil {
		return []error{fmt.Errorf("fulfillment order request (%v): %w", orderExternalID, err)}
	}

	respBody, err := shopify.FireRequest(ctx, connection.APIKey, req, http.StatusOK)
	if err != nil {
		return []error{fmt.Errorf("could not fetch fulfillment order (%v): %w", orderExternalID, err)}
	}

	var foBody ordermodels.FulfillmentOrders
	err = json.Unmarshal(respBody, &foBody)
	if err != nil {
		return []error{fmt.Errorf("could not unmarshal fulfillment order (%v): %w", orderExternalID, err)}
	}

	orderLines, err := shipment.QueryColli().
		QueryOrderLines().
		WithProductVariant().
		All(ctx)
	if err != nil {
		return []error{fmt.Errorf("could not query order lines: %w", err)}
	}

	fulfilledOrderLines := make(map[string]int64)
	for _, ol := range orderLines {
		fulfilledOrderLines[ol.Edges.ProductVariant.ExternalID] = int64(ol.Units)
	}

	// In practice, there should only be 1 FO
	for _, fo := range foBody.FulfillmentOrders {

		foundOrderLines := make([]ordermodels.CreateFulfillmentOrderLineItem, 0)
		alreadyFulfilledOrderLines := make(map[string]bool)
		for _, li := range fo.LineItems {
			orderLineID := strconv.FormatInt(li.ID, 10)
			externalLineItemVariantID := strconv.FormatInt(li.VariantID, 10)
			quantityFulfilled := fulfilledOrderLines[externalLineItemVariantID]
			if quantityFulfilled > 0 && li.FulfillableQuantity == 0 {
				alreadyFulfilledOrderLines[orderLineID] = true
			} else if quantityFulfilled > 0 && li.FulfillableQuantity > 0 {
				// Should be an edge case that this is true, but
				// we want to prevent failure in case the two
				// systems are out of sync
				if quantityFulfilled > li.FulfillableQuantity {
					foundOrderLines = append(foundOrderLines, ordermodels.CreateFulfillmentOrderLineItem{
						ID:       li.ID,
						Quantity: li.Quantity,
					})
				} else {
					foundOrderLines = append(foundOrderLines, ordermodels.CreateFulfillmentOrderLineItem{
						ID:       li.ID,
						Quantity: quantityFulfilled,
					})
				}
			} else {
				// If quantityFulfilled == 0 (the remaining case), then we end up
				// marking as fulfilled and not firing anything to Shopify?
			}
		}

		// In case there was a network interruption before the Shopify response
		// could be saved
		if len(foundOrderLines) == 0 && len(alreadyFulfilledOrderLines) > 0 {
			err = shipment.Update().
				SetFulfillmentSyncedAt(time.Now()).
				Exec(ctx)
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("save sync fulfilled at time: %w", err))
			}
			continue
		}

		carrierBrand, err := shipment.QueryColli().
			QueryDeliveryOption().
			QueryCarrier().
			QueryCarrierBrand().
			Only(ctx)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("shopify: fulfillment sync: could not find carrier brand: %w", err))
			continue
		}

		shopifyCompany := ""
		switch carrierBrand.InternalID {
		case carrierbrand.InternalIDGLS:
			shopifyCompany = "GLS"
		case carrierbrand.InternalIDPostNord:
			shopifyCompany = "PostNord DK"
		case carrierbrand.InternalIDUSPS:
			shopifyCompany = "USPS"
		case carrierbrand.InternalIDEasyPost:
			// TODO: update when we support more Easy Post carriers
			shopifyCompany = "USPS"
		default:
			allErrors = append(allErrors, fmt.Errorf("shopify: fulfillment sync: unexpected carrier: %v", carrierBrand.InternalID))
			continue

		}

		fulfillment := ordermodels.CreateFulfillment{
			Fulfillment: ordermodels.Fulfillment{
				Message:        "Order fulfilled",
				NotifyCustomer: false,
				TrackingInfo: ordermodels.TrackingInfo{
					Company: shopifyCompany,
					Number:  shipment.ItemID,
				},
				LineItemsByFulfillmentOrder: []ordermodels.LineItemsByFulfillmentOrder{
					{
						FulfillmentOrderID:        fo.ID,
						FulfillmentOrderLineItems: foundOrderLines,
					},
				},
			},
		}

		postBody, err := json.Marshal(fulfillment)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("could not marshal fulfillment: %w", err))
			continue
		}

		createURL := fmt.Sprintf(createFulfillmentURL, connection.StoreURL)
		req, err := http.NewRequest(http.MethodPost, createURL, bytes.NewReader(postBody))
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("could not create fulfillment request: %w", err))
			continue
		}

		out, _ := httputil.DumpRequest(req, true)
		fmt.Println(string(out))

		_, err = shopify.FireRequest(ctx, connection.APIKey, req, http.StatusCreated)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("fire fulfillment request: %w", err))
			continue
		}

		err = shipment.Update().
			SetFulfillmentSyncedAt(time.Now()).
			Exec(ctx)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("save sync fulfilled at time: %w", err))
		}

	}

	return allErrors

}

func fetchShopifyOrdersPage(url string, key string) (*ordermodels.Orders, error) {
	time.Sleep(shopify.RequestWait(time.Now(), key))
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Shopify-Access-Token", key)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	shopify.RecordHeader(key, res)

	if res.StatusCode != 200 {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(string(msg))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var orders ordermodels.Orders
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil

}

func ProcessShopifyOrderSync(ctx context.Context, connectShopify *ent.ConnectionShopify, shopSyncFrom time.Time, lastSyncEvent time.Time, evt pulid.ID) {
	db := ent.FromContext(ctx)
	var sinceID uint64 = 0

	fetchURL := fmt.Sprintf(
		// When updating a fulfillment, we are updating the update_at value...
		"%s/admin/api/%s/orders.json?created_at_min%s&updated_at_min=%s&since_id=%v",
		connectShopify.StoreURL,
		apiDate,
		// UTC() is a hack because the Shopify API isn't working as documented???
		shopSyncFrom.UTC().Format(time.RFC3339),
		lastSyncEvent.UTC().Format(time.RFC3339),
		sinceID,
	)

	orders, err := fetchShopifyOrdersPage(fetchURL, connectShopify.APIKey)
	if err != nil {
		err := db.SystemEvents.Update().
			SetStatus(systemevents.StatusFail).
			SetDescription(fmt.Sprintf("Failed fetching orders after %v orders retrieved", 0)).
			SetData(fmt.Sprintf("%v -> %v", fetchURL, err.Error())).
			Where(systemevents.ID(evt)).
			Exec(ctx)
		if err != nil {
			if err != nil {
				log.Printf("processShopifyOrderSync: systemevent: %v", err)
			}
		}
		return
	}

	connect, err := connectShopify.QueryConnection().
		Only(ctx)
	if err != nil {
		if err != nil {
			log.Printf("processShopifyOrderSync: query connection: %v", err)
		}
		return
	}

	orderCount := len(orders.Orders)
	orderSaveErrors := make([]string, 0)

	for len(orders.Orders) > 0 {

		select {
		case <-ctx.Done():
			return
		default:
			for _, o := range orders.Orders {
				if o.ID > sinceID {
					sinceID = o.ID
				}

				err = saveShopifyOrder(ctx, connect, o)
				if err != nil {
					orderSaveErrors = append(orderSaveErrors, err.Error())
				}
			}
		}

		fetchURL := fmt.Sprintf(
			"%s/admin/api/%s/orders.json?created_at_min%s&updated_at_min=%s&since_id=%v",
			connectShopify.StoreURL,
			apiDate,
			// UTC() is a hack because the Shopify API isn't working as documented???
			shopSyncFrom.UTC().Format(time.RFC3339),
			lastSyncEvent.UTC().Format(time.RFC3339),
			sinceID,
		)
		orders, err = fetchShopifyOrdersPage(fetchURL, connectShopify.APIKey)
		if err != nil {
			err = db.SystemEvents.Update().
				SetStatus(systemevents.StatusFail).
				SetDescription(fmt.Sprintf("Failed fetching orders after %v orders retrieved", orderCount)).
				SetData(fmt.Sprintf("%v -> %v", fetchURL, err.Error())).
				Where(systemevents.ID(evt)).
				Exec(ctx)
			if err != nil {
				log.Println(2, err)
			}
			return
		}

		orderCount += len(orders.Orders)

	}

	if len(orderSaveErrors) > 0 {
		err = db.SystemEvents.Update().
			SetStatus(systemevents.StatusFail).
			SetDescription(fmt.Sprintf("Failed saving order after %v orders retrieved", orderCount)).
			SetData(fmt.Sprintf("%v", strings.Join(orderSaveErrors, ", "))).
			Where(systemevents.ID(evt)).
			Exec(ctx)
		if err != nil {
			log.Printf("processShopifyOrderSync: systemevent: %v", err)
		}
		return
	}

	err = db.SystemEvents.Update().
		SetStatus(systemevents.StatusSuccess).
		SetDescription(fmt.Sprintf("Successfully synced %v orders", orderCount)).
		Where(systemevents.ID(evt)).
		Exec(ctx)
	if err != nil {
		log.Printf("processShopifyOrderSync: systemevent: %v", err)
	}
}

type deliveryOptionServicePoint struct {
	DeliveryOption  *ent.DeliveryOption
	ServicePoint    *ent.ParcelShop
	Location        *ent.Location
	RequestTracking *ent.HypothesisTestDeliveryOptionRequest
}

func processShippingOptions(ctx context.Context, shippingLines []ordermodels.ShippingLine) ([]deliveryOptionServicePoint, error) {
	allShippingOptions := make([]deliveryOptionServicePoint, 0)
	for _, sl := range shippingLines {
		option, err := processShippingOption(ctx, sl)
		if err != nil {
			return nil, err
		}
		if option != nil {
			allShippingOptions = append(allShippingOptions, *option)
		}
	}
	return allShippingOptions, nil
}

func processShippingOption(ctx context.Context, sl ordermodels.ShippingLine) (*deliveryOptionServicePoint, error) {
	// Otherwise not a DELIVRIO rate
	if sl.Code == nil {
		return nil, nil
	}

	do, ps, loc, htr, err := shopifyDeliveryOption(ctx, *sl.Code)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Println("DELIVERY OPTION NOT FOUND" + *sl.Code)
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return &deliveryOptionServicePoint{
		DeliveryOption:  do,
		ServicePoint:    ps,
		Location:        loc,
		RequestTracking: htr,
	}, nil
}

// Requires-shipping = false products are not synced by default
func orderLineItemsRequireNoShipping(ctx context.Context, ord ordermodels.Order) (bool, error) {
	cli := ent.FromContext(ctx)

	for _, ol := range ord.LineItems {
		// TODO: optimize this if query count ever becomes an issue
		tagFound, err := cli.ProductVariant.Query().
			Where(productvariant.ExternalID(strconv.FormatUint(ol.VariantID, 10))).
			QueryProduct().
			Where(product.HasProductTagsWith(producttag.NameEqualFold("order-sync"))).
			Count(ctx)
		if err != nil {
			return false, err
		}

		if (ol.RequiresShipping != nil && *ol.RequiresShipping) || tagFound > 0 {
			return false, nil
		}
	}
	return true, nil
}

// processOrderTags accepts a comma+space separated list of tags and normalizes them to
// a lower case slice
func processOrderTags(shopifyTags string) []string {
	output := make([]string, 0)
	splitTags := strings.Split(shopifyTags, ",")
	for _, t := range splitTags {
		output = append(output, strings.ToLower(strings.TrimSpace(t)))
	}
	return output
}

func hasOrderTags(ctx context.Context, connect *ent.Connection, ord ordermodels.Order) (bool, bool, error) {

	shopifyConnect, err := connect.QueryConnectionShopify().
		Only(ctx)
	if err != nil {
		return false, false, err
	}

	if len(shopifyConnect.FilterTags) == 0 {
		return true, false, nil
	}

	orderTags := processOrderTags(ord.Tags)

	for _, tag := range shopifyConnect.FilterTags {
		// Order tags should already be lower
		if slices.Contains(orderTags, strings.ToLower(tag)) {
			return false, true, nil
		}
	}

	return false, false, nil

}

func saveShopifyOrder(ctx context.Context, connect *ent.Connection, ord ordermodels.Order) error {
	view := viewer.FromContext(ctx)
	db := ent.FromContext(ctx)

	xid := ord.ID

	orderAlreadyExists, err := db.Order.Query().
		Where(order.And(
			order.ExternalID(
				strconv.FormatUint(xid, 10),
			),
			order.HasConnectionWith(connection.ID(connect.ID)),
		)).
		Exist(ctx)
	if err != nil {
		return fmt.Errorf("save order shopify: query order: %w", err)
	}

	// We don't do updates since the order should not have been modified since
	// creation, and it disturbs the fulfillment collis

	noTagFilters, hasTags, err := hasOrderTags(ctx, connect, ord)
	if err != nil {
		return err
	}

	// Support for not syncing digital products since we don't handle non-fulfill-able products
	// like events
	shippingNotRequired, err := orderLineItemsRequireNoShipping(ctx, ord)
	if err != nil {
		return err
	}

	if orderAlreadyExists || shippingNotRequired || noTagFilters || !hasTags {
		return nil
	}

	allShippingOptions, err := processShippingOptions(ctx, ord.ShippingLines)
	if err != nil {
		return err
	}

	defaultDeliveryOption, err := connect.DefaultDeliveryOption(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("shopify: save order: default delivery option: %w", err)
	}

	tx, err := db.Tx(ctx)
	if err != nil {
		return utils.Rollback(tx, fmt.Errorf("shopify: save order: open tx: %w", err))
	}
	defer tx.Rollback()

	txCtx := ent.NewTxContext(ctx, tx)

	create, err := tx.Order.Create().
		SetOrderPublicID(ord.Name).
		SetExternalID(strconv.FormatUint(ord.ID, 10)).
		SetConnection(connect).
		SetStatus(order.StatusPending).
		SetTenantID(view.TenantID()).
		SetNillableCommentExternal(ord.Note).
		SetNoteAttributes(noteAttributes(ord)).
		Save(txCtx)
	if err != nil {
		return utils.Rollback(tx, fmt.Errorf("shopify: save order: create: %w", err))
	}

	var colliID pulid.ID
	createPackage, err := newPackageFromOrder(txCtx, create.ID, ord)
	if err != nil {
		return utils.Rollback(tx, fmt.Errorf("shopify: save order: new package: %w", err))
		// Only create package when shopify has address info
	} else if createPackage != nil {
		createPackage, historyDescription, err := configurePackage(txCtx, createPackage, create, allShippingOptions, defaultDeliveryOption)
		if err != nil {
			return err
		}

		pack, err := createPackage.Save(
			history.NewConfig(txCtx).
				SetDescription(historyDescription).
				SetOrigin(changehistory.OriginBackground).
				Ctx(),
		)
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("shopify: save order: create package: %w", err))
		}

		// Create remaining packages without order lines
		// Whereby order lines can be manually placed in these packages
		err = createAdditionalPackages(txCtx, create, ord, allShippingOptions, defaultDeliveryOption)
		if err != nil {
			return utils.Rollback(tx, err)
		}

		err = createOrderLines(txCtx, ord, pack)
		if err != nil {
			return utils.Rollback(tx, err)
		}
		colliID = pack.ID
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return processAutoPrint(ctx, connect, colliID)
}

func noteAttributes(ord ordermodels.Order) fieldjson.NoteAttributes {
	output := make(fieldjson.NoteAttributes)
	for _, na := range ord.NoteAttributes {
		output[na.Name] = na.Value
	}
	return output
}

func processAutoPrint(ctx context.Context, conn *ent.Connection, colliID pulid.ID) error {
	if !conn.AutoPrintParcelSlip {
		return nil
	}

	cli := ent.FromContext(ctx)
	col, err := cli.Colli.Query().
		Where(colli.ID(colliID)).
		All(ctx)
	if err != nil {
		return err
	}

	ws, err := cli.Workstation.Query().
		Where(workstation.AutoPrintReceiver(true)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	} else if ent.IsNotFound(err) {
		return nil
	}

	historyCtx := history.NewConfig(ctx).
		SetDescription(fmt.Sprintf("Auto-print parcel slip to %v", ws.Name)).
		SetOrigin(changehistory.OriginBackground).
		Ctx()

	err = printerutils.CreatePackingSlipPrintJobs(historyCtx, col, ws)
	if err != nil {
		return err
	}

	return nil
}

func configurePackage(ctx context.Context, createPackage *ent.ColliCreate, create *ent.Order, allShippingOptions []deliveryOptionServicePoint, defaultDeliveryOption *ent.DeliveryOption) (*ent.ColliCreate, string, error) {
	historyDescription := "Add colli from Shopify"

	if len(allShippingOptions) == 0 {
		return createPackage, historyDescription, nil
	}

	firstOption := allShippingOptions[0]

	if firstOption.DeliveryOption != nil {
		createPackage = createPackage.SetDeliveryOption(firstOption.DeliveryOption)
	} else if defaultDeliveryOption != nil {
		historyDescription = "Add colli from Shopify with default connection delivery option"
		createPackage = createPackage.SetDeliveryOption(defaultDeliveryOption)
	}

	if firstOption.ServicePoint != nil {
		createPackage = createPackage.SetParcelShop(firstOption.ServicePoint)
	}

	if firstOption.Location != nil {
		createPackage = createPackage.SetClickCollectLocation(firstOption.Location)
	}

	if firstOption.RequestTracking != nil {
		err := create.Update().
			SetHypothesisTestDeliveryOptionRequest(firstOption.RequestTracking).
			Exec(ctx)
		if err != nil {
			log.Println("ignored tracking error", err)
		}
	}

	return createPackage, historyDescription, nil
}

func createAdditionalPackages(ctx context.Context, create *ent.Order, ord ordermodels.Order, allShippingOptions []deliveryOptionServicePoint, defaultDeliveryOption *ent.DeliveryOption) error {
	if len(allShippingOptions) <= 1 {
		return nil
	}

	for sli, sl := range allShippingOptions {
		if sli <= 1 {
			continue
		}

		err := createAdditionalPackage(ctx, create, ord, sl, defaultDeliveryOption)
		if err != nil {
			return fmt.Errorf("shopify: save order: create additional package: %w", err)
		}
	}

	return nil
}

func createAdditionalPackage(ctx context.Context, create *ent.Order, ord ordermodels.Order, sl deliveryOptionServicePoint, defaultDeliveryOption *ent.DeliveryOption) error {
	createExtraPack, err := newPackageFromOrder(ctx, create.ID, ord)
	if err != nil {
		return fmt.Errorf("shopify: save order: new package: %w", err)
	}

	historyDescription := "Add colli from Shopify"

	if sl.DeliveryOption != nil {
		createExtraPack = createExtraPack.SetDeliveryOption(sl.DeliveryOption)
	} else if defaultDeliveryOption != nil {
		historyDescription = "Add colli from Shopify with default connection delivery option"
		createExtraPack = createExtraPack.SetDeliveryOption(defaultDeliveryOption)
	}

	if sl.ServicePoint != nil {
		createExtraPack = createExtraPack.SetParcelShop(sl.ServicePoint)
	}

	if sl.Location != nil {
		createExtraPack = createExtraPack.SetClickCollectLocation(sl.Location)
	}

	_, err = createExtraPack.Save(
		history.NewConfig(ctx).
			SetDescription(historyDescription).
			SetOrigin(changehistory.OriginBackground).
			Ctx(),
	)
	if err != nil {
		return fmt.Errorf("shopify: save order: create extra package: %w", err)
	}

	return nil
}

func createOrderLines(ctx context.Context, ord ordermodels.Order, pack *ent.Colli) error {
	for _, li := range ord.LineItems {
		err := createOrderLine(ctx, li, ord, pack)
		if err != nil {
			return err
		}
	}
	return nil
}

func createOrderLine(ctx context.Context, li ordermodels.LineItem, ord ordermodels.Order, pack *ent.Colli) error {
	tx := ent.TxFromContext(ctx)
	shopifyLineItemID := strconv.FormatUint(li.ID, 10)
	price, err := strconv.ParseFloat(li.Price, 64)
	if err != nil {
		return fmt.Errorf("shopify: save order: parse float: %w", err)
	}

	quantity := li.Quantity - orderLineHasRefundedQuantity(li.ID, ord)
	if quantity < 0 {
		quantity = 0
	}

	shopifyVariantID := strconv.FormatUint(li.VariantID, 10)
	product, err := tx.ProductVariant.Query().
		Where(productvariant.ExternalIDEqualFold(shopifyVariantID)).
		OnlyID(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Printf("shopify: save order: query product variant (order: %v): skipping non-existing product", ord.ID)
			return nil
		}
		return fmt.Errorf("shopify: save order: query product variant (%v): %w", ord.ID, err)
	}

	cur, err := utils.ToCurrency(ctx, tx.Client(), li.PriceSet.ShopMoney.CurrencyCode)
	if err != nil {
		return fmt.Errorf("shopify: save order (%v): currency lookup: %w", ord.ID, err)
	}

	totalLineDiscount := calculateTotalLineDiscount(li.DiscountAllocations)

	err = tx.OrderLine.Create().
		SetExternalID(shopifyLineItemID).
		SetUnits(quantity).
		SetUnitPrice(price).
		SetDiscountAllocationAmount(totalLineDiscount).
		SetColli(pack).
		SetCurrency(cur).
		SetProductVariantID(product).
		SetTenantID(viewer.FromContext(ctx).TenantID()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("shopify: save order: create order line: %w", err)
	}

	return nil
}

func calculateTotalLineDiscount(discountAllocations []ordermodels.DiscountAllocation) float64 {
	var totalLineDiscount float64
	for _, d := range discountAllocations {
		amount, err := strconv.ParseFloat(d.Amount, 64)
		if err != nil {
			log.Printf("Error parsing discount amount: %v", err)
			continue
		}
		totalLineDiscount += amount
	}
	return totalLineDiscount
}

func shopifyDeliveryOption(ctx context.Context, input string) (*ent.DeliveryOption, *ent.ParcelShop, *ent.Location, *ent.HypothesisTestDeliveryOptionRequest, error) {
	pieces := strings.Split(input, "-")

	if len(pieces) <= 0 {
		return nil, nil, nil, nil, fmt.Errorf("shipping method not defined in Shopify")
	}

	if len(pieces) != 1 && len(pieces) != 2 {
		return nil, nil, nil, nil, fmt.Errorf("shipping method syntax does not match expected")
	}

	tx := ent.FromContext(ctx)

	var htr *ent.HypothesisTestDeliveryOptionRequest

	deliveryOptionID := pulid.ID(pieces[0])
	if val, err := pulid_server_prefix.IDToType(ctx, deliveryOptionID); err == nil && val == hypothesistestdeliveryoptionlookup.Table {
		htl, err := tx.HypothesisTestDeliveryOptionLookup.Query().
			WithHypothesisTestDeliveryOptionRequest().
			WithDeliveryOption().
			Where(hypothesistestdeliveryoptionlookup.ID(deliveryOptionID)).
			Only(ctx)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		deliveryOptionID = htl.Edges.DeliveryOption.ID
		htr = htl.Edges.HypothesisTestDeliveryOptionRequest
	} else {
		// This true??
		log.Printf("Could not find lookup %v, %v", deliveryOptionID, err)
	}

	do, err := tx.DeliveryOption.Query().
		WithCarrier().
		Where(deliveryoption.ID(deliveryOptionID)).
		Only(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	cb, err := do.Edges.Carrier.QueryCarrierBrand().
		Only(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch cb.InternalID {
	case carrierbrand.InternalIDPostNord:
		servicePointRequired := false

		// Perhaps a hook or also validating country?
		servicePointService, err := do.QueryCarrierService().
			QueryCarrierServicePostNord().
			QueryCarrierAddServPostNord().
			Where(carrieradditionalservicepostnord.And(
				carrieradditionalservicepostnord.InternalIDEQ("optional_service_point"),
			)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, nil, nil, nil, err
		} else if ent.IsNotFound(err) {
			servicePointRequired = false
		} else {
			servicePointRequired = servicePointService.Mandatory
		}

		if len(pieces) == 2 && servicePointService != nil {
			pnPUDOID := pieces[1]
			ps, err := tx.ParcelShop.Query().Where(
				parcelshop.HasParcelShopPostNordWith(parcelshoppostnord.Pudoid(pnPUDOID)),
			).Only(ctx)
			if err != nil {
				// Should not be not-found at this point
				// force manual action by not setting an invalid delivery option
				// with the required service point
				return nil, nil, nil, nil, err
			}
			return do, ps, nil, htr, nil
		} else if servicePointRequired {
			return nil, nil, nil, nil, fmt.Errorf("missing required PN service point ID from Shopify")
		} else if do.ClickCollect {
			// Should have a hook to exclude all service point services from being C&C
			locationID := pulid.ID(pieces[1])
			loc, err := tx.Location.Query().
				Where(location.ID(locationID)).
				Only(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return nil, nil, nil, nil, err
			}
			return do, nil, loc, nil, err
		}

		return do, nil, nil, htr, nil
	case carrierbrand.InternalIDGLS:
		servicePointRequired := false

		// Perhaps a hook or also validating country?
		servicePointService, err := do.QueryCarrierService().
			QueryCarrierServiceGLS().
			QueryCarrierAdditionalServiceGLS().
			Where(carrieradditionalservicegls.And(
				carrieradditionalservicegls.InternalIDEQ(seed.GLSDeliveryPointServiceInternalID),
			)).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, nil, nil, nil, err
		} else if ent.IsNotFound(err) {
			servicePointRequired = false
		} else {
			servicePointRequired = servicePointService.Mandatory
		}

		if len(pieces) == 2 && servicePointService != nil {
			glsID := pieces[1]
			ps, err := tx.ParcelShop.Query().Where(
				parcelshop.HasParcelShopGLSWith(parcelshopgls.GLSParcelShopIDEQ(glsID)),
			).Only(ctx)
			if err != nil {
				// Should not be not-found at this point
				// force manual action by not setting an invalid delivery option
				// with the required service point
				return nil, nil, nil, nil, err
			}
			return do, ps, nil, htr, nil
		} else if servicePointRequired {
			return nil, nil, nil, nil, fmt.Errorf("missing required GLS service point ID from Shopify")
		} else if do.ClickCollect {
			locationID := pulid.ID(pieces[1])
			loc, err := tx.Location.Query().
				Where(location.ID(locationID)).
				Only(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return nil, nil, nil, nil, err
			}
			return do, nil, loc, nil, err
		}

		return do, nil, nil, htr, nil
	case carrierbrand.InternalIDUSPS:
		// Service point delivery not supported
		return do, nil, nil, htr, nil
	}

	log.Println("unsupported carrier for Shopify order product sync")

	return nil, nil, nil, nil, fmt.Errorf("unsupported carrier for Shopify order product sync")

}

func newPackageFromOrder(ctx context.Context, orderID pulid.ID, ord ordermodels.Order) (*ent.ColliCreate, error) {
	view := viewer.FromContext(ctx)
	tx := ent.TxFromContext(ctx)

	colliCreate := tx.Colli.Create().
		SetOrderID(orderID).
		SetStatus(colli.StatusPending).
		SetTenantID(view.TenantID())

	if ord.ShippingAddress != nil {
		recipientAddress, err := recipientFromAddress(ctx, ord.Email, *ord.ShippingAddress)
		if err != nil {
			return nil, fmt.Errorf("recipientFromAddress: shipping: %w", err)
		}
		colliCreate = colliCreate.SetRecipient(recipientAddress)
	} else if ord.BillingAddress != nil {
		// Support for digital only orders
		recipientAddress, err := recipientFromAddress(ctx, ord.Email, *ord.BillingAddress)
		if err != nil {
			return nil, fmt.Errorf("recipientFromAddress: billing: %w", err)
		}
		colliCreate = colliCreate.SetRecipient(recipientAddress)
	} else {
		recipientAddress, err := recipientFromAddress(ctx, ord.Email, ordermodels.Address{
			CountryName: "DK",
		})
		if err != nil {
			return nil, fmt.Errorf("recipientFromAddress: empty: %w", err)
		}
		colliCreate = colliCreate.SetRecipient(recipientAddress)
	}

	senderLocation, err := tx.Connection.Query().
		QuerySenderLocation().
		WithAddress().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("new package: from order: %w", err)
	}

	senderLocationCountry, err := senderLocation.Edges.Address.QueryCountry().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("new package: from order: %w", err)
	}

	createColliSenderAddress, err := senderLocation.Edges.Address.CloneEntity(tx).
		SetCountryID(senderLocationCountry.ID).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("new package: from order: %w", err)
	}

	return colliCreate.SetSender(createColliSenderAddress), nil
}

func recipientFromAddress(ctx context.Context, email string, addr ordermodels.Address) (*ent.Address, error) {
	view := viewer.FromContext(ctx)
	tx := ent.TxFromContext(ctx)

	recipientCountry, err := tx.Country.Query().
		Where(country.Alpha2EqualFold(*addr.CountryCode)).
		OnlyID(ctx)
	if err != nil {
		return nil, err
	}

	return tx.Address.Create().
		SetAddressOne(addr.Address1).
		// TODO: ensure nulls don't break anything
		SetAddressTwo(addr.Address2).
		SetCompany(addr.Company).
		SetZip(addr.Zip).
		SetCity(addr.City).
		SetState(addr.ProvinceCode).
		SetFirstName(addr.FirstName).
		SetLastName(addr.LastName).
		SetEmail(email).
		SetPhoneNumber(addr.Phone).
		SetCountryID(recipientCountry).
		SetTenantID(view.TenantID()).
		Save(ctx)
}

func orderLineHasRefundedQuantity(shopifyOrderLineID uint64, order ordermodels.Order) int {
	output := 0
	for _, r := range order.Refunds {
		for _, rli := range r.RefundLineItems {
			if rli.LineItemID == shopifyOrderLineID {
				output += rli.Quantity
			}
		}
	}

	return output
}
