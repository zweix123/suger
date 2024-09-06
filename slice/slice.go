package slice

func Contain[T comparable](l []T, e T) bool {
	// 只读, 不需要检测l是否是nil, 可以对nil的slice进行便利
	for _, v := range l {
		if v == e {
			return true
		}
	}
	return false
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Reverse[T any](src []T) (dst []T) {
	dst = make([]T, len(src))
	for i, j := 0, len(src)-1; i <= j; i, j = i+1, j-1 {
		dst[i], dst[j] = src[j], src[i]
	}
	return
}
