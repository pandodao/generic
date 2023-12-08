package generic

import (
	"cmp"
	"sort"
)

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

// Reverse reverses the elements of s.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func CopySlice[T any](s []T) []T {
	d := make([]T, len(s))
	copy(d, s)
	return d
}

func HashSlice[K comparable, T any](s []T, fn func(T) K) map[K]T {
	m := make(map[K]T, len(s))

	for idx := range s {
		v := s[idx]
		m[fn(v)] = v
	}

	return m
}

func MapSlice[T, V any](s []T, fn func(T) V) []V {
	d := make([]V, len(s))

	for i, t := range s {
		d[i] = fn(t)
	}

	return d
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func SortMapKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := MapKeys(m)
	sort.Slice(keys, func(i, j int) bool {
		return cmp.Less(keys[i], keys[j])
	})

	return keys
}

func IndexSlice[T comparable](s []T, v T) int {
	for i, t := range s {
		if t == v {
			return i
		}
	}

	return -1
}

func ContainSlice[T comparable](s []T, v T) bool {
	return IndexSlice(s, v) >= 0
}
