package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func init() {
	maxIdle := 10
	maxActive := 100
	var idleTimeout time.Duration = 0
	if rc.MaxIdle != 0 {
		maxIdle = rc.MaxIdle
	}
	if rc.MaxActive != 0 {
		maxActive = rc.MaxActive
	}
	if rc.IdleTimeout != 0 {
		idleTimeout = time.Duration(rc.IdleTimeout)
	}
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return RConn()
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Wait:        true,
	}
}
