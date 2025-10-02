package main

import (
	"context"
	"embed"
	"entgo.io/ent/dialect/sql"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"io"
	"log"
	"os"
	"path/filepath"

	"delivrio.io/print-client/ent"
	"delivrio.io/shared-utils/models/printer"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"gopkg.in/natefinch/lumberjack.v2"

	// Schema migrations
	migrate2 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed frontend/dist/browser
var assets embed.FS
var client *ent.Client

//go:embed build/icons/hicolor/512x512/delivrio.png
var icon []byte

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		// TODO display something to the user
		log.Fatal(err)
	}

	logDirectory := filepath.Join(homeDir, "delivrio")
	err = os.MkdirAll(logDirectory, 0777)
	if err != nil {
		log.Fatal(err)
	}

	lumberjackLogger := &lumberjack.Logger{
		// Log file absolute path, os agnostic
		Filename:   filepath.ToSlash(filepath.Join(logDirectory, "errors.log")),
		MaxSize:    5, // MB
		MaxBackups: 10,
		MaxAge:     30,    // days
		Compress:   false, // disabled by default
	}

	mw := io.MultiWriter(os.Stdout, lumberjackLogger)
	log.SetOutput(mw)

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "DELIVRIO",
		Width:  724,
		Height: 724,
		// MinWidth:          720,
		// MinHeight:         570,
		// MaxWidth:          1280,
		// MaxHeight:         740,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		LogLevel:          logger.DEBUG,
		Logger:            nil,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
			&USBDevice{},
			&NetworkDevice{},
			&SubnetSearch{},
			&RemoteConnectionData{},
			&RecentScan{},
			&printer.PrintClientPing{},
			&printer.PrintClientPingResponse{},
		},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId:               "c9c8fd93-6758-4144-87d1-34bdb0a8bd60",
			OnSecondInstanceLaunch: app.onSecondInstanceLaunch,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "DELIVRIO",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

//go:embed ent/migrate/migrations
var migrations embed.FS

func migrateDB(ctx context.Context, sqlDriver *sql.Driver) {
	drv, err := sqlite3.WithInstance(sqlDriver.DB(), &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mFS, err := iofs.New(migrations, "ent/migrate/migrations")
	if err != nil {
		log.Fatal("[db-migrations] fs: ", err)
	}

	m, err := migrate2.NewWithInstance(
		"iofs",
		mFS,
		"sqlite3",
		drv,
	)
	if err != nil {
		log.Fatal("[db-migrations] db instance: ", err)
	}

	// Disable all FKs since go-migrate executes in a TX and
	// PRAGMA is no-op in a TX
	// https://github.com/mattn/go-sqlite3/issues/377#issuecomment-275882103
	_, err = sqlDriver.DB().Exec(`PRAGMA foreign_keys = off;`)
	if err != nil {
		log.Fatal("[db-migrations] disable FKs: ", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate2.ErrNoChange) {
		log.Fatal("[db-migrations] apply changes: ", err)
	}

	_, err = sqlDriver.DB().Exec(`PRAGMA foreign_keys = on;`)
	if err != nil {
		log.Fatal("[db-migrations] enable FKs: ", err)
	}

	v, e, p := m.Version()
	log.Println("[db-migrations] database up-to-date", v, e, p)

}
