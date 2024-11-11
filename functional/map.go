package functional

import (
	"sync"

	"github.com/zweix123/suger/common"
)

// copy from https://github.com/samber/lo/blob/master/slice.go
// Map manipulates a slice and transforms it to a slice of another type.
func MapSerial[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
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

func MapParallelWithGoroutineUpperLimit[T any, R any](collection []T, iteratee func(item T, index int) R, goroutineNum int) (result []R) {
	if goroutineNum <= 0 {
		goroutineNum = 1
	}

	result = make([]R, len(collection))

	goroutineLimit := make(chan struct{}, goroutineNum-1)
	functionResult := make(chan struct{}, goroutineNum-1) // Currently only used for synchronization, it can be expanded if needed in the future

	go func() {
		defer close(goroutineLimit)
		defer common.HandlePanic(func(_ string, _ int, _ any) {})

		for i, item := range collection {
			goroutineLimit <- struct{}{}
			go func(_item T, _i int) {
				defer common.HandlePanic(func(_ string, _ int, _ any) {})
				defer func() {
					functionResult <- struct{}{}
				}()
				result[_i] = iteratee(_item, _i)
			}(item, i)
		}
	}()

	for range goroutineLimit {
		<-functionResult
	}

	return result
}
