package Sifu

import (
	"reflect"
)

type MutableExpression[T any] struct {
	result func(item T) T
}

func (curr *PropExpression[T]) SetString(value string) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.String}); success {

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
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetInt(value int64) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Int, reflect.Int16, reflect.Int32, reflect.Int8, reflect.Int64}); success {

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

			f.SetInt(int64(value))

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetUint(value uint64) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}); success {

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

			f.SetUint(value)

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetFloat(value float64) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Float64, reflect.Float32}); success {

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

			f.SetFloat(value)

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetBool(value bool) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Bool}); success {

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
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}

	}

}

func (curr *PropExpression[T]) StrApp(value string) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.String}); success {
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
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) AppStruct(target any) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Slice}); success {

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
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetStruct(target any) MutableExpression[T] {

	if success, fieldIndex := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Struct}); success {

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
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}
