package redis

import (
	"gin_work/wrap/config"
	"gin_work/wrap/driver"
	"time"

	"github.com/gomodule/redigo/redis"
)

type poolConf struct {
	enable         bool
	maxIdle        int
	maxActive      int
	idleTimeout    time.Duration
	maxConnTimeout time.Duration
}

var (
	pool     *redis.Pool
	PoolConf *poolConf
	useReads *driver.Reads
)

func init() {
	conf := PoolConf.getPoolConfig()
	// redis pool关闭
	if conf.enable == false {
		return
	}
	rc := PoolConf.getConfig()
	useReads = driver.NewReads(rc.Host, rc.Port, rc.Password, rc.Prefix, rc.Select)
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := useReads.Connection()
			return conn.(redis.Conn), err
		},
		MaxIdle:         conf.maxIdle,
		MaxActive:       conf.maxActive,
		IdleTimeout:     conf.idleTimeout,
		MaxConnLifetime: conf.maxConnTimeout,
		Wait:            true,
	}
}

func (*poolConf) getPoolConfig() *poolConf {
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

func (*poolConf) getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func (*poolConf) setFullName(name string) string {
	return useReads.Prefix + name
}

func (p *poolConf) Set(name, value string) (string, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return "", err
	}
	return useReads.Set(p.setFullName(name), value)
}

func (p *poolConf) Get(name string) (string, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return "", err
	}
	return useReads.Get(p.setFullName(name))
}

func (p *poolConf) HSet(name string, value interface{}) (string, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return "", err
	}
	return useReads.HSet(p.setFullName(name), value)
}

func (p *poolConf) HGet(name string) (interface{}, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return nil, err
	}
	return useReads.HGet(p.setFullName(name))
}

func (p *poolConf) LPush(name string, args ...interface{}) (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.LPush(p.setFullName(name), args...)
}

func (p *poolConf) LPop(name string) (interface{}, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.LPop(p.setFullName(name))
}

func (p *poolConf) LRange(name string, start int, stop int) ([]interface{}, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return nil, err
	}
	return useReads.LRange(p.setFullName(name), start, stop)
}

func (p *poolConf) SAdd(name string, args ...interface{}) (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.SAdd(p.setFullName(name), args...)
}

func (p *poolConf) SPop(name string) (interface{}, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.SPop(p.setFullName(name))
}

func (p *poolConf) ZAdd(name string, key string, value string) (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.ZAdd(p.setFullName(name), key, value)
}

func (p *poolConf) ZRange(name string, start int, stop int, args ...bool) ([]interface{}, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return nil, err
	}
	return useReads.ZRange(p.setFullName(name), start, stop, args...)
}

func (p *poolConf) ZScore(name string, key string) (interface{}, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return nil, err
	}
	return useReads.ZScore(p.setFullName(name), key)
}

func (p *poolConf) Exists(name string) (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.Exists(p.setFullName(name))
}

func (p *poolConf) Del(name string) (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.Del(p.setFullName(name))
}

func (p *poolConf) Clear() (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.Clear()
}

func (p *poolConf) Multi() (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.Multi()
}

func (p *poolConf) Exec() (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.Exec()
}

func (p *poolConf) Discard() (bool, error) {
	_, err := useReads.Connection(pool.Get())
	if err != nil {
		return false, err
	}
	return useReads.Discard()
}
