package Thor

import "github.com/malikhan-dev/zenq/contracts"

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

type CollectionCompiledQueryable[T any] struct {
	contracts.CompiledQueryable[T]
}

type AssertCompiledQueryable[T any] struct {
	contracts.CompiledQueryable[T]
}
type GroupCompiledQueryable[K comparable, T any] struct {
	contracts.CompiledQueryable[T]
	PropLocator func(T) K
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
