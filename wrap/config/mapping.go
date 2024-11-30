package config

import (
	"encoding/json"
	"encoding/xml"
	"gin_work/wrap/utils"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type configMapping struct {
	App      appConfiguration      `yaml:"app" bson:"app" json:"app" xml:"app"`
	Database databaseConfiguration `yaml:"database" bson:"database" json:"database" xml:"database"`
	Cache    CacheConfiguration    `yaml:"cache" bson:"cache" json:"cache" xml:"cache"`
	Redis    RedisConfiguration    `yaml:"redis" bson:"redis" json:"redis" xml:"redis"`
	Cookie   CookieConfiguration   `yaml:"cookie" bson:"cookie" json:"cookie" xml:"cookie"`
	Session  SessionConfiguration  `yaml:"session" bson:"session" json:"session" xml:"session"`
	View     viewConfiguration     `yaml:"view" bson:"view" json:"view" xml:"view"`
}

type appConfiguration struct {
	Name    string `yaml:"name" bson:"name" json:"name" xml:"name"`
	Version string `yaml:"version" bson:"version" json:"version" xml:"version"`
	Port    int    `yaml:"port" bson:"port" json:"port" xml:"port"`
}

type databaseConfiguration struct {
	DBType   string `yaml:"db_type" bson:"db_type" json:"db_type" xml:"db_type"`
	Host     string `yaml:"host" bson:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" bson:"port" json:"port" xml:"port"`
	Username string `yaml:"username" bson:"username" json:"username" xml:"username"`
	Password string `yaml:"password" bson:"password" json:"password" xml:"password"`
	Name     string `yaml:"name" bson:"name" json:"name" xml:"name"`
	Charset  string `yaml:"charset" bson:"charset" json:"charset" xml:"charset"`
}

type CacheConfiguration struct {
	CType    string `yaml:"cache_type" bson:"cache_type" json:"cache_type" xml:"cache_type"`
	Host     string `yaml:"host" bson:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" bson:"port" json:"port" xml:"port"`
	Password string `yaml:"password" bson:"password" json:"password" xml:"password"`
	Prefix   string `yaml:"prefix" bson:"prefix" json:"prefix" xml:"prefix"`
	Timeout  int    `yaml:"timeout" bson:"timeout" json:"timeout" xml:"timeout"`
}

type RedisConfiguration struct {
	Host     string `yaml:"host" bson:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" bson:"port" json:"port" xml:"port"`
	Password string `yaml:"password" bson:"password" json:"password" xml:"password"`
	Prefix   string `yaml:"prefix" bson:"prefix" json:"prefix" xml:"prefix"`
	Select   int    `yaml:"select" bson:"select" json:"select" xml:"select"`
	Pool     RedisPoolConfiguration
}

type RedisPoolConfiguration struct {
	MaxIdle        int `yaml:"max_idle" bson:"max_idle" json:"max_idle" xml:"max_idle"`
	MaxActive      int `yaml:"max_active" bson:"max_active" json:"max_active" xml:"max_active"`
	IdleTimeout    int `yaml:"idle_timeout" bson:"idle_timeout" json:"idle_timeout" xml:"idle_timeout"`
	MaxConnTimeout int `yaml:"max_conn_timeout" bson:"max_conn_timeout" json:"max_conn_timeout" xml:"max_conn_timeout"`
}

type CookieConfiguration struct {
	Expire   int    `yaml:"expire" bson:"expire" json:"expire" xml:"expire"`
	Path     string `yaml:"path" bson:"path" json:"path" xml:"path"`
	Domain   string `yaml:"domain" bson:"domain" json:"domain" xml:"domain"`
	Secure   bool   `yaml:"secure" bson:"secure" json:"secure" xml:"secure"`
	HttpOnly bool   `yaml:"http_only" bson:"http_only" json:"http_only" xml:"http_only"`
}

type SessionConfiguration struct {
	Secret      string `yaml:"secret" bson:"secret" json:"secret" xml:"secret"`
	Expire      int    `yaml:"expire" bson:"expire" json:"expire" xml:"expire"`
	SessionName string `yaml:"session_name" bson:"session_name" json:"session_name" xml:"session_name"`
}

type viewConfiguration struct {
	TempPath   string `yaml:"temp_path" bson:"temp_path" json:"temp_path" xml:"temp_path"`
	StaticPath string `yaml:"static_path" bson:"static_path" json:"static_path" xml:"static_path"`
	DelimBegin string `yaml:"delim_begin" bson:"delim_begin" json:"delim_begin" xml:"delim_begin"`
	DelimEnd   string `yaml:"delim_end" bson:"delim_end" json:"delim_end" xml:"delim_end"`
}

type configurationFile struct {
	path     string
	filename string
	ext      string
}

var configFile configurationFile

var Mapping configMapping

func init() {
	Load("config.yaml")
}

func (configurationFile) formatFilename(file string) {
	file = strings.TrimSpace(file)
	pathArr := strings.Split(file, string(os.PathSeparator))
	var filename string
	if len(pathArr) > 1 {
		filename = pathArr[len(pathArr)-1]
	} else {
		filename = file
	}
	filenameArr := strings.Split(filename, ".")
	if len(filenameArr) > 1 {
		configFile.filename = filenameArr[len(filenameArr)-2]
		configFile.ext = filenameArr[len(filenameArr)-1]
	} else {
		configFile.filename = filenameArr[len(filenameArr)-1]
	}
}

func (configMapping) Parse(file string) error {
	fileContent, _ := utils.GetSmallFileContent(file)
	var err error = nil
	switch configFile.ext {
	case "yaml", "yml":
		err = yaml.Unmarshal(fileContent, &Mapping)
	case "json":
		err = json.Unmarshal(fileContent, &Mapping)
	case "bson":
	// 空着先
	case "xml":
		err = xml.Unmarshal(fileContent, &Mapping)
	}
	return err
}

func (configMapping) ParamsToConfig() map[string]any {
	return utils.GetParams(Mapping, "yaml")
}

func Load(file string) error {
	configFile.formatFilename(file)
	var err error
	if _, err = os.Stat(file); err != nil {
		file = configFile.path + string(os.PathSeparator) + configFile.filename + "." + configFile.ext
		if _, err = os.Stat(file); err != nil {
			return err
		}
	}
	return Mapping.Parse(file)
}
