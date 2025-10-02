package productmodels

import "time"

type Products struct {
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	ID                uint64     `json:"id"`
	Title             *string    `json:"title,omitempty"`
	BodyHTML          *string    `json:"body_html,omitempty"`
	Vendor            *string    `json:"vendor,omitempty"`
	ProductType       *string    `json:"product_type,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	Handle            *string    `json:"handle,omitempty"`
	UpdatedAt         *string    `json:"updated_at,omitempty"`
	PublishedAt       *string    `json:"published_at,omitempty"`
	TemplateSuffix    *string    `json:"template_suffix,omitempty"`
	Status            *string    `json:"status,omitempty"`
	PublishedScope    *string    `json:"published_scope,omitempty"`
	Tags              string     `json:"tags"`
	AdminGraphqlAPIID *string    `json:"admin_graphql_api_id,omitempty"`
	Variants          []Variant  `json:"variants,omitempty"`
	Options           []Option   `json:"options,omitempty"`
	Images            []Image    `json:"images,omitempty"`
	Image             *Image     `json:"image"`
}

type Image struct {
	ID                *int64      `json:"id,omitempty"`
	ProductID         *int64      `json:"product_id,omitempty"`
	Position          *int64      `json:"position,omitempty"`
	CreatedAt         *string     `json:"created_at,omitempty"`
	UpdatedAt         *string     `json:"updated_at,omitempty"`
	Alt               interface{} `json:"alt"`
	Width             *int64      `json:"width,omitempty"`
	Height            *int64      `json:"height,omitempty"`
	Src               *string     `json:"src,omitempty"`
	VariantIDS        []int64     `json:"variant_ids,omitempty"`
	AdminGraphqlAPIID *string     `json:"admin_graphql_api_id,omitempty"`
}

type Option struct {
	ID        *int64   `json:"id,omitempty"`
	ProductID *int64   `json:"product_id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Position  *int64   `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type Variant struct {
	ID                   uint64             `json:"id"`
	ProductID            *int64             `json:"product_id,omitempty"`
	Title                *string            `json:"title,omitempty"`
	Price                *string            `json:"price,omitempty"`
	Sku                  *string            `json:"sku,omitempty"`
	Position             *int64             `json:"position,omitempty"`
	InventoryPolicy      *string            `json:"inventory_policy,omitempty"`
	CompareAtPrice       interface{}        `json:"compare_at_price"`
	FulfillmentService   *string            `json:"fulfillment_service,omitempty"`
	InventoryManagement  *string            `json:"inventory_management,omitempty"`
	Option1              *string            `json:"option1,omitempty"`
	Option2              interface{}        `json:"option2"`
	Option3              interface{}        `json:"option3"`
	CreatedAt            *string            `json:"created_at,omitempty"`
	UpdatedAt            *string            `json:"updated_at,omitempty"`
	Taxable              *bool              `json:"taxable,omitempty"`
	Barcode              *string            `json:"barcode,omitempty"`
	Grams                *int64             `json:"grams,omitempty"`
	ImageID              *int64             `json:"image_id"`
	Weight               *float64           `json:"weight,omitempty"`
	WeightUnit           *string            `json:"weight_unit,omitempty"`
	InventoryItemID      *int64             `json:"inventory_item_id,omitempty"`
	InventoryQuantity    *int64             `json:"inventory_quantity,omitempty"`
	OldInventoryQuantity *int64             `json:"old_inventory_quantity,omitempty"`
	PresentmentPrices    []PresentmentPrice `json:"presentment_prices,omitempty"`
	RequiresShipping     *bool              `json:"requires_shipping,omitempty"`
	AdminGraphqlAPIID    *string            `json:"admin_graphql_api_id,omitempty"`
}

type PresentmentPrice struct {
	Price          *Price      `json:"price,omitempty"`
	CompareAtPrice interface{} `json:"compare_at_price"`
}

type Price struct {
	Amount       *string `json:"amount,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
}

type InventoryItems struct {
	InventoryItems []InventoryItem `json:"inventory_items,omitempty"`
}

type InventoryItem struct {
	Cost string `json:"cost,omitempty"`
	// Not always set
	CountryCodeOfOrigin          *string                       `json:"country_code_of_origin,omitempty"`
	CountryHarmonizedSystemCodes []CountryHarmonizedSystemCode `json:"country_harmonized_system_codes,omitempty"`
	CreatedAt                    string                        `json:"created_at,omitempty"`
	HarmonizedSystemCode         *string                       `json:"harmonized_system_code,omitempty"`
	ID                           int64                         `json:"id,omitempty"`
	ProvinceCodeOfOrigin         string                        `json:"province_code_of_origin,omitempty"`
	Sku                          *string                       `json:"sku,omitempty"`
	Tracked                      bool                          `json:"tracked,omitempty"`
	UpdatedAt                    string                        `json:"updated_at,omitempty"`
	RequiresShipping             bool                          `json:"requires_shipping,omitempty"`
}

type CountryHarmonizedSystemCode struct {
	HarmonizedSystemCode string `json:"harmonized_system_code,omitempty"`
	CountryCode          string `json:"country_code,omitempty"`
}
