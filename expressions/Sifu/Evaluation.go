package Sifu

import (
	"reflect"
)

type TypeExpression[T any] struct {
	op []PropExpression[T]
}

type PropExpression[T any] struct {
	FieldName string
}

type ExpressionEvaluation[T any] struct {
	Result func(item T) bool
}

func (op ExpressionEvaluation[T]) Gen() func(T) bool {
	return op.Result
}

func (op ExpressionEvaluation[T]) And(operation ...ExpressionEvaluation[T]) ExpressionEvaluation[T] {
	return ExpressionEvaluation[T]{
		Result: func(item T) bool {
			result := op.Result(item)
			for _, v := range operation {
				result = result && v.Result(item)
			}
			return result
		},
	}
}

func (op ExpressionEvaluation[T]) Or(operation ...ExpressionEvaluation[T]) ExpressionEvaluation[T] {
	return ExpressionEvaluation[T]{
		Result: func(item T) bool {
			result := op.Result(item)
			for _, v := range operation {
				result = result || v.Result(item)
			}
			return result
		},
	}
}

func OfType[T any]() *TypeExpression[T] {
	return &TypeExpression[T]{}
}

func Expr[T any]() *TypeExpression[T] {
	return OfType[T]()
}

func (curr *TypeExpression[T]) Prop(name string) *PropExpression[T] {

	operation := PropExpression[T]{FieldName: name}
	curr.op = append(curr.op, operation)
	return &operation
}

func (curr *PropExpression[T]) Btint(num int) ExpressionEvaluation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	index := field.Index

	fnc := func(item T) bool {
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

		if f.Kind() != reflect.Int && f.Kind() != reflect.Int64 && f.Kind() != reflect.Int32 && f.Kind() != reflect.Int16 && f.Kind() != reflect.Int8 && f.Kind() != reflect.Uint {
			return false
		}
		return f.Int() > int64(num)
	}
	return ExpressionEvaluation[T]{Result: fnc}
}

func (curr *PropExpression[T]) Stint(num int) ExpressionEvaluation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	index := field.Index

	fnc := func(item T) bool {
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

		if f.Kind() != reflect.Int && f.Kind() != reflect.Int64 && f.Kind() != reflect.Int32 && f.Kind() != reflect.Int16 && f.Kind() != reflect.Int8 && f.Kind() != reflect.Uint {
			return false
		}

		return f.Int() < int64(num)
	}
	return ExpressionEvaluation[T]{Result: fnc}
}

func (curr *PropExpression[T]) EqStr(value string) ExpressionEvaluation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)

	if !ok || field.Type.Kind() != reflect.String {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	index := field.Index

	fnc := func(item T) bool {
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
	return ExpressionEvaluation[T]{Result: fnc}
}

func (curr *PropExpression[T]) NotEqStr(value string) ExpressionEvaluation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.String {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	index := field.Index

	fnc := func(item T) bool {
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

		return f.String() != value
	}
	return ExpressionEvaluation[T]{Result: fnc}
}

func (curr *PropExpression[T]) True() ExpressionEvaluation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.Bool {
		return ExpressionEvaluation[T]{Result: func(item T) bool {
			return false
		}}
	}

	index := field.Index

	fnc := func(item T) bool {
		v := reflect.ValueOf(item)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return false
			}
			v = v.Elem()
		}

		f := v.FieldByIndex(index)
		return f.Bool() == true
	}
	return ExpressionEvaluation[T]{Result: fnc}
}

func (curr *PropExpression[T]) False() ExpressionEvaluation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.Bool {
		return ExpressionEvaluation[T]{Result: func(item T) bool {
			return false
		}}
	}

	index := field.Index

	fnc := func(item T) bool {
		v := reflect.ValueOf(item)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return false
			}
			v = v.Elem()
		}

		f := v.FieldByIndex(index)
		return f.Bool() == false
	}
	return ExpressionEvaluation[T]{Result: fnc}
}

func (curr *PropExpression[T]) Any(expr any) ExpressionEvaluation[T] {

	evaluated, ok := expr.(interface {
		evalAny(item any) bool
	})

	if !ok {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	var zero T
	typ := reflect.TypeOf(zero)
	if typ == nil {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	if field.Type.Kind() != reflect.Slice && field.Type.Kind() != reflect.Array {
		return ExpressionEvaluation[T]{Result: func(item T) bool { return false }}
	}

	index := field.Index

	return ExpressionEvaluation[T]{
		Result: func(item T) bool {
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
func (curr ExpressionEvaluation[T]) evalAny(item any) bool {
	typed, ok := item.(T)
	if !ok {
		return false
	}
	return curr.Result(typed)
}

type CompareOperation[T any] struct {
	Result func(a T, b T) bool
}

func (curr CompareOperation[T]) Gen() func(T, T) bool {
	return curr.Result
}

func (curr *PropExpression[T]) Less() CompareOperation[T] {
	fieldName := curr.FieldName

	return CompareOperation[T]{
		Result: func(a T, b T) bool {
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
	Result func(item T) K
}

func (curr KeySelectorExpression[T, K]) Gen() func(T) K {
	return curr.Result
}

func KeyAs[T any, K comparable](operation *PropExpression[T]) KeySelectorExpression[T, K] {
	var zeroKey K

	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return KeySelectorExpression[T, K]{Result: func(item T) K { return zeroKey }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return KeySelectorExpression[T, K]{Result: func(item T) K { return zeroKey }}
	}

	field, ok := typ.FieldByName(operation.FieldName)
	if !ok {
		return KeySelectorExpression[T, K]{Result: func(item T) K { return zeroKey }}
	}

	expectedType := reflect.TypeOf(zeroKey)
	if expectedType == nil || field.Type != expectedType {
		return KeySelectorExpression[T, K]{Result: func(item T) K { return zeroKey }}
	}

	index := field.Index

	return KeySelectorExpression[T, K]{
		Result: func(item T) K {
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
