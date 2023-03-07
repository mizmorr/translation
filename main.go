package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

type Java_word struct {
	JavaName string `json:"java_name"`
	Sign     string `json:"sign"`
}

func processing(s string) {
	var (
		current_string string
		value          string
		flag           bool
	)
	is_sep_eq, is_quotes := false, false
	current_char, main_string := pop_front(string_maker(s))
	for current_char != "end" {
		main_string, value, flag = is_separator(current_char, main_string)
		if !is_sep_eq {
			is_sep_eq = value == "O8"
		}
		if current_char == "\"" {
			is_quotes = !is_quotes
		}
		if flag {
			if !is_quotes {
				if current_string != "" && string(current_string[len(current_string)-1]) == "\"" {
					C := get_const(current_string)
					make_list(C)
					break
				}
				if is_sep_eq && value != "O8" {
					C := get_const(current_string)
					make_list(C)
					is_sep_eq = false
				}
				if current_string != "" {
					word := get_word(current_string)
					make_list(word)
				}

				make_list(value)
				current_string = ""
			} else {
				current_string += current_char
			}
		} else {
			current_string += current_char

		}
		current_char, main_string = pop_front(main_string)
	}
}

func is_separator(testing_char, main_string string) (string, string, bool) {
	flag, value := finder(testing_char, 2) //where 2 - separators
	if flag {
		//get separtor's num from json add to list
		if value == "R1" { //add value condition
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

	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(" " + s + " "); err != nil {
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

func get_word(test_word string) string {
	flag, value := finder(test_word, 1)
	if flag {
		return value
	}
	return appender(test_word, 0)
}
func get_const(const_value string) string {
	_, err := strconv.ParseFloat(const_value, 32)
	if err != nil {
		return appender(const_value, 2)
	}
	return appender(const_value, 1)
}

func appender(added_string string, num int) string {
	json_s := []string{"identifiers", "numeric_const", "symbol_const"}
	symbol := []string{"I", "N", "C"}
	count := gojsonq.New().File(json_s[num] + ".json").Count()
	file, err := os.ReadFile(json_s[num] + ".json")
	if err != nil {
		panic(err)
	}
	current_sign := symbol[num] + fmt.Sprint(count+1)
	data := []Java_word{}
	json.Unmarshal(file, &data)
	if num == 2 {
		added_string = string(added_string[1 : len(added_string)-1])
	}
	newIdentifier := &Java_word{
		JavaName: added_string,
		Sign:     current_sign,
	}
	data = append(data, *newIdentifier)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(json_s[num]+".json", dataBytes, 0644)
	if err != nil {
		panic(err)
	}
	return current_sign
}

func pop_front(s string) (char, str string) {
	if s == "" {
		return "end", ""
	}
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
func cleaner() {
	err, _, _, _ := os.WriteFile("result.txt", []byte(""), 06664), os.WriteFile("identifiers.json", []byte(""), 06664), os.WriteFile("numeric_const.json", []byte(""), 06664), os.WriteFile("symbol_const.json", []byte(""), 06664)
	if err != nil {
		panic(err)
	}
}
func string_maker(s string) string {
	if strings.HasSuffix(s, "\n") || strings.HasSuffix(s, "\t") {
		return s
	}
	return s + " "
}

func main() {
	cleaner()
	// s := "string k=\"1\""
	// s2 := "public static void main(String[] args) {"
	s2 := "\"hello world\" boolean"
	// s3 := "System.out.println(\"Hello World\");"
	// file, err := os.ReadFile("operators.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, k := range file {
	// 	fmt.Print(string(k))
	// }
	processing(s2)
	// test := true
	// fmt.Println(test)
	// test = !test
	// fmt.Println(test)
}
