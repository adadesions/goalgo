package tree

import (
	"fmt"
	"testing"
)

var myTree = []Node{
	Node{Data: 1},
	Node{Data: 2},
	Node{Data: 3},
	Node{Data: 4},
	Node{Data: 5},
}

func TestInorder(t *testing.T) {
	want := []int{4, 2, 5, 1, 3}
	gotArray := []int{}

	myTree[0].Left = &myTree[1]
	myTree[0].Right = &myTree[2]
	myTree[1].Left = &myTree[3]
	myTree[1].Right = &myTree[4]

	ch := make(chan int)
	go func() {
		Inorder(&myTree[0], ch)
		close(ch)
	}()

	for {
		v, ok := <- ch
		if !ok {
			break
		}
		gotArray = append(gotArray, v)
		fmt.Printf("%d, ", v)
	}

	fmt.Println()

	for i, got := range gotArray {
		if want[i] != got {
			t.Fatalf("Error at Inorder: want %v but got %v", want[i], got)
		}
	}
}

func TestPreorder(t *testing.T) {
	want := []int{1, 2, 4, 5, 3}
	gotArray := []int{}

	myTree[0].Left = &myTree[1]
	myTree[0].Right = &myTree[2]
	myTree[1].Left = &myTree[3]
	myTree[1].Right = &myTree[4]

	ch := make(chan int)
	go func() {
		Preorder(&myTree[0], ch)
		close(ch)
	}()

	for {
		v, ok := <- ch
		if !ok {
			break
		}
		gotArray = append(gotArray, v)
		fmt.Printf("%d, ", v)
	}

	fmt.Println()

	for i, got := range gotArray {
		if want[i] != got {
			t.Fatalf("Error at Inorder: want %v but got %v", want[i], got)
		}
	}
}

func TestPostorder(t *testing.T) {
	want := []int{4, 5, 2, 3, 1}
	gotArray := []int{}

	myTree[0].Left = &myTree[1]
	myTree[0].Right = &myTree[2]
	myTree[1].Left = &myTree[3]
	myTree[1].Right = &myTree[4]

	ch := make(chan int)
	go func() {
		Postorder(&myTree[0], ch)
		close(ch)
	}()

	for {
		v, ok := <- ch
		if !ok {
			break
		}
		gotArray = append(gotArray, v)
		fmt.Printf("%d, ", v)
	}

	fmt.Println()

	for i, got := range gotArray {
		if want[i] != got {
			t.Fatalf("Error at Inorder: want %v but got %v", want[i], got)
		}
	}
}



