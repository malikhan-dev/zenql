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

	for _, item := range op.Collect() {
		result = append(result, mapper(item))
	}

	return result

}

func (op *CollectionCompiledQueryable[T]) CollectMapped[U any](mapper func(T) U) []U {
	return Project[T, U](op, mapper)
}
