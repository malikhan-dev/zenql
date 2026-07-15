package Sifu

import (
	"reflect"
)

func canReflect[T any](fieldName string, fieldKind []reflect.Kind) (bool, []int) {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return false, []int{0}
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return false, []int{0}
	}

	field, ok := typ.FieldByName(fieldName)

	if !ok {
		return false, []int{0}
	}

	for _, kind := range fieldKind {

		if field.Type.Kind() == kind {
			return true, field.Index

		}
	}

	return false, []int{0}
}
