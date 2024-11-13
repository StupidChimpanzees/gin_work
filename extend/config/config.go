package config

import (
	"maps"
	"os"
	"path"
	"runtime"
	"strings"
	"workspace/src/framework/extend/utils"
)

type Params map[string]any

type Config struct {
	Params
	path     string
	filename string
	ext      string
}

var Info *Config

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	Info = &Config{Params: Params{}, path: path.Dir(path.Dir(path.Dir(fileName))), ext: ".xml"}
}

func (Config) SetPath(path string) {
	Info.path = path
}

func (Config) SetExt(ext string) {
	Info.ext = ext
}

func (Config) GetExt() string {
	return Info.ext
}

func (Config) formatFilename(file string) {
	strings.TrimSpace(file)
	pathArr := strings.Split(file, string(os.PathSeparator))
	var filename string
	if len(pathArr) > 1 {
		filename = pathArr[len(pathArr)-1]
	} else {
		filename = file
	}
	filenameArr := strings.Split(filename, ".")
	if len(filenameArr) > 1 {
		Info.filename = filenameArr[len(filenameArr)-2]
		Info.SetExt("." + filenameArr[len(filenameArr)-1])
	} else {
		Info.filename = filenameArr[len(filenameArr)-1]
	}
}

func (Config) Load(file string) error {
	Info.formatFilename(file)
	var err error
	if _, err = os.Stat(file); err != nil {
		file = Info.path + string(os.PathSeparator) + Info.filename + Info.ext
		if _, err = os.Stat(file); err != nil {
			return err
		}
	}
	err = configInstance.Parse(file)
	if err != nil {
		return err
	}
	Info.Params = configInstance.ParamsToConfig()
	return nil
}

func (Config) Set(name string, value any) Params {
	strArr := strings.Split(name, ".")
	configValue := utils.StrArrToMultiMap(strArr, value)
	maps.Copy(Info.Params, configValue)
	return Info.Params
}

func (Config) Get(name string) any {
	strArr := strings.Split(name, ".")
	for _, v := range strArr {
		if value, ok := Info.Params[v]; !ok {
			return nil
		} else if typeValue, ok := value.(map[string]any); ok {
			Info.Params = typeValue
		} else {
			return Info.Params[v]
		}
	}
	return nil
}
