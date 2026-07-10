package Sifu

import (
	"reflect"
)

func (curr *PropExpression[T]) NumBigger(num any) ExpressionEvaluation[T] {

	return curr.NumCmp(true, num)
}

func (curr *PropExpression[T]) NumSmaller(num any) ExpressionEvaluation[T] {
	return curr.NumCmp(false, num)
}

func (curr *PropExpression[T]) NumCmp(isBigger bool, num any) ExpressionEvaluation[T] {
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

		kind := f.Kind()

		isInt := kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 ||
			kind == reflect.Int32 || kind == reflect.Int64 || kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64

		isFloat := kind == reflect.Float32 || kind == reflect.Float64

		if !isInt && !isFloat {
			return false
		}

		return CompareNum(num, f, isBigger)
	}
	return ExpressionEvaluation[T]{Result: fnc}
}

func CompareNum(num any, dest reflect.Value, bigger bool) bool {
	switch v := num.(type) {

	case int:
		return compareInt(dest.Int(), int64(v), bigger)
	case int8:
		return compareInt(dest.Int(), int64(v), bigger)
	case int16:
		return compareInt(dest.Int(), int64(v), bigger)
	case int32:
		return compareInt(dest.Int(), int64(v), bigger)
	case int64:
		return compareInt(dest.Int(), v, bigger)
	case uint8:
		return compareUint(dest.Uint(), uint64(v), bigger)
	case uint16:
		return compareUint(dest.Uint(), uint64(v), bigger)
	case uint32:
		return compareUint(dest.Uint(), uint64(v), bigger)
	case uint64:
		return compareUint(dest.Uint(), v, bigger)

	case float32:
		return compareFloat(dest.Float(), float64(v), bigger)
	case float64:
		return compareFloat(dest.Float(), v, bigger)
	}

	return false
}

func compareInt(a, b int64, bigger bool) bool {
	if bigger {
		return a > b
	}
	return a < b
}

func compareUint(a, b uint64, bigger bool) bool {
	if bigger {
		return a > b
	}
	return a < b
}

func compareFloat(a, b float64, bigger bool) bool {
	if bigger {
		return a > b
	}
	return a < b
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
