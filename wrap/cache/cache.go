package cache

import (
	"encoding/json"
	"gin_work/wrap/config"
	raede "gin_work/wrap/redis"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	rc   config.CacheConfiguration
	conn redis.Conn
}

var (
	rCache Redis
	r      raede.Raede
)

func init() {
	rCache.getConfig()
	rCache.connRedis()
	r.SelectDB(rCache.conn, 0)
}

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

func Has(name string) bool {
	rCache.connRedis()
	return r.Exists(rCache.conn, rCache.rc.Prefix+name)
}

func Get(name string, value *interface{}) bool {
	rCache.connRedis()
	str := r.Get(rCache.conn, rCache.rc.Prefix+name)
	if str == "" {
		return false
	}
	b := []byte(str)
	instance := *value
	err := json.Unmarshal(b, instance)
	if err != nil {
		return false
	}
	return true
}

func Set(name string, value interface{}, args ...int) bool {
	var timeout int = rCache.rc.Timeout
	if args != nil {
		timeout = args[0]
	}
	b, err := json.Marshal(&value)
	rCache.connRedis()
	if timeout == 0 {
		_, err = r.Set(rCache.conn, rCache.rc.Prefix+name, string(b))
	} else {
		_, err = r.Set(rCache.conn, rCache.rc.Prefix+name, string(b), rCache.rc.Timeout)
	}
	if err != nil {
		return false
	}
	return true
}

func Delete(name string) bool {
	rCache.connRedis()
	return r.Del(rCache.conn, rCache.rc.Prefix+name)
}

func Clear() bool {
	rCache.connRedis()
	r.FlushDB(rCache.conn)
}
