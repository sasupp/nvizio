package generic

func Filter[T any](ss []T, fn func(T) bool) (ret []T) {
	for _, s := range ss {
		if fn(s) {
			ret = append(ret, s)
		}
	}
	return
}

func First[T any](ss []T, fn func(T) bool) (ret []T) {
	for _, s := range ss {
		if fn(s) {
			ret = append(ret, s)
			return
		}
	}
	return
}

func All[T any](ss *[]T, fn func(T)) {
	for _, s := range *ss {
		fn(s)
	}
	return
}
