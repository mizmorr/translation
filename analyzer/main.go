package main

import (
	design "gojson/design"
	parse "gojson/parser"
)

func main() {
	parse.Cleaner()
	design.Design()
}
