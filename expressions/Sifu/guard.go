package Sifu

import (
	"reflect"
)

func canReflect[T any](fieldName string, fieldKind []reflect.Kind) (bool, reflect.StructField) {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return false, reflect.StructField{}
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return false, reflect.StructField{}
	}

	field, ok := typ.FieldByName(fieldName)

	if !ok {
		return false, reflect.StructField{}
	}

	for _, kind := range fieldKind {

		if field.Type.Kind() == kind {
			return true, field

		}
	}

	return false, reflect.StructField{}
}
