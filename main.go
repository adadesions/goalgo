package main

import (
	"fmt"

	"github.com/adadesions/goalgo/stack"
	f "github.com/adadesions/goalgo/fibonacci"
)

func main() {
	s := stack.Stack{}
	s.Push(100)
	s.Push(200)
	fmt.Println(s.Data)

	for i := 0; i < 100; i++ {
		fmt.Printf("%d, ", f.DyFibo(i))
	}

}
