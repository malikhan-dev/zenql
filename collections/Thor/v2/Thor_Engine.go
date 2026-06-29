package collections

import (
	"container/heap"
	"github.com/malikhan-dev/zenql/contracts/v2"
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
	SkipCollection     = 8
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

	skipLimit, takeLimit := extractLimits(op.Operators)

	skipCount := 0
	count := 0

	for _, item := range *op.Items {

		keep := true
		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)
			if !keep {
				break
			}
		}

		hasTake := takeLimit != -1
		hasSkip := skipLimit != -1

		if keep {
			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}

			if hasTake {
				if len(result) == takeLimit {
					return result
				}
				result = append(result, item)
				count++

			} else {
				result = append(result, item)
				count++
			}
		}
	}
	return result
}

func (op *CollectionCompiledQueryable[T]) CollectSorted(less func(T, T) bool, desc bool) []T {

	HeapInitializer := NewSortable[T](less, desc)
	heap.Init(HeapInitializer)

	skipLimit, takeLimit := extractLimits(op.Operators)

	skipCount := 0
	count := 0

	for _, item := range *op.Items {

		keep := true

		for _, op := range op.Operators {

			keep = CoreFilter(op, item)

			if !keep {
				break
			}

		}

		hasTake := takeLimit != -1
		hasSkip := skipLimit != -1

		if keep {
			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}

			if hasTake {
				if HeapInitializer.Len() == takeLimit {
					break
				}
				heap.Push(HeapInitializer, item)
				count++
			} else {
				heap.Push(HeapInitializer, item)
				count++
			}

		}

	}

	result := contracts.AllocateSlice[T](len(*op.Items))

	for HeapInitializer.Len() > 0 {

		item := heap.Pop(HeapInitializer).(T)

		result = append(result, item)

	}
	return result
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

func (op *GroupCompiledQueryable[K, T]) Collect() *GroupedQueryable[K, T] {

	var result GroupedQueryable[K, T]

	result.Items = contracts.AllocateMap[K, T](len(*op.Items))

	skipLimit, takeLimit := extractLimits(op.Operators)

	var LocatedKey K

	skipCount := 0
	count := 0

	for _, item := range *op.Items {

		LocatedKey = op.PropLocator(item)

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		hasTake := takeLimit != -1
		hasSkip := skipLimit != -1

		if keep {
			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}
			if hasTake {
				if len(result.Items) == takeLimit {
					return &result
				}
				result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
				count++

			} else {
				result.Items[LocatedKey] = append(result.Items[LocatedKey], item)
				count++
			}

		}

	}

	return &result
}

func Project[T any, M any](op *CollectionCompiledQueryable[T], mapper func(T) M) []M {
	var result []M
	result = contracts.AllocateSlice[M](len(*op.Items))

	skipLimit, takeLimit := extractLimits(op.Operators)

	skipCount := 0
	count := 0

	for _, item := range *op.Items {

		keep := true
		for _, operator := range op.Operators {
			keep = CoreFilter(operator, item)
			if !keep {
				break
			}
		}

		hasTake := takeLimit != -1
		hasSkip := skipLimit != -1

		if keep {

			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}

			if hasTake {
				if len(result) == takeLimit {
					return result
				}

				result = append(result, mapper(item))
				count++

			} else {
				result = append(result, mapper(item))
				count++
			}

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

func (op *CollectionCompiledQueryable[T]) Skip(count int) *CollectionCompiledQueryable[T] {
	op.Operators = append(op.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SkipCollection,
		Skip:         count,
	})
	return op
}
