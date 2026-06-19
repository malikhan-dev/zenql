package collections

import (
	"container/heap"

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
	TakeCollection     = 7
)

func From[T any](items *[]T) *CollectionCompiledQueryable[T] {

	initiateOperator := make([]contracts.ZenqlOperator[T], 0)
	initiateOperator = append(initiateOperator, contracts.ZenqlOperator[T]{
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

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "where",
			Function: function,
		},
	})
	return op
}

func (op *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
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

	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
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

func CoreFilter[T any](Operator contracts.ZenqlOperator[T], item T) bool {

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
	result = contracts.AllocateSlice[T](len(*op.Items))

	takeLimit := -1
	for _, operator := range op.Operators {
		if operator.OperatorType == TakeCollection {
			takeLimit = operator.Limit
			break
		}
	}

	count := 0
	for _, item := range *op.Items {
		if takeLimit != -1 && count >= takeLimit {
			break
		}

		keep := true
		for _, operator := range op.Operators {
			if operator.OperatorType == TakeCollection {
				continue
			}
			keep = CoreFilter(operator, item)
			if !keep {
				break
			}
		}

		if keep {
			result = append(result, item)
			count++
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

	result := contracts.AllocateSlice[T](len(*op.Items))

	for HeapInitializer.Len() > 0 {

		item := heap.Pop(HeapInitializer).(T)

		result = append(result, item)

	}
	return result
}

func (op *GroupCompiledQueryable[K, T]) Collect() *GroupedQueryable[K, T] {
	var result GroupedQueryable[K, T]
	result.Items = contracts.AllocateMap[K, T](len(*op.Items))

	hasDistinct := false
	for _, operator := range op.Operators {
		if operator.OperatorType == DistinctCollection {
			hasDistinct = true
			break
		}
	}

    var seen map[any]struct{}
    if hasDistinct {
        capacity := contracts.Guard(contracts.Alloc[any](len(*op.Items)))
        seen = make(map[any]struct{}, capacity)
    }

	for _, item := range *op.Items {
		keep := true
		for _, operator := range op.Operators {
			if operator.OperatorType == DistinctCollection {
				continue
			}
			keep = CoreFilter(operator, item)
			if !keep {
				break
			}
		}

		if keep {
			if hasDistinct {
				if _, exists := seen[any(item)]; exists {
					continue
				}
				seen[any(item)] = struct{}{}
			}
			LocatedKey := op.PropLocator(item)
			result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
		}
	}
	return &result
}

func Project[T any, M any](op *CollectionCompiledQueryable[T], mapper func(T) M) []M {
	var result []M
	result = contracts.AllocateSlice[M](len(*op.Items))

	takeLimit := -1
	for _, operator := range op.Operators {
		if operator.OperatorType == TakeCollection {
			takeLimit = operator.Limit
			break
		}
	}

	count := 0
	for _, item := range *op.Items {
		if takeLimit != -1 && count >= takeLimit {
			break
		}

		keep := true
		for _, operator := range op.Operators {
			if operator.OperatorType == TakeCollection {
				continue
			}
			keep = CoreFilter(operator, item)
			if !keep {
				break
			}
		}

		if keep {
			result = append(result, mapper(item))
			count++
		}
	}
	return result
}

func (op *CollectionCompiledQueryable[T]) Take(count int) *CollectionCompiledQueryable[T] {
	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: TakeCollection,
		Limit:        count,
	})
	return op
}
