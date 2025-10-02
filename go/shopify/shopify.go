package shopify

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"sync"
	"time"
)

var RequestCount = make(map[string]time.Time)
var RequestBucketCount = make(map[string]float64)
var RequestCountMutex = sync.Mutex{}

const maxRequestInterval = time.Millisecond * 500
const baselineRequestInterval = time.Millisecond * 250

func RecordHeader(key string, resp *http.Response) {
	RequestCountMutex.Lock()
	defer RequestCountMutex.Unlock()

	limitHeader := resp.Header.Get("X-Shopify-Shop-Api-Call-Limit")

	parts := strings.Split(limitHeader, "/")
	used, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("error: parsing call limit part 1 %v", err)
		return
	}
	limit, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("error: parsing call limit part 1 %v", err)
		return
	}

	RequestBucketCount[key] = float64(used) / float64(limit)
}

// Avoid the rate limiting
func RequestWait(now time.Time, apiKey string) time.Duration {
	RequestCountMutex.Lock()
	defer RequestCountMutex.Unlock()

	if lastTime, ok := RequestCount[apiKey]; ok {

		// Backoff when bucket is filling too quickly
		waitInterval := baselineRequestInterval
		if bucketThreshold, ok := RequestBucketCount[apiKey]; ok {
			if bucketThreshold >= 0.75 {
				waitInterval = maxRequestInterval * 3
			} else if bucketThreshold >= 0.5 {
				waitInterval = maxRequestInterval * 2
			}
		}

		if lastTime.Add(waitInterval).Before(now) {
			RequestCount[apiKey] = now
			return 0
		}

		RequestCount[apiKey] = now
		return lastTime.Add(waitInterval).Sub(now)
	}

	RequestCount[apiKey] = now
	return 0
}

func FireRequest(ctx context.Context, key string, req *http.Request, expectedResponse int) ([]byte, error) {

	time.Sleep(RequestWait(time.Now(), key))

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req.Header.Set("X-Shopify-Access-Token", key)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	o, _ := httputil.DumpResponse(res, true)
	fmt.Println(string(o))

	RecordHeader(key, res)

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("shopify error response: %w", err)
	}

	if res.StatusCode == expectedResponse {
		return respBody, nil
	}

	return respBody, fmt.Errorf("shopify unexpected response: %v", res.StatusCode)
}
