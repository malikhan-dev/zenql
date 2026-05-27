package streams

import (
	contracts "github.com/malikhan-dev/zenq/contracts"
)

type OperatorType int

const (
	BuildQueryable    = 1
	FilterQueryable   = 2
	ThrottleQueryable = 3
)

func Filter[T any](currentOps *contracts.CompiledQueryable[T], filter func(item T) bool) *contracts.CompiledQueryable[T] {

	currentOps.Operators = append(currentOps.Operators, contracts.ZenqOperator[T]{
		MetaData: contracts.OpData[T]{
			MetaData: "FFilter",
			Function: filter,
		},
		OperatorType: 2,
	})

	return currentOps
}
