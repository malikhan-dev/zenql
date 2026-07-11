package Sifu

import (
	"reflect"
)

func Expr[T any]() *TypeExpression[T] {
	return &TypeExpression[T]{}
}

func (curr ExpressionEvaluation[T]) evalAny(item any) bool {
	typed, ok := item.(T)
	if !ok {
		return false
	}
	return curr.result(typed)
}

type CompareOperation[T any] struct {
	result func(a T, b T) bool
}

func (curr *PropExpression[T]) Less() CompareOperation[T] {
	fieldName := curr.FieldName

	return CompareOperation[T]{
		result: func(a T, b T) bool {
			av := reflect.ValueOf(a)
			bv := reflect.ValueOf(b)

			if av.Kind() == reflect.Ptr {
				if av.IsNil() {
					return false
				}
				av = av.Elem()
			}

			if bv.Kind() == reflect.Ptr {
				if bv.IsNil() {
					return false
				}
				bv = bv.Elem()
			}

			af := av.FieldByName(fieldName)
			bf := bv.FieldByName(fieldName)

			if !af.IsValid() || !bf.IsValid() {
				return false
			}

			switch af.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return af.Int() < bf.Int()

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return af.Uint() < bf.Uint()

			case reflect.Float32, reflect.Float64:
				return af.Float() < bf.Float()

			case reflect.String:
				return af.String() < bf.String()

			default:
				return false
			}
		},
	}
}

type KeySelectorExpression[T any, K comparable] struct {
	result func(item T) K
}

func KeyAs[T any, K comparable](operation *PropExpression[T]) KeySelectorExpression[T, K] {
	var zeroKey K

	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return KeySelectorExpression[T, K]{result: func(item T) K { return zeroKey }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return KeySelectorExpression[T, K]{result: func(item T) K { return zeroKey }}
	}

	field, ok := typ.FieldByName(operation.FieldName)
	if !ok {
		return KeySelectorExpression[T, K]{result: func(item T) K { return zeroKey }}
	}

	expectedType := reflect.TypeOf(zeroKey)
	if expectedType == nil || field.Type != expectedType {
		return KeySelectorExpression[T, K]{result: func(item T) K { return zeroKey }}
	}

	index := field.Index

	return KeySelectorExpression[T, K]{
		result: func(item T) K {
			v := reflect.ValueOf(item)

			if !v.IsValid() {
				return zeroKey
			}

			if v.Kind() == reflect.Ptr {
				if v.IsNil() {
					return zeroKey
				}
				v = v.Elem()
			}

			if v.Kind() != reflect.Struct {
				return zeroKey
			}

			fieldValue := v.FieldByIndex(index)
			if !fieldValue.IsValid() || !fieldValue.CanInterface() {
				return zeroKey
			}

			value, ok := fieldValue.Interface().(K)
			if !ok {
				return zeroKey
			}

			return value
		},
	}
}

func (curr *PropExpression[T]) Any(expr any) ExpressionEvaluation[T] {

	evaluated, ok := expr.(interface {
		evalAny(item any) bool
	})

	if !ok {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	var zero T
	typ := reflect.TypeOf(zero)
	if typ == nil {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	if field.Type.Kind() != reflect.Slice && field.Type.Kind() != reflect.Array {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	index := field.Index

	return ExpressionEvaluation[T]{
		result: func(item T) bool {
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
			if f.Kind() != reflect.Slice && f.Kind() != reflect.Array {
				return false
			}

			for i := 0; i < f.Len(); i++ {
				if evaluated.evalAny(f.Index(i).Interface()) {
					return true
				}
			}

			return false
		},
	}
}

func (curr *TypeExpression[T]) Prop(name string) *PropExpression[T] {

	operation := PropExpression[T]{FieldName: name}
	curr.op = append(curr.op, operation)
	return &operation
}
