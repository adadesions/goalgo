package main

import (
	"fmt"
	"github.com/adadesions/goalgo/array/arraymod"
)

func main() {
	data := []int{1}
	newData := arraymod.Add(data, 2)

	fmt.Printf("data: %v\nnewData: %v\n", data, newData)
}
