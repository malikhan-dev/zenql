package collections

/*
 * Author: Mohammadreza Malikhan
 * License: MIT
 */

func DeepCopy[T any](items []T) *[]T {

	destination := make([]T, len(items))

	destIndex := 0

	for _, v := range items {
		destination[destIndex] = v
		destIndex++
	}

	return &destination
}

func ToSlice[k int | float64 | float32 | int16 | int32 | uint | int64, v any](item map[k]v) *[]v {

	var destinationSlice []v

	destIndex := 0

	destinationSlice = make([]v, len(item))

	for _, val := range item {
		destinationSlice[destIndex] = val
		destIndex++
	}
	return &destinationSlice

}

func Page(allItems int, pageSize int) int {

	return int((allItems + pageSize - 1) / pageSize)

}
