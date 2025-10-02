package productsync

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/countryharmonizedcode"
	"delivrio.io/go/ent/inventoryitem"
	"delivrio.io/go/ent/product"
	"delivrio.io/go/ent/productimage"
	"delivrio.io/go/ent/producttag"
	"delivrio.io/go/ent/productvariant"
	"delivrio.io/go/ent/systemevents"
	"delivrio.io/go/shopify"
	"delivrio.io/go/shopify/productmodels"
	"delivrio.io/go/utils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"errors"
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

const apiDate = "2024-04"

func RealTimeShopifyProductSync(ctx context.Context, storeURL string, apiKey string, productIDs []string) error {

	if len(productIDs) == 0 {
		return nil
	}

	fetchURL := fmt.Sprintf(
		"%s/admin/api/%s/products.json?ids=%s",
		storeURL,
		apiDate,
		strings.Join(productIDs, ","),
	)

	products, err := fetchShopifyProductsPage(fetchURL, apiKey)
	if err != nil {
		return err
	}

	for _, p := range products.Products {
		err = saveShopifyProduct(ctx, storeURL, apiKey, p)
		if err != nil {
			return err
		}
	}

	return nil

}

func ProcessShopifyProductSync(ctx context.Context, connection ent.ConnectionShopify, lastSyncEvent time.Time, evt pulid.ID) {
	db := ent.FromContext(ctx)
	var sinceID uint64 = 0

	fetchURL := fmt.Sprintf(
		"%s/admin/api/%s/products.json?updated_at_min=%s&since_id=%v",
		connection.StoreURL,
		apiDate,
		// UTC() is a hack because the Shopify API isn't working as documented???
		lastSyncEvent.UTC().Format(time.RFC3339),
		sinceID,
	)

	products, err := fetchShopifyProductsPage(fetchURL, connection.APIKey)
	if err != nil {
		err := db.SystemEvents.Update().
			SetStatus(systemevents.StatusFail).
			SetDescription(fmt.Sprintf("Failed fetching products after %v products retrieved", 0)).
			SetData(fmt.Sprintf("%v -> %v", fetchURL, err.Error())).
			Where(systemevents.ID(evt)).
			Exec(ctx)
		if err != nil {
			fmt.Println(4, err)
		}
		return
	}

	productCount := len(products.Products)

	for len(products.Products) > 0 {

		select {
		case <-ctx.Done():
			return
		default:
			for _, p := range products.Products {
				if p.ID > sinceID {
					sinceID = p.ID
				}

				err = saveShopifyProduct(ctx, connection.StoreURL, connection.APIKey, p)
				if err != nil {
					err = db.SystemEvents.Update().
						SetStatus(systemevents.StatusFail).
						SetDescription(fmt.Sprintf("Failed saving product after %v products retrieved", productCount)).
						SetData(fmt.Sprintf("%v", err.Error())).
						Where(systemevents.ID(evt)).
						Exec(ctx)
					if err != nil {
						log.Printf("error saving system event: %v", err)
					}
					return
				}

			}
		}

		fetchURL := fmt.Sprintf(
			"%s/admin/api/%s/products.json?updated_at_min=%s&since_id=%v",
			connection.StoreURL,
			apiDate,
			// UTC() is a hack because the Shopify API isn't working as documented???
			lastSyncEvent.UTC().Format(time.RFC3339),
			sinceID,
		)
		products, err = fetchShopifyProductsPage(fetchURL, connection.APIKey)
		if err != nil {
			err = db.SystemEvents.Update().
				SetStatus(systemevents.StatusFail).
				SetDescription(fmt.Sprintf("Failed fetching products after %v products retrieved", productCount)).
				SetData(fmt.Sprintf("%v -> %v", fetchURL, err.Error())).
				Where(systemevents.ID(evt)).
				Exec(ctx)
			if err != nil {
				log.Printf("error saving system event: %v", err)
			}
			return
		}

		productCount += len(products.Products)

	}

	err = db.SystemEvents.Update().
		SetStatus(systemevents.StatusSuccess).
		SetDescription(fmt.Sprintf("Successfully synced %v products", productCount)).
		Where(systemevents.ID(evt)).
		Exec(ctx)
	if err != nil {
		log.Printf("error saving system event: %v", err)
	}

}

func inventoryItemURL(storeURL string, inventoryItemShopifyID []int64) (*url.URL, error) {
	u, err := url.Parse(storeURL)
	if err != nil {
		return nil, err
	}

	allIDs := make([]string, 0)
	for _, id := range inventoryItemShopifyID {
		allIDs = append(allIDs, strconv.FormatInt(id, 10))
	}

	formValues := &url.Values{}
	formValues.Set("limit", "250")
	formValues.Set("ids", strings.Join(allIDs, ","))
	u.RawQuery = formValues.Encode()
	u = u.JoinPath(fmt.Sprintf(`/admin/api/%s/inventory_items.json`, apiDate))

	return u, nil
}

func fetchInventoryItems(ctx context.Context, storeURL string, key string, variants []productmodels.Variant) (map[int64]productmodels.InventoryItem, error) {

	inventoryToLookup := make([]int64, 0)
	for _, v := range variants {
		if v.InventoryItemID != nil && *v.InventoryItemID > 0 {
			inventoryToLookup = append(inventoryToLookup, *v.InventoryItemID)
		}
	}

	if len(inventoryToLookup) > 250 {
		return nil, fmt.Errorf("attempting to lookup >250 inventory; pagination not implemented")
	}

	u, err := inventoryItemURL(storeURL, inventoryToLookup)
	if err != nil {
		return nil, err
	}

	var reqOutput productmodels.InventoryItems
	err = genericShopifyRequest(u.String(), key, &reqOutput)
	if err != nil {
		return nil, err
	}

	output := make(map[int64]productmodels.InventoryItem)
	for _, ii := range reqOutput.InventoryItems {
		output[ii.ID] = ii
	}

	return output, nil
}

func saveInventoryItem(ctx context.Context, variantID pulid.ID, ii productmodels.InventoryItem) error {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	existing, err := tx.InventoryItem.Query().
		Where(inventoryitem.HasProductVariantWith(productvariant.ID(variantID))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	var coo *pulid.ID
	if ii.CountryCodeOfOrigin != nil {
		c, err := tx.Country.Query().
			Where(country.Alpha2EqualFold(*ii.CountryCodeOfOrigin)).
			Only(ctx)
		if err != nil {
			return err
		}
		coo = &c.ID
	}

	var iiID pulid.ID
	if existing == nil {
		saved, err := tx.InventoryItem.Create().
			SetTenantID(view.TenantID()).
			SetNillableCountryOfOriginID(coo).
			SetNillableCode(ii.HarmonizedSystemCode).
			SetNillableSku(ii.Sku).
			SetProductVariantID(variantID).
			Save(ctx)
		if err != nil {
			return err
		}
		iiID = saved.ID
	} else {
		err := tx.InventoryItem.Update().
			SetTenantID(view.TenantID()).
			SetNillableCountryOfOriginID(coo).
			SetNillableCode(ii.HarmonizedSystemCode).
			SetNillableSku(ii.Sku).
			Where(inventoryitem.HasProductVariantWith(productvariant.ID(variantID))).
			// Don't set variant again since it's 1:1
			Exec(ctx)
		if err != nil {
			return err
		}
		iiID = existing.ID
	}

	_, err = tx.CountryHarmonizedCode.Delete().
		Where(countryharmonizedcode.HasInventoryItemWith(inventoryitem.ID(iiID))).
		Exec(ctx)
	if err != nil {
		return err
	}

	allHSCountry := make([]*ent.CountryHarmonizedCodeCreate, 0)
	for _, hs := range ii.CountryHarmonizedSystemCodes {
		c, err := tx.Country.Query().
			Where(country.Alpha2EqualFold(hs.CountryCode)).
			Only(ctx)
		if err != nil {
			return err
		}

		allHSCountry = append(allHSCountry,
			tx.CountryHarmonizedCode.Create().
				SetCode(hs.HarmonizedSystemCode).
				SetCountry(c).
				SetTenantID(view.TenantID()))
	}

	return tx.CountryHarmonizedCode.CreateBulk(allHSCountry...).
		Exec(ctx)

}

func saveShopifyProduct(ctx context.Context, storeURL string, key string, prod productmodels.Product) error {

	view := viewer.FromContext(ctx)
	db := ent.FromContext(ctx)

	inventoryItems, err := fetchInventoryItems(ctx, storeURL, key, prod.Variants)
	if err != nil {
		return fmt.Errorf("shopify: save product: fetch inventory: %w", err)
	}

	tx, err := db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("shopify: save product: %w", err)
	}
	defer tx.Rollback()

	ctx = ent.NewTxContext(ctx, tx)

	status := product.StatusActive
	if *prod.Status == "archived" {
		status = product.StatusArchived
	} else if *prod.Status == "draft" {
		status = product.StatusDraft
	}

	xid := prod.ID
	tags, err := shopifyTagsToDelivrioTags(ctx, prod.Tags)
	if err != nil {
		return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
	}

	current, err := tx.Product.Query().
		Where(product.ExternalID(strconv.FormatUint(prod.ID, 10))).
		All(ctx)
	if err != nil {
		return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
	}

	var prodID pulid.ID
	if len(current) == 0 {
		create, err := tx.Product.Create().
			SetStatus(status).
			SetNillableBodyHTML(prod.BodyHTML).
			SetNillableCreatedAt(prod.CreatedAt).
			SetExternalID(strconv.FormatUint(xid, 10)).
			SetTitle(fmt.Sprintf("%v", func() string {
				if prod.Title != nil {
					return *prod.Title
				}
				return "<unknown>"
			}())).
			AddProductTags(tags...).
			SetTenantID(view.TenantID()).
			Save(ctx)
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
		}

		prodID = create.ID
	} else {
		_, err := tx.Product.Update().
			SetStatus(status).
			SetNillableBodyHTML(prod.BodyHTML).
			SetTitle(fmt.Sprintf("%v", func() string {
				if prod.Title != nil {
					return *prod.Title
				}
				return "<unknown>"
			}())).
			ClearProductTags().
			AddProductTags(tags...).
			SetTenantID(view.TenantID()).
			Where(product.ID(current[0].ID)).
			Save(ctx)
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
		}

		prodID = current[0].ID
	}

	currentVariant, err := tx.ProductVariant.Query().
		Where(productvariant.HasProductWith(product.ID(prodID))).
		All(ctx)
	if err != nil {
		return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
	}

	allVariantXIDs := make(map[string]pulid.ID)

	for _, v := range currentVariant {
		allVariantXIDs[strings.ToLower(v.ExternalID)] = v.ID
	}

	for _, v := range prod.Variants {

		currentVariantID := allVariantXIDs[strings.ToLower(strconv.FormatUint(v.ID, 10))]
		if len(currentVariantID.String()) > 0 {
			_, err := tx.ProductVariant.Update().
				SetNillableDescription(v.Title).
				SetWeightG(int(*v.Grams)).
				Where(productvariant.ID(currentVariantID)).
				Save(ctx)
			if err != nil {
				return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
			}

		} else {
			pv, err := tx.ProductVariant.Create().
				SetExternalID(strconv.FormatUint(v.ID, 10)).
				SetNillableDescription(v.Title).
				SetWeightG(int(*v.Grams)).
				SetProductID(prodID).
				SetTenantID(view.TenantID()).
				Save(ctx)
			if err != nil {
				return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
			}
			currentVariantID = pv.ID
		}

		if v.InventoryItemID != nil {
			if ii, ok := inventoryItems[*v.InventoryItemID]; ok {
				err = saveInventoryItem(ctx, currentVariantID, ii)
				if err != nil {
					return utils.Rollback(tx, fmt.Errorf("shopify: save product: inventory: %w", err))
				}
			}
		}

	}

	for _, i := range prod.Images {

		prodXID := fmt.Sprintf("%v", strconv.FormatInt(*i.ProductID, 10))
		prod, err := tx.Product.Query().
			Where(product.ExternalID(prodXID)).
			Only(ctx)
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("shopify: save product image: query product ID: %w", err))
		}

		varXIDs := make([]pulid.ID, 0)
		for _, vid := range i.VariantIDS {
			varXIDs = append(varXIDs, pulid.ID(strconv.FormatInt(vid, 10)))
		}

		variantIDs, err := tx.ProductVariant.Query().
			Where(productvariant.IDIn(varXIDs...)).
			All(ctx)
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("shopify: save product image: query variant: %w", err))
		}

		xid := fmt.Sprintf("%v", strconv.FormatInt(*i.ID, 10))

		err = tx.ProductImage.Create().
			SetExternalID(xid).
			SetURL(*i.Src).
			SetProduct(prod).
			// TODO: Check these gets cleared on upsert
			AddProductVariant(variantIDs...).
			SetTenantID(view.TenantID()).
			OnConflictColumns(productimage.FieldTenantID, productimage.FieldExternalID).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return utils.Rollback(tx, fmt.Errorf("shopify: save product: %w", err))
		}

	}

	return tx.Commit()

}

func shopifyTagsToDelivrioTags(ctx context.Context, tags string) ([]*ent.ProductTag, error) {
	tx := ent.TxFromContext(ctx)

	if strings.TrimSpace(tags) == "" {
		return make([]*ent.ProductTag, 0), nil
	}

	view := viewer.FromContext(ctx)
	tagsSplit := strings.Split(tags, ",")

	allTags, err := tx.ProductTag.Query().
		All(ctx)
	if err != nil {
		return nil, err
	}

	existingTagsLookup := make(map[string]*ent.ProductTag)
	for _, t := range allTags {
		clean := strings.TrimSpace(strings.ToLower(t.Name))
		existingTagsLookup[clean] = t
	}

	output := make([]*ent.ProductTag, 0)

	for _, t := range tagsSplit {

		clean := strings.TrimSpace(strings.ToLower(t))
		if _, ok := existingTagsLookup[clean]; ok {
			output = append(output, existingTagsLookup[clean])
		} else {
			tagID, err := tx.ProductTag.Create().
				SetName(clean).
				SetTenantID(view.TenantID()).
				OnConflictColumns(producttag.FieldTenantID, producttag.FieldName).
				UpdateName().
				ID(ctx)
			if err != nil {
				return nil, err
			}

			tag, err := tx.ProductTag.Query().
				Where(producttag.ID(tagID)).
				Only(ctx)
			if err != nil {
				return nil, err
			}

			existingTagsLookup[clean] = tag
			output = append(output, tag)
		}

	}

	return output, nil

}

func genericShopifyRequest(url string, key string, bodyOutput interface{}) error {
	sleep := shopify.RequestWait(time.Now(), key)
	time.Sleep(sleep)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Shopify-Access-Token", key)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	dp, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(dp))

	resDP, _ := httputil.DumpResponse(res, true)
	fmt.Println(string(resDP))

	shopify.RecordHeader(key, res)

	if res.StatusCode != 200 {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return errors.New(string(msg))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &bodyOutput)
	if err != nil {
		return err
	}

	return nil

}

func fetchShopifyProductsPage(url string, key string) (productmodels.Products, error) {
	time.Sleep(shopify.RequestWait(time.Now(), key))
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return productmodels.Products{}, err
	}
	req.Header.Set("X-Shopify-Access-Token", key)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return productmodels.Products{}, err
	}
	defer res.Body.Close()

	shopify.RecordHeader(key, res)

	if res.StatusCode != 200 {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return productmodels.Products{}, err
		}

		return productmodels.Products{}, errors.New(string(msg))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return productmodels.Products{}, err
	}

	var products productmodels.Products
	err = json.Unmarshal(body, &products)
	if err != nil {
		return productmodels.Products{}, err
	}

	return products, nil

}
