package pg

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/config"
	"app/pkg"
)

type gormConnection struct {
	DB *gorm.DB
}

var Gorm gormConnection

/* --------------------------------- Connect -------------------------------- */
// Connect to database and fill connection.DB
/* -------------------------------------------------------------------------- */
func (c *gormConnection) Connect() {
	appConfig := config.AppConfig

	var dsn string
	if len(appConfig.DB_DSN) > 0 {
		dsn = appConfig.DB_DSN
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			appConfig.DB_HOST,
			appConfig.DB_USER,
			appConfig.DB_PASSWORD,
			appConfig.DB_NAME,
			appConfig.DB_PORT,
			appConfig.DB_SSL_MODE,
			appConfig.DB_TIMEZONE,
		)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		pkg.Logger.Critical(err.Error())
		log.Panic(err)
	}

	c.DB = db
}

/* --------------------------------- Migrate -------------------------------- */
// Create Table If Not Exist
// Before use c.Migrate you have to run c.Connect() to fill c.DB
/* -------------------------------------------------------------------------- */
func (c *gormConnection) Migrate(modelStruct interface{}) {
	err := c.DB.AutoMigrate(&modelStruct)
	if err != nil {
		return
	}
}
