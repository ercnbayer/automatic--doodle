package utils

import (
	"reflect"
)

func IsJson(t reflect.Type) bool {
	if !IsStruct(t) {
		return false
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if _, ok := f.Tag.Lookup("json"); ok {
			return true
		}
	}

	return false
}
