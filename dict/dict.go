/*
dict is map : )
*/

package dict

func Contain[T comparable, V any](d map[T]V, k T) bool {
	_, ok := d[k]
	return ok
}

func Keys[T comparable, V any](d map[T]V) []T {
	keys := make([]T, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}

func Values[T comparable, V any](d map[T]V) []V {
	values := make([]V, 0, len(d))
	for _, v := range d {
		values = append(values, v)
	}
	return values
}
