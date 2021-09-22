package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateQuene(t *testing.T) {
	q := Quene{Data: []int{1, 2, 3}}
	want := []int{1, 2, 3}
	got := q.Data

	if !cmp.Equal(want, got) {
		t.Fatalf("Error create quene: want %v but got %v", want, got)
	}
}

func TestEnquene(t *testing.T) {
	q := Quene{Data: []int{1, 2, 3}}
	want := 1
	got, err := q.Enquene()

	if err != nil {
		t.Fatalf("Error at Enquene: %v", err)
	}

	if want != got {
		t.Fatalf("Error Enquene: want %v but got %v", want, got)
	}	
	
	want = 2
	got, _ = q.Enquene()

	if want != got {
		t.Fatalf("Error Enquene: want %v but got %v", want, got)
	}
}

func TestEnqueneEmpty(t *testing.T) {
	q := Quene{Data: []int{}}
	want := -1
	gotVal, gotErr := q.Enquene()

	if gotErr == nil {
		t.Fatalf("Need Error: Quene is empty.")
	}

	if want != gotVal {
		t.Fatalf("Error at Enquene: want %v but got %v", want, gotVal)
	}
}