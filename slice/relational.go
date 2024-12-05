package slice

// copy from https://github.com/samber/lo/blob/master/slice.go
// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
func GroupBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T) U) map[U]Slice {
	result := map[U]Slice{}

	for i := range collection {
		key := iteratee(collection[i])

		result[key] = append(result[key], collection[i])
	}

	return result
}
