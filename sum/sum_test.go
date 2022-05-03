package sum_test

import (
	"sum"
	"testing"
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
