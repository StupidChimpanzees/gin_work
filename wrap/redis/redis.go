package redis

import (
	"gin_work/wrap/config"

	"github.com/gomodule/redigo/redis"
)

type Raede struct {
	rc   config.RedisConfiguration
	conn redis.Conn
}

func (Raede) rSet(conn redis.Conn, command string, name string, value interface{}, args ...interface{}) (bool, error) {
	defer conn.Close()
	if args != nil {
		return redis.Bool(conn.Do(command, name, value, "EX", args[0]))
	} else {
		return redis.Bool(conn.Do(command, name, value))
	}
}

func (Raede) rGet(conn redis.Conn, command string, name string) interface{} {
	defer conn.Close()
	if value, err := conn.Do(command, name); err == nil {
		return value
	}
	return nil
}

func (Raede) SelectDB(conn redis.Conn, num int) error {
	_, err := redis.String(conn.Do("SELECT", num))
	return err
}

func (Raede) Set(conn redis.Conn, name, value string, args ...interface{}) (bool, error) {
	return r.rSet(conn, "SET", name, value, args...)
}

func (Raede) Get(conn redis.Conn, name string) string {
	return r.rGet(conn, "GET", name).(string)
}

func (Raede) HSet(conn redis.Conn, name string, value interface{}, args ...interface{}) (bool, error) {
	return r.rSet(conn, "HSET", name, value, args...)
}

func (Raede) HGet(conn redis.Conn, name string) interface{} {
	return r.rGet(conn, "HGET", name)
}

func (Raede) LPush(conn redis.Conn, name string, args ...interface{}) bool {
	defer conn.Close()
	slice := make([]interface{}, 1)
	slice[0] = name
	values := append(slice, args...)
	if b, err := redis.Bool(conn.Do("LPUSH", values...)); err == nil {
		return b
	}
	return false
}

func (Raede) LPop(conn redis.Conn, name string) interface{} {
	return r.rGet(conn, "LPOP", name)
}

func (Raede) LRange(conn redis.Conn, name string, start int, stop int) []interface{} {
	defer conn.Close()
	if v, err := redis.Values(conn.Do("LRANGE", name, start, stop)); err == nil {
		return v
	}
	return nil
}

func (Raede) SAdd(conn redis.Conn, name string, args ...interface{}) bool {
	defer conn.Close()
	slice := make([]interface{}, 1)
	slice[0] = name
	values := append(slice, args...)
	if b, err := redis.Bool(conn.Do("SADD", values...)); err == nil {
		return b
	}
	return false
}

func (Raede) SPop(conn redis.Conn, name string) interface{} {
	defer conn.Close()
	if v, err := redis.Values(conn.Do("SPOP", name)); err == nil {
		return v
	}
	return nil
}

func (Raede) ZAdd(conn redis.Conn, name string, key string, value string) bool {
	defer conn.Close()
	if b, err := redis.Bool(conn.Do("ZADD", name, key, value)); err == nil {
		return b
	}
	return false
}

func (Raede) ZRange(conn redis.Conn, name string, start int, stop int, args ...bool) []interface{} {
	defer conn.Close()
	var v []interface{}
	var err error
	if args != nil && args[0] == true {
		v, err = redis.Values(conn.Do("ZRANGE", name, start, stop, "WITHSCORES"))
	} else {
		v, err = redis.Values(conn.Do("ZRANGE", name, start, stop))
	}
	if err != nil {
		return nil
	}
	return v
}

func (Raede) ZScore(conn redis.Conn, name string, key string) interface{} {
	defer conn.Close()
	if v, err := redis.Values(conn.Do("ZSCORE", name, key)); err == nil {
		return v
	}
	return nil
}

func (Raede) Exists(conn redis.Conn, name string) bool {
	defer conn.Close()
	if exists, err := redis.Bool(conn.Do("EXISTS", name)); err == nil {
		return exists
	}
	return false
}

func (Raede) Del(conn redis.Conn, name string) bool {
	defer conn.Close()
	if b, err := redis.Bool(conn.Do("DEL", name)); err == nil {
		return b
	}
	return false
}

func (Raede) FlushDB(conn redis.Conn) bool {
	defer conn.Close()
	b, err := redis.Bool(conn.Do("FLUSHDB"))
	if err != nil {
		return false
	}
	return b
}
