package glsapis

import (
	"context"
	"delivrio.io/go/carrierapis/glsapis/glsrequest"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	shipment2 "delivrio.io/go/ent/shipment"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Test_limitString(t *testing.T) {
	s := "000000000_000000000_000000000_000000000_000000000_"
	res := limitString(s, 40)
	require.Equal(t, 40, len(res))

	s = "000000000_000000000_"
	res = limitString(s, 40)
	require.Equal(t, 20, len(res))
}

type SetupTest struct {
	ctx    context.Context
	client *ent.Client
	tx     *ent.Tx
	col    *ent.Colli
	ship   *ent.Shipment
	config RequestConfig
}

func setup_generatev1CreateShipment(t *testing.T) *SetupTest {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	ctx := ent.NewContext(context.Background(), client)

	tx, _ := client.Tx(ctx)
	ctx = ent.NewTxContext(ctx, tx)
	seed.Base(ctx)
	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Admin,
		Context: pulid.MustNew("CH"),
		Tenant:  seed.GetTenantID(),
	})
	seed.DeliveryOption(ctx)
	seed.Products(ctx, 2)
	seed.SeedOrders(ctx, 1)
	seed.ExtraOrderLines(ctx)

	ctx = viewer.NewContext(ctx, viewer.UserViewer{
		Role:    viewer.Admin,
		Context: pulid.MustNew("CH"),
		Tenant:  seed.GetTenantID(),
	})

	car := tx.Carrier.Query().FirstX(ctx)

	ship := tx.Shipment.Create().
		SetInput(ent.CreateShipmentInput{
			ShipmentPublicID:   "ØÆÅ",
			ShipmentPostNordID: nil,
			ShipmentGLSID:      nil,
			ShipmentParcelIDs:  nil,
		}).
		SetCarrier(car).
		SetStatus(shipment2.StatusPending).
		SetTenantID(seed.GetTenantID()).
		SaveX(ctx)

	col := tx.Colli.Query().WithOrderLines().FirstX(ctx)
	tx.ShipmentParcel.Create().
		SetColli(col).
		SetStatus(shipmentparcel.StatusPrinted).
		SetShipment(ship).
		SetTenantID(seed.GetTenantID()).
		ExecX(ctx)

	return &SetupTest{
		ctx:    ctx,
		client: client,
		tx:     tx,
		col:    col,
		ship:   ship,
		config: RequestConfig{
			GLSAPIAuth: GLSAPIAuth{
				UserName:   "12",
				Password:   "34",
				ContactID:  "56",
				CustomerID: "78",
			},
			Shipment: ShipmentConfig{
				GLSAPIAuth:        GLSAPIAuth{},
				ParcelShopAddress: nil,
				AdditionalServices: []AdditionalService{
					{Key: "ShopReturn", Value: "Y"},
				},
			},
		},
	}

}

func Test_generate1CreateShipment_invalidService(t *testing.T) {

	r := setup_generatev1CreateShipment(t)
	defer r.client.Close()

	sender := r.tx.Address.Query().FirstX(r.ctx)
	receiver := r.tx.Address.Query().FirstX(r.ctx)

	products := make([]*ent.ProductVariant, 0)
	for _, ol := range r.col.Edges.OrderLines {
		pv, err := ol.ProductVariant(r.ctx)
		require.NoError(t, err)
		products = append(products, pv)
	}

	pc := make([]PackageConfig, 0)
	pc = append(pc, PackageConfig{
		DelivrioShipmentID: r.ship.ID,
		DelivrioColliID:    r.col.ID,
		Items:              products,
	})

	config := r.config
	config.Shipment.Packages = pc
	config.Shipment.ConsignorAddress = sender
	config.Shipment.ConsigneeAddress = receiver
	config.Shipment.AdditionalServices = []AdditionalService{
		{Key: "Shipping", Value: "Y"},
	}
	_, _, err := generateV1CreateShipment(r.ctx, config)
	require.EqualError(t, err, ErrServiceNotSupported.Error())

}

func Test_generatev1CreateShipment_ShopReturn(t *testing.T) {

	r := setup_generatev1CreateShipment(t)
	defer r.client.Close()

	sender := r.tx.Address.Query().FirstX(r.ctx)
	receiver := r.tx.Address.Query().FirstX(r.ctx)

	products := make([]*ent.ProductVariant, 0)
	for _, ol := range r.col.Edges.OrderLines {
		pv, err := ol.ProductVariant(r.ctx)
		require.NoError(t, err)
		products = append(products, pv)
	}

	pc := make([]PackageConfig, 0)
	pc = append(pc, PackageConfig{
		DelivrioShipmentID: r.ship.ID,
		DelivrioColliID:    r.col.ID,
		Items:              products,
	})

	config := r.config
	config.Shipment.Packages = pc
	config.Shipment.ConsignorAddress = sender
	config.Shipment.ConsigneeAddress = receiver
	config.Shipment.AdditionalServices = []AdditionalService{
		{Key: "ShopReturn", Value: "Y"},
	}

	u, res, err := generateV1CreateShipment(r.ctx, config)
	require.NoError(t, err)
	require.Equal(t, 1, len(res.Parcels))
	require.Equal(t, glsrequest.Parcel{
		Weight: 0.154,
	}, res.Parcels[0])
	require.Equal(t, "Y", *res.Services.ShopReturn)
	// DELIVRIO -> GLS conversion tested elsewhere
	require.Equal(t, sender.Email, res.Addresses.Delivery.Email)
	require.Equal(t, receiver.Email, res.Addresses.Pickup.Email)
	require.Nil(t, res.Addresses.AlternativeShipper)
	uExpected, _ := url.Parse("http://api.gls.dk/ws/DK/V1/CreateShipment")
	require.Equal(t, uExpected, u)

}

func Test_generatev1CreateShipment_PrivateDelivery(t *testing.T) {

	r := setup_generatev1CreateShipment(t)
	defer r.client.Close()

	sender := r.tx.Address.Query().FirstX(r.ctx)
	receiver := r.tx.Address.Query().FirstX(r.ctx)

	products := make([]*ent.ProductVariant, 0)
	for _, ol := range r.col.Edges.OrderLines {
		pv, err := ol.ProductVariant(r.ctx)
		require.NoError(t, err)
		products = append(products, pv)
	}

	pc := make([]PackageConfig, 0)
	pc = append(pc, PackageConfig{
		DelivrioShipmentID: r.ship.ID,
		DelivrioColliID:    r.col.ID,
		Items:              products,
	})

	config := r.config
	config.Shipment.Packages = pc
	config.Shipment.ConsignorAddress = sender
	config.Shipment.ConsigneeAddress = receiver
	config.Shipment.AdditionalServices = []AdditionalService{
		{Key: "PrivateDelivery", Value: "Y"},
	}

	u, res, err := generateV1CreateShipment(r.ctx, config)
	require.NoError(t, err)
	require.Equal(t, 1, len(res.Parcels))
	require.Equal(t, glsrequest.Parcel{
		Weight: 0.154,
	}, res.Parcels[0])
	require.Nil(t, res.Services.ShopReturn)
	// DELIVRIO -> GLS conversion tested elsewhere
	require.Equal(t, receiver.Email, res.Addresses.Delivery.Email)
	require.Equal(t, sender.Email, res.Addresses.AlternativeShipper.Email)
	require.Nil(t, res.Addresses.Pickup)
	uExpected, _ := url.Parse("http://api.gls.dk/ws/DK/V1/CreateShipment")
	require.Equal(t, uExpected, u)
	require.Equal(t, time.Now().Format(glsTimestampFormat), res.ShipmentDate)

}

func Test_delivrioRecipientToGLSAddress(t *testing.T) {

	r := setup_generatev1CreateShipment(t)
	defer r.client.Close()

	receiver := r.tx.Address.Query().Offset(2).FirstX(r.ctx)
	adr, err := delivrioRecipientToGLSAddress(r.ctx, receiver)
	require.NoError(t, err)
	require.Equal(t, &glsrequest.Address{
		Name1:      "Pam",
		Name2:      "Armstrong",
		Name3:      "",
		Street1:    "999 main st. Apt 12",
		CountryNum: "208",
		ZipCode:    "8000",
		City:       "Aarhus",
		Contact:    "Pam Armstrong",
		Email:      "pam@example.com",
		Phone:      "+45 22 22 22 22",
		Mobile:     "+45 22 22 22 22",
	}, adr)
}

func Test_delivrioPackageSenderToGLSAddress(t *testing.T) {

	r := setup_generatev1CreateShipment(t)
	defer r.client.Close()

	sender := r.tx.Address.Query().Offset(3).FirstX(r.ctx)
	adr, err := delivrioPackageSenderToGLSAddress(r.ctx, sender)
	require.NoError(t, err)
	require.Equal(t, &glsrequest.Address{
		Name1:      "Returns",
		Name2:      "Department",
		Name3:      "",
		Street1:    "Shipper avenue Department 1",
		CountryNum: "208",
		ZipCode:    "8000",
		City:       "Aarhus",
		Contact:    "Returns Department",
		Email:      "support@example.com",
		Phone:      "+45 11 11 11 11",
		Mobile:     "+45 11 11 11 11",
	}, adr)
}

// ChatGPT generated
func TestGenerateRequestContentType(t *testing.T) {
	// Create a test payload
	payload := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, World!",
	}

	// Create a test server to capture the request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the "Content-Type" header
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
		}

		// Handle the request
		// ...

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// Generate the request using the test payload and URL
	u, _ := url.Parse(ts.URL)
	req, err := generateRequest(context.Background(), u, payload)
	if err != nil {
		t.Fatalf("Error generating request: %v", err)
	}

	// Perform the request (optional)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Add assertions for the response (if needed)
	// ...

	// You can also check the response status code and other aspects here
	// ...
}

func TestFireRequestTimeout(t *testing.T) {
	// Create a test HTTP server with a handler that sleeps longer than the client timeout
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Simulate a long-running request
	}))
	defer testServer.Close()

	// Create a request to the test server with a short client timeout
	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Agent", "test")

	// Set a short timeout for the client
	timeout = 1 * time.Second

	// Call the fireRequest function with the test request
	startTime := time.Now()
	_, err = fireRequest(req)
	elapsedTime := time.Since(startTime)

	require.Error(t, err)

	// Assert that the elapsed time is roughly equal to the client timeout
	require.True(t, elapsedTime >= timeout && elapsedTime <= (timeout+100*time.Millisecond),
		"elapsedTime should be within the timeout range")

	// You can add more specific assertions based on your use case
}
