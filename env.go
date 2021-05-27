package env

import (
	"os"
	"reflect"
	"strings"
)

func Unmarshal(p interface{}) {
	envMap := getEnvVars()
	tfMap := getTagFieldMap(p)
	for key, val := range tfMap {
		field := reflect.ValueOf(p).Elem().FieldByName(val)
		field.Set(reflect.ValueOf(envMap[key]))
	}
}

// 格式化环境变量
func getEnvVars() map[string]string {
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		vars[pair[0]] = pair[1]
	}
	return vars
}

func getTagFieldMap(t interface{}) map[string]string {
	tfMap := make(map[string]string)
	rType := reflect.TypeOf(t)
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
	}
	for i := 0; i < rType.NumField(); i++ {
		f := rType.Field(i)
		tag := f.Tag.Get("env")
		if tag == "" {
			continue
		}
		tfMap[tag] = f.Name
	}
	return tfMap
}
