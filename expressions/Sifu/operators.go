package Sifu

import "reflect"

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
