package Sifu

import (
	"reflect"
)

type MutableExpression[T any] struct {
	Result func(item T) T
}

func (curr *PropExpression[T]) SetString(value string) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, reflect.String); success {

		index := fieldIndex

		fnc := func(item T) T {

			v := reflect.ValueOf(&item).Elem()

			if !v.IsValid() {
				return item
			}

			f := v.FieldByIndex(index)

			if !f.CanSet() {
				return item
			}

			f.SetString(value)

			return item
		}
		return MutableExpression[T]{Result: fnc}
	} else {
		return MutableExpression[T]{Result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetBool(value bool) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, reflect.Bool); success {

		index := fieldIndex

		fnc := func(item T) T {

			v := reflect.ValueOf(&item).Elem()

			if !v.IsValid() {
				return item
			}

			f := v.FieldByIndex(index)

			if f.Kind() != reflect.Bool {
				return item
			}

			if !f.CanSet() {
				return item
			}

			f.SetBool(value)

			return item
		}
		return MutableExpression[T]{Result: fnc}
	} else {
		return MutableExpression[T]{Result: func(item T) T { return item }}

	}

}

func (curr *PropExpression[T]) StrApp(value string) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, reflect.String); success {
		index := fieldIndex

		fnc := func(item T) T {
			v := reflect.ValueOf(&item).Elem()

			f := v.FieldByIndex(index)

			if !f.CanSet() {
				return item
			}

			f.SetString(f.String() + value)
			return item
		}
		return MutableExpression[T]{Result: fnc}
	} else {
		return MutableExpression[T]{Result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) AppStruct(target any) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, reflect.Slice); success {

		index := fieldIndex

		fnc := func(item T) T {
			v := reflect.ValueOf(&item).Elem()

			f := v.FieldByIndex(index)

			if !f.CanSet() {
				return item
			}

			targetVal := reflect.ValueOf(target)

			if targetVal.Type() != f.Type().Elem() {
				return item
			}
			f.Set(reflect.Append(f, targetVal))

			return item
		}
		return MutableExpression[T]{Result: fnc}
	} else {
		return MutableExpression[T]{Result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetStruct(target any) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, reflect.Struct); success {

		index := fieldIndex

		fnc := func(item T) T {
			v := reflect.ValueOf(&item).Elem()

			f := v.FieldByIndex(index)

			if !f.CanSet() {
				return item
			}

			targetVal := reflect.ValueOf(target)

			if targetVal.Type() != f.Type() {
				return item
			}

			f.Set(targetVal)

			return item
		}
		return MutableExpression[T]{Result: fnc}
	} else {
		return MutableExpression[T]{Result: func(item T) T { return item }}
	}

}
