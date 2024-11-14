package config

import (
	"encoding/json"
	"encoding/xml"
	"go_custom/extend/utils"

	"gopkg.in/yaml.v3"
)

type configMapping struct {
	App      AppConfiguration      `yaml:"app" bson:"app" json:"app" xml:"app"`
	Database DatabaseConfiguration `yaml:"database" bson:"database" json:"database" xml:"database"`
	Cache    CacheConfiguration    `yaml:"cache" bson:"cache" json:"cache" xml:"cache"`
	Cookie   CookieConfiguration   `yaml:"cookie" bson:"cookie" json:"cookie" xml:"cookie"`
	Session  SessionConfiguration  `yaml:"session" bson:"session" json:"session" xml:"session"`
	View     ViewConfiguration     `yaml:"view" bson:"view" json:"view" xml:"view"`
}

type AppConfiguration struct {
	Name    string `yaml:"name" bson:"name" json:"name" xml:"name"`
	Version string `yaml:"version" bson:"version" json:"version" xml:"version"`
	Port    int    `yaml:"port" bson:"port" json:"port" xml:"port"`
}

type DatabaseConfiguration struct {
	Host     string `yaml:"host" bson:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" bson:"port" json:"port" xml:"port"`
	Username string `yaml:"username" bson:"username" json:"username" xml:"username"`
	Password string `yaml:"password" bson:"password" json:"password" xml:"password"`
}

type CacheConfiguration struct {
	CType    string `yaml:"cache_type" bson:"cache_type" json:"cache_type" xml:"cache_type"`
	Host     string `yaml:"host" bson:"host" json:"host" xml:"host"`
	Port     int    `yaml:"port" bson:"port" json:"port" xml:"port"`
	Password string `yaml:"password" bson:"password" json:"password" xml:"password"`
	Prefix   string `yaml:"prefix" bson:"prefix" json:"prefix" xml:"prefix"`
	Timeout  int    `yaml:"timeout" bson:"timeout" json:"timeout" xml:"timeout"`
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

type ViewConfiguration struct {
	TempPath   string `yaml:"temp_path" bson:"temp_path" json:"temp_path" xml:"temp_path"`
	StaticPath string `yaml:"static_path" bson:"static_path" json:"static_path" xml:"static_path"`
	DelimBegin string `yaml:"delim_begin" bson:"delim_begin" json:"delim_begin" xml:"delim_begin"`
	DelimEnd   string `yaml:"delim_end" bson:"delim_end" json:"delim_end" xml:"delim_end"`
}

var Mapping configMapping

func init() {
	Mapping = configMapping{}
}

func (configMapping) Parse(file string) error {
	fileContent, _ := utils.GetSmallFileContent(file)
	var err error = nil
	switch config.GetExt() {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(fileContent, &Mapping)
	case ".json":
		err = json.Unmarshal(fileContent, &Mapping)
	case ".bson":
	// 空着先
	case ".xml":
		err = xml.Unmarshal(fileContent, &Mapping)
	}
	if err != nil {
		return err
	}

	return nil
}

func (configMapping) ParamsToConfig() map[string]any {
	return utils.GetParams(Mapping, "yaml")
}
