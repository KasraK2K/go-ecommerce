package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type config struct {
	Port                 string   `json:"port"`
	Mode                 string   `json:"mode"`
	PreFork              bool     `json:"pre_fork"`
	BackendVersion       string   `json:"backend_version"`
	FrontendVersion      string   `json:"frontend_version"`
	AppVersion           string   `json:"app_version"`
	StdoutLogs           string   `json:"stdout_logs"`
	FileLogs             string   `json:"file_logs"`
	UploadPath           string   `json:"upload_path"`
	ValidUploadEndpoints []string `json:"valid_upload_endpoints"`
	FileExtensions       []string `json:"file_extensions"`
	FileSize             int64    `json:"file_size"`
	StaticFileUrl        string   `json:"static_file_url"`
	ServerDomain         string   `json:"server_domain"`
	DbDsn                string   `json:"db_dsn"`
	DbHost               string   `json:"db_host"`
	DbPort               string   `json:"db_port"`
	DbPassword           string   `json:"dn_password"`
	DbUser               string   `json:"db_user"`
	DbName               string   `json:"db_name"`
	DbTimezone           string   `json:"db_timezone"`
	DbSslMode            string   `json:"db_ssl_mode"`
	MongodbUri           string   `json:"mongodb_uri"`
	RedisAddress         string   `json:"redis_address"`
	RedisPassword        string   `json:"redis_password"`
	JwtSigningKey        string   `json:"jwt_signing_key"`
	MailgunPrivateApiKey string   `json:"mailgun_private_api_key"`
	MailgunDomain        string   `json:"mailgun_domain"`
	MailgunApiBase       string   `json:"mailgun_api_base"`
	MailgunSender        string   `json:"mailgun_sender"`
}

var AppConfig config

func SetConfig() {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	mode := os.Getenv("Mode")
	var envPath string
	if len(mode) > 0 && mode == "production" {
		envPath = "config/.env.production"
	} else {
		envPath = filepath.Join(wd, "config/.env.development")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal(err)
	}

	preFork, err := strconv.ParseBool(os.Getenv("PREFORK"))
	if err != nil {
		log.Fatal(err)
	}

	validUploadEndpoints := strings.Split(os.Getenv("VALID_UPLOAD_ENDPOINTS"), ",")
	fileExtensions := strings.Split(os.Getenv("FILE_EXTENSIONS"), ",")

	fileSize, err := strconv.ParseInt(os.Getenv("FILE_SIZE"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Application
	AppConfig.Port = os.Getenv("PORT")
	AppConfig.Mode = os.Getenv("MODE")
	AppConfig.PreFork = preFork
	AppConfig.BackendVersion = os.Getenv("BACKEND_VERSION")
	AppConfig.FrontendVersion = os.Getenv("FRONTEND_VERSION")
	AppConfig.AppVersion = os.Getenv("APP_VERSION")
	AppConfig.StdoutLogs = os.Getenv("STDOUT_LOGS")
	AppConfig.FileLogs = os.Getenv("FILE_LOGS")
	AppConfig.UploadPath = os.Getenv("UPLOAD_PATH")
	AppConfig.ValidUploadEndpoints = validUploadEndpoints
	AppConfig.FileExtensions = fileExtensions
	AppConfig.FileSize = fileSize
	AppConfig.StaticFileUrl = os.Getenv("STATIC_FILE_URL")
	AppConfig.ServerDomain = os.Getenv("SERVER_DOMAIN")
	// PostgreSQL Database
	AppConfig.DbDsn = os.Getenv("DB_DSN")
	AppConfig.DbHost = os.Getenv("DB_HOST")
	AppConfig.DbPort = os.Getenv("DB_PORT")
	AppConfig.DbPassword = os.Getenv("DB_PASSWORD")
	AppConfig.DbUser = os.Getenv("DB_USER")
	AppConfig.DbName = os.Getenv("DB_NAME")
	AppConfig.DbTimezone = os.Getenv("DB_TIMEZONE")
	AppConfig.DbSslMode = os.Getenv("DB_SSL_MODE")
	// MongoDB Database
	AppConfig.MongodbUri = os.Getenv("MONGODB_URI")
	// Redis Database
	AppConfig.RedisAddress = os.Getenv("REDIS_ADDRESS")
	AppConfig.RedisPassword = os.Getenv("REDIS_PASSWORD")
	// JWT
	AppConfig.JwtSigningKey = os.Getenv("JWT_SIGNING_KEY")
	// MailGun
	AppConfig.MailgunPrivateApiKey = os.Getenv("MAILGUN_PRIVATE_API_KEY")
	AppConfig.MailgunDomain = os.Getenv("MAILGUN_DOMAIN")
	AppConfig.MailgunApiBase = os.Getenv("MAILGUN_API_BASE")
	AppConfig.MailgunSender = os.Getenv("MAILGUN_SENDER")
}
