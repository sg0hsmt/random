package random_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/sg0hsmt/random"
)

func testSample[T comparable](t *testing.T, in []T) {
	t.Helper()

	const count = 10000
	const threshold = 10 // means 10%
	const size = 5

	cnt := make(map[T]int, len(in))

	for i := 0; i < count; i++ {
		v, err := random.Sample(in, size)
		if err != nil {
			t.Fatalf("sample failed, count %d, %v", i+1, err)
		}

		if len(v) != size {
			t.Fatalf("invalid size, count %d, want %d, got %d", i+1, size, len(v))
		}

		found := make(map[T]struct{}, size)

		for i := range v {
			if _, ok := found[v[i]]; ok {
				t.Fatalf("non unique value, count %d, value %v", i+1, v[i])
			}

			found[v[i]] = struct{}{}
			cnt[v[i]] = cnt[v[i]] + 1
		}
	}

	avg := count * size / len(in)
	min := avg - (avg / threshold)
	max := avg + (avg / threshold)

	for k, v := range cnt {
		if v < min {
			t.Errorf("value %v, count is too small", k)
		}
		if v > max {
			t.Errorf("value %v, count is too large", k)
		}
	}
}

func TestSample(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		testSample(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	t.Run("int (empty)", func(t *testing.T) {
		_, err := random.Sample([]int{}, 10)
		if !errors.Is(err, random.ErrEmptySlice) {
			t.Errorf("want %v, got %v", random.ErrEmptySlice, err)
		}
	})

	t.Run("int (negative size)", func(t *testing.T) {
		_, err := random.Sample([]int{1, 2, 3}, -1)
		if !errors.Is(err, random.ErrNegativeSize) {
			t.Errorf("want %v, got %v", random.ErrNegativeSize, err)
		}
	})

	t.Run("int (large size)", func(t *testing.T) {
		_, err := random.Sample([]int{1, 2, 3}, 5)
		if !errors.Is(err, random.ErrSizeTooLarge) {
			t.Errorf("want %v, got %v", random.ErrSizeTooLarge, err)
		}
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		testSample(t, []string{"a", "b", "c", "d", "e", "f"})
	})

	t.Run("string (empty)", func(t *testing.T) {
		_, err := random.Sample([]string{}, 10)
		if !errors.Is(err, random.ErrEmptySlice) {
			t.Errorf("want %v, got %v", random.ErrEmptySlice, err)
		}
	})

	t.Run("string (negative size)", func(t *testing.T) {
		_, err := random.Sample([]string{"a", "b", "c"}, -1)
		if !errors.Is(err, random.ErrNegativeSize) {
			t.Errorf("want %v, got %v", random.ErrNegativeSize, err)
		}
	})

	t.Run("string (large size)", func(t *testing.T) {
		_, err := random.Sample([]string{"a", "b", "c"}, 5)
		if !errors.Is(err, random.ErrSizeTooLarge) {
			t.Errorf("want %v, got %v", random.ErrSizeTooLarge, err)
		}
	})
}

func ExampleSample() {
	// Example for int slice.
	{
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		out, err := random.Sample(in, 3)
		if err != nil {
			fmt.Println("sample failed")
			return
		}

		fmt.Println(out) // print three unique numbers.
	}

	// Example for string slice.
	{
		in := []string{"a", "b", "c", "d", "e", "f"}

		out, err := random.Sample(in, 3)
		if err != nil {
			fmt.Println("sample failed")
			return
		}

		fmt.Println(out) // print three unique strings.
	}
}
