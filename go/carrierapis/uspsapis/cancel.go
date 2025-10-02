package uspsapis

import (
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/ent"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"log"
	"net/http"
	"net/url"
)

func deleteRequest(token string, requestURL string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Payment-Authorization-Token", token)

	return req, nil

}

func CancelByTrackingCode(ctx context.Context, carrierUSPS *ent.CarrierUSPS, code string) error {

	toggledURL := baseURLTest
	if !carrierUSPS.IsTestAPI {
		toggledURL = baseURL
	}

	tokenURL, err := url.JoinPath(toggledURL, "/oauth2/v3/token")
	if err != nil {
		return err
	}

	cliWithBasicAuth := &http.Client{
		Transport: &BasicAuthRoundTripper{
			Transport:    http.DefaultTransport,
			ClientID:     carrierUSPS.ConsumerKey,
			ClientSecret: carrierUSPS.ConsumerSecret,
		},
	}

	cliWithLogger := &http.Client{
		Transport: &common.LoggerAuthRoundTripper{
			Transport: http.DefaultTransport,
		},
	}

	config := &clientcredentials.Config{
		ClientID:     carrierUSPS.ConsumerKey,
		ClientSecret: carrierUSPS.ConsumerSecret,
		TokenURL:     tokenURL,
		Scopes:       []string{"payments", "labels"},
		// I think? Check here if not working.
		AuthStyle: oauth2.AuthStyleInParams,
	}

	ctxWithBasic := context.WithValue(ctx, oauth2.HTTPClient, cliWithBasicAuth)
	ctxWithLogger := context.WithValue(ctx, oauth2.HTTPClient, cliWithLogger)
	ts := config.TokenSource(ctxWithBasic)
	uspsClient := oauth2.NewClient(ctxWithLogger, ts)

	token, err := paymentTokenFromTestCredentials(ctx, uspsClient, carrierUSPS, toggledURL)
	if err != nil {
		return fmt.Errorf("cancel label: %w", err)
	}

	deleteUrl, err := url.JoinPath(toggledURL, "/labels/v3/label/", code)
	if err != nil {
		return fmt.Errorf("cancel label: %w", err)
	}

	req, err := deleteRequest(token, deleteUrl)
	if err != nil {
		return fmt.Errorf("cancel label: %w", err)
	}

	resp, err := uspsClient.Do(req)
	if err != nil {
		return fmt.Errorf("cancel label: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cancel label: %w", err)
	}

	log.Printf("cancel label: error body: %v", string(bod))

	return fmt.Errorf("cancel label: expected 200, got %v", resp.StatusCode)

}
