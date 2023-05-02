package stack

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
		stack.elements[len(stack.elements)-1] = &Operator{Name: name, Priority: priority}
	}
}
func Get_Stack() *Stack_operators {
	return &Stack_operators{elements: map[int]*Operator{}}
}
