package dlvshopify

import (
	"bytes"
	"context"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/shopify"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/connectionshopify"
	"delivrio.io/go/ent/hook"
	"entgo.io/ent"
	"github.com/google/uuid"
)

const ShopifyLookupEndpoint = `/api/shopify-lookup`

// Handle globals better
var BaseURL *url.URL

type FetchCarrierServices struct {
	CarrierServices []CarrierService `json:"carrier_services,omitempty"`
}

type CarrierService struct {
	ID                 *int64  `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Active             *bool   `json:"active,omitempty"`
	ServiceDiscovery   *bool   `json:"service_discovery,omitempty"`
	CarrierServiceType *string `json:"carrier_service_type,omitempty"`
	AdminGraphqlAPIID  *string `json:"admin_graphql_api_id,omitempty"`
	Format             *string `json:"format,omitempty"`
	CallbackURL        *string `json:"callback_url,omitempty"`
}

const DelivrioCarrierConnectionName = `DELIVRIO`

func CreateShopifyConnection() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ConnectionShopifyFunc(func(ctx context.Context, m *ent2.ConnectionShopifyMutation) (ent.Value, error) {

			_, exists := m.ID()
			if !exists {
				return nil, errors.New("missing ConnectionShopify ID")
			}

			/*			shop, err := m.Client().ConnectionShopify.Query().Where(connectionshopify.ID(id)).Only(ctx)
						if err != nil {
							return nil, err
						}*/
			// Causing seeding issues
			/*			shopifyBaseURL, _ := m.StoreURL()
						shopifyAPIKey, _ := m.APIKey()
						lookupKey := uuid.New().String()

						m.SetLookupKey(lookupKey)

						err := handleShopifyIntegration(ctx, shopifyBaseURL, shopifyAPIKey, lookupKey)
						if err != nil {
							return nil, err
						}*/

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			return v, err
		})
	}

	return hook.On(hk, ent.OpCreate)
}

func UpdateShopifyConnection() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ConnectionShopifyFunc(func(ctx context.Context, m *ent2.ConnectionShopifyMutation) (ent.Value, error) {
			ids, err := m.IDs(ctx)
			if err != nil {
				return nil, err
			}

			if len(ids) != 1 {
				return nil, errors.New("expected to update exactly 1 shopify connection")
			}

			mutatedFields := m.Fields()
			shopAPIUpdated := false
			for _, f := range mutatedFields {
				if f == connectionshopify.FieldStoreURL {
					shopAPIUpdated = true
				}
			}

			if !shopAPIUpdated {
				return next.Mutate(ctx, m)
			}

			shopifyBaseURL, _ := m.StoreURL()
			shopifyAPIKey, _ := m.APIKey()
			lookupKey := uuid.New().String()
			m.SetLookupKey(lookupKey)

			rateIntegration, exists := m.RateIntegration()
			if exists && rateIntegration {
				err = handleShopifyIntegration(ctx, shopifyBaseURL, shopifyAPIKey, lookupKey)
				if err != nil {
					return nil, err
				}
			}

			v, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			return v, err
		})
	}

	return hook.On(hk, ent.OpUpdate|ent.OpUpdateOne)
}

func handleShopifyIntegration(ctx context.Context, shopifyBaseURL string, shopifyAPIKey string, lookupKey string) error {

	requestURL := fmt.Sprintf("%s/admin/api/2024-04/carrier_services.json", shopifyBaseURL)
	time.Sleep(shopify.RequestWait(time.Now(), shopifyAPIKey))
	client := &http.Client{
		Timeout: time.Second * 4,
	}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Shopify-Access-Token", shopifyAPIKey)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 200 {

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var output FetchCarrierServices
		err = json.Unmarshal(resBody, &output)
		if err != nil {
			return err
		}

		var foundID int64 = 0
		for _, o := range output.CarrierServices {
			if *o.Name == DelivrioCarrierConnectionName {
				foundID = *o.ID
			}
		}

		callbackURL := BaseURL.JoinPath(ShopifyLookupEndpoint)
		params := callbackURL.Query()
		params.Set(delivrioroutes.QueryParamLookupID, lookupKey)
		callbackURL.RawQuery = params.Encode()

		if foundID > 0 {
			err = updateCarrierConnection(callbackURL.String(), shopifyBaseURL, foundID, shopifyAPIKey)
			if err != nil {
				return err
			}
		} else {
			err = createCarrierConnection(callbackURL.String(), shopifyBaseURL, shopifyAPIKey)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createCarrierConnection(callbackURL string, baseURLShopify string, token string) error {

	type Post struct {
		CarrierService struct {
			Name             string `json:"name,omitempty"`
			CallbackURL      string `json:"callback_url,omitempty"`
			ServiceDiscovery bool   `json:"service_discovery,omitempty"`
		} `json:"carrier_service,omitempty"`
	}

	body, err := json.Marshal(&Post{
		CarrierService: struct {
			Name             string `json:"name,omitempty"`
			CallbackURL      string `json:"callback_url,omitempty"`
			ServiceDiscovery bool   `json:"service_discovery,omitempty"`
		}{
			Name:             DelivrioCarrierConnectionName,
			CallbackURL:      callbackURL,
			ServiceDiscovery: true,
		},
	})
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)

	createURL := fmt.Sprintf("%s/admin/api/2024-04/carrier_services.json", baseURLShopify)
	time.Sleep(shopify.RequestWait(time.Now(), token))
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("POST", createURL, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("X-Shopify-Access-Token", token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil

}

func updateCarrierConnection(callbackURL string, baseURL string, id int64, token string) error {

	type Post struct {
		CarrierService struct {
			ID               int64  `json:"id,omitempty"`
			Name             string `json:"name,omitempty"`
			CallbackURL      string `json:"callback_url,omitempty"`
			ServiceDiscovery bool   `json:"service_discovery,omitempty"`
		} `json:"carrier_service,omitempty"`
	}

	body, err := json.Marshal(&Post{
		CarrierService: struct {
			ID               int64  `json:"id,omitempty"`
			Name             string `json:"name,omitempty"`
			CallbackURL      string `json:"callback_url,omitempty"`
			ServiceDiscovery bool   `json:"service_discovery,omitempty"`
		}{
			ID:               id,
			Name:             DelivrioCarrierConnectionName,
			CallbackURL:      callbackURL,
			ServiceDiscovery: true,
		},
	})
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)

	updateCreateURL := fmt.Sprintf("%s/admin/api/2024-04/carrier_services/%v.json", baseURL, id)

	time.Sleep(shopify.RequestWait(time.Now(), token))
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("PUT", updateCreateURL, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("X-Shopify-Access-Token", token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil

}
