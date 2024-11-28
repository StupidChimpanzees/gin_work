package redis

import (
	"gin_work/wrap/config"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var r Raede

func init() {
	r.rc = getConfig()
}

func getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func RConn() (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if r.rc.Password == "" {
		conn, err = redis.Dial("tcp", r.rc.Host+":"+strconv.Itoa(r.rc.Port))
	} else {
		conn, err = redis.Dial("tcp", r.rc.Host+":"+strconv.Itoa(r.rc.Port), redis.DialPassword(r.rc.Password))
	}
	if err != nil {
		panic("redis connection error: " + err.Error())
	}
	SelectDB()
	return conn, err
}

func SelectDB() error {
	r.conn, _ = RConn()
	return r.SelectDB(r.conn, r.rc.Select)
}

func Set(name, value string) (bool, error) {
	r.conn, _ = RConn()
	return r.Set(r.conn, r.rc.Prefix+name, value)
}

func Get(name string) string {
	r.conn, _ = RConn()
	return r.Get(r.conn, r.rc.Prefix+name)
}
