package tree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNode(t *testing.T) {
	want := Node{Data: 1234, Left: nil, Right: nil}
	got := Node{}
	got.Data = 1234
	got.Left = nil
	got.Right = nil

	if !cmp.Equal(want, got) {
		t.Fatalf("Error Node: want %v but got %v", want, got)
	}

}

func TestLeftNode(t *testing.T) {
	nodes := [7]Node{}

	for i := 0; i < 7; i++ {
		nodes[i] = Node{Data: i, Left: nil, Right: nil}
	}

	nodes[0].Left = &nodes[1]
	nodes[0].Right = &nodes[2]

	nodes[1].Left = &nodes[3]
	nodes[1].Right = &nodes[4]

	nodes[2].Left = &nodes[5]
	nodes[2].Right = &nodes[6]

	// Test Left Level 1
	want := &nodes[1]
	got := nodes[0].Left

	if want != got {
		t.Fatalf("Error Left Node LV1: want %v but got %v", want, got)
	}

	// Test Left Level 2.1
	want = &nodes[3]
	got = nodes[1].Left

	if want != got {
		t.Fatalf("Error Left Node LV2.1: want %v but got %v", want, got)
	}

	// Test Left Level 2.2
	want = &nodes[5]
	got = nodes[2].Left

	if want != got {
		t.Fatalf("Error Left Node LV2.2: want %v but got %v", want, got)
	}
}
