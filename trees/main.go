package main

import (
	"fmt"

	tree "github.com/adadesions/goalgo/trees/tree"
)

func main() {
	rootNode := tree.Node{Data: 92}
	left1 := tree.Node{Data: 100}
	right1 := tree.Node{Data: 108}
	
	rootNode.Left = &left1
	rootNode.Right = &right1

	fmt.Println(rootNode)

	

}
