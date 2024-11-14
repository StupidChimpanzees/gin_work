package config

import (
	"go_custom/extend/utils"
	"os"
	"path"
	"runtime"
	"strings"
)

type Params map[string]any

type configuration struct {
	Params
	path     string
	filename string
	ext      string
}

var config *configuration

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	config = &configuration{Params: Params{}, path: path.Dir(path.Dir(path.Dir(fileName))), ext: ".xml"}
}

func (configuration) SetPath(path string) {
	config.path = path
}

func (configuration) SetExt(ext string) {
	config.ext = ext
}

func (configuration) GetExt() string {
	return config.ext
}

func (configuration) formatFilename(file string) {
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
		config.filename = filenameArr[len(filenameArr)-2]
		config.SetExt("." + filenameArr[len(filenameArr)-1])
	} else {
		config.filename = filenameArr[len(filenameArr)-1]
	}
}

func Load(file string) error {
	config.formatFilename(file)
	var err error
	if _, err = os.Stat(file); err != nil {
		file = config.path + string(os.PathSeparator) + config.filename + config.ext
		if _, err = os.Stat(file); err != nil {
			return err
		}
	}
	err = Mapping.Parse(file)
	if err != nil {
		return err
	}
	config.Params = Mapping.ParamsToConfig()
	return nil
}

func Set(name string, value any) Params {
	strArr := strings.Split(name, ".")
	configValue := utils.StrArrToMultiMap(strArr, value)
	config.Params = utils.MergeMaps(config.Params, configValue)
	return config.Params
}

func Get(name string) any {
	strArr := strings.Split(name, ".")
	for _, v := range strArr {
		if value, ok := config.Params[v]; !ok {
			return nil
		} else if typeValue, ok := value.(map[string]any); ok {
			config.Params = typeValue
		} else {
			return config.Params[v]
		}
	}
	return nil
}
