package arraymod

import "testing"

func TestAddElement(t *testing.T) {
	data := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4, 5}
	got := Add(data, 5)

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Fatalf("Error want %v but got %v", want[i], got[i])
		}
	}
}


func TestNewArrayList(t *testing.T) {
	want := &ArrayList{}
	got := NewArrayList()

	if want != got {
		t.Fatalf("Error want %v but got %v", *want, *got)
	}
}
