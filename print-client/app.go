package main

import (
	"context"
	"delivrio.io/go/endpoints/delivrioroutes"
	"delivrio.io/print-client/ent/printjob"
	"delivrio.io/shared-utils/pulid"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"delivrio.io/print-client/ent"
	"delivrio.io/print-client/ent/localdevice"
	"delivrio.io/shared-utils/models/printer"
	"delivrio.io/shared-utils/printers"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const pingRequestTimeout = 5 * time.Second

// App struct
type App struct {
	ctx           context.Context
	appDone       context.CancelFunc
	scannerCancel context.CancelFunc
	isRegistered  bool
	pingMux       sync.Mutex

	pingOnce      sync.Once
	printJobsOnce sync.Once
	scannerOnce   sync.Once
	cleanupOnce   sync.Once
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		pingMux: sync.Mutex{},
	}
}

func NewComputerID(ctx context.Context) {
	cli := ent.FromContext(ctx)
	IDs, err := cli.UniqueComputer.Query().All(ctx)
	if err != nil {
		log.Printf("querying unique computer ID failed: %v", err)
		return
	}

	if len(IDs) == 1 {
		log.Printf("found computer ID: %v", IDs[0].ID)
		return
	}

	unique, err := cli.UniqueComputer.Create().
		Save(ctx)
	if err != nil {
		log.Printf("error creating new computer ID: %v", err)
		return
	}

	log.Printf("created new computer ID: %v", unique.ID)

}

func OpenDSN(dsn string) (*ent.Client, *sql.Driver, error) {
	drv, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(time.Minute * 10)

	return ent.NewClient(ent.Driver(drv)), drv, nil
}

func handlePanic(funcName string) {
	if r := recover(); r != nil {
		log.Printf("Panic recovered in %s: %v", funcName, r)
	}
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {

	ctxCancel, cancel := context.WithCancel(a.ctx)
	a.appDone = cancel
	a.ctx = ctxCancel

	a.pingOnce.Do(func() {
		go func(ctx context.Context, app *App) {
			defer handlePanic("LoopBackgroundPing")
			// Give the other pings time to finish
			time.Sleep(time.Second * 3)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					app.LoopBackgroundPing()
					time.Sleep(15 * time.Second)
				}
			}
		}(a.ctx, a)
	})

	a.printJobsOnce.Do(func() {
		go func(ctx context.Context) {
			defer handlePanic("processPrintJobs")
			for {
				select {
				case <-ctx.Done():
					return
				default:
					processPrintJobs(ctx, a)
					time.After(1 * time.Second)
				}
			}
		}(ctxCancel)
	})

	ctxCancelScanner, cancel := context.WithCancel(a.ctx)
	a.scannerCancel = cancel

	a.scannerOnce.Do(func() {
		go func(ctx context.Context, cancel context.CancelFunc) {
			defer handlePanic("ReadScannerData")
			ReadScannerData(ctx, cancel)
		}(ctxCancelScanner, cancel)
	})

	a.scannerOnce.Do(func() {
		go func(ctx context.Context) {
			defer handlePanic("cleanup")
			for {
				select {
				case <-ctx.Done():
					return
				default:
					cleanup(ctx)
					time.After(15 * time.Second)
				}
			}
		}(ctxCancelScanner)
	})

}

const cleanupHorizon = -time.Hour * 24 * 30

func cleanup(ctx context.Context) {
	cli := ent.FromContext(ctx)
	_, err := cli.PrintJob.Delete().
		Where(printjob.CreatedAtLTE(time.Now().Add(cleanupHorizon))).
		Exec(ctx)
	if err != nil {
		log.Printf("cleanup job failed: %v", err)
	}
}

func (a *App) LoopBackgroundPing() {

	cli := ent.FromContext(a.ctx)

	availablePrinters, err := printersToSync(a.ctx)
	if err != nil {
		log.Printf("could not query available printers: %v", err)
		return
	}

	jobsToCancel, err := cli.PrintJob.Query().
		Where(printjob.StatusEQ(printjob.StatusPendingCancel)).
		All(a.ctx)
	if err != nil {
		log.Printf("could not query jobs to cancel: %v", err)
		return
	}
	jobsToSuccess, err := cli.PrintJob.Query().
		Where(printjob.StatusEQ(printjob.StatusPendingSuccess)).
		All(a.ctx)
	if err != nil {
		log.Printf("could not query jobs to success: %v", err)
		return
	}

	jobsToSuccessOutput := make([]printer.StatusChangeRequest, 0)
	allJobSuccessIDs := make([]pulid.ID, 0)
	for _, job := range jobsToSuccess {
		jobsToSuccessOutput = append(jobsToSuccessOutput, printer.StatusChangeRequest{
			ID:       job.ID,
			Messages: job.Messages,
		})
		allJobSuccessIDs = append(allJobSuccessIDs, job.ID)
	}
	jobsToCancelOutput := make([]printer.StatusChangeRequest, 0)
	allJobCancelIDs := make([]pulid.ID, 0)
	for _, job := range jobsToCancel {
		jobsToCancelOutput = append(jobsToCancelOutput, printer.StatusChangeRequest{
			ID:       job.ID,
			Messages: job.Messages,
		})
		allJobCancelIDs = append(allJobCancelIDs, job.ID)
	}

	ping := printer.PrintClientPing{
		Printers:         availablePrinters,
		CancelPrintJobs:  jobsToCancelOutput,
		SuccessPrintJobs: jobsToSuccessOutput,
	}

	u, err := RequestURL(a.ctx, delivrioroutes.PrintClientPing)
	if err != nil {
		log.Printf("error constructing remote URL: %v", err)
		return
	}

	success, workstationName := a.FirePrintPing(u.String(), ping)
	if success {
		a.isRegistered = true
		err = cli.RemoteConnection.Update().
			SetWorkstationName(workstationName).
			SetLastPing(time.Now()).
			Exec(a.ctx)
		if err != nil {
			log.Printf("error updating workstation name: %v", err)
		}

		err = cli.PrintJob.Update().
			Where(printjob.IDIn(allJobCancelIDs...)).
			SetStatus(printjob.StatusCanceled).
			Exec(a.ctx)
		if err != nil {
			log.Printf("error updating from pending_cancel -> cancel: %v", err)
		}
		err = cli.PrintJob.Update().
			Where(printjob.IDIn(allJobSuccessIDs...)).
			SetStatus(printjob.StatusSuccess).
			Exec(a.ctx)
		if err != nil {
			log.Printf("error updating from pending_success -> success: %v", err)
		}
	}

}

func RequestURL(ctx context.Context, uri string) (*url.URL, error) {
	cli := ent.FromContext(ctx)
	remote, err := cli.RemoteConnection.Query().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query remote connections: %v", err)
	}

	computerID, err := cli.UniqueComputer.Query().
		OnlyID(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query computer ID: %v", err)
	}

	u, err := url.Parse(remote.RemoteURL)
	if err != nil {
		return nil, fmt.Errorf("invalid remote URL: %v", err)
	}

	// Remote URL is stored with path, so we remove it to apply other URI's
	u.Path = ""
	u = u.JoinPath(delivrioroutes.API, uri)

	queryParams := url.Values{}
	queryParams.Add("id", remote.RegistrationToken)
	queryParams.Add("computer-id", computerID.String())
	queryParams.Add("device-type", "label_station")
	u.RawQuery = queryParams.Encode()

	return u, nil
}

func (a *App) FirePrintPing(url string, ping printer.PrintClientPing) (bool, string) {
	haveLock := a.pingMux.TryLock()
	if !haveLock {
		return false, fmt.Sprintf("ping is already running")
	}
	defer a.pingMux.Unlock()

	cli := ent.FromContext(a.ctx)

	body, err := json.Marshal(ping)
	if err != nil {
		log.Printf("could not marshall ping data: %v", err)
		return false, ""
	}

	httpClient := http.Client{
		Timeout: pingRequestTimeout,
	}

	log.Printf("pinging URL: %v -> %v", url, string(body))
	resp, err := httpClient.Post(
		url,
		"application/json",
		strings.NewReader(string(body)),
	)
	if err != nil {
		log.Printf("ping request failed: %v", err)
		return false, ""
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("could not read body: %v", err)
	}

	if resp.StatusCode != 200 {
		runtime.EventsEmit(a.ctx, "token-not-recognized")
		log.Printf("ping remote failed, expected status 200, got: %v: %v", resp.StatusCode, string(respBody))
		return false, ""
	}

	var bodyData printer.PrintClientPingResponse
	err = json.Unmarshal(respBody, &bodyData)
	if err != nil {
		log.Printf("could not unmarshall body: %v", err)
		return false, ""
	}

	// Assume we only have 1
	err = cli.RemoteConnection.Update().
		SetLastPing(time.Now()).
		SetWorkstationName(bodyData.WorkstationName).
		Exec(a.ctx)
	if err != nil {
		log.Printf("could not update remote connection: %v", err)
	}

	err = savePrintJobs(a.ctx, bodyData.PrintJobs)
	if err != nil {
		log.Printf("could not save print jobs: %v", err)
		return false, ""
	}

	return true, bodyData.WorkstationName
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	a.appDone()
	a.scannerCancel()
}

func (a *App) Register(tryURL string, token string) string {
	cli := ent.FromContext(a.ctx)

	availablePrinters, err := printersToSync(a.ctx)
	if err != nil {
		return err.Error()
	}

	computerID, err := cli.UniqueComputer.Query().OnlyID(a.ctx)
	if err != nil {
		return fmt.Errorf("error querying computer ID: %w", err).Error()
	}

	ping := printer.PrintClientPing{
		Printers: availablePrinters,
	}

	u, err := url.Parse(tryURL)
	if err != nil {
		return fmt.Sprintf("invalid remote ping URL")
	}

	queryParams := url.Values{}
	queryParams.Add("id", token)
	queryParams.Add("computer-id", computerID.String())
	queryParams.Add("device-type", "label_station")
	u.RawQuery = queryParams.Encode()

	isRegistered, workstationName := a.FirePrintPing(u.String(), ping)
	if !isRegistered {
		return fmt.Sprintf("not registered")
	}

	_, err = cli.RemoteConnection.Delete().Exec(a.ctx)
	if err != nil {
		return err.Error()
	}

	err = cli.RemoteConnection.Create().
		SetRegistrationToken(token).
		SetRemoteURL(tryURL).
		SetWorkstationName(workstationName).
		OnConflict().
		UpdateNewValues().
		Exec(a.ctx)
	if err != nil {
		return err.Error()
	}

	a.isRegistered = true
	runtime.EventsEmit(a.ctx, "registration-changed")

	return ""
}

func (a *App) IsRegistered() string {
	tx := ent.FromContext(a.ctx)
	remote, err := tx.RemoteConnection.Query().Only(a.ctx)
	if err != nil {
		return err.Error()
	}

	computerID, err := client.UniqueComputer.Query().OnlyID(a.ctx)
	if err != nil {
		return fmt.Errorf("error querying computer ID: %w", err).Error()
	}

	availablePrinters, err := printersToSync(a.ctx)
	if err != nil {
		return err.Error()
	}

	ping := printer.PrintClientPing{
		Printers: availablePrinters,
	}
	u, err := url.Parse(remote.RemoteURL)
	if err != nil {
		return fmt.Sprintf("invalid remote ping URL")
	}

	queryParams := url.Values{}
	queryParams.Add("id", remote.RegistrationToken)
	queryParams.Add("computer-id", computerID.String())
	queryParams.Add("device-type", "label_station")
	u.RawQuery = queryParams.Encode()

	isRegistered, _ := a.FirePrintPing(u.String(), ping)
	if !isRegistered {
		return fmt.Sprintf("not registered")
	}

	return ""
}

func (a *App) RegistrationData() (*RemoteConnectionData, error) {
	cli := ent.FromContext(a.ctx)
	rc, err := cli.RemoteConnection.Query().
		First(a.ctx)
	if err != nil {
		return nil, err
	}

	return &RemoteConnectionData{
		WorkstationName: rc.WorkstationName,
		URL:             rc.RemoteURL,
		LastPing:        rc.LastPing.Format(time.ANSIC),
	}, nil
}

func (a *App) RemoveRegistrationData() error {
	cli := ent.FromContext(a.ctx)
	_, err := cli.RemoteConnection.Delete().
		Exec(a.ctx)
	if err != nil {
		return err
	}

	_, err = cli.PrintJob.Delete().
		Exec(a.ctx)
	if err != nil {
		return err
	}

	_, err = cli.LocalDevice.Delete().
		Exec(a.ctx)
	if err != nil {
		return err
	}

	a.isRegistered = false

	return nil
}

func printersToSync(ctx context.Context) ([]printers.Printer, error) {
	availablePrinters, err := client.LocalDevice.Query().
		Where(
			localdevice.And(
				localdevice.Active(true),
				localdevice.Archived(false),
				localdevice.CategoryEQ(localdevice.CategoryPrinter),
			),
		).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying printers")
	}

	output := make([]printers.Printer, 0)
	for _, d := range availablePrinters {
		output = append(output, printers.Printer{
			ID:   d.ID,
			Name: d.Name,
		})
	}

	return output, nil
}

func (a *App) WorkstationName() string {
	cli := ent.FromContext(a.ctx)
	ws, err := cli.RemoteConnection.Query().
		First(a.ctx)
	if err != nil && ent.IsNotFound(err) {
		return ""
	}

	return ws.WorkstationName
}

func (a *App) ShowDialog(title string, message string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   title,
		Message: message,
	})

	if err != nil {
		panic(err)
	}
}
