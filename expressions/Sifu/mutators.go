package Sifu


/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"reflect"
	"strings"
	"unsafe"
)

type MutableExpression[T any] struct {
	result func(item T) T
}

func (curr *PropExpression[T]) SetString(value string) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.String}); success {

		offset := field.Offset

		fnc := func(item T) T {

			targetString := (*string)(unsafe.Add(unsafe.Pointer(&item), offset))

			*targetString = value

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetInt(value int64) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Int, reflect.Int16, reflect.Int32, reflect.Int8, reflect.Int64}); success {

		offSet := field.Offset

		fnc := func(item T) T {

			targetField := (*int64)(unsafe.Add(unsafe.Pointer(&item), offSet))

			*targetField = value

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetUint(value uint64) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}); success {

		offSet := field.Offset

		fnc := func(item T) T {

			targetField := (*uint64)(unsafe.Add(unsafe.Pointer(&item), offSet))

			*targetField = value

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetFloat(value float64) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Float64, reflect.Float32}); success {

		offSet := field.Offset

		fnc := func(item T) T {

			targetField := (*float64)(unsafe.Add(unsafe.Pointer(&item), offSet))

			*targetField = value

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) SetBool(value bool) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Bool}); success {

		offSet := field.Offset

		fnc := func(item T) T {

			targetVal := (*bool)(unsafe.Add(unsafe.Pointer(&item), offSet))

			*targetVal = value

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}

	}

}

func (curr *PropExpression[T]) StrApp(value string) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.String}); success {
		offSet := field.Offset

		fnc := func(item T) T {

			targetVal := (*string)(unsafe.Add(unsafe.Pointer(&item), offSet))

			var b strings.Builder

			b.WriteString(*targetVal)

			b.WriteString(value)

			*targetVal = b.String()

			return item
		}
		return MutableExpression[T]{result: fnc}
	} else {
		return MutableExpression[T]{result: func(item T) T { return item }}
	}

}

func (curr *PropExpression[T]) AppStruct(target any) MutableExpression[T] {

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Slice}); success {

		index := field.Index

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

	if success, field := canReflect[T](curr.FieldName, []reflect.Kind{reflect.Struct}); success {

		index := field.Index

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
