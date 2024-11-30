package driver

type Driver interface {
	NewDriver(args ...interface{}) interface{}
	Connection() (interface{}, error)
	Exists(name string) bool
	Set(name, value string, args ...interface{}) (bool, error)
	Get(name string) string
	Del(name string) bool
	Clear() bool
}
