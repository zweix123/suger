package functional

import (
	"sync"

	"github.com/zweix123/suger/common"
)

// copy from https://github.com/samber/lo/blob/master/slice.go
// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i := range collection {
		result[i] = iteratee(collection[i], i)
	}

	return result
}

// copy from https://github.com/samber/lo/blob/master/parallel/slice.go
// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
// tips:
// 1. If the execution time of iterate is very short and slower than synchronization, because creating goroutines takes time.
// 2. If the length of the collection is very large, there may be problems. This function is not a coroutine pool
func MapParallel[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) { // nolint
			defer wg.Done()
			defer common.HandlePanic(func(_ string, _ int, _ any) {})
			res := iteratee(_item, _i)

			result[_i] = res
		}(item, i)
	}

	wg.Wait()

	return result
}

// copy from https://github.com/samber/lo/blob/master/slice.go
// Filter iterates over elements of collection, returning an array of all elements predicate returns truthy for.
func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice {
	result := make(Slice, 0, len(collection))

	for i := range collection {
		if predicate(collection[i], i) {
			result = append(result, collection[i])
		}
	}

	return result
}
