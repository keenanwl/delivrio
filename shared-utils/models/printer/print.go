package printer

import (
	"delivrio.io/shared-utils/printers"
	"delivrio.io/shared-utils/pulid"
)

type PrintClientLabelRequest struct {
	Token           string `json:"token"`
	ComputerID      string `json:"computer_id"`
	InternalBarcode string `json:"internal_barcode"` // String for more control over the random barcodes that may get scanned
}

type StatusChangeRequest struct {
	ID       pulid.ID `json:"id"`
	Messages []string `json:"messages"`
}

type PrintClientPing struct {
	Printers         []printers.Printer    `json:"printers"`
	LabelID          string                `json:"label_id"`
	CancelPrintJobs  []StatusChangeRequest `json:"cancel_print_jobs"`
	SuccessPrintJobs []StatusChangeRequest `json:"success_print_jobs"`
}

type PrintJob struct {
	ID            pulid.ID `json:"id"`
	PrinterID     pulid.ID `json:"printer_id"`
	Base64Data    string   `json:"base64_data"`
	FileExtension string   `json:"file_extension"`
	UseShell      bool     `json:"use_shell"`
}

type PrintClientLabelResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type PrintClientPingResponse struct {
	Success         bool       `json:"success"`
	Message         string     `json:"message"`
	WorkstationName string     `json:"workstation_name"`
	PrintJobs       []PrintJob `json:"print_jobs"`
}
