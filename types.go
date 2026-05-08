package lingo

import "errors"

type Queryable[T any] struct {
	Items []T
	err   []OpError
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
}
