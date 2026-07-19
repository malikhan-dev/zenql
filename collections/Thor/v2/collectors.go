package collections

import (
	"github.com/malikhan-dev/zenql/contracts/v2"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

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
