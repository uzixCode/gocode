package utils

import (
	"reflect"
	"strings"
)

func StructToMap(i interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(i)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		key := strings.ToLower(t.Field(i).Name)
		m[key] = field.Interface()
	}

	return m
}

type OnLoopType func(*map[string]interface{})

func ListStructToListMap[T interface{}](listStruct []T, onLoop OnLoopType) []map[string]interface{} {
	var returnData []map[string]interface{}
	for _, v := range listStruct {
		mapData := StructToMap(v)
		onLoop(&mapData)
		returnData = append(returnData, mapData)
	}
	return returnData
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
