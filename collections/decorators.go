package collections

import (
	"github.com/malikhan-dev/lingo/contracts"
	"reflect"
)

func From[T any](items []T) *Queryable[T] {

	return FFrom[T](items)
}

func Any[T any](Items []T, Condition func(T) bool) bool {

	return AAny(Items, Condition)
}

func (query *Queryable[T]) Filter(predicate func(T) bool) *Queryable[T] {

	return FFilter(query, predicate)
}

func (query *Queryable[T]) Where(fieldName string, fieldValue any) *Queryable[T] {
	return WWhere(query, fieldName, fieldValue)
}

func (query *Queryable[T]) All() *Queryable[T] {
	if len(query.Items) > 0 {
		return query
	} else {
		panic(ErrFactory(4, ""))
	}
}

func (query *Queryable[T]) First() *Queryable[T] {
	return FFirst(query)
}

func (query *Queryable[T]) AllOrDefault() *Queryable[T] {
	return AAllOrDefault(query)
}

func (query *Queryable[T]) FirstOrDefault() *Queryable[T] {
	return FFirstOrDefault(query)
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

func (query *GroupedQueryable[K, T]) CollectGroup() (map[K][]T, []contracts.OpError) {
	return CCollectGroup[K, T](query)
}

func (query *Queryable[T]) CCount() int {
	return len(query.Items)
}

func (query *Queryable[T]) EErrCount() int {
	return len(query.Err)
}

func (query *Queryable[T]) Collect() ([]T, []contracts.OpError) {
	return CCollect[T](query)

}

func (query *Queryable[T]) CollectRange(cnt int) ([]T, []contracts.OpError) {

	return CCollectRange(query, cnt)
}

func (query *Queryable[T]) Pipe(BufferSize int) <-chan contracts.CollectStream[T] {

	return PPipe(query, BufferSize)
}
