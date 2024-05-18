package utils

import (
	"go-fiber-clean-architecture/application/utils/helper2"
	"log"
	"reflect"
)

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func GetApplicationPath() string {
	return helper2.GetRootPath() + "/application"
}

func GetStoragePath() string {
	return GetApplicationPath() + "/storage"
}

func GetStoragePublicPath() string {
	return GetStoragePath() + "/public"
}

func GetStoragePrivatePath() string {
	return GetStoragePath() + "/private"
}

func Recover()  {
	if r := recover(); r != nil {
		log.Print("Recovered from ", r)
	}
}