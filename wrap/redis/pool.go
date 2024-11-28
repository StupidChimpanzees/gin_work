package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	maxIdle := 10
	maxActive := 100
	idleTimeout := time.Duration(3)
	if r.rc.MaxIdle != 0 {
		maxIdle = r.rc.MaxIdle
	}
	if r.rc.MaxActive != 0 {
		maxActive = r.rc.MaxActive
	}
	if r.rc.IdleTimeout != 0 {
		idleTimeout = time.Duration(r.rc.IdleTimeout)
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
