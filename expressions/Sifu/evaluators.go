package Sifu

import (
	"reflect"
	"unsafe"
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
func (curr *PropExpression[T]) link(linkProp string, eval int8) CompareOperation[T] {
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
			bf := bv.FieldByName(linkProp)

			if !af.IsValid() || !bf.IsValid() {
				return false
			}

			switch af.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

				if eval == 1 {
					return af.Int() > bf.Int()
				} else if eval == 2 {
					return af.Int() < bf.Int()
				} else if eval == 3 {
					return af.Int() == bf.Int()
				}
				return false

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				if eval == 1 {
					return af.Uint() > bf.Uint()
				} else if eval == 2 {
					return af.Uint() < bf.Uint()
				} else if eval == 3 {
					return af.Uint() == bf.Uint()
				}
				return false

			case reflect.Float32, reflect.Float64:

				if eval == 1 {
					return af.Float() > bf.Float()
				} else if eval == 2 {
					return af.Float() < bf.Float()
				} else if eval == 3 {
					return af.Float() == bf.Float()
				}
				return false
			case reflect.String:
				if eval == 1 {
					return af.String() > bf.String()
				} else if eval == 2 {
					return af.String() < bf.String()
				} else if eval == 3 {
					return af.String() == bf.String()
				}
				return false

			default:
				return false
			}
		},
	}
}

func (curr *PropExpression[T]) LinkEq(linkProp string) CompareOperation[T] {
	return curr.link(linkProp, 3)
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

	destType := dest.Type()

	switch destType.Name() {
	case "int", "int8", "int16", "int32", "int64":
		n, ok := num.(int)
		if ok {
			return compareInt(dest.Int(), int64(n), eval)
		} else {
			return false
		}

	case "uint", "uint8", "uint16", "uint32", "uint64":
		n, ok := num.(uint)
		if ok {
			return compareUint(dest.Uint(), uint64(n), eval)
		} else {
			return false
		}

	case "float32", "float64":
		n, ok := num.(float64)
		if ok {
			return compareFloat(dest.Float(), n, eval)
		} else {
			return false
		}
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

	offset := field.Offset

	fnc := func(item T) bool {

		ptr := (*string)(unsafe.Add(unsafe.Pointer(&item), offset))
		return *ptr == value
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) StrEqNot(value string) ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	offset := field.Offset

	fnc := func(item T) bool {

		ptr := (*string)(unsafe.Add(unsafe.Pointer(&item), offset))
		return *ptr != value
	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) True() ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	offset := field.Offset

	fnc := func(item T) bool {

		ptr := (*bool)(unsafe.Add(unsafe.Pointer(&item), offset))
		return *ptr == true

	}
	return ExpressionEvaluation[T]{result: fnc}
}

func (curr *PropExpression[T]) False() ExpressionEvaluation[T] {

	var zero T

	typ := reflect.TypeOf(zero)

	field, _ := typ.FieldByName(curr.FieldName)

	offset := field.Offset

	fnc := func(item T) bool {

		ptr := (*bool)(unsafe.Add(unsafe.Pointer(&item), offset))
		return *ptr == false

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
