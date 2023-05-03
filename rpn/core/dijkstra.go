package core

import (
	"fmt"
	stack "rpn/stack"

	// "strconv"
	"strings"
	// "github.com/thedevsaddam/gojsonq"
)

func Dijkstra(input_str []string) (result string) {
	stack := stack.Get_Stack()
	var (
		class_num   int = 1
		nesting_num int = 1
		is_in_class bool
	)

	for _, token := range input_str {
		if is_const(token) {
			result += token + " "
		} else {
			// var is_pushed bool
			if token == "R9" && strings.HasPrefix(stack.Top().Name, "BC") {
				_, elem := stack.Pop()
				result += elem.Name
				is_in_class = true
				nesting_num++
			}
			if !stack.Is_empty() && strings.HasPrefix(stack.Top().Name, "class") {
				result += token + " "
				if flag, class_name := stack.Pop(); flag {
					_, need, _ := strings.Cut(class_name.Name, " ")
					stack.Push(fmt.Sprint("BC:", need), 10)
				}
			}
			if token == "R10" && nesting_num == 1 && is_in_class {
				class_num++
				result += fmt.Sprint("EC:", class_num, ",", 1, " ")
				is_in_class = false
				nesting_num = 1
			}
			//if func{
			// result+=token
			// stack<-1F

			// }
			if token == "W9" {
				stack.Push(fmt.Sprint("class ", class_num, ",", 1, " "), 10)
			}
			if token == "R11" {
				stack.Push("2AEA", 0)
			}
			if token == "R12" {
				if flag, oper := stack.Pop(); flag {
					result += oper.Name + " "
				}
			}
			if flag, name := is_init(token); flag {
				stack.Push("1:"+name, 8)
			}
			if token == "R4" {
				if !stack.Is_empty() {
					elems := stack.Pop_till(-1)
					for j := len(elems) - 1; j >= 0; j-- {
						result += elems[j].Name
					}
				}
			}
			priority, flag := get_priority(token)
			if flag {
				if !stack.Is_empty() && priority <= stack.Top().Priority {
					elems := stack.Pop_till(priority)
					for i := len(elems) - 1; i >= 0; i-- {
						result += elems[i].Name + " "
					}
				}
				stack.Push(token, priority)
			}
		}
	}
	if !stack.Is_empty() {
		elems := stack.Pop_till(-1)
		for j := len(elems) - 1; j >= 0; j-- {
			result += elems[j].Name
		}
	}
	return
}

func is_const(token string) bool {
	if strings.HasPrefix(token, "C") || strings.HasPrefix(token, "N") || strings.HasPrefix(token, "I") {
		return true
	}
	return false
}

func is_init(token string) (bool, string) {
	init_list := map[string]string{"W50": "STR", "W20": "FLT", "W27": "INT"}
	result, flag := init_list[token]
	return flag, result
}

// func get_priority(name string) (int, bool) {
// 	if gojsonq.New().File("data/priorities.json").Count() > 0 {
// 		jq, err := gojsonq.New().File("data/priorities.json").WhereEqual("element", name).PluckR("priority")
// 		if err != nil {
// 			panic(err)
// 		}
// 		pri, _ := strconv.Atoi(jq)
// 		return pri, true
// 	} else {
// 		return 0, false
// 	}
// }

// func is_operator(token string) bool {

// }
// func get_priority(str string) int {
// 	switch str {
// 	case "O1":

// 	}

// }

func get_priority(str string) (int, bool) {
	prities := map[string]int{"O1": 6, "O2": 6, "O22": 5, "O8": 2}
	if pri, ok := prities[str]; ok {
		return pri, true
	}
	return 0, false
}
