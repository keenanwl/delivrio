package main

import (
	"context"
	"delivrio.io/print-client/ent"
	"delivrio.io/print-client/ent/localdevice"
	"delivrio.io/print-client/ent/printjob"
	"delivrio.io/shared-utils/models/printer"
	"delivrio.io/shared-utils/printers"
	"delivrio.io/shared-utils/pulid"
	b64 "encoding/base64"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"net"
	"strings"
	"time"
)

func isValidIPPort(ipPort string) bool {
	_, err := net.ResolveTCPAddr("tcp", ipPort)
	return err == nil
}

type PrintJobDisplay struct {
	ID            pulid.ID        `json:"id"`
	FileExtension string          `json:"file_extension"`
	PrinterName   string          `json:"printer_name"`
	Created       time.Time       `json:"created"`
	Status        printjob.Status `json:"status"`
	Messages      []string        `json:"messages"`
}

type AllJobs struct {
	CurrentJobs []PrintJobDisplay `json:"current_jobs"`
	RecentJobs  []PrintJobDisplay `json:"recent_jobs"`
}

func (a *App) ActivePrintJobs() *AllJobs {
	cli := ent.FromContext(a.ctx)

	jobs, err := cli.PrintJob.Query().
		WithLocalDevice().
		Where(printjob.StatusEQ(printjob.StatusPending)).
		Limit(5).
		Order(printjob.ByID()).
		All(a.ctx)
	if err != nil {
		log.Println("could not query print jobs")
		return nil
	}

	output := make([]PrintJobDisplay, 0)
	for _, j := range jobs {
		output = append(output, PrintJobDisplay{
			ID:            j.ID,
			FileExtension: j.FileExtension.String(),
			PrinterName:   j.Edges.LocalDevice.Name,
			Created:       j.CreatedAt,
			Status:        j.Status,
			Messages:      j.Messages,
		})
	}

	recentJobs, err := cli.PrintJob.Query().
		WithLocalDevice().
		Where(printjob.StatusNEQ(printjob.StatusPending)).
		Limit(5).
		Order(printjob.ByID(sql.OrderDesc())).
		All(a.ctx)
	if err != nil {
		log.Println("could not query print jobs")
		return nil
	}

	outputRecent := make([]PrintJobDisplay, 0)
	for _, j := range recentJobs {
		outputRecent = append(outputRecent, PrintJobDisplay{
			ID:            j.ID,
			FileExtension: j.FileExtension.String(),
			PrinterName:   j.Edges.LocalDevice.Name,
			Created:       j.CreatedAt,
			Status:        j.Status,
			Messages:      j.Messages,
		})
	}

	return &AllJobs{
		CurrentJobs: output,
		RecentJobs:  outputRecent,
	}
}

func (a *App) JobPendingCancel(id string, msg []string) string {
	cli := ent.FromContext(a.ctx)

	err := cli.PrintJob.Update().
		Where(printjob.ID(pulid.ID(id))).
		SetStatus(printjob.StatusPendingCancel).
		SetMessages(msg).
		Exec(a.ctx)
	if err != nil {
		log.Println("could not update print job")
		return err.Error()
	}

	return ""
}

func (a *App) SaveNetworkPrinter(addr string) string {
	cli := ent.FromContext(a.ctx)

	if isValidIPPort(addr) {
		noWhitespace := strings.ReplaceAll(addr, " ", "")
		err := cli.LocalDevice.Create().
			SetName(noWhitespace).
			SetSystemName(noWhitespace).
			SetAddress(noWhitespace).
			SetActive(true).
			SetArchived(false).
			SetCategory(localdevice.CategoryPrinter).
			OnConflict().
			UpdateActive().
			Exec(a.ctx)
		if err != nil {
			return err.Error()
		}
		return ""
	}
	return "invalid address format"
}

func (a *App) Printers() []printers.Printer {
	cli := ent.FromContext(a.ctx)
	cliLocal := printers.NewClient(a.ctx)
	listPrinters, err := cliLocal.ListPrinters()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, p := range listPrinters {
		err = cli.LocalDevice.Create().
			SetName(p.Name).
			SetSystemName(p.Name).
			SetCategory(localdevice.CategoryPrinter).
			SetActive(false).
			SetArchived(false).
			OnConflict().
			UpdateArchived().
			Exec(a.ctx)
		if err != nil {
			log.Printf("creating local device printers: %v", err)
		}
	}

	allPrinters, err := cli.LocalDevice.Query().
		Where(localdevice.And(
			localdevice.CategoryEQ(localdevice.CategoryPrinter),
			localdevice.Archived(false),
		)).
		All(a.ctx)
	if err != nil {
		panic(err)
	}

	output := make([]printers.Printer, 0)
	for _, p := range allPrinters {
		output = append(output, printers.Printer{
			ID:      p.ID,
			Name:    p.Name,
			Active:  p.Active,
			Network: len(p.Address) > 0,
		})
	}

	return output
}

func (a *App) ArchivePrinter(id string) string {
	cli := ent.FromContext(a.ctx)

	err := cli.LocalDevice.Update().
		SetArchived(true).
		Where(localdevice.ID(pulid.ID(id))).
		Exec(a.ctx)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) TogglePrinter(name string, active bool) string {
	cli := ent.FromContext(a.ctx)
	if cli == nil {
		return fmt.Errorf("could not extract cli from ctx").Error()
	}

	nonPrinterNameCount, err := cli.LocalDevice.Query().
		Where(localdevice.And(
			localdevice.SystemNameEQ(name),
			localdevice.CategoryNEQ(localdevice.CategoryPrinter),
		)).
		Count(a.ctx)
	if err != nil {
		return err.Error()
	}

	if nonPrinterNameCount > 0 {
		return fmt.Sprintf("Scanner and Printer names may not overlap")
	}

	err = cli.LocalDevice.Create().
		SetName(name).
		SetSystemName(name).
		SetCategory(localdevice.CategoryPrinter).
		SetActive(active).
		OnConflict().
		UpdateActive().
		Exec(a.ctx)
	if err != nil {
		return err.Error()
	}

	// Side effects
	a.IsRegistered()

	return ""
}

type DYMOPngInput struct {
	Base64Data string `json:"base64Data"`
	ID         string `json:"id"`
}
type DYMOPngOutput struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	ID      string `json:"id"`
}

func (a *App) PrintDYMOPng(id string, base64String string) {
	runtime.EventsEmit(a.ctx, "dymo-png", DYMOPngInput{
		Base64Data: base64String,
		ID:         id,
	})
	// Prevent bad things from happening in case the JS message
	// doesn't come back
	cli := ent.FromContext(a.ctx)
	cli.PrintJob.Update().
		Where(printjob.ID(pulid.ID(id))).
		SetStatus(printjob.StatusPendingSuccess).
		Exec(a.ctx)
}

func (a *App) DYMOPngRegisterStatus(outputString string) {
	cli := ent.FromContext(a.ctx)
	var output DYMOPngOutput
	err := json.Unmarshal([]byte(outputString), &output)
	if err != nil {
		log.Println("error print png: " + err.Error())
		return
	}

	if output.Success {
		cli.PrintJob.Update().
			Where(printjob.ID(pulid.ID(output.ID))).
			SetStatus(printjob.StatusPendingSuccess).
			Exec(a.ctx)
	} else {
		cli.PrintJob.Update().
			Where(printjob.ID(pulid.ID(output.ID))).
			SetStatus(printjob.StatusPendingCancel).
			SetMessages([]string{output.Msg}).
			Exec(a.ctx)
	}

}

func savePrintJobs(ctx context.Context, printJobs []printer.PrintJob) error {
	cli := ent.FromContext(ctx)

	for _, pj := range printJobs {
		err := cli.PrintJob.Create().
			SetID(pj.ID).
			SetBase64PrintData(pj.Base64Data).
			SetFileExtension(printjob.FileExtension(pj.FileExtension)).
			SetStatus(printjob.StatusPending).
			SetLocalDeviceID(pj.PrinterID).
			SetUseShell(pj.UseShell).
			OnConflict().
			Ignore().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil

}

func processPrintJobs(ctx context.Context, a *App) {

	cli := ent.FromContext(ctx)
	pendingJobs, err := cli.PrintJob.Query().
		WithLocalDevice().
		Where(printjob.StatusEQ(printjob.StatusPending)).
		All(ctx)
	if err != nil {
		log.Printf("error fetching pending print jobs: %v", err)
	}

	for _, pj := range pendingJobs {

		active := pj.Edges.LocalDevice.Active
		adr := pj.Edges.LocalDevice.Address
		if active {

			if pj.FileExtension == printjob.FileExtensionPng {
				// Not the cleanest implementation.
				// Move if this works.
				a.PrintDYMOPng(pj.ID.String(), pj.Base64PrintData)
				continue
			}

			printData, err := b64.StdEncoding.DecodeString(pj.Base64PrintData)
			if err != nil {
				log.Println("error: decoding print data: ", err)
				continue
			}

			if adr != "" {
				err = printers.FireNetworkPrintJob(adr, printData)
				if err != nil {
					pj.Update().
						SetStatus(printjob.StatusPendingCancel).
						SetMessages([]string{err.Error()}).
						Exec(ctx)
					log.Println("Error:", err)
					continue
				}
			} else {
				err = printers.FireUSBPrintJob(pj.Edges.LocalDevice.SystemName, printData, pj.FileExtension, pj.UseShell)
				if err != nil {
					pj.Update().
						SetStatus(printjob.StatusPendingCancel).
						SetMessages([]string{err.Error()}).
						Exec(ctx)
					log.Println("Error:", err)
					continue
				}
			}

			err = pj.Update().
				SetStatus(printjob.StatusPendingSuccess).
				Exec(ctx)
			if err != nil {
				pj.Update().
					SetStatus(printjob.StatusPendingCancel).
					Exec(ctx)
				log.Println("Error:", err)
				return
			}
		}

	}

}
