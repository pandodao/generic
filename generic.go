package generic

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func Try[T any](v T, err error) T {
	return v
}

func Ptr[T any](v T) *T {
	return &v
}
