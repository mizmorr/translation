package main

import (
	"fmt"
	stack "rpn/stack"
)

func main() {
	stack := stack.Get_Stack()
	fmt.Println(stack.Is_empty())
	stack.Push("O31", 1)
	stack.Push("W21", 2)
}
