package redis

import (
	"github.com/gomodule/redigo/redis"
	"go_custom/wrap/config"
	"strconv"
)

var rc config.RedisConfiguration

func init() {
	rc = getConfig()
}

func getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func RConn() (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if rc.Password == "" {
		conn, err = redis.Dial("tcp", rc.Host+":"+strconv.Itoa(rc.Port))
	} else {
		conn, err = redis.Dial("tcp", rc.Host+":"+strconv.Itoa(rc.Port), redis.DialPassword(rc.Password))
	}
	if err != nil {
		panic("redis connection error: " + err.Error())
	}
	return conn, err
}

func rSet(conn redis.Conn, command string, name string, value interface{}, args ...interface{}) (bool, error) {
	defer conn.Close()
	if args != nil {
		return redis.Bool(conn.Do(command, rc.Prefix+name, value, "EX", args[0]))
	} else {
		return redis.Bool(conn.Do(command, rc.Prefix+name, value))
	}
}

func rGet(conn redis.Conn, command string, name string) interface{} {
	defer conn.Close()
	if value, err := conn.Do(command, rc.Prefix+name); err == nil {
		return value
	}
	return nil
}

func SelectDB(conn redis.Conn, num int) error {
	defer conn.Close()
	_, err := redis.String(conn.Do("SELECT", num))
	return err
}

func Set(conn redis.Conn, name, value string, args ...interface{}) (bool, error) {
	return rSet(conn, "SET", name, value, args)
}

func Get(conn redis.Conn, name string) string {
	return rGet(conn, "GET", name).(string)
}

func HSet(conn redis.Conn, name string, value interface{}, args ...interface{}) (bool, error) {
	return rSet(conn, "HSET", name, value, args)
}

func HGet(conn redis.Conn, name string) interface{} {
	return rGet(conn, "HGET", name)
}

func LPush(conn redis.Conn, name string, args ...interface{}) bool {
	defer conn.Close()
	slice := make([]interface{}, 1)
	slice[0] = rc.Prefix + name
	values := append(slice, args...)
	if b, err := redis.Bool(conn.Do("LPUSH", values...)); err == nil {
		return b
	}
	return false
}

func LPop(conn redis.Conn, name string) interface{} {
	return rGet(conn, "LPOP", name)
}

func LRange(conn redis.Conn, name string, start int, stop int) []interface{} {
	defer conn.Close()
	if v, err := redis.Values(conn.Do("LRANGE", rc.Prefix+name, start, stop)); err == nil {
		return v
	}
	return nil
}

func SAdd(conn redis.Conn, name string, args ...interface{}) bool {
	defer conn.Close()
	slice := make([]interface{}, 1)
	slice[0] = rc.Prefix + name
	values := append(slice, args...)
	if b, err := redis.Bool(conn.Do("SADD", values...)); err == nil {
		return b
	}
	return false
}

func SPop(conn redis.Conn, name string) interface{} {
	defer conn.Close()
	if v, err := redis.Values(conn.Do("SPOP", rc.Prefix+name)); err == nil {
		return v
	}
	return nil
}

func ZAdd(conn redis.Conn, name string, key string, value string) bool {
	defer conn.Close()
	if b, err := redis.Bool(conn.Do("ZADD", rc.Prefix+name, key, value)); err == nil {
		return b
	}
	return false
}

func ZRange(conn redis.Conn, name string, start int, stop int, args ...bool) []interface{} {
	defer conn.Close()
	var v []interface{}
	var err error
	if args != nil {
		v, err = redis.Values(conn.Do("ZRANGE", rc.Prefix+name, start, stop, "WITHSCORES"))
	} else {
		v, err = redis.Values(conn.Do("ZRANGE", rc.Prefix+name, start, stop))
	}
	if err != nil {
		return nil
	}
	return v
}

func ZScore(conn redis.Conn, name string, key string) interface{} {
	defer conn.Close()
	if v, err := redis.Values(conn.Do("ZSCORE", rc.Prefix+name, key)); err == nil {
		return v
	}
	return nil
}

func Exists(conn redis.Conn, name string) bool {
	defer conn.Close()
	if exists, err := redis.Bool(conn.Do("EXISTS", name)); err == nil {
		return exists
	}
	return false
}
