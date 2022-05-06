package main

import (
	"fmt"

	"github.com/ssksameer56/Generics/stacks"
)

func main() {
	stack := stacks.Stack[int]{}
	err := stack.Push(3)
	if err != nil {
		fmt.Println(err)
	}
}
