package main

import (
	"bytes"
	"delivrio.io/shared-utils/printers"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type PrintClientPing struct {
	Token      string             `json:"token"`
	ComputerID string             `json:"computer_id"`
	Printers   []printers.Printer `json:"printers"`
	LabelID    string             `json:"label_id"`
}

func main() {
	// Define the endpoint URL
	endpointURL := "http://localhost:8080/api/register-scan?device-type=app&computer-id=12345&id=cVB5ZkJIWU1BdUxwUmVjNFZsaDJYVmQ5TDNxRmxHWnZhQlRqeDlQbnQ3UXlXQlh4bDBxTlEzMzBJbXFjWXFlQTpXUzAxSDVNQU45WlFBWFJKNFJIRDhRNjBXMkszOlRFMDFINFJBS0ZRSEM2OUJWUVY0WU1LMUpSV1A="
	//endpointURL := "http://localhost:8080/api/register-scan"

	// Number of simultaneous requests
	numRequests := 300000

	successCount := 0
	failureCount := 0

	// Maximum number of workers in the pool
	maxWorkers := 750

	// Create a mutex to synchronize access to the counters
	var counterMutex sync.Mutex

	start := time.Now()

	requests := make(chan struct{}, maxWorkers)

	// Create a WaitGroup to wait for all requests to finish
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Send simultaneous requests
	for i := 0; i < numRequests; i++ {
		i := i
		requests <- struct{}{}
		go func(count int) {
			defer func() {
				// Release the worker back to the pool
				<-requests
				wg.Done()
			}()

			extra := ""
			if count%2 == 0 {
				extra = "1"
			}

			// Create a new request body
			ping := PrintClientPing{
				//Token:      "YmVJdlExQUVsSW5GNXdqb3hUVTN6dkJ1Ymp5ekNhSkpwdGRYakc1Wnpnd0JmelVKeXZzRjRCQkE5OWx1Y3JIRTpXUzAxSDJGRjFGUEQ1TTlLM1NZVlY1MVQ2MzVLOlRFMDFIMVRZRUFZWlk5WDBIRzkzSzlORllGMUo=",
				Token:      "cVB5ZkJIWU1BdUxwUmVjNFZsaDJYVmQ5TDNxRmxHWnZhQlRqeDlQbnQ3UXlXQlh4bDBxTlEzMzBJbXFjWXFlQTpXUzAxSDVNQU45WlFBWFJKNFJIRDhRNjBXMkszOlRFMDFINFJBS0ZRSEM2OUJWUVY0WU1LMUpSV1A=",
				ComputerID: "12345",
				Printers:   []printers.Printer{ /* populate your printer data */ },
				LabelID:    fmt.Sprintf("sdfsfsdfds%v", extra),
			}

			// Convert the request body to JSON
			jsonData, err := json.Marshal(ping)
			if err != nil {
				log.Fatal(err)
			}

			// Send the HTTP POST request
			resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()

			d, _ := io.ReadAll(resp.Body)
			fmt.Println(string(d))
			// Check the response status code
			if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotFound {
				fmt.Println("Request succeeded", count)
				counterMutex.Lock()
				successCount++
				counterMutex.Unlock()
			} else {
				fmt.Println("Request failed with status code:", resp.StatusCode)
				counterMutex.Lock()
				failureCount++
				counterMutex.Unlock()
			}
		}(i)
	}

	// Wait for all requests to finish
	wg.Wait()
	// Print the final metrics
	fmt.Println("Total requests:", numRequests)
	fmt.Println("Successful requests:", successCount)
	fmt.Println("Failed requests:", failureCount)
	fmt.Println("Total seconds:", time.Now().Sub(start).Seconds())
}
