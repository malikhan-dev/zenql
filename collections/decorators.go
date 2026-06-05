package collections

import (
	"github.com/malikhan-dev/zenql/contracts"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

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
	return AAll(query)
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

	return GGroupBy[K, T](query, fieldName)
}

func (query *GroupedQueryable[K, T]) CollectGroup() (map[K][]T, []contracts.OpError) {
	return CCollectGroup[K, T](query)
}

func (query *Queryable[T]) Count() int {
	return CCount(query)
}

func (query *Queryable[T]) ErrCount() int {
	return EErrCount(query)
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
