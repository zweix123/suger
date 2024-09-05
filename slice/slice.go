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
