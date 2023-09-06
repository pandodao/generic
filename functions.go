package generic

// Must panics if err is not nil, otherwise returns v.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// Try returns v regardless of err.
func Try[T any](v T, err error) T {
	return v
}

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}
