package collections

import (
	"errors"

	"github.com/malikhan-dev/zenql/contracts"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

type FilterChan[T any] struct {
	Keep bool
	Item T
}
type CollectionCompiledQueryable[T any] struct {
	contracts.CompiledQueryable[T]
}

func (collection *CollectionCompiledQueryable[T]) Iterate() *[]T {
	return collection.Items
}

func (collection *CollectionCompiledQueryable[T]) IterateOperators() *[]contracts.ZenqlOperator[T] {
	return &collection.Operators
}

type AssertCompiledQueryable[T any] struct {
	contracts.CompiledQueryable[T]
}
type GroupCompiledQueryable[K comparable, T any] struct {
	contracts.CompiledQueryable[T]
	PropLocator func(T) K
}

func (collection *AssertCompiledQueryable[T]) Iterate() *[]T {
	return collection.Items
}

func (collection *CollectionCompiledQueryable[T]) GetLen() int {
	return len(*collection.Items)
}
func (collection *GroupCompiledQueryable[K, T]) Iterate() *[]T {
	return collection.Items
}

type Sortable[T any] struct {
	Items []T
	less  func(a, b T) bool
	desc  bool
}

func NewSortable[T any](less func(a, b T) bool, desc bool) *Sortable[T] {
	return &Sortable[T]{
		Items: make([]T, 0),
		less:  less,
		desc:  desc,
	}
}

func (h Sortable[T]) Len() int {
	return len(h.Items)
}

func (h Sortable[T]) Less(i, j int) bool {

	if h.desc {
		return h.less(h.Items[j], h.Items[i])
	}

	return h.less(h.Items[i], h.Items[j])
}

func (h Sortable[T]) Swap(i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
}

func (h *Sortable[T]) Push(x any) {
	h.Items = append(h.Items, x.(T))
}

func (h *Sortable[T]) Pop() any {

	old := h.Items
	n := len(old)

	item := old[n-1]
	h.Items = old[:n-1]

	return item
}

func ErrFactory(Code int, MetaData string) contracts.OpError {
	errStr := OpErrors[Code]

	return contracts.OpError{
		Code:     Code,
		Err:      errors.New(errStr),
		MetaData: MetaData,
	}
}

type Queryable[T any] struct {
	Items []T
	Err   []contracts.OpError
}

type GroupedQueryable[K comparable, T any] struct {
	Items map[K][]T
	Err   []contracts.OpError
}

var OpErrors = map[int]string{
	1: "unable to fetch result based on given criteria.",
	2: "property does not exist on type.",
	3: "unsupported type. a struct expected.",
	4: "cant query on empty slice.",
	5: "index is out of range.",
	6: "specified type is not comparable.",
}
