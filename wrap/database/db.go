package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInterface interface {
	dsn(username string, password string, host string, port int,
		dbname string, charset string, parseTime bool, local string) string
	Open()
}

var GConfig gorm.Config

var DB *gorm.DB

func init() {
	LoadDB(MysqlInstance)
}

func LoadDB(db dbInterface) {
	db.Open()
}

func SetDbLog() {
	GConfig.Logger = logger.Default.LogMode(logger.Info)
}
