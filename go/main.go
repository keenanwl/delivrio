package main

import (
	"context"
	dbsql "database/sql"
	"delivrio.io/go/apm"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/background"
	"delivrio.io/go/seed"
	"delivrio.io/go/utils"
	"errors"
	"flag"
	"fmt"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/toqueteos/webbrowser"
	"go.opentelemetry.io/otel"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"syscall"
	"time"

	"delivrio.io/go/ent"
	"delivrio.io/go/viewer"

	_ "go.opentelemetry.io/otel/trace"

	// Required by Ent/Privacy
	_ "delivrio.io/go/ent/runtime"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	client    *ent.Client
	Clock     jwt.ClockFunc = time.Now
	tracer                  = otel.Tracer("main")
	tokenAuth *jwtauth.JWTAuth

	// These variables will be set at build time using ldflags
	AppVersion string
	BuildTime  string

	lastVacuum time.Time = time.Now()
)

func main() {

	ctx := context.Background()
	seedBase := flag.Bool("init-seed", false, "Run once on initial production installation")
	seedDemo := flag.Bool("seed", false, "Run once for initial demo data")
	configPath := flag.String("config", "delivrio_config.yaml", "Path to config file, otherwise uses default")
	testClock := flag.String("test-clock", "", fmt.Sprintf("Timestamp for testing, expect format: '%v'", time.RFC3339))
	testSeed := flag.Bool("test-seed", false, "Adds state to DB for e2e tests")
	createConfig := flag.Bool("create-config", false, "Writes the default config file in the working directory if it doesn't exist")
	flag.Parse()

	if *createConfig {
		err := os.WriteFile("delivrio_config.yaml", appconfig.Default(), 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created default config file")
		os.Exit(0)
	}

	conf, err := appconfig.Parse(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(*testClock) > 0 && !conf.Production {
		val, err := time.Parse(time.RFC3339, *testClock)
		if err != nil {
			log.Fatal("setting test clock: ", err)
		}
		Clock = func() time.Time { return val }
	}

	chiLogger := Logging(conf.LoggingPolicy)
	setConf(conf)

	cleanup := apm.InitTracer(conf)
	defer func() {
		if err := cleanup(ctx); err != nil {
			log.Println("cleanup APM: " + err.Error())
		}
	}()
	log.Println("Open DB")
	db, sqlDriver, err := openDatabase(*conf)
	if err != nil {
		log.Fatal(err)
	}

	client = db.Debug()
	defer client.Close()
	defer db.Bucket.Close()

	client.Intercept(apm.QueryTracer(conf.ServerID))

	ctx = ent.NewContext(ctx, client)
	migrateDatabase(*conf, sqlDriver)

	err = utils.InitBarcodeCounter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := db.CarrierBrand.Query().
		Count(viewer.NewBackgroundContext(ctx))
	if err != nil {
		log.Fatal(err)
	}

	if *seedBase || (isDesktop && rowCount == 0) {
		seed.ProductionBaseTx(ctx)
		log.Println("Successfully initialized production seed data.")

		if !conf.Production && *testSeed {
			seed.E2E(ctx)
			log.Println("Successfully initialized e2e seed data.")
		}

	} else if *seedDemo && !conf.Production {
		seed.DemoData(ctx)
	} else if *seedDemo && conf.Production {
		log.Println("error: invalid config: production config cannot be toggled with demo data")
	} else {
		migrateData(ctx, sqlDriver, client)
	}

	if isDesktop {
		url := fmt.Sprintf("http://localhost:%s", conf.GoBindPort)
		err = webbrowser.Open(url)
		if err != nil {
			log.Println("Error opening browser:", err)
		} else {
			log.Println("Successfully opened browser to", url)
		}
	}

	router, err := configureRouter(conf.GoBindPort, *conf, chiLogger)
	if err != nil {
		log.Fatalf("configureRouter(): %v\n", err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.GoBindPort),
		Handler: router,
	}

	clientCTX := ent.NewContext(viewer.NewBackgroundContext(context.Background()), client)
	go startBackgroundJobs(clientCTX, sqlDriver, conf.Database == appconfig.DatabaseTypePostgres, conf.BackgroundJobPolicy)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	log.Printf("Server is ready to handle requests at :%v\n", conf.GoBindPort)
	<-quit
	log.Println("Server is shutting down...")

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Should support a graceful shutdown
	// (1) Finish up any existing requests with a 5s max timeout
	// (2) Nginx holds any new connections until we are back up again
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Recieved shutdown request: %v", err)
	}

	log.Println("Server exiting")

}

func openDatabase(conf appconfig.DelivrioConfig) (*ent.Client, *dbsql.DB, error) {
	switch conf.Database {
	case appconfig.DatabaseTypePostgres:
		return OpenPostgres(conf)
	}
	return OpenSQLite(conf)
}

func migrateDatabase(conf appconfig.DelivrioConfig, db *dbsql.DB) {
	switch conf.Database {
	case appconfig.DatabaseTypePostgres:
		migrateDBPostgres(db)
		return
	case appconfig.DatabaseTypeSQLite:
		migrateDBSQLite(db)
		return
	}
	log.Fatal("database type not supported", conf.Database)
}

func Logging(policy appconfig.LoggingPolicy) *httplog.Logger {
	logDirectory := "./logs"
	err := os.MkdirAll(logDirectory, 0777)
	if err != nil {
		log.Fatal(err)
	}

	lumberjackLogger := &lumberjack.Logger{
		// Log file absolute path, os agnostic
		Filename:   filepath.ToSlash(filepath.Join(logDirectory, "delivrio.log")),
		MaxSize:    5, // MB
		MaxBackups: 10,
		MaxAge:     30,    // days
		Compress:   false, // disabled by default
	}

	logWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	if policy == appconfig.LoggingPolicyBackground {
		logWriter = lumberjackLogger
	} else if policy == appconfig.LoggingPolicyNone {
		logWriter = io.Discard
	} else if policy == appconfig.LoggingPolicyConsole {
		logWriter = io.MultiWriter(os.Stdout)
	}
	log.SetOutput(logWriter)
	return httplog.NewLogger("delivrio-server", httplog.Options{
		// JSON:             true,
		LogLevel:         slog.LevelDebug,
		Writer:           logWriter,
		Concise:          true,
		RequestHeaders:   false,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})
}

func startBackgroundJobs(ctx context.Context, driver *dbsql.DB, isPG bool, policy appconfig.BackgroundJobPolicy) {
	for {
		finished := func() bool {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered from panic: %v", r)
					log.Println("stacktrace from panic: \n" + string(debug.Stack()))
					time.Sleep(time.Second * 60)
				}
			}()

			if isPG {
				lastVacuum = background.HandleDBVacuum(ctx, driver, lastVacuum)
			}

			// Reset the change history ID on each run
			ctx = viewer.NewBackgroundContext(ctx)

			if policy == appconfig.BackgroundJobPolicyLoop {
				background.HandleBackgroundJobs(ctx, false)
				return false
			} else if policy == appconfig.BackgroundJobPolicyOnce {
				background.HandleBackgroundJobs(ctx, true)
				return true
			}

			return true
		}()
		if finished {
			break
		}
	}
}
