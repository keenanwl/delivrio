package printerutils

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/printjob"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/viewer"
	b64 "encoding/base64"
)

// Selects the first document printer, otherwise selects an arbitrary PDF printer
// finally, selects the first printer in the list
// Could still implement page size matching between document and printer
func defaultPackingSlipPrinter(ctx context.Context, workstation *ent.Workstation) (*ent.Printer, error) {
	allPrinters, err := workstation.QueryPrinter().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var firstPrinter, pdfPrinter *ent.Printer

	for i, printer := range allPrinters {
		if i == 0 {
			firstPrinter = printer
		}

		if printer.Document {
			return printer, nil // Return immediately if a document printer is found
		} else if printer.LabelPdf && pdfPrinter == nil {
			pdfPrinter = printer // Store the first PDF printer found
		}
	}

	// Return PDF printer if found, otherwise return the first printer
	if pdfPrinter != nil {
		return pdfPrinter, nil
	}
	return firstPrinter, nil
}

func CreatePackingSlipPrintJobs(ctx context.Context, collis []*ent.Colli, workstation *ent.Workstation) error {

	cli := ent.FromContext(ctx)
	view := viewer.FromContext(ctx)

	selectedPrinter, err := defaultPackingSlipPrinter(ctx, workstation)
	if err != nil {
		return err
	}

	for _, c := range collis {

		if selectedPrinter.LabelZpl {

			packingSlipZPLLabels, err := mergeutils.QueryOrFetchPackingSlipZPL(ctx, c)
			if err != nil {
				return err
			}

			for _, l := range packingSlipZPLLabels {
				encodedPrintDataZPL := b64.StdEncoding.EncodeToString([]byte(l))
				err := cli.PrintJob.Create().
					SetTenantID(view.TenantID()).
					SetDocumentType(printjob.DocumentTypePackingList).
					SetColli(c).
					SetStatus(printjob.StatusPending).
					SetPrinter(selectedPrinter).
					SetFileExtension(printjob.FileExtensionZpl).
					SetBase64PrintData(encodedPrintDataZPL).
					Exec(ctx)
				if err != nil {
					return err
				}
			}
		} else if selectedPrinter.LabelPng {
			packingSlipPng, err := mergeutils.QueryOrFetchPackingSlipPng(ctx, c)
			if err != nil {
				return err
			}

			for _, l := range packingSlipPng {
				err := cli.PrintJob.Create().
					SetTenantID(view.TenantID()).
					SetDocumentType(printjob.DocumentTypePackingList).
					SetColli(c).
					SetStatus(printjob.StatusPending).
					SetPrinter(selectedPrinter).
					SetFileExtension(printjob.FileExtensionPng).
					SetBase64PrintData(l).
					Exec(ctx)
				if err != nil {
					return err
				}
			}
		} else if selectedPrinter.LabelPdf {
			packingSlipPDF, err := mergeutils.QueryOrFetchPackingSlip(ctx, c)
			if err != nil {
				return err
			}

			for _, l := range packingSlipPDF {
				err := cli.PrintJob.Create().
					SetTenantID(view.TenantID()).
					SetDocumentType(printjob.DocumentTypePackingList).
					SetColli(c).
					SetStatus(printjob.StatusPending).
					SetPrinter(selectedPrinter).
					SetFileExtension(printjob.FileExtensionPdf).
					SetBase64PrintData(l).
					Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
