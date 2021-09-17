package dp

import (
	"fmt"
	"testing"
)

func TestFibo(t *testing.T) {
	result := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, val := range result {
		want := val
		got := Fibo(i)

		if want != got {
			t.Fatalf("want %v but got %v\n", want, got)
		}
	}
}

func TestFibo2(t *testing.T) {
	result := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, val := range result {
		want := val
		got := Fibo2(i)

		if want != got {
			t.Fatalf("want %v but got %v\n", want, got)
		}

		fmt.Printf("%v, ", got)
	}
}

func TestFactorialTabulation(t *testing.T) {
	expects := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	for i, want := range expects {
		got := fac(i)

		if want != got {
			t.Fatalf("want %v but got %v\n", want, got)
		}
	}
}

func TestFactorialMemoization(t *testing.T) {
	expects := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	for i, want := range expects {
		got := fac2(i)

		if want != got {
			t.Fatalf("want %v but got %v\n", want, got)
		}

		fmt.Printf("got = %v\n", got)
	}
}

func TestMin(t *testing.T) {
	set1 := [3]int{-1, 0, 1}
	set2 := [4]int{-3, -2, -1, -4}
	set3 := [3]int{3, 2, 1}

	want := -1
	got := Min(set1[0], set1[1], set1[2])

	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}

	want = -4
	got = Min(set2[0], set2[1], set2[2], set2[3])

	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}

	want = 1
	got = Min(set3[0], set3[1], set3[2])

	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestUglyNumber(t *testing.T) {
	inputs := []int{7, 10, 15, 150}
	expects :=  []int{8, 12, 24, 5832}

	for i, val := range inputs {
		want := expects[i]
		got := UglyNumber(val)

		if want != got {
			t.Fatalf("want %v but got %v", want, got)
		}
	}
}


func TestBST(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputs := []int{0, 1, 3, 5, 8, 10, 99, 2}
	outputs := []bool{false, true, true, true, true, true, false, true}

	for i, input := range inputs{
		want := outputs[i]
		got := BST(data, input)

		if want != got {
			t.Fatalf("Error At %v: want %v but got %v", i, want, got)
		}
	}
}
