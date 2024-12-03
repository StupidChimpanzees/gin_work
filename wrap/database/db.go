package database

import (
	"gin_work/wrap/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInterface interface {
	dsn(username string, password string, host string, port int,
		dbname string, charset string, parseTime bool, local string) string
	Open()
}

type dbConfig struct {
	DBType    string
	DBName    string
	Username  string
	Password  string
	Host      string
	Port      int
	Charset   string
	ParseTime bool
	Loc       string
}

var DBConfig *dbConfig

var DB *gorm.DB

func init() {
	DBConfig = getConfig()
	if DBConfig.DBType == "mysql" {
		MysqlInstance.Open()
	}
}

func getConfig() *dbConfig {
	dbc := config.Mapping.Database
	return &dbConfig{
		DBType:    dbc.DBType,
		DBName:    dbc.Name,
		Username:  dbc.Username,
		Password:  dbc.Password,
		Host:      dbc.Host,
		Port:      dbc.Port,
		Charset:   dbc.Charset,
		ParseTime: true,
		Loc:       "Local",
	}
}

func SetDbLog(conf *gorm.Config) {
	conf.Logger = logger.Default.LogMode(logger.Info)
}
