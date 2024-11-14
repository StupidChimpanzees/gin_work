package config

import (
	"os"
	"path"
	"runtime"
	"strings"
)

type File struct {
	path     string
	filename string
	ext      string
}

var Info *File

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	Info = &File{path: path.Dir(path.Dir(path.Dir(fileName))), ext: ".xml"}
}

func (File) SetPath(path string) {
	Info.path = path
}

func (File) SetExt(ext string) {
	Info.ext = ext
}

func (File) GetExt() string {
	return Info.ext
}

func (File) formatFilename(file string) {
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

func (File) Load(file string) error {
	Info.formatFilename(file)
	var err error
	if _, err = os.Stat(file); err != nil {
		file = Info.path + string(os.PathSeparator) + Info.filename + Info.ext
		if _, err = os.Stat(file); err != nil {
			return err
		}
	}
	err = Instance.ParamsToConfig(file)
	if err != nil {
		return err
	}
	return nil
}
