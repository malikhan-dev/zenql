package collections

import (
	"github.com/malikhan-dev/zenql/contracts/v2"
)

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

// Project can be use to map a type to a destination type. its a function that requires a mapping function to the destination type
func Project[T any, M any](op *CollectionCompiledQueryable[T], mapper func(T) M) []M {

	var result []M

	result = contracts.AllocateSlice[M](len(*op.Items))

	for _, item := range op.Collect() {
		result = append(result, mapper(item))
	}

	return result

}

// Project can be use to map a type to a destination type. its a generic method that requires a mapping function to the destination type
func (op *CollectionCompiledQueryable[T]) Project[U any](mapper func(T) U) []U {
	return Project[T, U](op, mapper)
}
