package functional

// copy from https://github.com/samber/lo/blob/master/slice.go
// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i := range collection {
		initial = accumulator(initial, collection[i], i)
	}

	return initial
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
