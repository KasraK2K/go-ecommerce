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
	if len(appConfig.DbDsn) > 0 {
		dsn = appConfig.DbDsn
	} else {
		dsn = fmt.Sprintf("host=%s users=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			appConfig.DbHost,
			appConfig.DbUser,
			appConfig.DbPassword,
			appConfig.DbName,
			appConfig.DbPort,
			appConfig.DbSslMode,
			appConfig.DbTimezone,
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
