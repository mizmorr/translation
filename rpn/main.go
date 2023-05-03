package main

import (
	"fmt"
	dij "rpn/core"
)

func main() {
	// stack := stack.Get_Stack()
	fmt.Println(dij.Dijkstra([]string{"I1", "O8", "N1", "R4"}))
}
