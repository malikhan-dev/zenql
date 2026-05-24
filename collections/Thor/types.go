package Thor

import "github.com/malikhan-dev/zenq/contracts"

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
