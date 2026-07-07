package contracts

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

type OpData[T any] struct {
	Function func(T) bool
}

type CompiledQueryable[T any] struct {
	Operators []ZenqlOperator[T]
	Items     *[]T
}
type ZenqlOperator[T any] struct {
	MetaData     OpData[T]
	OperatorType int
	Limit        int
	Skip         int
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

type StreamConf struct {
	FilePath string

	BufferSize int

	ParseErrorCallback func([]error, int)

	ItemCount int
}
type CsvStreamConf[T any] struct {
	Parser        func(row []string) (T, []error)
	StreamHeaders bool
	StreamConf
}

type JsonStreamConf struct {
	StreamConf
}

type DbParam[T any] struct {
	Name  string
	Value T
}
