package slice

import "testing"

func TestSlice(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Slice(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of 5 numbers", func(t *testing.T) {
		n := []int{1, 2, 3, 4, 5}
		want := 15
		got := Slice(n)
		if want != got {
			t.Errorf("want %d but got %d given, %v", want, got, n)
		}
	})

}
