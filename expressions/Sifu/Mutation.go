package Sifu

import "reflect"

type MutableOperation[T any] struct {
	Result func(item T) T
}

func (curr *Operation[T]) SetString(value string) MutableOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return MutableOperation[T]{Result: func(item T) T { return item }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return MutableOperation[T]{Result: func(item T) T { return item }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.String {
		return MutableOperation[T]{Result: func(item T) T { return item }}
	}

	index := field.Index

	fnc := func(item T) T {
		v := reflect.ValueOf(&item).Elem()

		if !v.IsValid() {
			return item
		}

		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return item
			}
			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return item
		}

		f := v.FieldByIndex(index)

		if f.Kind() != reflect.String {
			return item
		}

		if !f.CanSet() {
			return item
		}

		f.SetString(value)
		return item
	}
	return MutableOperation[T]{Result: fnc}
}

func (curr *Operation[T]) AppendString(value string) MutableOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return MutableOperation[T]{Result: func(item T) T { return item }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return MutableOperation[T]{Result: func(item T) T { return item }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.String {
		return MutableOperation[T]{Result: func(item T) T { return item }}
	}

	index := field.Index

	fnc := func(item T) T {
		v := reflect.ValueOf(&item).Elem()

		if !v.IsValid() {
			return item
		}

		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return item
			}
			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return item
		}

		f := v.FieldByIndex(index)

		if f.Kind() != reflect.String {
			return item
		}

		if !f.CanSet() {
			return item
		}

		f.SetString(f.String() + value)
		return item
	}
	return MutableOperation[T]{Result: fnc}
}

func (op MutableOperation[T]) Exec() func(item T) T {
	return op.Result
}
