package collections

import "github.com/malikhan-dev/zenql/contracts/v2"

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */
type OperatorType int

const (
	FromItems          = 1
	WhereCollection    = 2
	AnyCollection      = 4
	GroupCollection    = 5
	DistinctCollection = 6
	TakeCollection     = 7
	SkipCollection     = 8
)

func CoreFilter[T any](Operator contracts.ZenqlOperator[T], item T) bool {

	ShouldKeep := true

	switch Operator.OperatorType {

	case WhereCollection:
		if !Operator.MetaData.Function(item) {
			ShouldKeep = false
			break
		}

	}
	return ShouldKeep
}

func extractLimits[T any](op []contracts.ZenqlOperator[T]) (int, int) {

	skipLimit := -1
	takeLimit := -1
	for _, operator := range op {

		if operator.OperatorType == SkipCollection {
			skipLimit = operator.Skip
			continue
		}

		if operator.OperatorType == TakeCollection {
			takeLimit = operator.Limit
			continue
		}
	}
	return skipLimit, takeLimit

}
func (op *CollectionCompiledQueryable[T]) Take(count int) *CollectionCompiledQueryable[T] {
	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: TakeCollection,
		Limit:        count,
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Skip(count int) *CollectionCompiledQueryable[T] {
	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SkipCollection,
		Skip:         count,
	})
	return op
}
func Group[K comparable, T any](op *CollectionCompiledQueryable[T], locator func(T) K) *GroupCompiledQueryable[K, T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: GroupCollection,
		MetaData: contracts.OpData[T]{
			Function: func(t T) bool {
				return true
			},
		},
	})
	return &GroupCompiledQueryable[K, T]{
		CompiledQueryable: op.CompiledQueryable,
		PropLocator:       locator,
	}
}
func (op *AssertCompiledQueryable[T]) Assert() bool {

	for _, item := range *op.Items {

		for _, op := range op.Operators {

			switch op.OperatorType {

			case AnyCollection:
				if op.MetaData.Function(item) {
					return true

				}

			}

		}
	}
	return false
}
func From[T any](items *[]T) *CollectionCompiledQueryable[T] {

	initiateOperator := make([]contracts.ZenqlOperator[T], 0)
	initiateOperator = append(initiateOperator, contracts.ZenqlOperator[T]{
		OperatorType: FromItems,
		MetaData: contracts.OpData[T]{
			Function: func(t T) bool {
				return true
			},
		},
	})
	queryData := contracts.CompiledQueryable[T]{
		Items:     items,
		Operators: initiateOperator,
	}

	return &CollectionCompiledQueryable[T]{
		queryData,
	}
}
func (op *CollectionCompiledQueryable[T]) Where(function func(T) bool) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		MetaData: contracts.OpData[T]{
			Function: function,
		},
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: AnyCollection,
		MetaData: contracts.OpData[T]{
			Function: function,
		},
	})
	return &AssertCompiledQueryable[T]{
		op.CompiledQueryable,
	}
}
