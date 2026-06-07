package Thor

import (
	"container/heap"

	"github.com/malikhan-dev/zenql/collections"
	"github.com/malikhan-dev/zenql/contracts"
)

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
)

func From[T any](items *[]T) *CollectionCompiledQueryable[T] {

	initiateOperator := make([]contracts.ZenqOperator[T], 0)
	initiateOperator = append(initiateOperator, contracts.ZenqOperator[T]{
		OperatorType: FromItems,
		MetaData: contracts.OpData[T]{
			MetaData: "from",
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

	op.Operators = append(op.Operators, contracts.ZenqOperator[T]{
		OperatorType: WhereCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "where",
			Function: function,
		},
	})
	return op
}

func (op *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqOperator[T]{
		OperatorType: AnyCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "any",
			Function: function,
		},
	})
	return &AssertCompiledQueryable[T]{
		op.CompiledQueryable,
	}
}

func Group[K comparable, T any](op *CollectionCompiledQueryable[T], locator func(T) K) *GroupCompiledQueryable[K, T] {

	op.Operators = append(op.Operators, contracts.ZenqOperator[T]{
		OperatorType: GroupCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "group",
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

	Any := false

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
	return Any
}

func CoreFilter[T any](Operator contracts.ZenqOperator[T], item T) bool {

	ShouldKeep := true

	switch Operator.OperatorType {

	case FromItems:
		if !Operator.MetaData.Function(item) {
			ShouldKeep = false
			break
		}

	case WhereCollection:
		if !Operator.MetaData.Function(item) {
			ShouldKeep = false
			break
		}

	}
	return ShouldKeep
}

func (op *CollectionCompiledQueryable[T]) Collect() []T {

	var result []T

	result = make([]T, 0, len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, op := range op.Operators {

			keep = CoreFilter(op, item)

			if !keep {
				break
			}

		}

		if keep {

			result = append(result, item)

		}

	}
	return result
}

func (op *CollectionCompiledQueryable[T]) CollectSorted(less func(T, T) bool, desc bool) []T {

	HeapInitializer := NewSortable[T](less, desc)
	heap.Init(HeapInitializer)

	for _, item := range *op.Items {

		keep := true

		for _, op := range op.Operators {

			keep = CoreFilter(op, item)

			if !keep {
				break
			}

		}

		if keep {
			heap.Push(HeapInitializer, item)
		}

	}

	result := make([]T, 0)

	for HeapInitializer.Len() > 0 {

		item := heap.Pop(HeapInitializer).(T)

		result = append(result, item)

	}
	return result
}

func (op *GroupCompiledQueryable[K, T]) Collect() *collections.GroupedQueryable[K, T] {

	var result collections.GroupedQueryable[K, T]

	result.Items = make(map[K][]T)

	var LocatedKey K

	for _, item := range *op.Items {

		LocatedKey = op.PropLocator(item)

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if !keep {
			continue
		}

		result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
	}

	return &result
}
