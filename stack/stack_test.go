package stack

import (
	"testing"
)

func TestPop(t *testing.T) {
	s := Stack{Data: []int{1, 2, 3}}
	want := 3
	got := s.Pop()

	if want != got {
		t.Fatalf("POP Testing: want %d but got %d\n", want, got)
	}

	t.Logf("Passed: POP")
}

func TestPopAll(t *testing.T) {
	s := Stack{Data: []int{1, 2, 3}}
	want := []int{3, 2, 1}

	for i := 0; i < len(s.Data)-1; i++ {
		got := s.Pop()
		if got != want[i] {
			t.Fatalf("want %d but got %d", want[i], got)
		}
	}
}

func TestPush(t *testing.T) {
	s := Stack{Data: []int{1, 2, 3}}
	want := []int{1, 2, 3, 4}
	s.Push(4)

	for i := 0; i < len(s.Data)-1; i++ {
		got := s.Data[i]
		if want[i] != got {
			t.Fatalf("want %v but got %v", want, got)
		}
	}
}

func TestSeek(t *testing.T) {
	s := Stack{Data: []int{1, 2, 3}}
	want := 3
	got := s.Seek()

	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}