package contracts

import (
	"runtime"
	"unsafe"
)

var Max_Alloc_Guard = 5000000

func SetMaxAllocGuard(max int) {
	Max_Alloc_Guard = max
}
func Alloc[T any](itemCount int) int {

	if itemCount <= 0 {
		return 0
	}

	if itemCount < 1024 {
		return itemCount
	}

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	freeHeap := m.HeapSys - m.HeapAlloc

	var t T

	itemSize := unsafe.Sizeof(t)

	estimated := uint64(itemCount) * uint64(itemSize)

	if estimated > freeHeap/2 {
		return itemCount / 4
	}
	return itemCount / 2
}

func Guard(EstimatedSize int) int {
	if EstimatedSize > Max_Alloc_Guard {
		EstimatedSize = Max_Alloc_Guard
	}
	return EstimatedSize
}

func AllocateSlice[T any](itemCount int) []T {
	return make([]T, 0, Guard(Alloc[T](itemCount)))
}

func AllocateMap[K comparable, T any](keyCount int) map[K][]T {
	capacity := Guard(keyCount / 2)
	if keyCount < 100 {
		capacity = keyCount
	}
	return make(map[K][]T, capacity)
}
