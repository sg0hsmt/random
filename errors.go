package random

import "errors"

// ErrEmptySlice slice is empty.
var ErrEmptySlice = errors.New("slice is empty")

// ErrSizeTooLarge size is greater than slice.
var ErrSizeTooLarge = errors.New("size is greater than slice")

// ErrNegativeSize size is negative.
var ErrNegativeSize = errors.New("size is negative")
