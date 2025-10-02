package restcustomer

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/product"
	"delivrio.io/go/ent/producttag"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"net/url"
	"strings"
)

type ProductsCreateRequest struct {
	Products []ProductCreate `json:"products"`
}

type ProductsCreateResponse struct {
	// 1:1 output corresponding to posted Orders
	Products []ProductsSaved `json:"products"`
}

type ProductsSaved struct {
	// Whether the order was saved.
	Success bool `json:"success"`
	// Save error. Present when Success is false.
	Error string `json:"error"`
	// Internal DELIVRIO ID. Present when save is successful.
	ProductID string `json:"product_id"`
}

type ImageCreate struct {
	// The full URL to the image hosted on an external system
	// example: https://cdn.example.com/a/image/path/12345.jpg
	URL string `json:"url" validate:"required"`
	// ID from system external to DELIVRIO
	// example: 99999911111888888222222
	ExternalID string `json:"external_id" validate:"required"`
	// The array indexes of the variants this image should be associated with.
	// example: [1, 5, 9]
	VariantIndexes *[]int `json:"variant_indexes"`
}

type ProductVariantCreate struct {
	// ID from system external to DELIVRIO
	// example: 99999911111888888222222
	ExternalID *string `json:"external_id"`
	// Description of the specific variant
	// example: A green t-shirt
	Description *string `json:"description"`
	// EAN Number for the variant
	// example: 1234567891234
	EAN *string `json:"ean"`
	// Length cm
	// example: 10
	Length *int `json:"length"`
	// Width cm
	// example: 10
	Width *int `json:"width"`
	// Height cm
	// example: 10
	Height *int `json:"height"`
}

type ProductCreate struct {
	// ID from system external to DELIVRIO
	// example: 99999911111888888222222
	ExternalID *string `json:"external_id"`
	// The title of the product variant group
	Title string `json:"title" validate:"required"`
	// A description of the product which may include HTML
	BodyHTML *string `json:"body_html"`
	// Visibility status of the product.
	// example: active
	Status string `json:"status"  validate:"required,oneof='active' 'draft' 'archived'"`
	// Product tags. Existing tags will have this product added.
	// Case-insensitive
	// example: ["T-shirts", "Branded", "normal-delivery"]
	Tags []string `json:"tags"`
	// Images of all product variants
	Images []ImageCreate `json:"images"`
	// ProductVariants of all product variants
	ProductVariants []ProductVariantCreate `json:"product_variants"  validate:"required"`
}

// HandleProductsCreate godoc
//
//	@Summary		Create products
//	@Description	Responds with the save info for each POSTed product. Be aware partial success will return status 207.
//	@Tags			products
//	@ID				create-products
//	@Accept			json
//	@Produce		json
//	@Param			body	body	ProductsCreateRequest	true	"Product creation request"
//	@Success		200		{object}	ProductsCreateResponse
//	@Success		207		{object}	ProductsCreateResponse
//	@Failure		400		{object}	GeneralError
//	@Failure		404		{object}	GeneralError
//	@Failure		500		{object}	GeneralError
//	@Security		api_key
//	@Router			/products [post]
func HandleProductsCreate(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "HandleProductsCreate")
	defer span.End()

	var input ProductsCreateRequest
	err := httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		span.SetStatus(codes.Error, "bind input")
		span.RecordError(err)
		httputils.JSONResponse(w, http.StatusBadRequest,
			GeneralError{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		return
	}

	span.SetAttributes(
		attribute.Int("productCount", len(input.Products)),
	)

	output := ProductsCreateResponse{
		Products: make([]ProductsSaved, 0),
	}

	cli := ent.FromContext(ctx)
	multiStatus := 0

	for _, prod := range input.Products {

		tx, err := cli.Tx(r.Context())
		if err != nil {
			span.SetStatus(codes.Error, "start tx")
			span.RecordError(err)
			httputils.JSONResponse(w, http.StatusInternalServerError, GeneralError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
			return
		}
		defer tx.Rollback()

		ctx := ent.NewTxContext(r.Context(), tx)

		orderID, err := createProduct(ctx, prod)
		if err != nil {
			span.SetStatus(codes.Error, "create product")
			span.RecordError(err)
			tx.Rollback()
			multiStatus++
			output.Products = append(output.Products, ProductsSaved{
				Success: false,
				Error:   err.Error(),
			})
		} else {
			err := tx.Commit()
			if err != nil {
				span.SetStatus(codes.Error, "commit")
				span.RecordError(err)
				output.Products = append(output.Products, ProductsSaved{
					Success: false,
					Error:   err.Error(),
				})
			} else {
				output.Products = append(output.Products, ProductsSaved{
					Success:   true,
					ProductID: orderID.String(),
				})
			}
		}
	}

	if multiStatus > 0 {
		span.AddEvent("multistatus", trace.WithAttributes(attribute.Int("failCreate", multiStatus)))
		httputils.JSONResponse(w, http.StatusMultiStatus, output)
		return
	}

	span.AddEvent("createdAllProducts")
	httputils.JSONResponse(w, http.StatusOK, output)
	return

}

func createProduct(ctx context.Context, prod ProductCreate) (pulid.ID, error) {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	tags, err := createTags(ctx, prod.Tags)
	if err != nil {
		return "", err
	}

	prodCreate := tx.Product.Create().
		AddProductTags(tags...).
		SetTenantID(view.TenantID()).
		SetNillableExternalID(prod.ExternalID).
		SetNillableBodyHTML(prod.BodyHTML)

	switch strings.TrimSpace(strings.ToLower(prod.Status)) {
	case strings.ToLower(product.StatusActive.String()):
		prodCreate = prodCreate.SetStatus(product.StatusActive)
		break
	case strings.ToLower(product.StatusActive.String()):
		prodCreate = prodCreate.SetStatus(product.StatusActive)
		break
	case strings.ToLower(product.StatusActive.String()):
		prodCreate = prodCreate.SetStatus(product.StatusActive)
		break
	default:
		return "", fmt.Errorf("createproduct: unrecognized product status: %v", prod.Status)
	}

	prodSaved, err := prodCreate.Save(ctx)
	if err != nil {
		return "", fmt.Errorf("createproduct: %w", err)
	}

	prodVariantsSaved, err := createProductVariants(ctx, prod.ProductVariants, prodSaved.ID)
	if err != nil {
		return "", fmt.Errorf("createproductvariants: %w", err)
	}

	for _, img := range prod.Images {
		variantWithImage := make([]*ent.ProductVariant, 0)
		if img.VariantIndexes != nil {
			for _, i := range *img.VariantIndexes {
				variantWithImage = append(variantWithImage, prodVariantsSaved[i:i+1]...)
			}
		}
		err = createImage(ctx, img, prodSaved.ID, variantWithImage)
		if err != nil {
			return "", fmt.Errorf("createproductvariants: %w", err)
		}
	}

	return "", nil
}

func createProductVariants(ctx context.Context, variants []ProductVariantCreate, prodID pulid.ID) ([]*ent.ProductVariant, error) {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	bulk := make([]*ent.ProductVariantCreate, 0)
	for _, v := range variants {
		bulk = append(
			bulk,
			tx.ProductVariant.Create().
				SetProductID(prodID).
				SetTenantID(view.TenantID()).
				SetNillableExternalID(v.ExternalID).
				SetNillableEanNumber(v.EAN).
				SetNillableDimensionHeight(v.Height).
				SetNillableDimensionWidth(v.Width).
				SetNillableDimensionLength(v.Length).
				SetNillableDescription(v.Description),
		)
	}

	return tx.ProductVariant.CreateBulk(bulk...).Save(ctx)

}

func createImage(ctx context.Context, image ImageCreate, prodID pulid.ID, variants []*ent.ProductVariant) error {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	u, err := url.Parse(image.URL)
	if err != nil {
		return fmt.Errorf("createimage: invalid image URL: %w", err)
	}

	err = tx.ProductImage.Create().
		SetTenantID(view.TenantID()).
		SetExternalID(image.ExternalID).
		SetProductID(prodID).
		AddProductVariant(variants...).
		SetURL(u.String()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("createimage: %w", err)
	}

	return nil
}

func createTags(ctx context.Context, tags []string) ([]*ent.ProductTag, error) {
	tx := ent.TxFromContext(ctx)
	view := viewer.FromContext(ctx)

	output := make([]*ent.ProductTag, 0)
	for _, t := range tags {
		tagSearch, err := tx.ProductTag.Query().
			Where(producttag.NameEqualFold(t)).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		} else if ent.IsNotFound(err) {
			tag, err := tx.ProductTag.Create().
				SetName(t).
				SetTenantID(view.TenantID()).
				Save(ctx)
			if err != nil {
				return nil, err
			}
			output = append(output, tag)
			continue
		}

		output = append(output, tagSearch)
	}

	return output, nil

}
