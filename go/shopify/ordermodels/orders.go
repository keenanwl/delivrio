package ordermodels

type Orders struct {
	Orders []Order `json:"orders,omitempty"`
}

type Order struct {
	ID                       uint64          `json:"id,omitempty"`
	AdminGraphqlAPIID        *string         `json:"admin_graphql_api_id,omitempty"`
	AppID                    *int64          `json:"app_id,omitempty"`
	BrowserIP                *string         `json:"browser_ip,omitempty"`
	BuyerAcceptsMarketing    *bool           `json:"buyer_accepts_marketing,omitempty"`
	CancelReason             interface{}     `json:"cancel_reason"`
	CancelledAt              interface{}     `json:"cancelled_at"`
	CartToken                interface{}     `json:"cart_token"`
	CheckoutID               *int64          `json:"checkout_id,omitempty"`
	CheckoutToken            *string         `json:"checkout_token,omitempty"`
	ClientDetails            *ClientDetails  `json:"client_details,omitempty"`
	ClosedAt                 interface{}     `json:"closed_at"`
	Company                  interface{}     `json:"company"`
	Confirmed                *bool           `json:"confirmed,omitempty"`
	ContactEmail             *string         `json:"contact_email,omitempty"`
	CreatedAt                *string         `json:"created_at,omitempty"`
	Currency                 string          `json:"currency,omitempty"`
	CurrentSubtotalPrice     *string         `json:"current_subtotal_price,omitempty"`
	CurrentSubtotalPriceSet  *Set            `json:"current_subtotal_price_set,omitempty"`
	CurrentTotalDiscounts    *string         `json:"current_total_discounts,omitempty"`
	CurrentTotalDiscountsSet *Set            `json:"current_total_discounts_set,omitempty"`
	CurrentTotalDutiesSet    interface{}     `json:"current_total_duties_set"`
	CurrentTotalPrice        *string         `json:"current_total_price,omitempty"`
	CurrentTotalPriceSet     *Set            `json:"current_total_price_set,omitempty"`
	CurrentTotalTax          *string         `json:"current_total_tax,omitempty"`
	CurrentTotalTaxSet       *Set            `json:"current_total_tax_set,omitempty"`
	CustomerLocale           *string         `json:"customer_locale,omitempty"`
	DeviceID                 interface{}     `json:"device_id"`
	DiscountCodes            []interface{}   `json:"discount_codes,omitempty"`
	Email                    string          `json:"email,omitempty"`
	EstimatedTaxes           *bool           `json:"estimated_taxes,omitempty"`
	FinancialStatus          *string         `json:"financial_status,omitempty"`
	FulfillmentStatus        interface{}     `json:"fulfillment_status"`
	Gateway                  *string         `json:"gateway,omitempty"`
	LandingSite              interface{}     `json:"landing_site"`
	LandingSiteRef           interface{}     `json:"landing_site_ref"`
	LocationID               interface{}     `json:"location_id"`
	MerchantOfRecordAppID    interface{}     `json:"merchant_of_record_app_id"`
	Name                     string          `json:"name,omitempty"`
	Note                     *string         `json:"note"` // May be null - verified
	NoteAttributes           []NoteAttribute `json:"note_attributes,omitempty"`
	Number                   *int64          `json:"number,omitempty"`
	OrderNumber              *int64          `json:"order_number,omitempty"`
	OrderStatusURL           *string         `json:"order_status_url,omitempty"`
	OriginalTotalDutiesSet   interface{}     `json:"original_total_duties_set"`
	PaymentGatewayNames      []string        `json:"payment_gateway_names,omitempty"`
	Phone                    string          `json:"phone"`
	PresentmentCurrency      string          `json:"presentment_currency,omitempty"`
	ProcessedAt              *string         `json:"processed_at,omitempty"`
	ProcessingMethod         *string         `json:"processing_method,omitempty"`
	Reference                *string         `json:"reference,omitempty"`
	ReferringSite            interface{}     `json:"referring_site"`
	SourceIdentifier         *string         `json:"source_identifier,omitempty"`
	SourceName               *string         `json:"source_name,omitempty"`
	SourceURL                interface{}     `json:"source_url"`
	SubtotalPrice            *string         `json:"subtotal_price,omitempty"`
	SubtotalPriceSet         *Set            `json:"subtotal_price_set,omitempty"`
	Tags                     string          `json:"tags,omitempty"`
	TaxLines                 []interface{}   `json:"tax_lines,omitempty"`
	TaxesIncluded            *bool           `json:"taxes_included,omitempty"`
	Test                     *bool           `json:"test,omitempty"`
	Token                    *string         `json:"token,omitempty"`
	TotalDiscounts           *string         `json:"total_discounts,omitempty"`
	TotalDiscountsSet        *Set            `json:"total_discounts_set,omitempty"`
	TotalLineItemsPrice      *string         `json:"total_line_items_price,omitempty"`
	TotalLineItemsPriceSet   *Set            `json:"total_line_items_price_set,omitempty"`
	TotalOutstanding         *string         `json:"total_outstanding,omitempty"`
	TotalPrice               *string         `json:"total_price,omitempty"`
	TotalPriceSet            *Set            `json:"total_price_set,omitempty"`
	TotalShippingPriceSet    *Set            `json:"total_shipping_price_set,omitempty"`
	TotalTax                 *string         `json:"total_tax,omitempty"`
	TotalTaxSet              *Set            `json:"total_tax_set,omitempty"`
	TotalTipReceived         *string         `json:"total_tip_received,omitempty"`
	TotalWeight              *int64          `json:"total_weight,omitempty"`
	UpdatedAt                *string         `json:"updated_at,omitempty"`
	UserID                   *int64          `json:"user_id,omitempty"`
	BillingAddress           *Address        `json:"billing_address,omitempty"`
	Customer                 *Customer       `json:"customer,omitempty"`
	DiscountApplications     []interface{}   `json:"discount_applications,omitempty"`
	Fulfillments             []interface{}   `json:"fulfillments,omitempty"`
	LineItems                []LineItem      `json:"line_items,omitempty"`
	PaymentTerms             *PaymentTerms   `json:"payment_terms,omitempty"`
	Refunds                  []Refund        `json:"refunds,omitempty"`
	ShippingAddress          *Address        `json:"shipping_address,omitempty"`
	ShippingLines            []ShippingLine  `json:"shipping_lines,omitempty"`
}

type NoteAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Address struct {
	FirstName    string  `json:"first_name,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	City         string  `json:"city,omitempty"`
	Zip          string  `json:"zip,omitempty"`
	Province     string  `json:"province"`
	Country      string  `json:"country,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Address2     string  `json:"address2"`
	Company      string  `json:"company,omitempty"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name,omitempty"`
	CountryCode  *string `json:"country_code,omitempty"`
	ProvinceCode string  `json:"province_code"`
	ID           int64   `json:"id,omitempty"`
	CustomerID   int64   `json:"customer_id,omitempty"`
	CountryName  string  `json:"country_name,omitempty"`
	Default      bool    `json:"default,omitempty"`
}

type ClientDetails struct {
	AcceptLanguage interface{} `json:"accept_language"`
	BrowserHeight  interface{} `json:"browser_height"`
	BrowserIP      *string     `json:"browser_ip,omitempty"`
	BrowserWidth   interface{} `json:"browser_width"`
	SessionHash    interface{} `json:"session_hash"`
	UserAgent      *string     `json:"user_agent,omitempty"`
}

type Set struct {
	ShopMoney        Money `json:"shop_money,omitempty"`
	PresentmentMoney Money `json:"presentment_money,omitempty"`
}

type Money struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}

type Customer struct {
	ID                        *int64                 `json:"id,omitempty"`
	Email                     *string                `json:"email,omitempty"`
	AcceptsMarketing          *bool                  `json:"accepts_marketing,omitempty"`
	CreatedAt                 *string                `json:"created_at,omitempty"`
	UpdatedAt                 *string                `json:"updated_at,omitempty"`
	FirstName                 *string                `json:"first_name,omitempty"`
	LastName                  *string                `json:"last_name,omitempty"`
	State                     *string                `json:"state,omitempty"`
	Note                      interface{}            `json:"note"`
	VerifiedEmail             *bool                  `json:"verified_email,omitempty"`
	MultipassIdentifier       interface{}            `json:"multipass_identifier"`
	TaxExempt                 *bool                  `json:"tax_exempt,omitempty"`
	Phone                     interface{}            `json:"phone"`
	EmailMarketingConsent     *EmailMarketingConsent `json:"email_marketing_consent,omitempty"`
	SMSMarketingConsent       interface{}            `json:"sms_marketing_consent"`
	Tags                      *string                `json:"tags,omitempty"`
	Currency                  string                 `json:"currency,omitempty"`
	AcceptsMarketingUpdatedAt *string                `json:"accepts_marketing_updated_at,omitempty"`
	MarketingOptInLevel       *string                `json:"marketing_opt_in_level,omitempty"`
	TaxExemptions             []interface{}          `json:"tax_exemptions,omitempty"`
	AdminGraphqlAPIID         *string                `json:"admin_graphql_api_id,omitempty"`
	DefaultAddress            *Address               `json:"default_address,omitempty"`
}

type EmailMarketingConsent struct {
	State            *string `json:"state,omitempty"`
	OptInLevel       *string `json:"opt_in_level,omitempty"`
	ConsentUpdatedAt *string `json:"consent_updated_at,omitempty"`
}

type LineItem struct {
	ID                         uint64               `json:"id,omitempty"`
	AdminGraphqlAPIID          *string              `json:"admin_graphql_api_id,omitempty"`
	FulfillableQuantity        *int64               `json:"fulfillable_quantity,omitempty"`
	FulfillmentService         *string              `json:"fulfillment_service,omitempty"`
	FulfillmentStatus          interface{}          `json:"fulfillment_status"`
	GiftCard                   *bool                `json:"gift_card,omitempty"`
	Grams                      *int64               `json:"grams,omitempty"`
	Name                       *string              `json:"name,omitempty"`
	Price                      string               `json:"price,omitempty"`
	PriceSet                   *Set                 `json:"price_set,omitempty"`
	ProductExists              *bool                `json:"product_exists,omitempty"`
	ProductID                  *int64               `json:"product_id,omitempty"`
	Properties                 []interface{}        `json:"properties,omitempty"`
	Quantity                   int                  `json:"quantity,omitempty"`
	RequiresShipping           *bool                `json:"requires_shipping,omitempty"`
	Sku                        *string              `json:"sku,omitempty"`
	Taxable                    *bool                `json:"taxable,omitempty"`
	Title                      *string              `json:"title,omitempty"`
	TotalDiscount              *string              `json:"total_discount,omitempty"`
	TotalDiscountSet           *Set                 `json:"total_discount_set,omitempty"`
	VariantID                  uint64               `json:"variant_id,omitempty"`
	VariantInventoryManagement *string              `json:"variant_inventory_management,omitempty"`
	VariantTitle               *string              `json:"variant_title,omitempty"`
	Vendor                     *string              `json:"vendor,omitempty"`
	TaxLines                   []interface{}        `json:"tax_lines,omitempty"`
	Duties                     []interface{}        `json:"duties,omitempty"`
	DiscountAllocations        []DiscountAllocation `json:"discount_allocations,omitempty"`
}

type DiscountAllocation struct {
	Amount                   string    `json:"amount,omitempty"`
	AmountSet                AmountSet `json:"amount_set,omitempty"`
	DiscountApplicationIndex int64     `json:"discount_application_index,omitempty"`
}

type AmountSet struct {
	ShopMoney        Money `json:"shop_money,omitempty"`
	PresentmentMoney Money `json:"presentment_money,omitempty"`
}

type Refund struct {
	ID                int64            `json:"id,omitempty"`
	AdminGraphqlAPIID *string          `json:"admin_graphql_api_id,omitempty"`
	CreatedAt         string           `json:"created_at,omitempty"`
	Note              interface{}      `json:"note"`
	OrderID           int64            `json:"order_id,omitempty"`
	ProcessedAt       *string          `json:"processed_at,omitempty"`
	Restock           *bool            `json:"restock,omitempty"`
	TotalDutiesSet    *Set             `json:"total_duties_set,omitempty"`
	UserID            *int64           `json:"user_id,omitempty"`
	OrderAdjustments  []interface{}    `json:"order_adjustments,omitempty"`
	Transactions      []interface{}    `json:"transactions,omitempty"`
	RefundLineItems   []RefundLineItem `json:"refund_line_items,omitempty"`
	Duties            []interface{}    `json:"duties,omitempty"`
}

type RefundLineItem struct {
	ID          uint64      `json:"id,omitempty"`
	LineItemID  uint64      `json:"line_item_id,omitempty"`
	LocationID  *int64      `json:"location_id,omitempty"`
	Quantity    int         `json:"quantity,omitempty"`
	RestockType *string     `json:"restock_type,omitempty"`
	Subtotal    interface{} `json:"subtotal,omitempty"`
	SubtotalSet *Set        `json:"subtotal_set,omitempty"`
	TotalTax    *float64    `json:"total_tax,omitempty"`
	TotalTaxSet *Set        `json:"total_tax_set,omitempty"`
	LineItem    *LineItem   `json:"line_item,omitempty"`
}

type PaymentTerms struct {
	ID               *int64            `json:"id,omitempty"`
	CreatedAt        *string           `json:"created_at,omitempty"`
	DueInDays        interface{}       `json:"due_in_days"`
	PaymentSchedules []PaymentSchedule `json:"payment_schedules,omitempty"`
	PaymentTermsName *string           `json:"payment_terms_name,omitempty"`
	PaymentTermsType *string           `json:"payment_terms_type,omitempty"`
	UpdatedAt        *string           `json:"updated_at,omitempty"`
}

type PaymentSchedule struct {
	ID          *int64      `json:"id,omitempty"`
	Amount      *string     `json:"amount,omitempty"`
	Currency    string      `json:"currency,omitempty"`
	IssuedAt    interface{} `json:"issued_at"`
	DueAt       interface{} `json:"due_at"`
	CompletedAt interface{} `json:"completed_at"`
	CreatedAt   *string     `json:"created_at,omitempty"`
	UpdatedAt   *string     `json:"updated_at,omitempty"`
}

type ShippingLine struct {
	ID                            *int64        `json:"id,omitempty"`
	CarrierIdentifier             *string       `json:"carrier_identifier,omitempty"`
	Code                          *string       `json:"code,omitempty"`
	DeliveryCategory              interface{}   `json:"delivery_category"`
	DiscountedPrice               *string       `json:"discounted_price,omitempty"`
	DiscountedPriceSet            *Set          `json:"discounted_price_set,omitempty"`
	Phone                         interface{}   `json:"phone"`
	Price                         *string       `json:"price,omitempty"`
	PriceSet                      *Set          `json:"price_set,omitempty"`
	RequestedFulfillmentServiceID interface{}   `json:"requested_fulfillment_service_id"`
	Source                        *string       `json:"source,omitempty"`
	Title                         *string       `json:"title,omitempty"`
	TaxLines                      []interface{} `json:"tax_lines,omitempty"`
	DiscountAllocations           []interface{} `json:"discount_allocations,omitempty"`
}
