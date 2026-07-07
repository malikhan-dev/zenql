package contracts

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

type Filterer[T any] struct {
	Function func(T) bool
}

type Updater[T any] struct {
	Function func(T) T
}

type OpData[T any, O any] struct {
	Function func(T) O
}

func (f Filterer[T]) Filter(item T) bool {
	return f.Function(item)
}

func (U Updater[T]) Update(item T) T {
	return U.Function(item)
}

type IFilter[T any] interface {
	Filter(T) bool
}

type IUpdater[T any] interface {
	Update(T) T
}

type CompiledQueryable[T any] struct {
	Operators []ZenqlOperator[T]
	Items     *[]T
}
type ZenqlOperator[T any] struct {
	Filter       IFilter[T]
	Update       IUpdater[T]
	OperatorType int8
}

type PageOption struct {
	Limit int32
	Skip  int32
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
