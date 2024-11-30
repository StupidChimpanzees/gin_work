package cache

import (
	"encoding/json"
	"gin_work/wrap/config"
	raede "gin_work/wrap/driver/redis"
)

type cache struct {
	CType    string
	Host     string
	Port     int
	Password string
	Prefix   string
	Timeout  int
}

var (
	rCache *cache
	device interface{}
)

func init() {
	rCache = getConfig()
	if rCache.CType == "redis" {
		device = raede.NewRaede(rCache.Host, rCache.Port, rCache.Password, rCache.Prefix, 0)
	}
}

func getConfig() *cache {
	c := config.Mapping.Cache
	return &cache{
		CType:    c.CType,
		Host:     c.Host,
		Port:     c.Port,
		Password: c.Password,
		Prefix:   c.Prefix,
		Timeout:  c.Timeout,
	}
}

func Has(name string) bool {
	return device.(*raede.Raede).Exists(rCache.Prefix + name)
}

func Get(name string) any {
	str := device.(*raede.Raede).Get(rCache.Prefix + name)
	if str == "" {
		return nil
	}
	b := []byte(str)
	var instance any
	err := json.Unmarshal(b, &instance)
	if err != nil {
		return nil
	}
	return instance
}

func Set(name string, value interface{}, args ...int) bool {
	timeout := rCache.Timeout
	if args != nil {
		timeout = args[0]
	}
	b, err := json.Marshal(&value)
	if timeout == 0 {
		_, err = device.(*raede.Raede).Set(rCache.Prefix+name, string(b))
	} else {
		_, err = device.(*raede.Raede).Set(rCache.Prefix+name, string(b), rCache.Timeout)
	}
	if err != nil {
		return false
	}
	return true
}

func Delete(name string) bool {
	return device.(*raede.Raede).Del(rCache.Prefix + name)
}

func Clear() bool {
	return device.(*raede.Raede).Clear()
}
