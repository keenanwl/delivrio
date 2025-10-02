package main

import (
	"bytes"
	"context"
	"delivrio.io/go/restcustomer"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	open := make(chan struct{}, 50)
	euCountries := []string{"DK", "DE", "FR", "IT", "ES"} // Add more EU countries if needed
	pvID := "40329440428092"

	start := time.Now()
	startInterval := time.Now()

	// Number of requests to fire simultaneously
	numRequests := 500

	for c := 0; c < numRequests; c++ {
		open <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() { <-open }()
			defer wg.Done()

			ordersRequest := restcustomer.OrdersCreateRequest{make([]restcustomer.OrderCreate, 0)}
			for i := 1; i <= 4; i++ {
				order := restcustomer.OrderCreate{
					PublicID:       fmt.Sprintf("SH%d%v", i, time.Now().Format(time.RFC3339Nano)),
					ConnectionName: "Shopify DK",
					DeliveryAddress: restcustomer.Address{
						FirstName:     "John",
						LastName:      "Doe",
						StreetOne:     "123 Main St",
						PostalCode:    "12345",
						City:          "SampleCity",
						CountryAlpha2: euCountries[rand.Intn(len(euCountries))],
						Email:         "john.doe@example.com",
					},
					OrderLines: []restcustomer.OrderLine{{
						ExternalProductVariantID: &pvID,
						ProductVariantID:         nil,
						Units:                    1,
						Price:                    1000,
						Currency:                 "DKK",
					}},
				}
				ordersRequest.Orders = append(ordersRequest.Orders, order)
			}

			// Convert orders to JSON
			ordersJSON, err := json.Marshal(ordersRequest)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Create an HTTP request and attach the span's context
			req, _ := http.NewRequest("POST", "http://localhost:8080/rest/v1/orders", bytes.NewReader(ordersJSON))
			req = req.WithContext(context.Background())
			req.Header.Set("X-DELIVRIO-Key", "<insert>")

			// Make the API request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			bod, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			fmt.Println(string(bod))
			fmt.Println("----INTERVAL TIME:", time.Now().UnixMilli()-startInterval.UnixMilli())
			startInterval = time.Now()

		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("----TOTAL TIME:", time.Now().UnixMilli()-start.UnixMilli())
}
