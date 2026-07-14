package Sifu

import (
	"reflect"
)

func (curr *PropExpression[T]) NumBigger(num any) ExpressionEvaluation[T] {
	return curr.numcmp(1, num)
}

func (curr *PropExpression[T]) NumSmaller(num any) ExpressionEvaluation[T] {
	return curr.numcmp(2, num)
}

func (curr *PropExpression[T]) NumEq(num any) ExpressionEvaluation[T] {
	return curr.numcmp(3, num)
}

func (curr *PropExpression[T]) numcmp(eval int8, num any) ExpressionEvaluation[T] {
	var zero T

	typ := reflect.TypeOf(zero)

	field, ok := typ.FieldByName(curr.FieldName)

	if !ok {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	index := field.Index

	fnc := func(item T) bool {
		v := reflect.ValueOf(item)

		f := v.FieldByIndex(index)

		kind := f.Kind()

		isInt := kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 ||
			kind == reflect.Int32 || kind == reflect.Int64 || kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64

		isFloat := kind == reflect.Float32 || kind == reflect.Float64

		if !isInt && !isFloat {
			return false
		}

		return castAndCompare(num, f, eval)
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func castAndCompare(num any, dest reflect.Value, eval int8) bool {
	switch v := num.(type) {

	case int:
		return compareInt(dest.Int(), int64(v), eval)
	case int8:
		return compareInt(dest.Int(), int64(v), eval)
	case int16:
		return compareInt(dest.Int(), int64(v), eval)
	case int32:
		return compareInt(dest.Int(), int64(v), eval)
	case int64:
		return compareInt(dest.Int(), v, eval)
	case uint8:
		return compareUint(dest.Uint(), uint64(v), eval)
	case uint16:
		return compareUint(dest.Uint(), uint64(v), eval)
	case uint32:
		return compareUint(dest.Uint(), uint64(v), eval)
	case uint64:
		return compareUint(dest.Uint(), v, eval)

	case float32:
		return compareFloat(dest.Float(), float64(v), eval)
	case float64:
		return compareFloat(dest.Float(), v, eval)
	}

	return false
}

func compareInt(a, b int64, eval int8) bool {
	if eval == 1 {
		return a > b
	} else if eval == 2 {
		return a < b
	} else if eval == 3 {
		return a == b
	} else {
		return false
	}

}

func compareUint(a, b uint64, eval int8) bool {
	if eval == 1 {
		return a > b
	} else if eval == 2 {
		return a < b
	} else if eval == 3 {
		return a == b
	} else {
		return false
	}
}

func compareFloat(a, b float64, eval int8) bool {
	if eval == 1 {
		return a > b
	} else if eval == 2 {
		return a < b
	} else if eval == 3 {
		return a == b
	} else {
		return false
	}
}

func (curr *PropExpression[T]) StrIn(value []string) ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, ok := typ.FieldByName(curr.FieldName)

	if !ok || field.Type.Kind() != reflect.String {
		return ExpressionEvaluation[T]{result: func(item T) bool { return false }}
	}

	index := field.Index

	fnc := func(item T) bool {

		v := reflect.ValueOf(item)

		f := v.FieldByIndex(index)

		if f.Kind() != reflect.String {
			return false
		}

		for _, v := range value {

			if f.String() == v {
				return true
			}
		}

		return false
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) StrEq(value string) ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	index := field.Index

	fnc := func(item T) bool {

		v := reflect.ValueOf(item)

		f := v.FieldByIndex(index)

		return f.String() == value
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) StrEqNot(value string) ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	index := field.Index

	fnc := func(item T) bool {

		v := reflect.ValueOf(item)

		f := v.FieldByIndex(index)

		/*		if f.Kind() != reflect.String {
					return false
				}
		*/
		return f.String() != value
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) True() ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	index := field.Index

	fnc := func(item T) bool {

		v := reflect.ValueOf(item)

		f := v.FieldByIndex(index)

		return f.Bool() == true
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) False() ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	index := field.Index

	fnc := func(item T) bool {

		v := reflect.ValueOf(item)

		f := v.FieldByIndex(index)

		return f.Bool() == false
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (op ExpressionEvaluation[T]) And(operation ...ExpressionEvaluation[T]) ExpressionEvaluation[T] {
	return ExpressionEvaluation[T]{
		result: func(item T) bool {
			result := op.result(item)
			for _, v := range operation {
				result = result && v.result(item)
			}
			return result
		},
	}
}

func (op ExpressionEvaluation[T]) Or(operation ...ExpressionEvaluation[T]) ExpressionEvaluation[T] {
	return ExpressionEvaluation[T]{
		result: func(item T) bool {
			result := op.result(item)
			if result {
				return true
			} else {
				for _, v := range operation {
					curr := v.result(item)
					if curr {
						return true
					}
					result = result || curr
				}
			}

			return result
		},
	}
}
