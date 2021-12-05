package random

import "math/rand"

// Choice return a random element from slice.
// Similar of Python's random.choice().
// see: https://docs.python.org/3/library/random.html#random.choice
func Choice[T any](in []T) (T, error) {
	if len(in) == 0 {
		var v T // zero value
		return v, ErrEmptySlice
	}

	return in[rand.Intn(len(in))], nil
}

// Choices returns multiple elements from slice.
// Similar of Python's random.choices(), But wight is not support.
// see: https://docs.python.org/3/library/random.html#random.choices
func Choices[T any](in []T, size int) ([]T, error) {
	if len(in) == 0 {
		return nil, ErrEmptySlice
	}

	if size < 0 {
		return nil, ErrNegativeSize
	}

	out := make([]T, 0, size)
	for i := 0; i < size; i++ {
		v, _ := Choice(in)
		out = append(out, v)
	}

	return out, nil
}
