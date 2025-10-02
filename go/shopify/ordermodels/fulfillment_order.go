package ordermodels

type FulfillmentOrders struct {
	FulfillmentOrders []FulfillmentOrder `json:"fulfillment_orders"`
}

type FulfillmentOrder struct {
	ID                  int64                      `json:"id"`
	ShopID              int64                      `json:"shop_id"`
	OrderID             int64                      `json:"order_id"`
	AssignedLocationID  int64                      `json:"assigned_location_id"`
	RequestStatus       string                     `json:"request_status"`
	Status              string                     `json:"status"`
	SupportedActions    []string                   `json:"supported_actions"`
	Destination         Destination                `json:"destination"`
	LineItems           []FulfillmentOrderLineItem `json:"line_items"`
	InternationalDuties interface{}                `json:"international_duties"`
	FulfillAt           string                     `json:"fulfill_at"`
	FulfillBy           interface{}                `json:"fulfill_by"`
	FulfillmentHolds    []interface{}              `json:"fulfillment_holds"`
	CreatedAt           string                     `json:"created_at"`
	UpdatedAt           string                     `json:"updated_at"`
	DeliveryMethod      DeliveryMethod             `json:"delivery_method"`
	AssignedLocation    AssignedLocation           `json:"assigned_location"`
	MerchantRequests    []interface{}              `json:"merchant_requests"`
}

type AssignedLocation struct {
	Address1    string      `json:"address1"`
	Address2    interface{} `json:"address2"`
	City        string      `json:"city"`
	CountryCode string      `json:"country_code"`
	LocationID  int64       `json:"location_id"`
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	Province    string      `json:"province"`
	Zip         string      `json:"zip"`
}

type DeliveryMethod struct {
	ID                  int64       `json:"id"`
	MethodType          string      `json:"method_type"`
	MinDeliveryDateTime interface{} `json:"min_delivery_date_time"`
	MaxDeliveryDateTime interface{} `json:"max_delivery_date_time"`
}

type Destination struct {
	ID        int64       `json:"id"`
	Address1  string      `json:"address1"`
	Address2  interface{} `json:"address2"`
	City      string      `json:"city"`
	Company   interface{} `json:"company"`
	Country   string      `json:"country"`
	Email     string      `json:"email"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Phone     interface{} `json:"phone"`
	Province  interface{} `json:"province"`
	Zip       string      `json:"zip"`
}

// Seems like we only get a subset of properties compared to /orders response?
type FulfillmentOrderLineItem struct {
	ID                  int64 `json:"id"`
	ShopID              int64 `json:"shop_id"`
	FulfillmentOrderID  int64 `json:"fulfillment_order_id"`
	Quantity            int64 `json:"quantity"`
	LineItemID          int64 `json:"line_item_id"`
	InventoryItemID     int64 `json:"inventory_item_id"`
	FulfillableQuantity int64 `json:"fulfillable_quantity"`
	VariantID           int64 `json:"variant_id"`
}

type CreateFulfillment struct {
	Fulfillment Fulfillment `json:"fulfillment"`
}

type Fulfillment struct {
	Message                     string                        `json:"message"`
	NotifyCustomer              bool                          `json:"notify_customer"`
	TrackingInfo                TrackingInfo                  `json:"tracking_info"`
	LineItemsByFulfillmentOrder []LineItemsByFulfillmentOrder `json:"line_items_by_fulfillment_order"`
}

type LineItemsByFulfillmentOrder struct {
	FulfillmentOrderID        int64                            `json:"fulfillment_order_id"`
	FulfillmentOrderLineItems []CreateFulfillmentOrderLineItem `json:"fulfillment_order_line_items"`
}

type CreateFulfillmentOrderLineItem struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
}

// https://shopify.dev/docs/api/admin-rest/2023-04/resources/fulfillment
// For list of companies
type TrackingInfo struct {
	Company string `json:"company,omitempty"`
	Number  string `json:"number"`
	URL     string `json:"url,omitempty"`
}
