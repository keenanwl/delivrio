package common

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type LoggerAuthRoundTripper struct {
	Transport http.RoundTripper
}

// Add the required USPS basic auth header to the initial
// oauth2 request
func (rt *LoggerAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	bod, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bod))

	resp, err := rt.Transport.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("logger: %w", err)
	}

	bod, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bod))

	return resp, nil
}
