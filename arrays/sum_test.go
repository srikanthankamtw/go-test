package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("should be right for normal cases", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}
		if !slices.Equal(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("should be right for edge cases", func(t *testing.T) {
		got := SumAll([]int{0, 0})
		want := []int{0}
		if !slices.Equal(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should be right for normal cases", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should be right for edge cases", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
