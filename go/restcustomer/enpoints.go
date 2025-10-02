package restcustomer

import (
	"go.opentelemetry.io/otel"
	"net/http"
)

var (
	tracer = otel.Tracer("restcustomer")
)

type ParcelLabel struct {
	// URL to request PDF label for this parcel
	// Can the link be authenticated?
	Link string `json:"link"`
	// Base64 encoded PDF label.
	LabelPDF string `json:"label_pdf"`
	// Error order lines may occur if the given order line is already part of a dispatched shipment
	ErrorOrderLines []ErrorOrderLines `json:"error_order_lines"`
}

type ErrorOrderLines struct {
	Message             string `json:"message"`
	OrderLineExternalID string `json:"order_line_external_id"`
}

// swagger:model GeneralError
type GeneralError struct {
	// Error description
	Message string `json:"message"`
	// HTTP Status of the error
	Status int `json:"status"`
}

func generateLabelsForShipmentsHandler(w http.ResponseWriter) {
	// swagger:route POST /shipments ShipmentRequestBody
	// Request shipment labels
	//
	// security:
	//	api_key:
	//
	// responses:
	//
	//	401: GeneralError
	//	500: GeneralError
	//	200: ShipmentsResponse
	//
	// Parameters:
	//   + name: ShipmentsResponse
	//     in: body
	//     required: true
	//     type: ShipmentsResponse

}
