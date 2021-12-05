package random_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/sg0hsmt/random"
)

func testChoice[T comparable](t *testing.T, in []T) {
	t.Helper()

	const count = 10000
	const threshold = 10 // means 10%

	cnt := make(map[T]int, len(in))

	for i := 0; i < count; i++ {
		v, err := random.Choice(in)
		if err != nil {
			t.Fatalf("choice failed, count %d, %v", i+1, err)
		}

		cnt[v] = cnt[v] + 1
	}

	avg := count / len(in)
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

func TestChoice(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		testChoice(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	t.Run("int (empty)", func(t *testing.T) {
		_, err := random.Choice([]int{})
		if !errors.Is(err, random.ErrEmptySlice) {
			t.Errorf("want %v, got %v", random.ErrEmptySlice, err)
		}
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		testChoice(t, []string{"a", "b", "c", "d", "e", "f"})
	})

	t.Run("string (empty)", func(t *testing.T) {
		_, err := random.Choice([]string{})
		if !errors.Is(err, random.ErrEmptySlice) {
			t.Errorf("want %v, got %v", random.ErrEmptySlice, err)
		}
	})
}

func ExampleChoice() {
	// Example for int slice.
	{
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		out, err := random.Choice(in)
		if err != nil {
			fmt.Println("choice failed")
			return
		}

		fmt.Println(out) // print one number.
	}

	// Example for string slice.
	{
		in := []string{"a", "b", "c", "d", "e", "f"}

		out, err := random.Choice(in)
		if err != nil {
			fmt.Println("choice failed")
			return
		}

		fmt.Println(out) // print one string.
	}
}

func testChoices[T comparable](t *testing.T, in []T) {
	t.Helper()

	const count = 10000
	const threshold = 10 // means 10%
	const size = 5

	cnt := make(map[T]int, len(in))

	for i := 0; i < count; i++ {
		v, err := random.Choices(in, size)
		if err != nil {
			t.Fatalf("choices failed, count %d, %v", i+1, err)
		}

		if len(v) != size {
			t.Fatalf("invalid size, count %d, want %d, got %d", i+1, size, len(v))
		}

		for i := range v {
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

func TestChoices(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		testChoices(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	t.Run("int (empty)", func(t *testing.T) {
		_, err := random.Choices([]int{}, 10)
		if !errors.Is(err, random.ErrEmptySlice) {
			t.Errorf("want %v, got %v", random.ErrEmptySlice, err)
		}
	})

	t.Run("int (negative size)", func(t *testing.T) {
		_, err := random.Choices([]int{1, 2, 3}, -1)
		if !errors.Is(err, random.ErrNegativeSize) {
			t.Errorf("want %v, got %v", random.ErrNegativeSize, err)
		}
	})

	t.Run("int (large size)", func(t *testing.T) {
		t.Parallel()
		testChoices(t, []int{1, 2, 3}) // size is 5
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		testChoices(t, []string{"a", "b", "c", "d", "e", "f"})
	})

	t.Run("string (empty)", func(t *testing.T) {
		_, err := random.Choices([]string{}, 10)
		if !errors.Is(err, random.ErrEmptySlice) {
			t.Errorf("want %v, got %v", random.ErrEmptySlice, err)
		}
	})

	t.Run("string (negative size)", func(t *testing.T) {
		_, err := random.Choices([]string{"a", "b", "c"}, -1)
		if !errors.Is(err, random.ErrNegativeSize) {
			t.Errorf("want %v, got %v", random.ErrNegativeSize, err)
		}
	})

	t.Run("string (large size)", func(t *testing.T) {
		t.Parallel()
		testChoices(t, []string{"a", "b", "c"}) // size is 5
	})
}

func ExampleChoices() {
	// Example for int slice.
	{
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		out, err := random.Choices(in, 3)
		if err != nil {
			fmt.Println("choices failed")
			return
		}

		fmt.Println(out) // print three numbers.
	}

	// Example for string slice.
	{
		in := []string{"a", "b", "c", "d", "e", "f"}

		out, err := random.Choices(in, 3)
		if err != nil {
			fmt.Println("choices failed")
			return
		}

		fmt.Println(out) // print three strings.
	}
}
