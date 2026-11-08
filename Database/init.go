package database

import (
	"database/sql"
	"fmt"
	config "hashbat-image-bot/Config"
	logger "hashbat-image-bot/Logger"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() bool {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Settings.DB_USER,
		config.Settings.DB_PASSWORD,
		config.Settings.DB_IP_ADDRESS,
		config.Settings.DB_PORT,
		config.Settings.DB_NAME)

	dbIn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Error("", err)
		return false
	}

	Db = dbIn
	return true
}
