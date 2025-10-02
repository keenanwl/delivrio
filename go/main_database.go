package main

import (
	"context"
	dbsql "database/sql"
	"delivrio.io/go/appconfig"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/migrate/migrationsdata"
	"embed"
	"entgo.io/ent/dialect/sql"
	"errors"
	"github.com/aws/aws-sdk-go-v2/credentials"
	s3v2 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/logging"
	migrate2 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
	"gocloud.dev/blob/s3blob"
	"log"
	"os"
	"time"

	awsv2cfg "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/lib/pq"
	_ "gocloud.dev/blob/s3blob"
)

func migrateData(ctx context.Context, sqlDriver *dbsql.DB, client *ent.Client) {
	_, err := sqlDriver.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations_data (
		version INTEGER,
		dirty INTEGER
	)`)
	if err != nil {
		log.Fatal(err)
	}

	err = migrationsdata.Run(ctx, sqlDriver, client, 2, migrationsdata.AddSEK(ctx))
	if err != nil {
		log.Fatal(err)
	}

}

//go:embed ent/migrate/migrations
var migrations embed.FS

func migrateDBPostgres(db *dbsql.DB) {
	drv, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mFS, err := iofs.New(migrations, "ent/migrate/migrations/postgres")
	if err != nil {
		log.Fatal("[db-migrations-postgres] fs: ", err)
	}

	m, err := migrate2.NewWithInstance(
		"iofs",
		mFS,
		"postgres",
		drv,
	)
	if err != nil {
		log.Fatal("[db-migrations-postgres] db instance: ", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate2.ErrNoChange) {
		log.Fatal("[db-migrations-postgres] apply changes: ", err)
	}

	v, e, p := m.Version()
	log.Println("[db-migrations-postgres] database up-to-date", v, e, p)

}

func migrateDBSQLite(sqlDriver *dbsql.DB) {
	drv, err := sqlite3.WithInstance(sqlDriver, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mFS, err := iofs.New(migrations, "ent/migrate/migrations/sqlite")
	if err != nil {
		log.Fatal("[db-migrations-sqlite] fs: ", err)
	}

	m, err := migrate2.NewWithInstance(
		"iofs",
		mFS,
		"sqlite3",
		drv,
	)
	if err != nil {
		log.Fatal("[db-migrations-postgres] db instance: ", err)
	}

	// Disable all FKs since go-migrate executes in a TX and
	// PRAGMA is no-op in a TX
	// https://github.com/mattn/go-sqlite3/issues/377#issuecomment-275882103
	_, err = sqlDriver.Exec(`PRAGMA foreign_keys = off;`)
	if err != nil {
		log.Fatal("[db-migrations-postgres] disable FKs: ", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate2.ErrNoChange) {
		log.Fatal("[db-migrations-postgres] apply changes: ", err)
	}

	_, err = sqlDriver.Exec(`PRAGMA foreign_keys = on;`)
	if err != nil {
		log.Fatal("[db-migrations-postgres] enable FKs: ", err)
	}

	v, e, p := m.Version()
	log.Println("[db-migrations-postgres] database up-to-date", v, e, p)

}

// BasicLogger is a basic implementation of the Logger interface.
type BasicLogger struct {
	logger *log.Logger
}

// NewBasicLogger creates a new BasicLogger.
func NewBasicLogger() *BasicLogger {
	return &BasicLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Logf logs messages with the given classification and format.
func (l *BasicLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	prefix := ""
	l.logger.Printf(prefix+format, v...)
}

func openBucket(config2 appconfig.DelivrioConfig) (*blob.Bucket, error) {
	switch config2.BlobStorage.Type {
	case appconfig.BlobStorageTypeBucketLocal:
		return fileblob.OpenBucket(config2.BlobStorage.LocalPath, &fileblob.Options{
			CreateDir: true,
			NoTempDir: true,
		})
	case appconfig.BlobStorageTypeBucketS3:
		cfg, err := awsv2cfg.LoadDefaultConfig(
			context.Background(),
			awsv2cfg.WithDefaultRegion(config2.BlobStorage.S3Region),
			awsv2cfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				config2.BlobStorage.S3AccessKey,
				config2.BlobStorage.S3SecretKey,
				"",
			)),
		)
		if err != nil {
			return nil, err
		}

		clientV2 := s3v2.NewFromConfig(cfg)
		return s3blob.OpenBucketV2(
			context.Background(),
			clientV2,
			config2.BlobStorage.S3BucketName,
			nil,
		)
	case appconfig.BlobStorageTypeDatabase:
		fallthrough
	default:
		return nil, nil
	}
}

func OpenPostgres(config2 appconfig.DelivrioConfig) (*ent.Client, *dbsql.DB, error) {
	drv, err := sql.Open("postgres", config2.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	bucket, err := openBucket(config2)
	if err != nil {
		return nil, nil, err
	}

	return ent.NewClient(ent.Driver(drv), ent.Bucket(bucket)), db, nil
}

func OpenSQLite(config2 appconfig.DelivrioConfig) (*ent.Client, *dbsql.DB, error) {
	drv, err := sql.Open("sqlite3", config2.DatabaseDSN)
	if err != nil {
		return nil, nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(time.Minute * 10)

	bucket, err := openBucket(config2)
	if err != nil {
		return nil, nil, err
	}

	return ent.NewClient(ent.Driver(drv), ent.Bucket(bucket)), db, nil
}
