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

func (op EvaluationOperation[T]) Eval() func(T) bool {
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

func (curr *Operation[T]) BiggerThanInt(num int) EvaluationOperation[T] {
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

func (curr *Operation[T]) EqualToString(value string) EvaluationOperation[T] {
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

func (curr *Operation[T]) NotEqualToString(value string) EvaluationOperation[T] {
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
