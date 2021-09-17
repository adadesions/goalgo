package fibonacci

import "testing"

func TestDyFibo10FirstTerms(t *testing.T) {
	result := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i, val := range result {
		want := val
		got := DyFibo(i)

		if want != got {
			t.Fatalf("want %v but got %v\n", want, got)
		}
	}
}
