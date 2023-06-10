package ptr

func Of[T any](t T) (ret *T) {
	ret = &t
	return
}

func Safe[T any](t *T) (ret T) {
	if t == nil {
		return ret
	}

	return *t
}

func Must[T any](t *T) (ret T) {
	if t == nil {
		panic("nil pointer")
	}

	return *t
}
