package cache

import (
	"encoding/json"
	"gin_work/wrap/config"
	driver2 "gin_work/wrap/driver"
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
	Cache  *cache
	driver driver2.Driver
)

func init() {
	Cache = getConfig()
	if Cache.CType == "redis" {
		driver = driver2.NewReads(Cache.Host, Cache.Port, Cache.Password, Cache.Prefix, 0)
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

func Has(name string) (bool, error) {
	return driver.Exists(Cache.Prefix + name)
}

func BindGet(name string, value interface{}) error {
	str, err := driver.Get(Cache.Prefix + name)
	if err != nil || str == "" {
		return err
	}
	b := []byte(str)
	err = json.Unmarshal(b, &value)
	return err
}

func AnyGet(name string) (interface{}, error) {
	str, err := driver.Get(Cache.Prefix + name)
	if err != nil || str == "" {
		return nil, err
	}
	b := []byte(str)
	var instance interface{}
	err = json.Unmarshal(b, &instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func Set(name string, value interface{}, args ...int) (bool, error) {
	timeout := Cache.Timeout
	if args != nil {
		timeout = args[0]
	}
	b, err := json.Marshal(&value)
	if timeout == 0 {
		_, err = driver.Set(Cache.Prefix+name, string(b))
	} else {
		_, err = driver.Set(Cache.Prefix+name, string(b), timeout)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func Del(name string) (bool, error) {
	return driver.Del(Cache.Prefix + name)
}

func Clear() (bool, error) {
	return driver.Clear()
}
