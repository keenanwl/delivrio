package main

import (
	"context"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/print-client/ent"
	"delivrio.io/print-client/ent/localdevice"
	"delivrio.io/print-client/ent/recentscan"
	printerUtils "delivrio.io/shared-utils/models/printer"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/pteich/hid"
	"github.com/pteich/usbsymbolreader/code"
	"github.com/pteich/usbsymbolreader/scanner"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var scannerMux = sync.Mutex{}

const labelRequestTimeout = 10 * time.Second

func (a *App) USBDevices() []USBDevice {

	cli := ent.FromContext(a.ctx)
	devices := hid.Enumerate(0, 0)

	for _, deviceInfo := range devices {
		name := fmt.Sprintf("%s - %s", deviceInfo.Manufacturer, deviceInfo.Product)
		uniqueName := fmt.Sprintf("%v - %v", deviceInfo.VendorID, deviceInfo.ProductID)
		err := cli.LocalDevice.Create().
			SetActive(false).
			SetArchived(false).
			SetCategory(localdevice.CategoryScanner).
			SetName(name).
			SetSystemName(uniqueName).
			SetVendorID(int(deviceInfo.VendorID)).
			SetProductID(int(deviceInfo.ProductID)).
			OnConflict().
			Ignore().
			Exec(a.ctx)
		if err != nil {
			log.Printf("saving local scanner device: %v", err)
		}
	}

	allDevices, err := cli.LocalDevice.Query().
		Where(localdevice.CategoryEQ(localdevice.CategoryScanner)).
		All(a.ctx)
	if err != nil {
		log.Printf("query local scanner device: %v", err)
	}

	deviceOutput := make([]USBDevice, 0)
	for _, d := range allDevices {
		deviceOutput = append(deviceOutput, USBDevice{
			ID:        d.ID,
			VendorID:  fmt.Sprintf("%v", d.VendorID),
			ProductID: fmt.Sprintf("%v", d.ProductID),
			Name:      d.Name,
			Active:    d.Active,
		})
	}

	return deviceOutput
}

func (a *App) SaveSelectedDevice(id string) string {
	cli := ent.FromContext(a.ctx)

	tx, err := cli.Tx(a.ctx)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	err = tx.LocalDevice.Update().
		SetActive(false).
		Where(localdevice.CategoryEQ(localdevice.CategoryScanner)).
		Exec(a.ctx)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	err = tx.LocalDevice.Update().
		SetActive(true).
		Where(localdevice.ID(pulid.ID(id))).
		Exec(a.ctx)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	a.scannerCancel()

	ctxCancelScanner, cancel := context.WithCancel(a.ctx)
	a.scannerCancel = cancel

	go func(ctx context.Context, cancel context.CancelFunc) {
		ReadScannerData(ctx, cancel)
	}(ctxCancelScanner, cancel)

	return ""
}

func (a *App) RecentScans() ([]RecentScan, error) {
	cli := ent.FromContext(a.ctx)
	scans, err := cli.RecentScan.Query().
		Order(recentscan.ByID(sql.OrderDesc())).
		Limit(10).
		All(a.ctx)
	if err != nil {
		return nil, err
	}
	output := make([]RecentScan, 0)
	for _, s := range scans {
		output = append(output, RecentScan{
			Created: s.CreatedAt.Format(time.DateTime),
			Code:    s.ScanValue,
			Result:  s.Response,
		})
	}

	return output, nil
}

func ReadScannerData(ctx context.Context, cancel context.CancelFunc) {
	cli := ent.FromContext(ctx)
	scan, err := cli.LocalDevice.Query().
		Where(localdevice.And(
			localdevice.Active(true),
			localdevice.CategoryEQ(localdevice.CategoryScanner),
		)).
		Only(ctx)
	if err != nil {
		log.Printf("error fetching active scanner: %v %v", err, scan)
		time.Sleep(5 * time.Second)
		return
	}

	devices := hid.Enumerate(uint16(scan.VendorID), uint16(scan.ProductID))
	if len(devices) == 1 {
		symbolScanner, err := scanner.New(devices[0])
		if err != nil {
			log.Print(err)
			return
		}

		log.Printf("connected to %s (Serial No. %s)", devices[0].Product, devices[0].Serial)

		codes := ReadCodes(ctx, symbolScanner.Device(), cancel)

	readLoop:
		for {
			select {
			case outputCode := <-codes:
				// TODO: this is blocking until he request finishes
				log.Printf("Read code: %v", outputCode)
				err := FireLabelRequest(ctx, outputCode.String())
				if err != nil {
					log.Printf("Error requesting label: %v", err)
				}
				break

			case <-ctx.Done():
				log.Printf("stopping background scanner listener")
				break readLoop
			}
		}

	}
}

func FireLabelRequest(ctx context.Context, internalBarcode string) error {
	// Probably need some sort of lock request timeout instead of insta-fail
	lock := scannerMux.TryLock()
	defer scannerMux.Unlock()
	if !lock {
		return fmt.Errorf("could not aquire lock: other remote requests still pending")
	}

	scan, err := createRecentScan(ctx, internalBarcode)
	if err != nil {
		return err
	}

	u, err := RequestURL(ctx, delivrioroutes.PrintClientRequestLabel)
	if err != nil {
		return updateScan(ctx, scan, fmt.Errorf("creating remote auth URL: %w", err))
	}

	req := printerUtils.PrintClientLabelRequest{
		InternalBarcode: internalBarcode,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return updateScan(ctx, scan, fmt.Errorf("could not marshall ping data: %v", err))
	}

	httpClient := http.Client{
		Timeout: labelRequestTimeout,
	}

	resp, err := httpClient.Post(
		u.String(),
		"application/json",
		strings.NewReader(string(body)),
	)
	if err != nil {
		return updateScan(ctx, scan, fmt.Errorf("could fetch label data: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return updateScan(ctx, scan, nil)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return updateScan(ctx, scan, fmt.Errorf("could not read body: %v", err))
	}

	return updateScan(ctx, scan, fmt.Errorf("server error: code %v: %v", resp.StatusCode, string(respBody)))

}

func updateScan(ctx context.Context, scan *ent.RecentScan, inputErr error) error {

	msg := "Success"
	if inputErr != nil {
		msg = inputErr.Error()
	}

	err := scan.Update().
		SetResponse(msg).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", inputErr, err)
	}
	return inputErr
}

func createRecentScan(ctx context.Context, code string) (*ent.RecentScan, error) {
	cli := ent.FromContext(ctx)
	scan, err := cli.RecentScan.Create().
		SetScanType(recentscan.ScanTypeLabelRequest).
		SetResponse("Pending...").
		SetScanValue(code).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating recent scan: %w", err)
	}

	fiveDays := time.Hour * 24 * -5
	_, err = cli.RecentScan.Delete().
		Where(recentscan.CreatedAtLT(time.Now().Add(fiveDays))).
		Exec(ctx)
	if err != nil {
		log.Printf("error: cleanup recent scans failed %s", err)
	}

	return scan, nil
}

// Adapted from external lib to allow greater control over cancel fn
func ReadCodes(ctx context.Context, device *hid.Device, cancel context.CancelFunc) <-chan *code.Code {

	codeChan := make(chan *code.Code)
	go func() {
		defer close(codeChan)
		for {
			buf := make([]byte, 255)
			n, err := device.ReadTimeout(buf, 500)
			if err != nil {
				cancel()
				return
			}

			if n > 0 {
				scannedCode, err := code.New(buf)
				if err != nil {
					continue
				}

				codeChan <- scannedCode
			}

			select {
			case <-ctx.Done():
				device.Close()
				return
			default:
			}
		}
	}()

	return codeChan
}
