package redis

import (
	"gin_work/wrap/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

type poolConf struct {
	maxIdle        int
	maxActive      int
	idleTimeout    time.Duration
	maxConnTimeout time.Duration
}

var pool *redis.Pool

func init() {
	conf := getPoolConfig()
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := raeder.Connection()
			return conn.(redis.Conn), err
		},
		MaxIdle:         conf.maxIdle,
		MaxActive:       conf.maxActive,
		IdleTimeout:     conf.idleTimeout,
		MaxConnLifetime: conf.maxConnTimeout,
		Wait:            true,
	}
}

func getPoolConfig() *poolConf {
	pConfig := config.Mapping.Redis.Pool
	conf := &poolConf{
		maxIdle:        10,
		maxActive:      100,
		idleTimeout:    time.Duration(120),
		maxConnTimeout: time.Duration(3),
	}
	if pConfig.MaxIdle != 0 {
		conf.maxIdle = pConfig.MaxIdle
	}
	if pConfig.MaxActive != 0 {
		conf.maxActive = pConfig.MaxActive
	}
	if pConfig.IdleTimeout != 0 {
		conf.idleTimeout = time.Duration(pConfig.IdleTimeout)
	}
	if pConfig.IdleTimeout != 0 {
		conf.maxConnTimeout = time.Duration(pConfig.MaxConnTimeout)
	}
	return conf
}

func PSet(name, value string) (string, error) {
	_, err := raeder.Connection(pool.Get())
	if err != nil {
		return "", err
	}
	return raeder.Set(name, value)
}

func PGet(name string) string {
	_, err := raeder.Connection(pool.Get())
	if err != nil {
		return ""
	}
	return raeder.Get(name)
}
