package main

import (
	"fmt"

	"github.com/adadesions/goalgo/stack"
)

func main() {
	s := stack.Stack{}
	s.Push(100)
	s.Push(200)
	fmt.Println(s.Data)

}
