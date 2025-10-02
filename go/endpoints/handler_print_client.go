package endpoints

import (
	"delivrio.io/go/carrierapis/labels"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/printer"
	"delivrio.io/go/ent/printjob"
	"delivrio.io/go/ent/shipmentparcel"
	"delivrio.io/go/ent/workstation"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/utils/httputils"
	"delivrio.io/go/viewer"
	"delivrio.io/shared-utils/pulid"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"delivrio.io/go/ent"
	printerUtils "delivrio.io/shared-utils/models/printer"
)

func PrintClientPing(w http.ResponseWriter, r *http.Request) {

	ws, ctx, err := validWorkstationRequest(w, r)
	if err != nil {
		return
	}

	var input printerUtils.PrintClientPing
	err = httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	cli := ent.FromContext(r.Context())

	for _, p := range input.Printers {
		err := cli.Printer.Create().
			SetName(p.Name).
			SetDeviceID(p.ID). // Unique ID
			SetWorkstation(ws).
			SetTenantID(ws.TenantID).
			SetLastPing(time.Now()).
			OnConflict().
			UpdateName(). // Should be the only editable field originating from the client
			Exec(ctx)
		if err != nil {
			httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	if input.CancelPrintJobs != nil {
		tx, err := cli.Tx(ctx)
		if err != nil {
			httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		defer tx.Rollback()

		txCtx := ent.NewTxContext(ctx, tx)

		for _, cpj := range input.CancelPrintJobs {
			err = tx.PrintJob.Update().
				Where(printjob.ID(cpj.ID)).
				SetStatus(printjob.StatusCanceled).
				Exec(txCtx)
			if err != nil {
				tx.Rollback()
				httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
					Success: false,
					Message: err.Error(),
				})
				return
			}
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

	}

	// Orders are queried to create history record for the print job
	accessRightOverrideCtx := viewer.EnableAnonymousOverride(ctx)

	if input.SuccessPrintJobs != nil {

		allErrors := make([]string, 0)

		for _, j := range input.SuccessPrintJobs {

			pj, err := cli.PrintJob.Query().
				WithColli().
				WithShipmentParcel().
				Where(printjob.ID(j.ID)).
				Only(ctx)
			if err != nil {
				allErrors = append(allErrors, err.Error())
				continue
			}

			err = pj.Update().
				SetStatus(printjob.StatusSuccess).
				Exec(accessRightOverrideCtx)
			if err != nil {
				allErrors = append(allErrors, err.Error())
				continue
			}

			switch pj.DocumentType {
			case printjob.DocumentTypePackingList:
				if c, err := pj.Edges.ColliOrErr(); err == nil {
					err = c.Update().
						SetSlipPrintStatus(colli.SlipPrintStatusPrinted).
						Exec(history.NewConfig(accessRightOverrideCtx).
							SetDescription("Packing list printed").
							SetOrigin(changehistory.OriginPrintClient).
							Ctx())
					if err != nil {
						allErrors = append(allErrors, err.Error())
						continue
					}
				}
				break
			case printjob.DocumentTypeParcelLabel:
				if s, err := pj.Edges.ShipmentParcelOrErr(); err == nil {
					err = s.Update().
						SetStatus(shipmentparcel.StatusPrinted).
						Exec(history.NewConfig(accessRightOverrideCtx).
							SetDescription("Shipment label printed").
							SetOrigin(changehistory.OriginPrintClient).
							Ctx())
					if err != nil {
						allErrors = append(allErrors, err.Error())
						continue
					}
					c, err := s.Colli(accessRightOverrideCtx)
					if err != nil {
						allErrors = append(allErrors, err.Error())
						continue
					}
					err = c.Update().
						SetStatus(colli.StatusDispatched).
						Exec(history.NewConfig(accessRightOverrideCtx).
							SetDescription("Shipment label printed").
							SetOrigin(changehistory.OriginPrintClient).
							Ctx())
					if err != nil {
						allErrors = append(allErrors, err.Error())
						continue
					}

				}
				break
			}

		}

		if len(allErrors) > 0 {
			httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
				Success: false,
				Message: strings.Join(allErrors, "; "),
			})
			return
		}

		ids := make([]pulid.ID, 0)
		for _, pj := range input.SuccessPrintJobs {
			ids = append(ids, pj.ID)
		}

		err = cli.PrintJob.Update().
			Where(printjob.IDIn(ids...)).
			SetStatus(printjob.StatusSuccess).
			Exec(accessRightOverrideCtx)
		if err != nil {
			httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	// Suffers from there being multiple printers where first is
	// backed up and the others could be waiting for jobs
	jobs, err := cli.PrintJob.Query().
		WithPrinter().
		Where(printjob.And(
			printjob.HasPrinterWith(printer.HasWorkstationWith(workstation.ID(ws.ID))),
			// Send the job until it has the status changed
			printjob.StatusIn(printjob.StatusPending, printjob.StatusAtPrinter),
		)).
		Order(printjob.ByID()).
		Limit(25).
		All(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	printJobs := make([]printerUtils.PrintJob, 0)
	jobIDs := make([]pulid.ID, 0)
	for _, pj := range jobs {
		jobIDs = append(jobIDs, pj.ID)
		printJobs = append(printJobs, printerUtils.PrintJob{
			ID:            pj.ID,
			PrinterID:     pj.Edges.Printer.DeviceID,
			Base64Data:    pj.Base64PrintData,
			FileExtension: pj.FileExtension.String(),
			UseShell:      pj.Edges.Printer.UseShell,
		})
	}

	err = cli.PrintJob.Update().
		SetStatus(printjob.StatusAtPrinter).
		Where(printjob.IDIn(jobIDs...)).
		Exec(accessRightOverrideCtx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientPingResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, printerUtils.PrintClientPingResponse{
		Success:         true,
		Message:         fmt.Sprintf("Saved %v printers", len(input.Printers)),
		WorkstationName: ws.Name,
		PrintJobs:       printJobs,
	})
	return

}

func PrintClientRequestLabel(w http.ResponseWriter, r *http.Request) {
	ws, ctx, err := validWorkstationRequest(w, r)
	if err != nil {
		return
	}
	cli := ent.FromContext(ctx)

	var input printerUtils.PrintClientLabelRequest
	err = httputils.UnmarshalRequestBody(r, &input)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientLabelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	printers, err := ws.QueryPrinter().
		All(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, printerUtils.PrintClientLabelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	barcode, err := strconv.ParseInt(input.InternalBarcode, 10, 64)
	if err != nil {
		httputils.JSONResponse(w, http.StatusBadRequest, printerUtils.PrintClientLabelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	col, err := cli.Colli.Query().
		Where(colli.InternalBarcode(barcode)).
		Only(ctx)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, printerUtils.PrintClientLabelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	shipments, err := labels.MustShipmentFromColliID(ctx, col.ID)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, printerUtils.PrintClientLabelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = labels.CreatePrintJobs(ctx, printers, shipments)
	if err != nil {
		httputils.JSONResponse(w, http.StatusInternalServerError, printerUtils.PrintClientLabelResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	httputils.JSONResponse(w, http.StatusOK, printerUtils.PrintClientLabelResponse{
		Success: true,
		Message: fmt.Sprintf("created print job from shipment"),
	})
	return

}
