package delivrioroutes

const (
	PrintClientPing         = "/print-client/ping"
	PrintClientRequestLabel = "/print-client/label"
	RegisterScan            = "/register-scan"
	ScanList                = "/scan-list"
	UpdateStatus            = "/update-status"
	SignatureUpload         = "/signature-upload"
	ReturnView              = "/return-view"
	CreateReturnOrder       = "/return-create"

	ReturnDeliveryOptions = "/return-delivery-options"
	QueryReturnPortalID   = "return-portal-id"
	QueryOrderPublicID    = "order-public-id"
	QueryEmail            = "email"

	ReturnLabel          = "/return-label"
	ReturnLabelDownload  = "/return-label/download"
	ReturnLabelPNG       = "/return-label/png"
	ReturnQRCode         = "/return-qr-code"
	ReturnQRCodeDownload = "/return-qr-code/download"
	QueryGoland          = "/query-goland"

	API   = "/api"
	Query = "/query"
	Graph = "/graph"

	NodeHealth                = "/health-check"
	ShopifyLookup             = "/shopify-lookup"
	ShopifyLookupPickupPoints = "/shopify-lookup-pickup-points"
	Restv1APIDocs             = "/rest/v1/api-docs"
	ReturnPolyfills           = "/return-portal-assets/polyfills.js"
	ReturnMain                = "/return-portal-assets/main.js"
	Uploads                   = "/uploads"
	Static                    = "/static"
	Images                    = "/images"

	AddressLookup   = "/addressLookup"
	RequestEmail    = "/requestEmail"
	RequestPassword = "/resetPassword"
	Register        = "/register"
	Login           = "/login"

	// Customer REST endpoints
	Label         = "/label"
	Order         = "/order"
	Orders        = "/orders"
	Products      = "/products"
	Shipment      = "/shipment"
	Shipments     = "/shipments"
	Return        = "/return"
	ReturnReasons = "/return/reasons"
)

const QueryParamReturnColliID = "return-colli-id"
const QueryParamOrderPublicID = "order-public-id"
const QueryParamLookupID = "lookup-id"
