package functional

// copy from https://github.com/samber/lo/blob/master/slice.go
// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
func Times[T any](count int, iteratee func(index int) T) []T {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		result[i] = iteratee(i)
	}

	return result
}

func All[T any](collection []T, predicate func(item T, index int) bool) bool {
	for i := range collection {
		if !predicate(collection[i], i) {
			return false
		}
	}
	return true
}

func Any[T any](collection []T, predicate func(item T, index int) bool) bool {
	for i := range collection {
		if predicate(collection[i], i) {
			return true
		}
	}
	return false
}
