package collections

import (
	"container/heap"
	"context"
	"sort"

	"github.com/malikhan-dev/zenql/contracts/v2"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

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

func (op *CollectionCompiledQueryable[T]) CollectUpdated(Updater func(T) T) []T {
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
				result = append(result, Updater(item))
				count++

			} else {
				result = append(result, Updater(item))
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

func (op *CollectionCompiledQueryable[T]) FindParentNode(NodeLocator func(T) bool, Criteria func(child T, parent T) bool) T {

	var result []T

	var TargetNode T

	result = contracts.AllocateSlice[T](len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if keep {

			if NodeLocator(item) {
				TargetNode = item
			}
			result = append(result, item)

		}
	}

	var value T
	for _, val := range result {
		if Criteria(TargetNode, val) {
			value = val
			break
		}
	}

	return value
}

func (op *CollectionCompiledQueryable[T]) FindRootNode(Start func(T) bool, Link func(child T, parent T) bool, Less func(T, T) bool) T {

	var result []T

	var TargetNode T

	result = contracts.AllocateSlice[T](len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if keep {

			if Start(item) {
				TargetNode = item
			}
			result = append(result, item)

		}
	}

	sort.Slice(result, func(i, j int) bool {
		return Less(result[j], result[i])
	})

	for _, val := range result {

		if Link(TargetNode, val) {
			TargetNode = val
		}

	}

	return TargetNode
}

func (op *CollectionCompiledQueryable[T]) TraverseRootNode(Start func(T) bool, Link func(child T, parent T) bool, Less func(T, T) bool, ctx context.Context) <-chan T {

	var result []T

	var TargetNode T

	result = contracts.AllocateSlice[T](len(*op.Items))

	for _, item := range *op.Items {

		keep := true

		for _, operator := range op.Operators {

			keep = CoreFilter(operator, item)

			if !keep {
				break
			}
		}

		if keep {

			if Start(item) {
				TargetNode = item
			}
			result = append(result, item)

		}
	}

	sort.Slice(result, func(i, j int) bool {
		return Less(result[j], result[i])
	})

	out := make(chan T, 1)

	go func() {

		for _, val := range result {

			if Link(TargetNode, val) {
				select {
				case <-ctx.Done():
					break
				case out <- val:
					TargetNode = val
				}
			}
		}
		defer close(out)

	}()

	return out
}
