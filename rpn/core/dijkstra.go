package core

import (
	stack "rpn/stack"
	"strings"
)

func dijkstra(input_str []string) (result string) {
	stack := stack.Get_Stack()
	for _, token := range input_str {
		if is_const(token) {
			result += token + " "
		} else {
			stack.Is_empty()
		}

	}
	return
}

func is_const(token string) bool {
	if strings.HasPrefix(token, "C") || strings.HasPrefix(token, "N") {
		return true
	}
	return false
}
