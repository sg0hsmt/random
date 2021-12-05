package random

import "math/rand"

// Sample returns unique elements from slice.
// Similar of Python's random.sample().
// see: https://docs.python.org/3/library/random.html#random.sample
func Sample[T any](in []T, size int) ([]T, error) {
	if len(in) == 0 {
		return nil, ErrEmptySlice
	}

	if size < 0 {
		return nil, ErrNegativeSize
	}

	if len(in) < size {
		return nil, ErrSizeTooLarge
	}

	idx := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		idx[i] = i
	}

	rand.Shuffle(len(idx), func(i, j int) {
		idx[i], idx[j] = idx[j], idx[i]
	})

	out := make([]T, 0, size)
	for i := 0; i < size; i++ {
		out = append(out, in[idx[i]])
	}

	return out, nil
}
