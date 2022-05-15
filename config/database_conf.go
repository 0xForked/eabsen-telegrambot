package config

import (
	"fmt"
	"github.com/aasumitro/svc-tgbotgo/src/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

const (
	mysqlDriver = "mysql"
	sqliteDriver = "sqlite"
)

var availableDatabaseDriver = []string{mysqlDriver, sqliteDriver}

var db *gorm.DB

func (config AppConfig) InitDatabaseConnection() {
	if !helpers.InArray(config.GetDatabaseDriver(), availableDatabaseDriver) {
		panic(fmt.Sprintf(
			helpers.DbDriverNotAvailable,
			config.GetDatabaseDriver()))
	}

	log.Println("Initialize database connection . . .")
	conn, err := openConnection(config)
	if err != nil {
		log.Panicln(fmt.Sprintf("database connection: %s", err.Error()))
	}

	log.Println("database connection: connected")
	setConnection(conn)
}

func openConnection(config AppConfig) (db *gorm.DB, err error) {
	if config.GetDatabaseDriver() == mysqlDriver {
		return gorm.Open(
			mysql.Open(config.GetMySqlDsnUrl()),
			&gorm.Config{})
	}

	return gorm.Open(
		sqlite.Open(config.GetSqlitePath()),
		&gorm.Config{})
}

func setConnection(conn *gorm.DB)  {
	db = conn
}

func (config AppConfig) GetDatabaseConnection() *gorm.DB {
	return db
}
