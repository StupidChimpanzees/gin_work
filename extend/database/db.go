package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInterface interface {
	dsn(username string, password string, host string, port int,
		dbname string, charset string, parseTime bool, local string) string
	Open() error
}

var GConfig gorm.Config

var DB *gorm.DB

func init() {
	err := LoadDB(MysqlInstance)
	if err != nil {
		return
	}
}

func SetDbLog() {
	GConfig.Logger = logger.Default.LogMode(logger.Info)
}

func LoadDB(db dbInterface) error {
	return db.Open()
}
