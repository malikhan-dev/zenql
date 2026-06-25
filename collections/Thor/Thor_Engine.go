package collections

import (
	"container/heap"
	"context"

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
func (collection *CollectionCompiledQueryable[T]) Where(function func(T) bool) *CollectionCompiledQueryable[T] {

	collection.Operators = append(collection.Operators, contracts.ZenqlOperator[T]{
		OperatorType: WhereCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "where",
			Function: function,
		},
	})
	return collection
}

func (collection *CollectionCompiledQueryable[T]) Any(function func(T) bool) *AssertCompiledQueryable[T] {

	collection.Operators = append(collection.Operators, contracts.ZenqlOperator[T]{
		OperatorType: AnyCollection,
		MetaData: contracts.OpData[T]{
			MetaData: "any",
			Function: function,
		},
	})
	return &AssertCompiledQueryable[T]{
		collection.CompiledQueryable,
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

func (collection *AssertCompiledQueryable[T]) Assert() bool {

	for _, item := range *collection.Items {

		for _, op := range collection.Operators {

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

func (collection *CollectionCompiledQueryable[T]) Collect() []T {
	var result []T

	result = contracts.AllocateSlice[T](len(*collection.Items))

	skipLimit, takeLimit := extractLimits(collection.Operators)
	hasTake := takeLimit != -1
	hasSkip := skipLimit != -1

	skipCount := 0
	count := 0

	ctx, cancel := context.WithCancel(context.Background())

	out := make(chan FilterChan[T], 2048)

	for value := range FilterArr[T](ctx, collection, out) {

		if value.Keep {
			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}

			if hasTake {
				if len(result) == takeLimit {
					cancel()
					return result
				}

				result = append(result, value.Item)
				count++

			} else {
				result = append(result, value.Item)
				count++
			}
		}
	}

	defer cancel()
	return result

}

func (collection *CollectionCompiledQueryable[T]) CollectSorted(less func(T, T) bool, desc bool) []T {

	HeapInitializer := NewSortable[T](less, desc)
	heap.Init(HeapInitializer)

	skipLimit, takeLimit := extractLimits(collection.Operators)

	skipCount := 0
	count := 0

	for _, item := range *collection.Items {

		keep := true

		for _, op := range collection.Operators {

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

	result := contracts.AllocateSlice[T](len(*collection.Items))

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

func (collection *GroupCompiledQueryable[K, T]) Collect() *GroupedQueryable[K, T] {

	var result GroupedQueryable[K, T]

	result.Items = contracts.AllocateMap[K, T](len(*collection.Items))

	skipLimit, takeLimit := extractLimits(collection.Operators)

	var LocatedKey K

	skipCount := 0
	count := 0

	for _, item := range *collection.Items {

		LocatedKey = collection.PropLocator(item)

		keep := true

		for _, operator := range collection.Operators {

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

func FilterArr[T any](ctx context.Context, input contracts.ZenqlOperable[T], out chan FilterChan[T]) <-chan FilterChan[T] {

	go func() {

		defer close(out)
		for _, item := range *input.Iterate() {

			keep := true

			for _, operator := range *input.IterateOperators() {
				keep = CoreFilter(operator, item)
				if !keep {
					break
				}
			}
			if keep {
				select {
				case <-ctx.Done():
					return
				case out <- FilterChan[T]{keep, item}:
				}
			}
		}

	}()
	return out

}

func Project[T any, M any](op *CollectionCompiledQueryable[T], mapper func(T) M) []M {

	var result []M
	result = contracts.AllocateSlice[M](len(*op.Items))

	skipLimit, takeLimit := extractLimits(op.Operators)

	skipCount := 0
	count := 0

	hasTake := takeLimit != -1
	hasSkip := skipLimit != -1

	ctx, cancel := context.WithCancel(context.Background())

	out := make(chan FilterChan[T], 4096)

	for value := range FilterArr[T](ctx, op, out) {

		if value.Keep {
			if skipCount == skipLimit {
				hasSkip = false
			}

			if hasSkip {
				skipCount++
				continue
			}

			if hasTake {
				if len(result) == takeLimit {
					cancel()
					return result
				}

				result = append(result, mapper(value.Item))
				count++

			} else {
				result = append(result, mapper(value.Item))
				count++
			}
		}
	}

	defer cancel()
	return result
}

func (collection *CollectionCompiledQueryable[T]) Take(count int) *CollectionCompiledQueryable[T] {
	collection.Operators = append(collection.Operators, contracts.ZenqlOperator[T]{
		OperatorType: TakeCollection,
		Limit:        count,
	})
	return collection
}

func (collection *CollectionCompiledQueryable[T]) Skip(count int) *CollectionCompiledQueryable[T] {
	collection.Operators = append(collection.Operators, contracts.ZenqlOperator[T]{
		OperatorType: SkipCollection,
		Skip:         count,
	})
	return collection
}
