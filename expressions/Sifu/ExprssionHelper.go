package Sifu

import "reflect"

type SifuExpr[T any] struct {
	fieldName string
}

func OfType[T any]() *SifuExpr[T] {
	return &SifuExpr[T]{}
}

func (curr *SifuExpr[T]) IsProp(name string) *SifuExpr[T] {
	curr.fieldName = name
	return curr
}

func (curr *SifuExpr[T]) EqualToString(value string) func(T) bool {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return func(T) bool { return false }
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return func(T) bool { return false }
	}

	field, ok := typ.FieldByName(curr.fieldName)
	if !ok || field.Type.Kind() != reflect.String {
		return func(T) bool { return false }
	}

	index := field.Index

	return func(item T) bool {
		v := reflect.ValueOf(item)

		if !v.IsValid() {
			return false
		}

		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return false
			}
			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return false
		}

		f := v.FieldByIndex(index)

		if f.Kind() != reflect.String {
			return false
		}

		return f.String() == value
	}
}

func (curr *SifuExpr[T]) Bool() func(T) bool {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return func(T) bool { return false }
	}

	field, ok := typ.FieldByName(curr.fieldName)
	if !ok || field.Type.Kind() != reflect.Bool {
		return func(T) bool { return false }
	}

	index := field.Index

	return func(item T) bool {
		v := reflect.ValueOf(item)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return false
			}
			v = v.Elem()
		}

		f := v.FieldByIndex(index)
		return f.Bool()
	}
}

func Not[T any](pred func(T) bool) func(T) bool {
	return func(item T) bool { return !pred(item) }
}

func And[T any](preds ...func(T) bool) func(T) bool {

	return func(item T) bool {
		for _, pred := range preds {
			if !pred(item) {
				return false
			}
		}
		return true
	}

}

func Or[T any](preds ...func(T) bool) func(T) bool {

	return func(item T) bool {
		for _, pred := range preds {
			if pred(item) {
				return true
			}
		}
		return false
	}

}
