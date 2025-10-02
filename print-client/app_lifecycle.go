package main

import (
	"context"
	"delivrio.io/print-client/ent"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var wailsContext *context.Context

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	wailsContext = &ctx

	log.Printf("initializing database")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		// TODO display something to the user
		log.Fatal(err)
	}

	dbPath := filepath.Join(homeDir, "delivrio")
	err = os.MkdirAll(dbPath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	dbFullPath := "file:" + filepath.Join(dbPath, "delivrio_desktop.db?cache=shared&_fk=1")

	log.Printf("opening Ent at path: %v", dbFullPath)

	db, drv, err := OpenDSN(dbFullPath)
	if err != nil {
		// TODO display something to the user
		log.Fatal(err)
	}

	client = db

	migrateDB(ctx, drv)

	err = startOnLogin()
	if err != nil {
		log.Printf("add start on login: %v", err)
	}

	_, err = client.RemoteConnection.Query().
		Only(ctx)
	if err != nil {
		log.Printf("client not registered: %v", err)
	} else {
		a.isRegistered = true
		runtime.EventsEmit(ctx, "registration-changed")
	}

	ctx = ent.NewContext(ctx, client)
	a.ctx = ctx

	NewComputerID(ctx)

}

func (a *App) onSecondInstanceLaunch(secondInstanceData options.SecondInstanceData) {
	log.Println("user opened second instance", strings.Join(secondInstanceData.Args, ","))
	log.Println("user opened second from", secondInstanceData.WorkingDirectory)
	runtime.WindowUnminimise(*wailsContext)
	runtime.Show(*wailsContext)
}
