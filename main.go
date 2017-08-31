package main

import (
	"fmt"

	Ast "github.com/wreulicke/gojg/ast"
	"github.com/wreulicke/gojg/parse"
)

func main() {
	parse.Parse("bool(test)")
	parse.Parse("true")
	parse.Parse("false")
	parse.Parse("null")
	parse.Parse(`"test"`)
	parse.Parse("{{test}}")
	parse.Parse(`"{{test}}"`)
	parse.Parse(`{}`)
	parse.Parse("[{{test}}]")
	parse.Parse(`[
		{{test}},
		"{{test}}",
		"test",
		2,
		3.5
	]`)
	ast := parse.Parse("false")
	switch t := ast.(type) {
	case *Ast.ValueNode:
		if t.Value == false {
			fmt.Println("5000兆円欲しい")
		} else {
			fmt.Println("test")
		}
	}
}
