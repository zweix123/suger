package slice

func Contain[T comparable](l []T, e T) bool {
	if len(l) == 0 { // l is nil or []T{}
		return false
	}
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
	for i, j := 0, len(src)-1; i < len(src); i, j = i+1, j-1 {
		dst[i] = src[j]
		dst[j] = src[i]
	}
	return
}
