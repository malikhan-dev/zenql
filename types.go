package lingo

import "errors"

type Queryable[T any] struct {
	Items []T
	Err   []OpError
}

type GroupedQueryable[K comparable, T any] struct {
	Items map[K][]T
	Err   []OpError
}

type CollectStream[T any] struct {
	Value T
	Err   OpError
}
type CollectGroupStream[K comparable, T any] struct {
	Value map[K][]T
	Err   OpError
}

type OpError struct {
	Code     int
	Err      error
	MetaData string
}

func ErrFactory(Code int, MetaData string) OpError {
	errStr := OpErrors[Code]

	return OpError{
		Code:     Code,
		Err:      errors.New(errStr),
		MetaData: MetaData,
	}
}

var OpErrors = map[int]string{
	1: "unable to fetch result based on given criteria.",
	2: "property does not exist on type.",
	3: "unsupported type. a struct expected.",
	4: "cant query on empty slice.",
	5: "index is out of range.",
	6: "specified type is not comparable.",
}
