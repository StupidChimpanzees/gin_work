package cache

import (
	"encoding/json"
	"go_custom/wrap/config"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type cache interface {
	has(name string)
	get(name string, value interface{}) bool
	set(name string, args ...interface{}) bool
	delete(name string) bool
	clear() bool
}

type Redis struct {
	cache
	rc   config.CacheConfiguration
	conn redis.Conn
}

var rCache Redis

func (Redis) getConfig() {
	rCache.rc = config.Mapping.Cache
}

func Load() (err error) {
	rCache.conn, err = redis.Dial("tcp", rCache.rc.Host+":"+strconv.Itoa(rCache.rc.Port))
	return err
}

func (Redis) Get(key string, value *interface{}) bool {
	str, err := redis.String(rCache.conn.Do("GET", key))
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

func (Redis) Set(key string, args ...interface{}) bool {
	if args == nil {
		return false
	}
	if len(args) == 2 {
		rCache.conn.Do("EXPIRE", key, args[1])
	}
	b, err := json.Marshal(&args[1])
	_, err = rCache.conn.Do("SET", key, string(b))
	if err != nil {
		return false
	}
	return true
}
