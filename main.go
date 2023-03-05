package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

type Java_word struct {
	JavaName string `json:"java_name"`
	Sign     string `json:"sign"`
}

// func processing(s string) {
// 	var current_string string
// 	current_char, main_string := pop_front(s)
// 	for range main_string {
// 		main_string, value, flag := is_separator(current_char, main_string)
// 		if flag {
// 			is_keyword(current_string)
// 			set_separator(value)
// 			current_string = ""
// 		} else {
// 			current_string += current_char

// 		}
// 		current_char, main_string = pop_front(main_string)
// 	}
// }

// func finder(s string) bool { return separat(s) }

func is_separator(testing_char, main_string string) (string, string, bool) {
	flag, value := finder(testing_char, 2) //where 2 - separators
	if flag {
		//get separtor's num from json add to list
		if testing_char == " " { //add value condition
			new := strings.TrimLeft(main_string, " ")
			return new, value, flag
		} else {
			return main_string, value, flag
		}
	}
	flag, value = is_operator(testing_char)

	return main_string, value, flag
}

func is_operator(testing_char string) (bool, string) {

	return finder(testing_char, 0) //where 0 - operators
}

func make_list(s string) {
	err := os.WriteFile("result.txt", []byte(s), 0666)
	if err != nil {
		panic(err)
	}
}

// func separat(s string) bool {
// 	if div_find(s) {
// 		//add separatr to json
// 		return true
// 	} else {
// 		return is_keyword(s)
// 	}
// }

// func is_keyword(test_word string) {
// 	flag, value := finder(test_word, 1)
// 	if flag {
// 		//add value to list
// 	} else {
// 		set_ident(test_word)
// 	}

// }
func set_ident(test_string string) {
	count := gojsonq.New().File("identifiers.json").From("identifiers").Count()
	file, err := ioutil.ReadFile("identifiers.json")
	if err != nil {
		panic(err)
	}
	data := []Java_word{}
	json.Unmarshal(file, &data)

}

// func set_ident(char string) {
// 	//add char to list
// }

// func what_divid(s string) {
// }
func pop_front(s string) (char, str string) {
	return string(s[0]), string(s[1:])

}

func finder(testing string, num int) (bool, string) {
	jsons := []string{"operators", "keywords", "separators"}
	jq, err := gojsonq.New().File(jsons[num]+".json").WhereEqual("java_name", testing).PluckR("sign")
	if err != nil {
		panic(err)
	}
	res, _ := jq.StringSlice()
	if len(res) == 0 {
		return false, ""
	}
	return true, res[0]
}
func main() {
	// s := "    testing"
	// m := Message{"Test+_name", "W1"}
	// b, err := json.Marshal(m)
	// if err != nil {
	// 	panic(err)ssssssssss
	// }
	// jsonString := `[{"name1":"w1","name2":"w2","name3":"w3"}]`
	// b := []byte(`{"name1":"w1","name2":"w2","name3":"w3"}`)
	// var test []map[string]interface{}
	// err2 := json.Unmarshal([]byte(jsonString), &test)

	// fmt.Println(k)
	// file, err := os.Open("operators.json")
	// filename, err := os.Open("operators.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer filename.Close()
	// data, err := ioutil.ReadAll(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// var result []Message
	// jsonErr := json.Unmarshal(data, &result)
	// if jsonErr != nil {
	// 	panic(jsonErr)
	// }
	// fmt.Println(result)
	// for _, m := range result {
	// 	fmt.Println(m.JavaName, m.Sign)
	// }
	// jq := gojsonq.New().File("operators.json")
	// js, err := jq.WhereEqual("java_name", "for1").PluckR("sign")
	// fmt.Println(js)
	// if err != nil {
	// 	panic(err)
	// }
	// res, err := js.StringSlice()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(len(res) == 0)

	file, err := ioutil.ReadFile("identifiers.json")
	if err != nil {
		panic(err)
	}
	data := []Java_word{}
	newIdentifier := &Java_word{
		JavaName: "test3",
		Sign:     "first3",
	}
	json.Unmarshal(file, &data)
	data = append(data, *newIdentifier)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("identifiers.json", dataBytes, 0644)
	if err != nil {
		panic(err)
	}
	jq := gojsonq.New().File("identifiers.json").Count()
	fmt.Println(jq)
	// if err2 != nil {
	// 	fmt.Println(2, err2)
	// 	return
	// }
	// fmt.Println(&test)
	// for _, mes := range test {
	// 	fmt.Println(mes)
	// }
	// current_char, current_string := pop_front(s)
	// for range s[1:] {
	// 	fmt.Println(current_char, current_string)
	// 	current_char, current_string = pop_front(current_string)
	// }
}
