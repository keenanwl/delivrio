package appconfig

import (
	"crypto/rand"
	"delivrio.io/go/carrierapis/dfapis/dfrequest"
	"encoding/base64"
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type DelivrioConfig struct {
	ConfigVersion int    `yaml:"config_version"`
	ServerID      string `yaml:"server_id"`
	LimitedSystem bool   `yaml:"limited_system"`

	BaseURL     string       `yaml:"base_url"`
	Production  bool         `yaml:"production"`
	Database    DatabaseType `yaml:"database"`
	DatabaseDSN string       `yaml:"database_dsn"`
	GoBindPort  string       `yaml:"go_bind_port"`

	// JWT
	JWTKey string `yaml:"jwt_key"`

	Email Email `yaml:"email"`

	Bring            BringConfig              `yaml:"bring"`
	PostNord         PostNordConfig           `yaml:"post_nord"`
	GLS              GLSConfig                `yaml:"gls"`
	DanskeFragtmaend dfrequest.APICredentials `yaml:"danske_fragtmaend"`

	// Background jobs
	BackgroundJobPolicy BackgroundJobPolicy `yaml:"background_job_policy"`
	LoggingPolicy       LoggingPolicy       `yaml:"logging_policy"`

	// APM
	APM SignozConfig `yaml:"apm"`

	// Path to the go-cli wrapper around PDFium
	// for converting PDF->PNG and subsequently a ZPL "image"
	PathPDFium string `yaml:"path_pdfium"`

	HTML2PDF HTML2PDF `yaml:"html_2_pdf"`

	BlobStorage BlobStorage `yaml:"blob_storage"`
}

type BlobStorageType string

const (
	BlobStorageTypeDatabase    BlobStorageType = "database"
	BlobStorageTypeBucketLocal BlobStorageType = "bucket_local"
	BlobStorageTypeBucketS3    BlobStorageType = "bucket_s3"
)

type BlobStorage struct {
	Type BlobStorageType `yaml:"type"`

	LocalPath string `yaml:"local_path"`

	S3Name       string `yaml:"s3_name"`
	S3Region     string `yaml:"s3_region"`
	S3AccessKey  string `yaml:"s3_access_key"`
	S3SecretKey  string `yaml:"s3_secret_key"`
	S3BucketName string `yaml:"s3_bucket_name"`
}

type HTML2PDF struct {
	Gotenberg *Gotenberg `yaml:"gotenberg"`
}

// Gotenberg may be run in any Docker supported environment
// (also localhost if DELIVRIO is run locally)
// https://gotenberg.dev/docs/routes
type Gotenberg struct {
	Base string `yaml:"base"`
	// Just copy and past the keyfile.json contents as a string value.
	// For our purposes, it is an API token to prevent external abuse
	// but can be used without limit for DELIVRIO activities
	CloudRunAuthJSON *string `yaml:"cloud_run_auth_json"`
}

type DatabaseType string

const (
	DatabaseTypeMySQL    DatabaseType = "mysql"
	DatabaseTypeSQLite   DatabaseType = "sqlite"
	DatabaseTypePostgres DatabaseType = "postgres"
)

type LoggingPolicy string

const (
	LoggingPolicyBackground LoggingPolicy = "background"
	LoggingPolicyConsole    LoggingPolicy = "console"
	LoggingPolicyBoth       LoggingPolicy = "both"
	LoggingPolicyNone       LoggingPolicy = "none"
)

type BackgroundJobPolicy string

const (
	BackgroundJobPolicyOnce     BackgroundJobPolicy = "once"
	BackgroundJobPolicyLoop     BackgroundJobPolicy = "loop"
	BackgroundJobPolicyDisabled BackgroundJobPolicy = "disabled"
)

type PostNordConfig struct {
	APIKey string `yaml:"api_key"`
}

type BringConfig struct {
	APIKey string `yaml:"api_key"`
	APIUID string `yaml:"api_uid"`
}

// Might need to move to APP if is per-client stuff
type GLSConfig struct {
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
}

type SignozConfig struct {
	Endpoint string `yaml:"endpoint"`
	// Http(s)
	Insecure bool `yaml:"insecure"`
	Enabled  bool `yaml:"enabled"`
}

type Email struct {
	DefaultFrom string        `yaml:"default_from"`
	Mailgun     *EmailMailgun `yaml:"mailgun"`
}

type EmailMailgun struct {
	MGAPIKey string `yaml:"mg_api_key"`
	MGDomain string `yaml:"mg_domain"`
	MGURL    string `yaml:"mg_url"`
}

func generateRandomSecretKey() string {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(key)
}

var defaultCOnfig = &DelivrioConfig{
	ConfigVersion:       1,
	ServerID:            "delivrio-1",
	LimitedSystem:       false,
	BaseURL:             "http://localhost:8080",
	Production:          true,
	Database:            DatabaseTypeSQLite,
	DatabaseDSN:         "file:./delivrio.db?cache=shared&_fk=1",
	GoBindPort:          "8080",
	JWTKey:              generateRandomSecretKey(),
	BackgroundJobPolicy: BackgroundJobPolicyLoop,
	LoggingPolicy:       LoggingPolicyBoth,
	APM:                 SignozConfig{},
	// TODO: figure out default path
	PathPDFium: "/usr/local/bin/pdfium",
	HTML2PDF:   HTML2PDF{},
}

func Default() []byte {
	out, _ := yaml.Marshal(defaultCOnfig)
	return out
}

// Parse looks for the config file, or returns
// the default config
func Parse(path string) (*DelivrioConfig, error) {

	_, err := os.Stat(path)
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		return defaultCOnfig, nil
	} else if err != nil {
		return nil, err
	}

	cnf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf DelivrioConfig
	err = yaml.Unmarshal(cnf, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
