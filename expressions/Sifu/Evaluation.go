package Sifu

import (
	"reflect"
)

type SifuExpr[T any] struct {
	op []Operation[T]
}

type Operation[T any] struct {
	FieldName string
}

type EvaluationOperation[T any] struct {
	Result func(item T) bool
}

func (op EvaluationOperation[T]) Gen() func(T) bool {
	return op.Result
}

func (op EvaluationOperation[T]) And(operation ...EvaluationOperation[T]) EvaluationOperation[T] {
	return EvaluationOperation[T]{
		Result: func(item T) bool {
			result := op.Result(item)
			for _, v := range operation {
				result = result && v.Result(item)
			}
			return result
		},
	}
}

func (op EvaluationOperation[T]) Or(operation ...EvaluationOperation[T]) EvaluationOperation[T] {
	return EvaluationOperation[T]{
		Result: func(item T) bool {
			result := op.Result(item)
			for _, v := range operation {
				result = result || v.Result(item)
			}
			return result
		},
	}
}

func OfType[T any]() *SifuExpr[T] {
	return &SifuExpr[T]{}
}

func Expr[T any]() *SifuExpr[T] {
	return OfType[T]()
}

func (curr *SifuExpr[T]) Prop(name string) *Operation[T] {

	operation := Operation[T]{FieldName: name}
	curr.op = append(curr.op, operation)
	return &operation
}

func (curr *Operation[T]) BInt(num int) EvaluationOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
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
	return EvaluationOperation[T]{Result: fnc}
}

func (curr *Operation[T]) SmallerThanInt(num int) EvaluationOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
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
	return EvaluationOperation[T]{Result: fnc}
}

func (curr *Operation[T]) EqStr(value string) EvaluationOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)

	if !ok || field.Type.Kind() != reflect.String {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
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
	return EvaluationOperation[T]{Result: fnc}
}

func (curr *Operation[T]) NotEqStr(value string) EvaluationOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.String {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
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
	return EvaluationOperation[T]{Result: fnc}
}

func (curr *Operation[T]) True() EvaluationOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.Bool {
		return EvaluationOperation[T]{Result: func(item T) bool {
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
	return EvaluationOperation[T]{Result: fnc}
}

func (curr *Operation[T]) False() EvaluationOperation[T] {
	var zero T
	typ := reflect.TypeOf(zero)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok || field.Type.Kind() != reflect.Bool {
		return EvaluationOperation[T]{Result: func(item T) bool {
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
	return EvaluationOperation[T]{Result: fnc}
}

func (curr *Operation[T]) Any(expr any) EvaluationOperation[T] {

	evaluated, ok := expr.(interface {
		evalAny(item any) bool
	})

	if !ok {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	var zero T
	typ := reflect.TypeOf(zero)
	if typ == nil {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	field, ok := typ.FieldByName(curr.FieldName)
	if !ok {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	if field.Type.Kind() != reflect.Slice && field.Type.Kind() != reflect.Array {
		return EvaluationOperation[T]{Result: func(item T) bool { return false }}
	}

	index := field.Index

	return EvaluationOperation[T]{
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
func (curr EvaluationOperation[T]) evalAny(item any) bool {
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

func (curr *Operation[T]) Less() CompareOperation[T] {
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

type KeySelectorOperation[T any, K comparable] struct {
	Result func(item T) K
}

func (curr KeySelectorOperation[T, K]) Gen() func(T) K {
	return curr.Result
}

func KeyAs[T any, K comparable](operation *Operation[T]) KeySelectorOperation[T, K] {
	var zeroKey K

	var zero T
	typ := reflect.TypeOf(zero)

	if typ == nil {
		return KeySelectorOperation[T, K]{Result: func(item T) K { return zeroKey }}
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return KeySelectorOperation[T, K]{Result: func(item T) K { return zeroKey }}
	}

	field, ok := typ.FieldByName(operation.FieldName)
	if !ok {
		return KeySelectorOperation[T, K]{Result: func(item T) K { return zeroKey }}
	}

	expectedType := reflect.TypeOf(zeroKey)
	if expectedType == nil || field.Type != expectedType {
		return KeySelectorOperation[T, K]{Result: func(item T) K { return zeroKey }}
	}

	index := field.Index

	return KeySelectorOperation[T, K]{
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
