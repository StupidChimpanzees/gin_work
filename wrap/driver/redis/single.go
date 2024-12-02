package redis

import (
	"gin_work/wrap/config"
	"gin_work/wrap/driver"
)

type Single struct {
	readser *driver.Reads
}

var RS *Single

func init() {
	conf := RS.getConfig()
	RS.readser = driver.NewReads(conf.Host, conf.Port, conf.Password, conf.Prefix, conf.Select)
}

func (*Single) getConfig() config.RedisConfiguration {
	return config.Mapping.Redis
}

func (*Single) setFullName(name string) string {
	return RS.readser.Prefix + name
}

func (s *Single) Set(name, value string) (string, error) {
	return RS.readser.Set(s.setFullName(name), value)
}

func (s *Single) Get(name string) (string, error) {
	return RS.readser.Get(s.setFullName(name))
}

func (s *Single) HSet(name string, value interface{}) (string, error) {
	return RS.readser.HSet(s.setFullName(name), value)
}

func (s *Single) HGet(name string) (interface{}, error) {
	return RS.readser.HGet(s.setFullName(name))
}

func (s *Single) LPush(name string, args ...interface{}) (bool, error) {
	return RS.readser.LPush(s.setFullName(name), args...)
}

func (s *Single) LPop(name string) (interface{}, error) {
	return RS.readser.LPop(s.setFullName(name))
}

func (s *Single) LRange(name string, start int, stop int) ([]interface{}, error) {
	return RS.readser.LRange(s.setFullName(name), start, stop)
}

func (s *Single) SAdd(name string, args ...interface{}) (bool, error) {
	return RS.readser.SAdd(s.setFullName(name), args...)
}

func (s *Single) SPop(name string) (interface{}, error) {
	return RS.readser.SPop(s.setFullName(name))
}

func (s *Single) ZAdd(name string, key string, value string) (bool, error) {
	return RS.readser.ZAdd(s.setFullName(name), key, value)
}

func (s *Single) ZRange(name string, start int, stop int, args ...bool) ([]interface{}, error) {
	return RS.readser.ZRange(s.setFullName(name), start, stop, args...)
}

func (s *Single) ZScore(name string, key string) (interface{}, error) {
	return RS.readser.ZScore(s.setFullName(name), key)
}

func (s *Single) Exists(name string) (bool, error) {
	return RS.readser.Exists(s.setFullName(name))
}

func (s *Single) Del(name string) (bool, error) {
	return RS.readser.Del(s.setFullName(name))
}

func (s *Single) Clear() (bool, error) {
	return RS.readser.Clear()
}

func (s *Single) Multi() (bool, error) {
	return RS.readser.Multi()
}

func (s *Single) Exec() (bool, error) {
	return RS.readser.Exec()
}

func (s *Single) Discard() (bool, error) {
	return RS.readser.Discard()
}
