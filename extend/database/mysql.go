package database

import (
	"fmt"
	"go_custom/extend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlInstance struct {
	dbname    string
	username  string
	password  string
	host      string
	port      int
	charset   string
	parseTime bool
	loc       string
}

func (mysqlInstance) getConfig() mysqlInstance {
	dbc := config.Mapping.Database
	return mysqlInstance{
		dbname:    dbc.Name,
		username:  dbc.Username,
		password:  dbc.Password,
		host:      dbc.Host,
		port:      dbc.Port,
		charset:   "utf8mb4",
		parseTime: true,
		loc:       "local",
	}
}

func (mysqlInstance) dsn(username string, password string, host string, port int, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=local", username, password, host, port, dbname)
}

func (mysqlInstance) Open() error {
	dbConfig := mysqlInstance.getConfig(mysqlInstance{})
	dsn := mysqlInstance.dsn(mysqlInstance{}, dbConfig.username, dbConfig.password, dbConfig.host, dbConfig.port, dbConfig.dbname)
	SetDbLog()
	dbInstance, err := gorm.Open(mysql.Open(dsn), &GConfig)
	DB = dbInstance
	return err
}
