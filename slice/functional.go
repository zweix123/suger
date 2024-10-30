package slice

import "errors"

var ChunkErr = errors.New("size must be greater than 0")

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
