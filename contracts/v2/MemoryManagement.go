package contracts

import (
	"runtime"
	"sync/atomic"
	"time"
	"unsafe"
)

var maxAllocGuard = 5000000

func SetMaxAllocGuard(max int) {
	maxAllocGuard = max
}

type HeapSnapshot struct {
	freeHeap   uint64
	capturedAt int64
}

var cachedSnapshot atomic.Pointer[HeapSnapshot]

const snapshotTTL = 1000 * time.Millisecond

func getFreeHeap() uint64 {
	now := time.Now().UnixNano()

	if s := cachedSnapshot.Load(); s != nil {
		if time.Duration(now-s.capturedAt) < snapshotTTL {
			return s.freeHeap
		}
	}

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	snapshot := &HeapSnapshot{
		freeHeap:   m.HeapSys - m.HeapAlloc,
		capturedAt: now,
	}
	cachedSnapshot.Store(snapshot)

	return snapshot.freeHeap
}

func Alloc[T any](itemCount int) int {
	if itemCount <= 0 {
		return 0
	}

	if itemCount < 1024 {
		return itemCount
	}

	freeHeap := getFreeHeap()

	var t T
	itemSize := unsafe.Sizeof(t)
	estimated := uint64(itemCount) * uint64(itemSize)

	if estimated > freeHeap/2 {
		return itemCount / 2
	}

	return itemCount
}

func Guard(estimatedSize int) int {
	if estimatedSize > maxAllocGuard {
		estimatedSize = maxAllocGuard
	}
	return estimatedSize
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
