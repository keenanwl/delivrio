package endpoints

import (
	"bytes"
	"context"
	"delivrio.io/go/carrierapis/common"
	"delivrio.io/go/carrierapis/glsapis"
	"delivrio.io/go/carrierapis/postnordapis"
	"delivrio.io/go/carrierapis/uspsapis"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/go/ent/address"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/language"
	"delivrio.io/go/ent/orderline"
	"delivrio.io/go/ent/returnorderline"
	"delivrio.io/go/i18n"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/utils"
	"delivrio.io/go/utils/httputils"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"delivrio.io/go/deliveryoptions"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/connection"
	"delivrio.io/go/ent/deliveryoption"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/ent/returnportal"
	"delivrio.io/go/ent/returnportalclaim"
	"delivrio.io/go/returns"
	pulid_server_prefix "delivrio.io/go/schema/pulid-server-prefix"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
)

func ReturnColliFileViewerHandler(fileResponseWriter func(w http.ResponseWriter, r *http.Request, returnColli *ent.ReturnColli) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		returnColliID := r.URL.Query().Get(delivrioroutes.QueryParamReturnColliID)
		orderPublicID := r.URL.Query().Get(delivrioroutes.QueryParamOrderPublicID)

		if len(returnColliID) == 0 || len(orderPublicID) == 0 {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "URL not understood"})
			return
		}

		_, err := pulid_server_prefix.IDToType(r.Context(), pulid.ID(returnColliID))
		if err != nil {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "invalid return colli ID"})
			return
		}

		db := ent.FromContext(r.Context())

		ctx := viewer.NewContext(r.Context(), viewer.UserViewer{
			Role: viewer.Background,
		})

		returnColli, err := db.ReturnColli.Query().
			Where(returncolli.And(
				returncolli.ID(pulid.ID(returnColliID)),
				// A little extra obscurity
				returncolli.HasOrderWith(order.OrderPublicIDEqualFold(orderPublicID)),
			)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			httputils.JSONResponse(w, http.StatusInternalServerError, httputils.Map{"message": err.Error()})
			return
		} else if ent.IsNotFound(err) {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": err.Error()})
			return
		}

		currentTime := time.Now()
		monthsDiff := int(currentTime.Month()) - int(returnColli.CreatedAt.Month()) + 12*(currentTime.Year()-returnColli.CreatedAt.Year())

		if monthsDiff > 2 {
			// Expired
			httputils.JSONResponse(w, http.StatusGone, httputils.Map{"message": "return label is no longer available"})
			return
		}

		err = fileResponseWriter(w, r, returnColli)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, httputils.Map{"message": err.Error()})
			return
		}
		return
	}
}

func ReturnLabelViewerWriter(w http.ResponseWriter, r *http.Request, returnColli *ent.ReturnColli) error {
	pdfData, err := base64.StdEncoding.DecodeString(returnColli.LabelPdf)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/pdf")

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(pdfData)
	if err != nil {
		return err
	}
	return nil
}

func ReturnQRCodeViewerWriter(w http.ResponseWriter, r *http.Request, returnColli *ent.ReturnColli) error {
	pngData, err := base64.StdEncoding.DecodeString(returnColli.QrCodePng)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "image/png")

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(pngData)
	if err != nil {
		return err
	}
	return nil
}

func ReturnLabelPNGViewerWriter(w http.ResponseWriter, r *http.Request, returnColli *ent.ReturnColli) error {
	labelPNG := returnColli.LabelPng
	if len(labelPNG) == 0 {
		png, err := utils.Base64PDFToPNG(returnColli.LabelPdf)
		if err != nil {
			return err
		}

		pngBytes, err := utils.EncodePNG(*png)
		if err != nil {
			return err
		}

		labelPNG = string(pngBytes)

		bgCtx := viewer.NewBackgroundContext(r.Context())

		tenant, err := returnColli.Tenant(bgCtx)
		if err != nil {
			return err
		}

		err = returnColli.Update().
			SetLabelPng(base64.StdEncoding.EncodeToString(pngBytes)).
			Exec(viewer.MergeViewerTenantID(bgCtx, tenant.ID))
		if err != nil {
			return err
		}
	} else {
		labelPNGBytes, err := base64.StdEncoding.DecodeString(labelPNG)
		if err != nil {
			return err
		}
		labelPNG = string(labelPNGBytes)
	}

	w.Header().Set("Content-Type", "image/png")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(labelPNG))
	if err != nil {
		return err
	}
	return nil
}

func ReturnColliFileDownloadHandler(fileResponseWriter func(w http.ResponseWriter, returnColli *ent.ReturnColli) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		returnColliID := r.URL.Query().Get(delivrioroutes.QueryParamReturnColliID)
		orderPublicID := r.URL.Query().Get(delivrioroutes.QueryParamOrderPublicID)

		if len(returnColliID) == 0 || len(orderPublicID) == 0 {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "URL not understood"})
			return
		}

		_, err := pulid_server_prefix.IDToType(r.Context(), pulid.ID(returnColliID))
		if err != nil {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "invalid return colli ID"})
			return
		}

		db := ent.FromContext(r.Context())

		ctx := viewer.NewContext(r.Context(), viewer.UserViewer{
			Role: viewer.Background,
		})

		returnColli, err := db.ReturnColli.Query().
			Where(returncolli.And(
				returncolli.ID(pulid.ID(returnColliID)),
				// A little extra obscurity
				returncolli.HasOrderWith(order.OrderPublicIDEqualFold(orderPublicID)),
			)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			httputils.JSONResponse(w, http.StatusInternalServerError, httputils.Map{"message": err.Error()})
			return
		} else if ent.IsNotFound(err) {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": err.Error()})
			return
		}

		currentTime := time.Now()
		monthsDiff := int(currentTime.Month()) - int(returnColli.CreatedAt.Month()) + 12*(currentTime.Year()-returnColli.CreatedAt.Year())

		if monthsDiff > 2 {
			// Expired
			httputils.JSONResponse(w, http.StatusGone, httputils.Map{"message": "return label is no longer available"})
			return
		}

		err = fileResponseWriter(w, returnColli)
		if err != nil {
			httputils.JSONResponse(w, http.StatusInternalServerError, httputils.Map{"message": err.Error()})
			return
		}
		return
	}
}

func ReturnLabelDownloadWriter(w http.ResponseWriter, returnColli *ent.ReturnColli) error {
	pdfData, err := base64.StdEncoding.DecodeString(returnColli.LabelPdf)
	if err != nil {
		return err
	}

	// Set the necessary headers for PDF download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=return-label-%s.pdf", returnColli.ID))

	// Create a bytes.Buffer and write the PDF data to it
	buffer := bytes.NewBuffer(pdfData)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func ReturnQRCodeDownloadWriter(w http.ResponseWriter, returnColli *ent.ReturnColli) error {
	pngData, err := base64.StdEncoding.DecodeString(returnColli.QrCodePng)
	if err != nil {
		return err
	}

	// Set the necessary headers for PNG download
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=return-qr-code-%s.png", returnColli.ID))

	// Create a bytes.Buffer and write the PDF data to it
	buffer := bytes.NewBuffer(pngData)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

type AddReturnOrderDeliveryOptionsInput struct {
	DeliveryOptions []common.ReturnOrderDeliveryOptionsColliIDs `json:"delivery_options"`
}
type AddReturnOrderDeliveryOptionsOutput struct {
	Success bool `json:"success"`
}

func viewerCtxFromRequest(ctx context.Context, r *http.Request) (context.Context, *ent.Order, error) {
	returnPortalID := pulid.ID(r.URL.Query().Get(delivrioroutes.QueryReturnPortalID))
	orderPublicID := r.URL.Query().Get(delivrioroutes.QueryOrderPublicID)
	email := r.URL.Query().Get(delivrioroutes.QueryEmail)

	if len(returnPortalID) == 0 || len(orderPublicID) == 0 || len(email) == 0 {
		return nil, nil, fmt.Errorf("validate: invalid endpoint params")
	}

	// Only allow Background ctx in limited scope...
	ctx = viewer.MergeViewerContextID(r.Context(), viewer.UserViewer{
		Role: viewer.Background,
	})

	db := ent.FromContext(r.Context())
	ord, err := db.Order.Query().
		Where(
			order.And(
				order.Or(
					order.OrderPublicIDEqualFold(orderPublicID),
					// Should be relatively unique with email and will
					// fail if results > 1 while solving prefix issue:
					// #1001 vs 1001
					order.OrderPublicIDContainsFold(orderPublicID),
				),
				order.HasConnectionWith(
					connection.HasReturnPortalWith(
						returnportal.ID(returnPortalID),
					),
				),
				order.HasColliWith(
					colli.HasRecipientWith(address.EmailEqualFold(email)),
				),
			),
		).Only(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("validate: %w", err)
	}

	return viewer.MergeViewerContextID(ctx, viewer.UserViewer{
		Role:   viewer.Anonymous,
		Tenant: ord.TenantID,
	}), ord, nil

}

func AddReturnOrderDeliveryOptionsAndRequestLabel(w http.ResponseWriter, r *http.Request) {
	ctx, _, err := viewerCtxFromRequest(r.Context(), r)
	if err != nil {
		// TODO: don't show the error publicly
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: query order: %s", err)})
		return
	}

	ctx = viewer.EnableAnonymousOverride(ctx)
	cli := ent.FromContext(ctx)

	var returnData AddReturnOrderDeliveryOptionsInput
	err = httputils.UnmarshalRequestBody(r, &returnData)
	if err != nil {
		// TODO: don't show the error publicly
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: unmarshal: %s", err)})
		return
	}

	for _, o := range returnData.DeliveryOptions {

		err := cli.ReturnColli.Update().
			SetDeliveryOptionID(o.DeliveryOptionID).
			Where(returncolli.ID(o.ReturnColliID)).
			Exec(ctx)
		if err != nil {
			// TODO: don't show the error publicly
			httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: set DO: %s", err)})
			return
		}

		cb, err := cli.DeliveryOption.Query().
			Where(deliveryoption.ID(o.DeliveryOptionID)).
			QueryCarrier().
			QueryCarrierBrand().
			Only(ctx)
		if err != nil {
			// TODO: don't show the error publicly
			httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: query do->cb: %s", err)})
			return
		}

		returnColliUpdate := cli.ReturnColli.Update()

		switch cb.InternalID {
		case carrierbrand.InternalIDPostNord:
			response, err := postnordapis.FetchLabelPostNord(ctx, o)
			if err != nil {
				// TODO: don't show the error publicly
				httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: pn label: %s", err)})
				return
			}
			returnColliUpdate.SetLabelPdf(response.LabelPDF)
			returnColliUpdate.SetQrCodePng(response.QRCodePNG)
			break
		case carrierbrand.InternalIDGLS:
			labelPDF, err := glsapis.FetchSingleLabelGLS(ctx, o)
			if err != nil {
				// TODO: don't show the error publicly
				httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: gls label: %s", err)})
				return
			}
			returnColliUpdate.SetLabelPdf(labelPDF)
			break
		case carrierbrand.InternalIDUSPS:
			labelResponse, err := uspsapis.FetchSingleLabel(ctx, o)
			if err != nil {
				// TODO: don't show the error publicly
				httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: USPS label: %s", err)})
				return
			}
			returnColliUpdate.SetLabelPdf(labelResponse.Responseb64PDF)
			break
		default:
			if err != nil {
				// TODO: don't show the error publicly
				httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": "unsupported carrier"})
				return
			}
		}

		err = returnColliUpdate.
			// Check this is the right ID since it is a return colli
			Where(returncolli.ID(o.ReturnColliID)).
			Exec(history.NewConfig(ctx).
				SetOrigin(changehistory.OriginWebClient).
				SetDescription("User selected return delivery option").
				Ctx())
		if err != nil {
			// TODO: don't show the error publicly
			httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: update colli: %s", err)})
			return
		}

		err = updateStatus(ctx, o)
		if err != nil {
			// TODO: don't show the error publicly
			httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": fmt.Sprintf("returns: add DO: update: %s", err)})
			return
		}

	}

	httputils.JSONResponse(w, http.StatusOK, AddReturnOrderDeliveryOptionsOutput{Success: true})
	return
}

func updateStatus(ctx context.Context, o common.ReturnOrderDeliveryOptionsColliIDs) error {
	db := ent.FromContext(ctx)

	current, err := db.ReturnColli.Query().
		Where(returncolli.ID(o.ReturnColliID)).
		Only(ctx)
	if err != nil {
		return err
	}

	err = db.ReturnColli.Update().
		Where(returncolli.ID(o.ReturnColliID)).
		SetStatus(returncolli.StatusPending).
		SetDeliveryOptionID(o.DeliveryOptionID).
		Exec(
			history.NewConfig(ctx).
				SetOrigin(changehistory.OriginWebClient).
				SetDescription(fmt.Sprintf("User updated status from %s to %s", current.Status, returncolli.StatusPending)).
				Ctx(),
		)
	if err != nil {
		return err
	}

	return nil
}

type CreateReturnOrderInput struct {
	PortalID   string                       `json:"portal_id"`
	OrderLines []CreateReturnOrderItemInput `json:"order_lines"`
	// Gets duplicated across all posted collis for now
	Comment string `json:"comment"`
}

type CreateReturnOrderItemInput struct {
	ClaimID     pulid.ID `json:"claim_id"`
	OrderLineID pulid.ID `json:"order_line_id"`
	Units       int      `json:"units"`
}

type CreateReturnColliDeliveryOption struct {
	DeliveryOptionID pulid.ID `json:"delivery_option_id"`
	LogoURL          string   `json:"logo_url"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	FormattedPrice   string   `json:"formatted_price"`
}

type CreateReturnColliOutput struct {
	ReturnColliID            pulid.ID                          `json:"return_colli_id"`
	SelectedDeliveryOptionID pulid.ID                          `json:"selected_delivery_option_id"`
	AvailableDeliveryOptions []CreateReturnColliDeliveryOption `json:"available_delivery_options"`
}

type CreateReturnOrderOuput struct {
	ReturnCollis []CreateReturnColliOutput `json:"return_collis"`
}

func CreateReturnOrder(w http.ResponseWriter, r *http.Request) {

	ctx, ord, err := viewerCtxFromRequest(r.Context(), r)
	if err != nil {
		httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "Order ID and Email required"})
		return
	}

	ctx = viewer.EnableAnonymousOverride(ctx)
	cli := ent.FromContext(ctx)

	var returnData CreateReturnOrderInput
	err = httputils.UnmarshalRequestBody(r, &returnData)
	if err != nil {
		// TODO: don't show the error publicly
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}

	allOrderLineIDs := make([]pulid.ID, 0)
	units := make(map[pulid.ID]returns.ReturnLineItem, 0)
	for _, o := range returnData.OrderLines {
		allOrderLineIDs = append(allOrderLineIDs, o.OrderLineID)
		units[o.OrderLineID] = returns.ReturnLineItem{
			Units:   o.Units,
			ClaimID: o.ClaimID,
		}
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		// TODO: don't show the error publicly
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}
	defer tx.Rollback()
	ctx = ent.NewTxContext(ctx, tx)

	// Delete all previous return collis where
	// the return was not finished (still Open)
	err = tx.ReturnColli.Update().
		SetStatus(returncolli.StatusDeleted).
		Where(returncolli.HasReturnOrderLineWith(
			returnorderline.HasOrderLineWith(
				orderline.IDIn(allOrderLineIDs...),
				orderline.HasReturnOrderLineWith(
					returnorderline.HasReturnColliWith(
						returncolli.StatusEQ(returncolli.StatusOpened),
					),
				),
			),
		)).Exec(
		history.NewConfig(ctx).
			SetOrigin(changehistory.OriginWebClient).
			SetDescription(fmt.Sprintf("User created new return, automatic status change to deleted for unfinished return")).
			Ctx(),
	)
	if err != nil {
		// TODO: don't show the error publicly
		tx.Rollback()
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}

	returnColliIDs, err := returns.CreateReturn(
		history.NewConfig(ctx).
			SetOrigin(changehistory.OriginWebClient).
			SetDescription("User selected return delivery option").
			Ctx(),
		ord.ID,
		pulid.ID(returnData.PortalID),
		units,
		returnData.Comment,
	)
	if err != nil {
		// TODO: don't show the error publicly
		tx.Rollback()
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		// TODO: don't show the error publicly
		tx.Rollback()
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}

	returnCollis := make([]CreateReturnColliOutput, 0)
	for _, colliID := range returnColliIDs {
		res, err := deliveryoptions.FromReturnColliID(ctx, colliID)
		if err != nil {
			// TODO: don't show the error publicly
			httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
			return
		}
		availableDeliveryOptions := make([]CreateReturnColliDeliveryOption, 0)
		for _, do := range res {

			/*			customerCountry, err := cli.ReturnColli.Query().
							Where(returncolli.ID(colliID)).
							QueryRecipient().
							QueryCountry().
							OnlyID(ctx)
						if err != nil {
							// TODO: don't show the error publicly
							httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
							return
						}

						price, err := cli.DeliveryOptionPrice.Query().
							Where(deliveryoptionprice.HasDeliveryOptionWith(deliveryoption.ID(do.DeliveryOptionID))).
							WithCurrency().
							Only(ctx)
						if err != nil {
							// TODO: don't show the error publicly
							httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
							return
						}

						calcPrice, currency, err := deliveryoptions.CalculatePrice(
							ctx,
							time.Now(),
							price,
							customerCountry,
							// TODO: fix this
							nil,
						)
						if err != nil {
							// TODO: don't show the error publicly
							httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
							return
						}

						// Maybe need some more localize formatting here?
						formattedPrice := fmt.Sprintf("%v %v", calcPrice, currency.Display)
						if calcPrice == 0 {
							// Localize
							formattedPrice = "Free"
						}*/

			if do.Status == deliveryoptions.DeliveryOptionBrandNameStatusNotAvailable {
				continue
			}

			availableDeliveryOptions = append(availableDeliveryOptions, CreateReturnColliDeliveryOption{
				DeliveryOptionID: do.DeliveryOptionID,
				LogoURL:          "/static/images/ship.svg",
				Name:             do.Name,
				Description:      do.Description,
				FormattedPrice:   fmt.Sprintf("%s %s", do.Price, do.Currency.Display),
			})
		}

		if len(availableDeliveryOptions) == 0 {
			httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": "No available return options could be found"})
			return
		}

		returnCollis = append(returnCollis, CreateReturnColliOutput{
			ReturnColliID:            colliID,
			AvailableDeliveryOptions: availableDeliveryOptions,
		})
	}

	httputils.JSONResponse(w, http.StatusOK, CreateReturnOrderOuput{
		ReturnCollis: returnCollis,
	})
	return
}

type ReturnOrderViewPackageItem struct {
	OrderLineID pulid.ID `json:"order_line_id"`
	Name        string   `json:"name"`
	VariantName string   `json:"variant_name"`
	Quantity    int      `json:"quantity"`
	ImageURL    string   `json:"image_url"`
}

type ReturnOrderViewPackage struct {
	Items []ReturnOrderViewPackageItem `json:"items"`
}

type ReturnReason struct {
	ID          pulid.ID `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
}

type ReturnOrderViewOutput struct {
	Packages      []ReturnOrderViewPackage `json:"packages"`
	OrderDate     time.Time                `json:"order_date"`
	OrderID       string                   `json:"order_id"`
	ReturnReasons []ReturnReason           `json:"return_reasons"`
}

func ReturnOrderView(w http.ResponseWriter, r *http.Request) {
	ctx, ord, err := viewerCtxFromRequest(r.Context(), r)
	if err != nil {
		msg := i18n.Value(language.InternalIDEN, i18n.ReturnsAuthenticationNotFound)
		if ent.IsNotFound(err) || ent.IsNotSingular(err) {
			msg = i18n.Value(language.InternalIDEN, i18n.ReturnsOrderNotFound)
		}

		httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": msg})
		return
	}

	ctx = viewer.EnableAnonymousOverride(ctx)
	cli := ent.FromContext(ctx)

	// TODO: check connection
	col, err := ord.QueryColli().
		WithOrder().
		WithOrderLines().
		All(ctx)
	if err != nil {
		// TODO: don't show the error publicly
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}

	if len(col) == 0 {
		httputils.JSONResponse(w, http.StatusNotFound, httputils.Map{"message": i18n.Value(language.InternalIDEN, i18n.ReturnsMissingCollis)})
		return
	}

	packages := make([]ReturnOrderViewPackage, 0)
	for _, col := range col {
		items := make([]ReturnOrderViewPackageItem, 0)
		for _, i := range col.Edges.OrderLines {
			variant, err := i.QueryProductVariant().
				WithProductImage().
				WithProduct().
				Only(ctx)
			if err != nil {
				// TODO: don't show the error publicly
				httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
				return
			}

			imgURL := ""
			if len(variant.Edges.ProductImage) > 0 {
				imgURL = variant.Edges.ProductImage[0].URL
			} else {
				productImg, err := variant.QueryProduct().
					QueryProductImage().
					First(ctx)
				if err != nil && !ent.IsNotFound(err) {
					// TODO: don't show the error publicly
					httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
					return
				}
				if productImg != nil {
					imgURL = productImg.URL
				}
			}

			items = append(items, ReturnOrderViewPackageItem{
				OrderLineID: i.ID,
				Name:        variant.Edges.Product.Title,
				VariantName: variant.Description,
				Quantity:    i.Units,
				ImageURL:    imgURL,
			})
		}
		packages = append(packages, ReturnOrderViewPackage{Items: items})
	}

	returnPortalID := pulid.ID(r.URL.Query().Get(delivrioroutes.QueryReturnPortalID))
	reasons, err := cli.ReturnPortalClaim.Query().
		Where(returnportalclaim.HasReturnPortalWith(returnportal.ID(returnPortalID))).
		All(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, httputils.Map{"message": err.Error()})
		return
	}

	returnReasons := make([]ReturnReason, 0)
	for _, r := range reasons {
		returnReasons = append(returnReasons, ReturnReason{
			ID:          r.ID,
			Name:        r.Name,
			Description: r.Description,
		})
	}

	httputils.JSONResponse(w, http.StatusOK, ReturnOrderViewOutput{
		Packages:      packages,
		OrderDate:     col[0].Edges.Order.CreatedAt,
		OrderID:       col[0].Edges.Order.OrderPublicID,
		ReturnReasons: returnReasons,
	})
	return

}
