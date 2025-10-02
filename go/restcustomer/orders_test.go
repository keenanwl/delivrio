package restcustomer

import (
	"bytes"
	"context"
	"delivrio.io/go/endpoints"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type setupConfig struct {
	cli     *ent.Client
	ctx     context.Context
	token   string
	request OrdersCreateRequest
	router  *chi.Mux
}

func setup(t *testing.T) setupConfig {
	cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := ent.NewContext(viewer.NewContext(context.Background(), viewer.UserViewer{Role: viewer.Background, Context: pulid.MustNew("CH")}), cli)

	tx, err := cli.Tx(ctx)
	require.NoError(t, err)
	ctx = ent.NewTxContext(ctx, tx)
	seed.Base(ctx)
	seed.Products(ctx, 6)
	token := seed.APICredentials(ctx)
	err = tx.Commit()
	require.NoError(t, err)

	r := chi.NewRouter()
	r.Use(endpoints.AddClient(cli))
	r.Use(endpoints.CheckRESTAPICredentials())
	r.Post("/orders", HandleOrderCreate)

	requestPayload := OrdersCreateRequest{
		Orders: []OrderCreate{
			{
				PublicID:       "SH123456789",
				ConnectionName: "Shopify DK `'*øæ~~~^@£€¤",
				OrderLines:     []OrderLine{},
				DeliveryAddress: Address{
					FirstName:     "PÅm",
					LastName:      "Armstrong",
					Company:       "Vandalay Industries",
					VATNumber:     "DK12121212",
					StreetOne:     "9999 Aarhus road",
					StreetTwo:     "Apt 1",
					PostalCode:    "8000",
					City:          "Aarhus",
					CountryAlpha2: "DK",
					State:         "",
					Email:         "pam@example.com",
					PhoneNumber:   "+45 11 11 11 11",
				},
			},
		},
	}

	return setupConfig{
		cli:     cli,
		ctx:     ctx,
		token:   token,
		request: requestPayload,
		router:  r,
	}

}

func Test_CreateOrdersNoOrderLines(t *testing.T) {

	set := setup(t)
	defer set.cli.Close()

	payload := set.request
	payload.Orders[0].SenderAddress = &Address{
		FirstName:     "A sender dept",
		LastName:      "In the side",
		Company:       "A company",
		VATNumber:     "DK0000000",
		StreetOne:     "A sender street",
		StreetTwo:     "Next sender street",
		PostalCode:    "2000",
		City:          "Aalborg",
		CountryAlpha2: "DE",
		State:         "",
		Email:         "sender@example.com",
		PhoneNumber:   "+45 00 00 00 00",
	}

	requestPayloadJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewReader(requestPayloadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(endpoints.DelivrioApiHeaderKey, set.token)

	w := httptest.NewRecorder()
	set.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	ord := set.cli.Order.Query().FirstX(set.ctx)

	expectedResponse := OrdersCreateResponse{
		Orders: []OrdersSaved{
			{
				Success: true,
				OrderID: ord.ID.String(),
			},
		},
	}

	expectedResponseJSON, _ := json.Marshal(expectedResponse)

	require.JSONEq(t, string(expectedResponseJSON), w.Body.String())
}

func Test_CreateOrdersMultiple(t *testing.T) {

	set := setup(t)
	defer set.cli.Close()

	payload := set.request
	externalID := fmt.Sprintf("%v", 1)
	pvID := seed.GetProductVariants()[0].ID.String()
	payload.Orders[0].OrderLines = []OrderLine{
		{
			ExternalProductVariantID: &externalID,
			ProductVariantID:         &pvID,
			Units:                    5,
			Price:                    10,
			Currency:                 "DKK",
		},
	}
	payload.Orders = append(payload.Orders, set.request.Orders[0])
	payload.Orders[1].OrderLines = []OrderLine{
		{
			ExternalProductVariantID: &externalID,
			ProductVariantID:         nil,
			Units:                    10,
			Price:                    100,
			Currency:                 "DKK",
		},
	}

	requestPayloadJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, delivrioroutes.Orders, bytes.NewReader(requestPayloadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(endpoints.DelivrioApiHeaderKey, set.token)

	w := httptest.NewRecorder()
	set.router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	require.Equal(t, http.StatusMultiStatus, w.Code)

	colli, err := set.cli.Colli.Query().
		WithOrder().
		WithSender().
		WithRecipient().
		WithOrderLines().
		Only(set.ctx)
	require.NoError(t, err)

	var resp OrdersCreateResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.False(t, resp.Orders[0].Success, "expect order was not saved")
	require.True(t, resp.Orders[1].Success, "expect order was saved")
	require.Len(t, resp.Orders[1].Error, 0, "expect no error")

	require.Equal(t, 1, len(colli.Edges.OrderLines), "expect order line count to match seeded")
	require.Equal(t, 10, colli.Edges.OrderLines[0].Units, "expect order line units to match input")
	require.Equal(t, float64(100), colli.Edges.OrderLines[0].UnitPrice, "expect order line price to match input")
	require.Equal(t, "Some first name", colli.Edges.Sender.FirstName, "expect connection location to be used as sender")

}

func Test_CreateOrdersNoSenderOrderLineFail(t *testing.T) {

	set := setup(t)
	defer set.cli.Close()

	lines := make([]OrderLine, 0)
	for pvIndex, pv := range seed.GetProductVariants() {
		pvID := pv.ID.String()
		lines = append(lines,
			OrderLine{
				Units:    pvIndex + 1,
				Price:    120.40 + float64(pvIndex),
				Currency: "DKK",
			},
		)

		if pvIndex%2 == 0 {
			id := fmt.Sprintf("%v", pvIndex)
			lines[pvIndex].ExternalProductVariantID = &id
		} else {
			lines[pvIndex].ProductVariantID = &pvID
		}
	}

	payload := set.request
	payload.Orders[0].OrderLines = lines

	requestPayloadJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewReader(requestPayloadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(endpoints.DelivrioApiHeaderKey, set.token)

	w := httptest.NewRecorder()
	set.router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	require.Equal(t, http.StatusMultiStatus, w.Code)

	_, err := set.cli.Colli.Query().
		WithOrder().
		WithSender().
		WithRecipient().
		WithOrderLines().
		Only(set.ctx)
	require.Error(t, err)

	var resp OrdersCreateResponse
	json.Unmarshal(w.Body.Bytes(), &resp)

	require.False(t, resp.Orders[0].Success, "expect order was not saved")
	require.Containsf(t, resp.Orders[0].Error, "(offset 0)", "expect error message")

}

func Test_CreateOrdersNoSender(t *testing.T) {

	set := setup(t)
	defer set.cli.Close()

	lines := make([]OrderLine, 0)
	for pvIndex, pv := range seed.GetProductVariants() {
		pvID := pv.ID.String()
		lines = append(lines,
			OrderLine{
				ExternalProductVariantID: nil,
				ProductVariantID:         &pvID,
				Units:                    pvIndex + 1,
				Price:                    120.40 + float64(pvIndex),
				Currency:                 "DKK",
			},
		)
	}

	payload := set.request
	payload.Orders[0].OrderLines = lines

	requestPayloadJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewReader(requestPayloadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(endpoints.DelivrioApiHeaderKey, set.token)

	w := httptest.NewRecorder()
	set.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	colli := set.cli.Colli.Query().
		WithOrder().
		WithSender().
		WithRecipient().
		WithOrderLines().
		OnlyX(set.ctx)

	expectedResponse := OrdersCreateResponse{
		Orders: []OrdersSaved{
			{
				Success: true,
				OrderID: colli.Edges.Order.ID.String(),
			},
		},
	}

	expectedResponseJSON, _ := json.Marshal(expectedResponse)
	require.JSONEq(t, string(expectedResponseJSON), w.Body.String())

	require.Equal(t, 6, len(colli.Edges.OrderLines), "expect order line count to match seeded")

	require.Equal(t, 1, colli.Edges.OrderLines[0].Units, "expect order line units to match input")
	require.Equal(t, 2, colli.Edges.OrderLines[1].Units, "expect order line units to match input")
	require.Equal(t, 3, colli.Edges.OrderLines[2].Units, "expect order line units to match input")

	require.Equal(t, 120.4, colli.Edges.OrderLines[0].UnitPrice, "expect order line price to match input")
	require.Equal(t, 121.4, colli.Edges.OrderLines[1].UnitPrice, "expect order line price to match input")
	require.Equal(t, 122.4, colli.Edges.OrderLines[2].UnitPrice, "expect order line price to match input")

	require.Equal(t, "Some first name", colli.Edges.Sender.FirstName, "expect connection location to be used as sender")
}

func Test_CreateOrdersNoOrderLinesNoSender(t *testing.T) {

	set := setup(t)
	defer set.cli.Close()

	requestPayloadJSON, _ := json.Marshal(set.request)

	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewReader(requestPayloadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(endpoints.DelivrioApiHeaderKey, set.token)

	w := httptest.NewRecorder()
	set.router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	colli := set.cli.Colli.Query().
		WithOrder().
		WithSender().
		WithRecipient().
		WithOrderLines().
		OnlyX(set.ctx)

	expectedResponse := OrdersCreateResponse{
		Orders: []OrdersSaved{
			{
				Success: true,
				OrderID: colli.Edges.Order.ID.String(),
			},
		},
	}

	expectedResponseJSON, _ := json.Marshal(expectedResponse)

	require.JSONEq(t, string(expectedResponseJSON), w.Body.String())
	require.Equal(t, "Some first name", colli.Edges.Sender.FirstName, "expect connection location to be used as sender")
}
