package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct{}

func (*MysqlConf) dsn(username string, password string, host string, port int, dbname string, charset string,
	parseTime bool, local string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, dbname, charset, parseTime, local)
}

func (m *MysqlConf) Open() {
	dsn := m.dsn(DBConfig.Username, DBConfig.Password, DBConfig.Host,
		DBConfig.Port, DBConfig.DBName, DBConfig.Charset, DBConfig.ParseTime, DBConfig.Loc)
	conf := gorm.Config{}
	SetDbLog(&conf)
	dbInstance, err := gorm.Open(mysql.Open(dsn), &conf)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	DB = dbInstance
}
