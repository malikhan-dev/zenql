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
	UpdateCollection   = 9
)

func CoreFilter[T any](Operator contracts.ZenqlOperator[T], item T) bool {

	ShouldKeep := true

	switch Operator.OperatorType {

	case WhereCollection:
		if !Operator.Filter.Filter(item) {
			ShouldKeep = false
			break
		}

	}
	return ShouldKeep
}

func (op *CollectionCompiledQueryable[T]) Take(count int32) *CollectionCompiledQueryable[T] {

	op.Page.Limit = count

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: TakeCollection,
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Skip(count int32) *CollectionCompiledQueryable[T] {

	op.Page.Skip = count

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SkipCollection,
	})
	return op
}
func Group[K comparable, T any](op *CollectionCompiledQueryable[T], locator func(T) K) *GroupCompiledQueryable[K, T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: GroupCollection,
	})
	return &GroupCompiledQueryable[K, T]{
		CompiledQueryable: op.CompiledQueryable,
		PropLocator:       locator,
		Page:              op.Page,
	}
}
func (op *AssertCompiledQueryable[T]) Assert() bool {

	for _, item := range *op.Items {

		for _, op := range op.Operators {

			switch op.OperatorType {

			case AnyCollection:
				if op.Filter.Filter(item) {
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
	})
	queryData := contracts.CompiledQueryable[T]{
		Items:     items,
		Operators: initiateOperator,
	}

	return &CollectionCompiledQueryable[T]{
		queryData,
		contracts.PageOption{
			Limit: -1,
			Skip:  -1,
		},
	}
}
func (op *CollectionCompiledQueryable[T]) Where(function func(T) bool) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		Filter:       contracts.Filterer[T]{Function: function},
	})
	return op
}
func (op *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: AnyCollection,
		Filter: contracts.Filterer[T]{
			Function: function,
		},
	})
	return &AssertCompiledQueryable[T]{
		op.CompiledQueryable,
	}
}

func (op *CollectionCompiledQueryable[T]) Update(Updater func(T) T) *CollectionCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: UpdateCollection,
		Update:       contracts.Updater[T]{Function: Updater},
	})

	return op

}
