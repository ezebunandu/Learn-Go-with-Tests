package sum_test

import (
	"sum"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := sum.Sum(numbers)
		want := 6
		if want != got {
			t.Errorf("want %d, got %d, given %#v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := sum.SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !cmp.Equal(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, want, got []int) {
		t.Helper()
		if !cmp.Equal(want, got) {
			t.Errorf("want %v, got %v", got, want)
		}

	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := sum.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, want, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := sum.SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, want, got)
	})
}
