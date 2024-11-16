package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInterface interface {
	dsn(username string, password string, host string, port int, dbname string) string
	Open() (*gorm.DB, error)
}

var GConfig gorm.Config

var DB *gorm.DB

func SetDbLog() {
	GConfig.Logger = logger.Default.LogMode(logger.Info)
}
