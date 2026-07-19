package collections

import (
	"container/heap"

	"github.com/malikhan-dev/zenql/contracts/v2"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

func (op *CollectionCompiledQueryable[T]) CollectUpdated(Updater func(T) T) []T {
	var result []T
	result = contracts.AllocateSlice[T](len(*op.Items))

	skipLimit, takeLimit := op.Page.Skip, op.Page.Limit

	var skipCount, count int32
	skipCount = 0
	count = 0

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
				if len(result) == int(takeLimit) {
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

	skipLimit, takeLimit := op.Page.Skip, op.Page.Limit

	var skipCount, count int32
	skipCount = 0
	count = 0

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
				if HeapInitializer.Len() == int(takeLimit) {
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
func Project[T any, M any](op *CollectionCompiledQueryable[T], mapper func(T) M) []M {
	var result []M
	result = contracts.AllocateSlice[M](len(*op.Items))

	skipLimit, takeLimit := op.Page.Skip, op.Page.Limit

	var skipCount, count int32
	skipCount = 0
	count = 0

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
				if len(result) == int(takeLimit) {
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
