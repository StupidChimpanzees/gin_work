package utils

import (
	"reflect"
)

func GetParams(object any, formatType string) map[string]any {
	var configMap = make(map[string]any)
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)

	for i := 0; i < objValue.NumField(); i++ {
		switch objValue.Field(i).Kind() {
		case reflect.Ptr:
			objType = objType.Elem()
			objValue = objValue.Elem()
			fallthrough
		case reflect.Struct:
			configMap[objType.Field(i).Tag.Get(formatType)] = GetParams(objValue.Field(i).Interface(), formatType)
		default:
			configMap[objType.Field(i).Tag.Get(formatType)] = objValue.Field(i)
		}
	}

	return configMap
}
