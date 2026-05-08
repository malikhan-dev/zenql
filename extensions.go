package lingo

import (
	"reflect"
)

func FindByPredicate[T any](items []T, predicate func(T) bool) *[]T {

	var result []T

	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return &result

}

func FindFirstByPredicate[T any](items []T, predicate func(T) bool) *T {

	var result T

	for _, v := range items {
		if predicate(v) {
			result = v
			break
		}
	}

	return &result
}

func RemoveFirstByPredicate[T any](items []T, predicate func(T) bool) *[]T {

	var result []T

	conditionMet := false

	for _, v := range items {

		if predicate(v) && !conditionMet {
			conditionMet = true
			continue

		} else {

			result = append(result, v)

		}

	}

	return &result
}

func RemoveByPredicate[T any](items []T, predicate func(T) bool) *[]T {

	var result []T

	for _, v := range items {
		if predicate(v) {
			continue
		} else {
			result = append(result, v)
		}

	}

	return &result
}

func (query *Queryable[T]) Filter(predicate func(T) bool) *Queryable[T] {

	var result []T
	result = make([]T, 0)

	for _, v := range query.Items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return &Queryable[T]{
		Items: result,
		err:   nil,
	}
}

func From[T any](items []T) *Queryable[T] {

	return &Queryable[T]{
		Items: items,
		err:   nil,
	}
}

func Any[T any](Items []T, Condition func(T) bool) bool {

	for _, v := range Items {

		if Condition(v) {
			return true
		}
	}
	return false
}

func (query *Queryable[T]) Where(fieldName string, fieldValue any) *Queryable[T] {

	var Out Queryable[T]

	strType := reflect.TypeFor[T]()

	if strType.Kind() != reflect.Struct {
		Out.err = append(Out.err, ErrFactory(3, strType.Name()))
	}

	if strType.Kind() == reflect.Ptr {
		strType = strType.Elem()
	}

	field, ok := strType.FieldByName(fieldName)

	newItems := make([]T, 0)

	if ok {

		for _, val := range query.Items {

			v := reflect.ValueOf(val)

			f := v.FieldByIndex(field.Index)

			if f.Interface() == fieldValue {
				newItems = append(newItems, val)
			}
		}

	} else {
		Out.err = append(Out.err, ErrFactory(2, fieldName))
	}
	for _, val := range query.err {

		Out.err = append(Out.err, val)
	}

	Out.Items = newItems
	return &Out
}

func (query *Queryable[T]) All() *Queryable[T] {
	if len(query.Items) > 0 {
		return query
	} else {
		panic(ErrFactory(4, ""))
	}
}

func (query *Queryable[T]) First() *Queryable[T] {
	if len(query.Items) > 0 {
		data := query.Items[0]
		query.Items = make([]T, 0)
		query.Items = append(query.Items, data)
		return query
	} else {
		panic(ErrFactory(4, ""))
	}
}

func (query *Queryable[T]) AllOrDefault() *Queryable[T] {
	if len(query.Items) > 0 {
		return query
	} else {
		query.err = append(query.err, ErrFactory(1, "AllOrDefault()"))
		return query
	}
}

func (query *Queryable[T]) FirstOrDefault() *Queryable[T] {

	if len(query.Items) > 0 {
		data := query.Items[0]
		query.Items = make([]T, 0)
		query.Items = append(query.Items, data)

	} else {
		query.err = append(query.err, ErrFactory(1, "FirstOrDefault()"))
	}
	return query
}
