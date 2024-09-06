/*
dict is map : )
*/

package dict

func Contain[T comparable, V any](d map[T]V, k T) bool {
	// 只读, 不需要检测d是否是nil, 不存在的键会返回值类型的零值以及false
	_, ok := d[k]
	return ok
}

func Keys[T comparable, V any](d map[T]V) []T {
	// 只读, 不需要检测d是否是nil, 可以对nil的map进行便利
	keys := make([]T, 0, len(d)) // make的cap参数可以是0, 所以当d是nil时, 也不会出问题
	for k := range d {
		keys = append(keys, k)
	}
	return keys
}

func Values[T comparable, V any](d map[T]V) []V { // 情况同Keys
	values := make([]V, 0, len(d))
	for _, v := range d {
		values = append(values, v)
	}
	return values
}
