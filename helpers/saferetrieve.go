package helpers

import (
	"reflect"
	"strings"

	"byvko.dev/repo/am-stats-dataprep-api/logs"
)

func SafeRetrieve(value interface{}, path string) interface{} {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("SafeRetrieve: path: %v, value: %v, err: %v", path, value, r)
			value = nil
		}
	}()

	paths := strings.Split(path, ".")
	for _, p := range paths {
		if value == nil {
			return nil
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			value = reflect.ValueOf(value).MapIndex(reflect.ValueOf(p)).Interface()
		case reflect.Slice:
			value = reflect.ValueOf(value).Index(0).Interface()
		case reflect.Struct:
			value = reflect.ValueOf(value).FieldByName(p).Interface()
		default:
			return nil
		}
	}

	return value
}
