package endpoints

import (
	"context"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/enttest"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/returns"
	"delivrio.io/go/schema/hooks/returncollihooks"
	"delivrio.io/go/seed"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type returnSetup struct {
	client               *ent.Client
	tx                   *ent.Tx
	ctx                  context.Context
	returnDeliveryOption *ent.DeliveryOption
	customer             *ent.Address
	orderLines           []*ent.OrderLine
	portal               *ent.ReturnPortal
	ord                  *ent.Order
}

func setupOrderForReturn(t *testing.T) returnSetup {
	u, _ := url.Parse("https://testurl.delivrio.io")
	BaseURL = u
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

	seed.Products(ctx, 2)
	seed.SeedOrders(ctx, 1)
	seed.ExtraOrderLines(ctx)
	seed.ReturnPortal(ctx)
	rdo := seed.ReturnDeliveryOption(ctx)

	ord := tx.Order.Query().
		WithColli().
		FirstX(ctx)
	cust := ord.Edges.Colli[0].
		QueryRecipient().
		OnlyX(ctx)

	orderLines := ord.Edges.Colli[0].QueryOrderLines().AllX(ctx)

	rp := tx.ReturnPortal.Query().
		WithReturnPortalClaim().
		OnlyX(ctx)

	return returnSetup{
		client:               client,
		tx:                   tx,
		ctx:                  ctx,
		returnDeliveryOption: rdo,
		customer:             cust,
		orderLines:           orderLines,
		portal:               rp,
		ord:                  ord,
	}
}

func Test_returnHalfUnits(t *testing.T) {

	r := setupOrderForReturn(t)
	defer r.client.Close()
	items := make(map[pulid.ID]returns.ReturnLineItem)

	// First line of same product variant, excluded
	items[r.orderLines[1].ID] = returns.ReturnLineItem{
		Units:   2,
		ClaimID: r.portal.Edges.ReturnPortalClaim[1].ID,
	}
	items[r.orderLines[2].ID] = returns.ReturnLineItem{
		Units:   25,
		ClaimID: r.portal.Edges.ReturnPortalClaim[2].ID,
	}

	ret, err := returns.CreateReturn(r.ctx, r.ord.ID, r.portal.ID, items, "some comment")
	require.NoError(t, err, "expect partial return to succeed")

	saved := r.tx.ReturnColli.Query().
		Where(returncolli.IDIn(ret...)).
		WithReturnOrderLine().
		OnlyX(r.ctx)

	require.Equalf(
		t,
		2,
		len(saved.Edges.ReturnOrderLine),
		"should only have two return lines, found %v",
		len(saved.Edges.ReturnOrderLine),
	)

	qty := saved.Edges.ReturnOrderLine[0].Units
	require.Equalf(t, 2, qty, "expect return quantity to match, got %v", qty)
	qty = saved.Edges.ReturnOrderLine[1].Units
	require.Equalf(t, 25, qty, "expect return quantity to match, got %v", qty)

	// Create return for remaining order lines + units
	items2 := make(map[pulid.ID]returns.ReturnLineItem)

	items2[r.orderLines[0].ID] = returns.ReturnLineItem{
		Units:   1,
		ClaimID: r.portal.Edges.ReturnPortalClaim[1].ID,
	}
	items2[r.orderLines[1].ID] = returns.ReturnLineItem{
		Units:   3,
		ClaimID: r.portal.Edges.ReturnPortalClaim[1].ID,
	}
	items2[r.orderLines[2].ID] = returns.ReturnLineItem{
		Units:   26,
		ClaimID: r.portal.Edges.ReturnPortalClaim[2].ID,
	}
	items2[r.orderLines[2].ID] = returns.ReturnLineItem{
		Units:   25,
		ClaimID: r.portal.Edges.ReturnPortalClaim[2].ID,
	}

	ret2, err := returns.CreateReturn(r.ctx, r.ord.ID, r.portal.ID, items2, "some comment")
	require.NoError(t, err, "expect not 2nd return to succeed")

	saved2 := r.tx.ReturnColli.Query().
		Where(returncolli.IDIn(ret2...)).
		WithReturnOrderLine().
		OnlyX(r.ctx)

	require.Equalf(
		t,
		3,
		len(saved2.Edges.ReturnOrderLine),
		"should only have two return lines, found %v",
		len(saved2.Edges.ReturnOrderLine),
	)

	qty2 := saved2.Edges.ReturnOrderLine[0].Units
	require.Equalf(t, 1, qty2, "expect return quantity to match, got %v", qty2)
	qty2 = saved2.Edges.ReturnOrderLine[1].Units
	require.Equalf(t, 3, qty2, "expect return quantity to match, got %v", qty2)
	qty2 = saved2.Edges.ReturnOrderLine[2].Units
	require.Equalf(t, 25, qty2, "expect return quantity to match, got %v", qty2)

}

func Test_returnMoreThanAvailable(t *testing.T) {

	r := setupOrderForReturn(t)
	defer r.client.Close()
	items := make(map[pulid.ID]returns.ReturnLineItem)
	for _, ol := range r.orderLines {
		items[ol.ID] = returns.ReturnLineItem{
			Units:   2,
			ClaimID: r.portal.Edges.ReturnPortalClaim[0].ID,
		}
	}

	_, err := returns.CreateReturn(r.ctx, r.ord.ID, r.portal.ID, items, "some comment")
	var validationErr returncollihooks.HookReturnOrderLineErr
	require.ErrorAsf(t, err, &validationErr, "expect more items than in order to fail")

}

func Test_returnIsEmpty(t *testing.T) {

	r := setupOrderForReturn(t)
	defer r.client.Close()
	items := make(map[pulid.ID]returns.ReturnLineItem)
	for _, ol := range r.orderLines {
		items[ol.ID] = returns.ReturnLineItem{
			Units:   0,
			ClaimID: r.portal.Edges.ReturnPortalClaim[0].ID,
		}
	}

	_, err := returns.CreateReturn(r.ctx, r.ord.ID, r.portal.ID, items, "some comment")
	var validationErr *ent.ValidationError
	require.ErrorAsf(t, err, &validationErr, "expect empty return to fail")

}

func Test_generateReturnEmailVariables(t *testing.T) {

	r := setupOrderForReturn(t)
	defer r.client.Close()
	items := make(map[pulid.ID]returns.ReturnLineItem)
	for _, ol := range r.orderLines {
		items[ol.ID] = returns.ReturnLineItem{
			Units:   1,
			ClaimID: r.portal.Edges.ReturnPortalClaim[0].ID,
		}
	}

	returnColliIDs, err := returns.CreateReturn(r.ctx, r.ord.ID, r.portal.ID, items, "some comment")
	require.NoError(t, err)

	returnColli, err := r.tx.ReturnColli.Query().
		Where(returncolli.ID(returnColliIDs[0])).
		Only(r.ctx)
	require.NoError(t, err)

	// C&P this to avoid TX vs client issues
	err = returnColli.Update().
		SetStatus(returncolli.StatusPending).
		SetDeliveryOptionID(r.returnDeliveryOption.ID).
		Exec(r.ctx)
	require.NoError(t, err)

	// Refresh after update
	returnColli, err = r.tx.ReturnColli.Query().
		Where(returncolli.ID(returnColliIDs[0])).
		Only(r.ctx)
	require.NoError(t, err)

	mergeutils.Init(&appconfig.DelivrioConfig{BaseURL: "https://testurl.delivrio.io"})

	returnEmail, err := mergeutils.ReturnColliConfirmationLabelMerge(r.ctx, returnColli)
	require.NoError(t, err)

	labelDownload, err := mergeutils.ReturnColliURL(
		"https://testurl.delivrio.io",
		delivrioroutes.ReturnLabelDownload,
		returnColliIDs[0],
		"1001",
	)
	require.NoError(t, err)
	labelView, err := mergeutils.ReturnColliURL(
		"https://testurl.delivrio.io",
		delivrioroutes.ReturnLabel,
		returnColliIDs[0],
		"1001",
	)
	require.NoError(t, err)

	viewPNGURL, err := mergeutils.ReturnColliURL(
		"https://testurl.delivrio.io",
		delivrioroutes.ReturnLabelPNG,
		returnColliIDs[0],
		"1001",
	)
	require.NoError(t, err)

	expected := &mergeutils.ReturnConfirmationLabel{
		ReturnBase: mergeutils.ReturnBase{
			CustomerAddress: mergeutils.CustomerAddress{
				CustomerFirstName: "Pam",
				CustomerLastName:  "Armstrong",
				CustomerCompany:   "Pam's Company Ghmb",
				CustomerEmail:     "pam@example.com",
				CustomerAddress1:  "999 main st.",
				CustomerAddress2:  "Apt 12",
				CustomerZip:       "8000",
				CustomerCity:      "Aarhus",
				CustomerState:     "Midtjylland",
				CustomerCountry:   "DK",
			},
			ReturnAddress: mergeutils.ReturnAddress{
				ReturnFirstName: "Returns department",
				ReturnLastName:  "Lænard",
				ReturnCompany:   "*",
				ReturnEmail:     "returns@example.com",
				ReturnAddress1:  "Return2me road",
				ReturnAddress2:  "Behind the restaurant on the corner",
				ReturnCity:      "Copenhagen",
				ReturnState:     "++++",
				ReturnCountry:   "DK",
				ReturnZip:       "2000",
			},
			OrderPublicID:    "1001",
			ReturnMethodName: "Return drop-off",
			TrackingID:       "",
			OrderLines: []mergeutils.ReturnOrderLine{
				{
					OrderLine: mergeutils.OrderLine{
						Quantity: "1",
						Price:    "900",
						Total:    "900",
						OrderLineProductInfo: mergeutils.OrderLineProductInfo{
							ProductName:          "Peanuts 1",
							ProductVariantName:   "Brown & Salty",
							ProductFirstImageURL: "",
						},
					},
					ReturnClaimName:        "Too big",
					ReturnClaimDescription: "",
				},
				{
					OrderLine: mergeutils.OrderLine{
						Quantity: "1",
						Price:    "200",
						Total:    "200",
						OrderLineProductInfo: mergeutils.OrderLineProductInfo{
							ProductName:          "Peanuts 1",
							ProductVariantName:   "Brown & Salty",
							ProductFirstImageURL: "",
						},
					},
					ReturnClaimName:        "Too big",
					ReturnClaimDescription: "",
				},
				{
					OrderLine: mergeutils.OrderLine{
						Quantity: "1",
						Price:    "1",
						Total:    "1",
						OrderLineProductInfo: mergeutils.OrderLineProductInfo{
							ProductName:          "Peanuts 2",
							ProductVariantName:   "Brown & Salty",
							ProductFirstImageURL: "",
						},
					},
					ReturnClaimName:        "Too big",
					ReturnClaimDescription: "",
				},
			},
		},
		LabelDownloadURL: labelDownload.String(),
		LabelURL:         labelView.String(),
		LabelPNGURL:      viewPNGURL.String(),
	}

	require.Equal(t, expected, returnEmail)

}

func Test_CreateReturnOrderAuthFail(t *testing.T) {
	r := chi.NewRouter()

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/%s?%s=&%s=&%s=",
			delivrioroutes.CreateReturnOrder,
			delivrioroutes.QueryReturnPortalID,
			delivrioroutes.QueryEmail,
			delivrioroutes.QueryOrderPublicID,
		),
		nil,
	)
	require.NoError(t, err)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Define the route handler
	r.Get(delivrioroutes.CreateReturnOrder, CreateReturnOrder)
	r.ServeHTTP(w, req)
	require.Equalf(t, w.Code, http.StatusNotFound, "Expected status code %d, got %d", http.StatusNotFound, w.Code)

}

func Test_AddReturnOrderDeliveryOptionsAuthFail(t *testing.T) {
	r := chi.NewRouter()

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/%s?%s=&%s=&%s=",
			delivrioroutes.ReturnDeliveryOptions,
			delivrioroutes.QueryReturnPortalID,
			delivrioroutes.QueryEmail,
			delivrioroutes.QueryOrderPublicID,
		),
		nil,
	)
	require.NoError(t, err)

	w := httptest.NewRecorder()

	r.Get(delivrioroutes.ReturnDeliveryOptions, AddReturnOrderDeliveryOptionsAndRequestLabel)
	r.ServeHTTP(w, req)

	require.Equalf(t, w.Code, http.StatusNotFound, "Expected status code %d, got %d", http.StatusNotFound, w.Code)

}
