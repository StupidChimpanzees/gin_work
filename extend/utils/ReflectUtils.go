package utils

import (
	"fmt"
	"reflect"
)

func GetParams(object any, objName string, formatType string) any {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	fmt.Println("+++++++++++++++")
	fmt.Println(objType)
	fmt.Println(objValue)

	for i := 0; i < objValue.NumField(); i++ {
		switch objValue.Field(i).Kind() {
		case reflect.Ptr:
			objType = objType.Elem()
			objValue = objValue.Field(i).Elem()
			fallthrough
		case reflect.Struct:
			if objName == objType.Field(i).Tag.Get(formatType) {
				return objValue.Field(i).Interface()
			}
		default:
			if objName == objType.Field(i).Tag.Get(formatType) {
				return objValue.Field(i)
			}
		}
	}

	return nil
}
