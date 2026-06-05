package collections

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

import (
	"errors"

	contracts "github.com/malikhan-dev/zenql/contracts"
)

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
