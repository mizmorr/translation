package stack

import (
	"fmt"
	"strconv"
)

type Stack interface {
	Pop()
	Push()
	Is_empty() bool
}

type Operator struct {
	Name     string
	Priority int
}
type Stack_operators struct {
	elements map[int]*Operator
}

func (stack *Stack_operators) Is_empty() bool {
	return len(stack.elements) == 0
}
func (stack *Stack_operators) last() *Operator {
	return stack.elements[len(stack.elements)-1]
}
func (stack *Stack_operators) Pop() (bool, *Operator) {
	if stack.Is_empty() {
		return false, nil
	}
	operator := stack.last()
	delete(stack.elements, len(stack.elements)-1)
	return true, operator
}
func (stack *Stack_operators) Push(name string, priority int) {
	if stack.Is_empty() {
		stack.elements[0] = &Operator{Name: name, Priority: priority}

	} else {
		stack.elements[len(stack.elements)] = &Operator{Name: name, Priority: priority}
	}
}
func Get_Stack() *Stack_operators {
	return &Stack_operators{elements: map[int]*Operator{}}
}
func (stack *Stack_operators) Increase_func_id() {
	for i := len(stack.elements) - 1; i >= 0; i-- {
		if string(stack.elements[i].Name[len(stack.elements[i].Name)-1]) == "F" {
			if new_name, err := strconv.Atoi(string(stack.elements[i].Name[0 : len(stack.elements[i].Name)-1])); err != nil {
				return
			} else {
				stack.elements[i].Name = fmt.Sprint(new_name+1) + "F"
				return
			}

		}
	}
}

func (stack *Stack_operators) Print() {
	for i := len(stack.elements) - 1; i >= 0; i-- {
		fmt.Println(stack.elements[i])
	}
}

func (stack *Stack_operators) Pop_till(priority int) (result []*Operator) {
	for {
		if flag, oper := stack.Pop(); flag && oper.Priority >= priority {
			result = append(result, oper)
		} else {
			return
		}
	}
}

func (stack *Stack_operators) Top() *Operator {
	if !stack.Is_empty() {
		return stack.last()
	}
	return nil
}
