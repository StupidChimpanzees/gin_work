package redis

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type Raede struct {
	Host     string
	Port     int
	Password string
	Prefix   string
	Select   int
	Conn     redis.Conn
}

var raeder *Raede

func NewRaede(host string, port int, password string, prefix string, libName int) *Raede {
	return &Raede{
		Host:     host,
		Port:     port,
		Password: password,
		Prefix:   prefix,
		Select:   libName,
	}
}

func (r *Raede) Connection(conn ...redis.Conn) (interface{}, error) {
	if conn != nil {
		return conn, nil
	}
	var err error
	if raeder.Password == "" {
		raeder.Conn, err = redis.Dial("tcp", raeder.Host+":"+strconv.Itoa(raeder.Port))
	} else {
		raeder.Conn, err = redis.Dial("tcp", raeder.Host+":"+strconv.Itoa(raeder.Port), redis.DialPassword(raeder.Password))
	}
	if err != nil {
		panic("redis connection error: " + err.Error())
	}
	return raeder.Conn, err
}

func (r *Raede) getConnection() redis.Conn {
	if raeder.Conn != nil {
		return raeder.Conn
	}
	conn, _ := r.Connection()
	raeder.Conn = conn.(redis.Conn)
	return raeder.Conn
}

func (r *Raede) Command(command string, args ...interface{}) (interface{}, error) {
	conn := r.getConnection()
	defer r.Close(conn)
	r.SelectDB(raeder.Select)
	return conn.Do(command, args...)
}

func (r *Raede) rSet(command string, name string, value interface{}, args ...interface{}) (string, error) {
	conn := r.getConnection()
	defer r.Close(conn)
	r.SelectDB(raeder.Select)
	if args != nil && args[0] != 0 {
		return redis.String(conn.Do(command, name, value, "EX", args[0]))
	} else {
		return redis.String(conn.Do(command, name, value))
	}
}

func (r *Raede) SelectDB(num int) error {
	conn := r.getConnection()
	_, err := redis.String(conn.Do("SELECT", num))
	return err
}

func (r *Raede) Set(name, value string, args ...interface{}) (string, error) {
	return r.rSet("SET", name, value, args...)
}

func (r *Raede) Get(name string) string {
	value, _ := redis.String(r.Command("GET", name))
	return value
}

func (r *Raede) HSet(name string, value interface{}, args ...interface{}) (string, error) {
	return r.rSet("HSET", name, value, args...)
}

func (r *Raede) HGet(name string) interface{} {
	value, _ := r.Command("HGET", name)
	return value
}

func (r *Raede) LPush(name string, args ...interface{}) bool {
	slice := make([]interface{}, 1)
	slice[0] = name
	values := append(slice, args...)
	if b, err := redis.Bool(r.Command("LPUSH", values...)); err == nil {
		return b
	}
	return false
}

func (r *Raede) LPop(name string) interface{} {
	value, _ := r.Command("LPOP", name)
	return value
}

func (r *Raede) LRange(name string, start int, stop int) []interface{} {
	if v, err := redis.Values(r.Command("LRANGE", name, start, stop)); err == nil {
		return v
	}
	return nil
}

func (r *Raede) SAdd(name string, args ...interface{}) bool {
	slice := make([]interface{}, 1)
	slice[0] = name
	values := append(slice, args...)
	if b, err := redis.Bool(r.Command("SADD", values...)); err == nil {
		return b
	}
	return false
}

func (r *Raede) SPop(name string) interface{} {
	if v, err := redis.Values(r.Command("SPOP", name)); err == nil {
		return v
	}
	return nil
}

func (r *Raede) ZAdd(name string, key string, value string) bool {
	if b, err := redis.Bool(r.Command("ZADD", name, key, value)); err == nil {
		return b
	}
	return false
}

func (r *Raede) ZRange(name string, start int, stop int, args ...bool) []interface{} {
	var v []interface{}
	var err error
	if args != nil && args[0] == true {
		v, err = redis.Values(r.Command("ZRANGE", name, start, stop, "WITHSCORES"))
	} else {
		v, err = redis.Values(r.Command("ZRANGE", name, start, stop))
	}
	if err != nil {
		return nil
	}
	return v
}

func (r *Raede) ZScore(name string, key string) interface{} {
	if v, err := redis.Values(r.Command("ZSCORE", name, key)); err == nil {
		return v
	}
	return nil
}

func (r *Raede) Exists(name string) bool {
	if exists, err := redis.Bool(r.Command("EXISTS", name)); err == nil {
		return exists
	}
	return false
}

func (r *Raede) Del(name string) bool {
	if b, err := redis.Bool(r.Command("DEL", name)); err == nil {
		return b
	}
	return false
}

func (r *Raede) Clear() bool {
	b, err := redis.Bool(r.Command("FLUSHDB"))
	if err != nil {
		return false
	}
	return b
}

func (r *Raede) Multi() bool {
	conn := r.getConnection()
	b, err := redis.Bool(conn.Do("MULTI"))
	if err != nil {
		return false
	}
	return b
}

func (r *Raede) Exec() bool {
	conn := r.getConnection()
	b, err := redis.Bool(conn.Do("EXEC"))
	if err != nil {
		return false
	}
	return b
}

func (r *Raede) Discard() bool {
	conn := r.getConnection()
	b, err := redis.Bool(conn.Do("DISCARD"))
	if err != nil {
		return false
	}
	return b
}

func (*Raede) Close(conn redis.Conn) {
	raeder.Conn = nil
	err := conn.Close()
	if err != nil {
		// Log
		return
	}
}
