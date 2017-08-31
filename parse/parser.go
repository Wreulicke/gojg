package parse

import (
	"fmt"
	"strings"

	"github.com/wreulicke/gojg/ast"
)

func Parse(str string) ast.AST {
	l := new(Lexer)
	l.Init(strings.NewReader(str))
	yyParse(l)
	if l.error != nil {
		fmt.Println(l.error)
		panic(l.error)
	}
	return l.result
}
