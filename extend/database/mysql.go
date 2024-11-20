package database

import (
	"fmt"
	"go_custom/extend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConfig struct {
	dbname    string
	username  string
	password  string
	host      string
	port      int
	charset   string
	parseTime bool
	loc       string
}

var MysqlInstance mysqlConfig

func (mysqlConfig) getConfig() mysqlConfig {
	dbc := config.Mapping.Database
	return mysqlConfig{
		dbname:    dbc.Name,
		username:  dbc.Username,
		password:  dbc.Password,
		host:      dbc.Host,
		port:      dbc.Port,
		charset:   dbc.Charset,
		parseTime: true,
		loc:       "Local",
	}
}

func (mysqlConfig) dsn(username string, password string, host string, port int,
	dbname string, charset string, parsetime bool, local string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, dbname, charset, parsetime, local)
}

func (mysqlConfig) Open() error {
	MysqlInstance = mysqlConfig.getConfig(mysqlConfig{})
	dsn := mysqlConfig.dsn(mysqlConfig{}, MysqlInstance.username, MysqlInstance.password, MysqlInstance.host,
		MysqlInstance.port, MysqlInstance.dbname, MysqlInstance.charset, MysqlInstance.parseTime, MysqlInstance.loc)
	SetDbLog()
	dbInstance, err := gorm.Open(mysql.Open(dsn), &GConfig)
	DB = dbInstance
	return err
}
