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
		Err:   nil,
	}
}

func From[T any](items []T) *Queryable[T] {

	return &Queryable[T]{
		Items: items,
		Err:   nil,
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
		Out.Err = append(Out.Err, ErrFactory(3, strType.Name()))
	}

	if strType.Kind() == reflect.Ptr {
		strType = strType.Elem()
	}

	TargetField, ok := strType.FieldByName(fieldName)

	newItems := make([]T, 0)

	if ok {

		for _, val := range query.Items {

			RowVale := reflect.ValueOf(val)

			RowField := RowVale.FieldByIndex(TargetField.Index)

			if RowField.Interface() == fieldValue {
				newItems = append(newItems, val)
			}
		}

	} else {
		Out.Err = append(Out.Err, ErrFactory(2, fieldName))
	}
	for _, val := range query.Err {

		Out.Err = append(Out.Err, val)
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
		query.Err = append(query.Err, ErrFactory(1, "AllOrDefault()"))
		return query
	}
}

func (query *Queryable[T]) FirstOrDefault() *Queryable[T] {

	if len(query.Items) > 0 {
		data := query.Items[0]
		query.Items = make([]T, 0)
		query.Items = append(query.Items, data)

	} else {
		query.Err = append(query.Err, ErrFactory(1, "FirstOrDefault()"))
	}
	return query
}

func GroupBy[K comparable, T any](query *Queryable[T], fieldName string) *GroupedQueryable[K, T] {

	var result GroupedQueryable[K, T]

	mapped := make(map[K][]T)

	strType := reflect.TypeFor[T]()

	for _, val := range query.Err {
		result.Err = append(result.Err, val)
	}

	if strType.Kind() == reflect.Ptr {
		strType = strType.Elem()
	}

	if strType.Kind() != reflect.Struct {
		result.Err = append(result.Err, ErrFactory(3, strType.Name()))
		return &result
	}

	targetField, ok := strType.FieldByName(fieldName)
	if !ok {
		result.Err = append(result.Err, ErrFactory(2, fieldName))
		return &result
	}

	for _, val := range query.Items {

		RowVal := reflect.ValueOf(val)
		if RowVal.Kind() == reflect.Ptr {
			RowVal = RowVal.Elem()
		}

		RowField := RowVal.FieldByIndex(targetField.Index)

		key := RowField.Interface()

		if !reflect.TypeOf(key).Comparable() {
			result.Err = append(result.Err, ErrFactory(6, fieldName))
			break
		}

		k, ok := key.(K)

		if !ok {

			result.Err = append(result.Err, ErrFactory(6, fieldName))
			break

		} else {
			mapped[k] = append(mapped[k], val)
		}

	}

	result.Items = mapped
	return &result

}

func (query *Queryable[T]) Count() int {
	return len(query.Items)
}

func (query *Queryable[T]) ErrCount() int {
	return len(query.Err)
}
