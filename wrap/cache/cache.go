package cache

import (
	"encoding/json"
	"go_custom/wrap/config"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	rc   config.CacheConfiguration
	conn redis.Conn
}

var rCache Redis

func (Redis) getConfig() {
	rCache.rc = config.Mapping.Cache
}

func (Redis) connRedis() error {
	var err error
	if rCache.rc.Password != "" {
		rCache.conn, err = redis.Dial("tcp", rCache.rc.Host+":"+strconv.Itoa(rCache.rc.Port), redis.DialPassword(rCache.rc.Password))
	} else {
		rCache.conn, err = redis.Dial("tcp", rCache.rc.Host+":"+strconv.Itoa(rCache.rc.Port))
	}
	return err
}

func Load() (err error) {
	rCache.getConfig()
	err = rCache.connRedis()

	return err
}

func Has(name string) bool {
	exists, err := redis.Bool(rCache.conn.Do("EXISTS", rCache.rc.Prefix+name))
	if err != nil {
		return false
	}
	return exists
}

func Get(name string, value *interface{}) bool {
	str, err := redis.String(rCache.conn.Do("GET", rCache.rc.Prefix+name))
	if err != nil {
		return false
	}
	b := []byte(str)
	instance := *value
	err = json.Unmarshal(b, instance)
	if err != nil {
		return false
	}
	return true
}

func Set(name string, args ...interface{}) bool {
	if args == nil {
		return false
	}
	var err error
	if len(args) == 2 {
		_, err = rCache.conn.Do("EXPIRE", rCache.rc.Prefix+name, args[1])
	} else {
		_, err = rCache.conn.Do("EXPIRE", rCache.rc.Prefix+name, rCache.rc.Timeout)
	}
	if err != nil {
		return false
	}
	var b []byte
	b, err = json.Marshal(&args[1])
	_, err = rCache.conn.Do("SET", rCache.rc.Prefix+name, string(b))
	if err != nil {
		return false
	}
	return true
}

func Delete(name string) bool {
	b, err := redis.Bool(rCache.conn.Do("DEL", rCache.rc.Prefix+name))
	if err != nil {
		return false
	}
	return b
}

func Clear() bool {
	b, err := redis.Bool(rCache.conn.Do("FLUSHDB"))
	if err != nil {
		return false
	}
	return b
}
