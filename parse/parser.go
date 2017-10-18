package parse

import (
	"strings"

	"github.com/wreulicke/gojg/ast"
)

func Parse(str string) (ast.AST, *Error) {
	l := &Lexer{}
	l.Init(strings.NewReader(str))
	yyParse(l)
	return l.result, l.error
}
