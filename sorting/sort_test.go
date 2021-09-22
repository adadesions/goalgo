package sorting

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	data := []int{64, 25, 12, 22, 11}
	want := []int{11, 12, 22, 25, 64}
	got := SelectSort(data)

	for i, v := range want {
		if got[i] != v {
			fmt.Println(got)
			t.Fatalf("Error: want %v but got %v", v, got[i])
		}

	}
}
