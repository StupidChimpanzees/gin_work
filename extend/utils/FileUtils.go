package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func GetFileType(file any) (string, error) {
	buffer := make([]byte, 512)
	if value, ok := file.(string); ok {
		fileStream, err := os.Open(value)
		if err != nil {
			return "", err
		}
		defer func(fileStream *os.File) {
			err := fileStream.Close()
			if err != nil {
				return
			}
		}(fileStream)
		_, err = fileStream.Read(buffer)
		if err != nil {
			return "", err
		}
	} else if value, ok := file.([]byte); ok {
		buffer = value[:512]
	} else {
		return "", errors.New("parameter type error")
	}

	return http.DetectContentType(buffer), nil
}

func GetSmallFileContent(file string) ([]byte, error) {
	fileStream, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(fileStream *os.File) {
		err := fileStream.Close()
		if err != nil {
			return
		}
	}(fileStream)
	return io.ReadAll(fileStream)
}
