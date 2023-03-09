package parser

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

// TODO: fix float numbers reading
func processing(mode bool, s string) {
	var (
		current_string string
		value          string
		flag           bool
		is_sep_eq      bool
		is_quotes      bool
	)
	current_char, main_string := pop_front(s)
	for current_char != "end" {
		main_string, value, flag = is_separator(current_char, main_string)
		if !is_sep_eq {
			is_sep_eq = value == "O8"
		}
		if current_char == "\"" {
			is_quotes = !is_quotes
		}
		if flag && !is_quotes {
			// fmt.Println(is_quotes, main_string)
			is_added := false
			if is_sep_eq && value != "O8" && !strings.ContainsAny(current_string, "\"") {
				C := get_const(current_string)
				make_list(C, true)
				is_sep_eq, is_added = false, true
			}
			if strings.Count(current_string, "\"") == 2 {
				C := get_const(current_string)
				make_list(C, true)
			} else {
				if current_string != "" && !is_added {
					word := get_word(current_string)
					make_list(word, true)

				}
			}
			if main_string != "" || mode {
				make_list(value, true)
			}
			current_string = ""

		} else {
			current_string += current_char

		}
		current_char, main_string = pop_front(main_string)
	}
	make_list("", false)
}

func is_separator(testing_char, main_string string) (string, string, bool) {
	flag, value := finder(testing_char, 2) //where 2 - separators
	if flag {
		if value == "R1" {
			new := strings.TrimLeft(main_string, " ")
			return new, value, flag
		} else {
			return main_string, value, flag
		}
	}
	flag, value = is_operator(testing_char)

	if flag {
		return strings.TrimLeft(main_string, " "), value, flag
	}

	return main_string, value, flag
}

func is_operator(testing_char string) (bool, string) {

	return finder(testing_char, 0) //where 0 - operators
}

func make_list(s string, if_inline bool) {

	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	switch if_inline {
	case true:
		if _, err = f.WriteString(s + " "); err != nil {
			panic(err)
		}
	case false:
		if _, err = f.WriteString("\n"); err != nil {
			panic(err)
		}
	}

}

func get_word(test_word string) string {
	flag, value := finder(test_word, 1)
	if flag {
		return value
	}
	flag, value = finder(test_word, 3)
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
	count := gojsonq.New().File("data/" + json_s[num] + ".json").Count()
	file, err := os.ReadFile("data/" + json_s[num] + ".json")
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

	err = os.WriteFile("data/"+json_s[num]+".json", dataBytes, 0644)
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
	jsons := []string{"operators", "keywords", "separators", "identifiers"}
	if gojsonq.New().File("data/"+jsons[num]+".json").Count() > 0 {
		jq, err := gojsonq.New().File("data/"+jsons[num]+".json").WhereEqual("java_name", testing).PluckR("sign")
		if err != nil {
			panic(err)
		}
		res, _ := jq.StringSlice()
		if len(res) == 0 {
			return false, ""
		}
		return true, res[0]
	} else {
		return false, ""
	}

}
func Cleaner() {
	err, _, _, _ := os.WriteFile("result.txt", []byte(""), 06664), os.WriteFile("data/identifiers.json", []byte(""), 06664), os.WriteFile("data/numeric_const.json", []byte(""), 06664), os.WriteFile("data/symbol_const.json", []byte(""), 06664)
	if err != nil {
		panic(err)
	}
}
func Parse(s string) {
	if strings.HasSuffix(s, "\n") || strings.HasSuffix(s, "\t") || strings.HasSuffix(s, " ") {
		processing(true, s)
		return
	}
	processing(false, s+" ")
}
func Get_data(num int) (result [][]string) {
	json_s := []string{"identifiers", "numeric_const", "symbol_const"}
	file, err := os.ReadFile("data/" + json_s[num] + ".json")
	if err != nil {
		panic(err)
	}
	data := []Java_word{}
	json.Unmarshal(file, &data)
	for _, val := range data {
		cur := []string{val.Sign, val.JavaName}
		result = append(result, cur)
	}
	return
}
func Get_result() string {
	file, err := os.ReadFile("result.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
