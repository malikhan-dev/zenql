package lingo

type Queryable[T any] struct {
	Items []T
	err   []error
}
