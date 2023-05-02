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
		current_op     string
		is_complex_op  bool
	)
	current_char, main_string := pop_front(s)
	for current_char != "end" {
		main_string, value, flag = is_separator(current_char, main_string)
		if !is_sep_eq {
			if is_sep_eq = strings.ContainsAny(value, "O"); is_sep_eq {
				current_op = value
			}

		}
		if current_char == "\"" {
			is_quotes = !is_quotes
		}
		if main_string != "" && flag {
			{
				_, cur_value, cur_flag := is_separator(current_char+string(main_string[0]), "")

				if cur_flag && (string(cur_value[0]) == "O" || string(cur_value[0]) == "S") {
					flag = false
					is_complex_op = true
				}
				if len(current_char) > 1 && !cur_flag {
					flag = true
					is_complex_op = false
					current_op = value
				}
			}
		}
		if flag && !is_quotes {
			is_added := false
			if _, err := strconv.Atoi(current_string); err == nil {
				N := get_const(current_string)
				make_list(N, true)
				is_added = true
			}
			if !is_added && is_sep_eq && value != current_op && !strings.ContainsAny(current_string, "\"") {
				C := get_const(current_string)
				make_list(C, true)
				is_sep_eq, is_added = false, true
			}
			if !is_added && !is_complex_op && strings.Count(current_string, "\"") == 2 {
				C := get_const(current_string)
				make_list(C, true)
			} else {
				if !is_complex_op && current_string != "" && !is_added {
					word := get_word(current_string)
					make_list(word, true)

				}
			}
			if main_string != "" || mode {
				make_list(value, true)
			}
			current_string = ""
			is_complex_op = false

		} else {
			if !is_complex_op {
				current_string += current_char
			}

		}
		if !is_complex_op {
			current_char, main_string = pop_front(main_string)
		} else {
			char, main := pop_front(main_string)
			current_char += char
			main_string = main
		}
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
	if flag, ident := finder(const_value, 3); flag {
		return ident
	}
	if flag, key_word := finder(const_value, 1); flag {
		return key_word
	}
	if const_value == "" {
		return ""
	}
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
	json_s := []string{"identifiers", "numeric_const", "symbol_const", "keywords", "operators", "separators"}
	file, err := os.ReadFile("data/" + json_s[num] + ".json")
	if err != nil {
		panic(err)
	}
	data := []Java_word{}
	json.Unmarshal(file, &data)
	if num == 4 {
		for _, val := range data {
			if strings.ContainsAny(val.Sign, "O") {
				cur := []string{val.Sign, val.JavaName}
				result = append(result, cur)
			}
		}

	} else {
		for _, val := range data {
			cur := []string{val.Sign, val.JavaName}
			result = append(result, cur)
		}
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

func GET() {
	file, err := os.ReadFile("data/keywords.json")
	if err != nil {
		panic(err)
	}
	result := []string{}
	data := []Java_word{}
	json.Unmarshal(file, &data)
	for _, val := range data {
		cur := val.JavaName
		result = append(result, cur)
	}
	f, err := os.OpenFile("t.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, s := range result {
		if _, err = f.WriteString(s + "\n"); err != nil {
			panic(err)
		}
	}

}
