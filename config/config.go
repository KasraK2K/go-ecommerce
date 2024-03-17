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
	PORT                    string   `json:"port"`
	MODE                    string   `json:"mode"`
	PREFORK                 bool     `json:"prefork"`
	BACKEND_VERSION         string   `json:"backend_version"`
	FRONTEND_VERSION        string   `json:"frontend_version"`
	APP_VERSION             string   `json:"app_version"`
	STDOUT_LOGS             string   `json:"stdout_logs"`
	FILE_LOGS               string   `json:"file_logs"`
	UPLOAD_PATH             string   `json:"upload_path"`
	VALID_UPLOAD_ENDPOINTS  []string `json:"valid_upload_endpoints"`
	FILE_EXTENSIONS         []string `json:"file_extensions"`
	FILE_SIZE               int64    `json:"file_size"`
	STATIC_FILE_URL         string   `json:"static_file_url"`
	SERVER_DOMAIN           string   `json:"server_domain"`
	DB_DSN                  string   `json:"db_dsn"`
	DB_HOST                 string   `json:"db_host"`
	DB_PORT                 string   `json:"db_port"`
	DB_PASSWORD             string   `json:"dn_password"`
	DB_USER                 string   `json:"db_user"`
	DB_NAME                 string   `json:"db_name"`
	DB_TIMEZONE             string   `json:"db_timezone"`
	DB_SSL_MODE             string   `json:"db_ssl_mode"`
	MONGODB_URI             string   `json:"mongodb_uri"`
	REDIS_ADDRESS           string   `json:"redis_address"`
	REDIS_PASSWORD          string   `json:"redis_password"`
	JWT_SIGNING_KEY         string   `json:"jwt_signing_key"`
	MAILGUN_PRIVATE_API_KEY string   `json:"mailgun_private_api_key"`
	MAILGUN_DOMAIN          string   `json:"mailgun_domain"`
	MAILGUN_API_BASE        string   `json:"mailgun_api_base"`
	MAILGUN_SENDER          string   `json:"mailgun_sender"`
}

var AppConfig config

func SetConfig() {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	mode := os.Getenv("MODE")
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

	prefork, err := strconv.ParseBool(os.Getenv("PREFORK"))
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
	AppConfig.PORT = os.Getenv("PORT")
	AppConfig.MODE = os.Getenv("MODE")
	AppConfig.PREFORK = prefork
	AppConfig.BACKEND_VERSION = os.Getenv("BACKEND_VERSION")
	AppConfig.FRONTEND_VERSION = os.Getenv("FRONTEND_VERSION")
	AppConfig.APP_VERSION = os.Getenv("APP_VERSION")
	AppConfig.STDOUT_LOGS = os.Getenv("STDOUT_LOGS")
	AppConfig.FILE_LOGS = os.Getenv("FILE_LOGS")
	AppConfig.UPLOAD_PATH = os.Getenv("UPLOAD_PATH")
	AppConfig.VALID_UPLOAD_ENDPOINTS = validUploadEndpoints
	AppConfig.FILE_EXTENSIONS = fileExtensions
	AppConfig.FILE_SIZE = fileSize
	AppConfig.STATIC_FILE_URL = os.Getenv("STATIC_FILE_URL")
	AppConfig.SERVER_DOMAIN = os.Getenv("SERVER_DOMAIN")
	// PostgreSQL Database
	AppConfig.DB_DSN = os.Getenv("DB_DSN")
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_USER = os.Getenv("DB_USER")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
	AppConfig.DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	// MongoDB Database
	AppConfig.MONGODB_URI = os.Getenv("MONGODB_URI")
	// Redis Database
	AppConfig.REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
	AppConfig.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	// JWT
	AppConfig.JWT_SIGNING_KEY = os.Getenv("JWT_SIGNING_KEY")
	// MailGun
	AppConfig.MAILGUN_PRIVATE_API_KEY = os.Getenv("MAILGUN_PRIVATE_API_KEY")
	AppConfig.MAILGUN_DOMAIN = os.Getenv("MAILGUN_DOMAIN")
	AppConfig.MAILGUN_API_BASE = os.Getenv("MAILGUN_API_BASE")
	AppConfig.MAILGUN_SENDER = os.Getenv("MAILGUN_SENDER")
}
