package tree

import (
	"fmt"
	"testing"
)

var rootNode Node = Node{Data: 92}
var leftNode Node = Node{Data: 101}
var rightNode Node = Node{Data: 102}

func TestCreateStack(t *testing.T) {
	s := Stack{Data: []Node{rootNode, leftNode, rightNode}}
	want := []Node{rootNode, leftNode, rightNode}
	gotNodes := s.Data

	for i, node := range gotNodes {
		if want[i].Data != node.Data {
			t.Fatalf("Error create Stack: want %v but got %v", want[i].Data, node.Data)
		}
	}
}

func TestPop(t *testing.T) {
	s := Stack{Data: []Node{rootNode, leftNode, rightNode}}
	want := rightNode.Data
	got, err := s.Pop()

	if err != nil {
		t.Fatalf("Error Pop: %v", err)
	}

	if want != got.Data {
		t.Fatalf("Error Pop: want %v but got %v", want, got)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := Stack{Data: []Node{}}
	want := Node{}
	got, gotErr := s.Pop()

	if gotErr == nil {
		t.Fatalf("Need Error at this Pop Empty Stackt: %v", gotErr)
	}

	if want != got {
		t.Fatalf("Error Pop: want -1 but got %v", got)
	}
}

func TestPushStack(t *testing.T) {
	s := Stack{Data: []Node{}}
	want := []Node{rootNode, rightNode, leftNode}

	s.Push(&rootNode)
	s.Push(&rightNode)
	s.Push(&leftNode)

	got := s.Data

	for i, ele := range want {
		if ele.Data != got[i].Data {
			t.Fatalf("Error PushStack: want %v but got %v", ele.Data, got[i].Data)
		}

		fmt.Printf("%d = %d\n", ele.Data, got[i].Data)
	}
}
