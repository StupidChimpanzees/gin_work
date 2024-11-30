package redis

import (
	"gin_work/wrap/config"
	"github.com/gomodule/redigo/redis"
)

func init() {
	conf := getConfig()
	raeder = NewRaede(conf.Host, conf.Port, conf.Password, conf.Prefix, conf.Select)
}

func getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func getConn() redis.Conn {
	if raeder.Conn != nil {
		return raeder.Conn
	}
	conn, _ := raeder.Connection()
	raeder.Conn = conn.(redis.Conn)
	return raeder.Conn
}

func Set(name, value string) (string, error) {
	return raeder.Set(raeder.Prefix+name, value)
}

func Get(name string) string {
	return raeder.Get(raeder.Prefix + name)
}

func HSet(name string, value interface{}) (string, error) {
	return raeder.HSet(raeder.Prefix+name, value)
}

func HGet(name string) interface{} {
	return raeder.HGet(raeder.Prefix + name)
}
