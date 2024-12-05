package slice

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
// Chunk returns an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
func Chunk[T any, Slice ~[]T](collection Slice, size int) []Slice {
	if size <= 0 {
		return nil //! In order not to return an error, it is compatible with unreasonable input parameters here, but the returned value is empty
	}

	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}

	result := make([]Slice, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last:last])
	}

	return result
}

// copy from https://github.com/samber/lo/blob/master/slice.go
// Flatten returns an array a single level deep.
// Same effect as ```
//
//	Reduce(collection, func(agg []T, item Slice, _ int) []T {
//		return append(agg, item...)
//	}, []T{})
//
// ```, but reduces a lot of copies
func Flatten[T any, Slice ~[]T](collection []Slice) Slice {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make(Slice, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}
