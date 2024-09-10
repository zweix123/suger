package functional

import "sync"

// Map manipulates a slice and transforms it to a slice of another type.
// `iteratee` is call in parallel. Result keep the same order.
func Map[T any, R any](collection []T, iteratee func(item T) R) []R {
	result := make([]R, len(collection))

	var wg sync.WaitGroup
	wg.Add(len(collection))

	for i, item := range collection {
		go func(_item T, _i int) {
			res := iteratee(_item)

			result[_i] = res

			wg.Done()
		}(item, i)
	}

	wg.Wait()

	return result
}
