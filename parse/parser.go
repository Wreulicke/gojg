package parse

import (
	"bufio"

	"github.com/wreulicke/gojg/ast"
)

func Parse(reader *bufio.Reader) (ast.AST, error) {
	l := &Lexer{}
	l.Init(reader)
	yyParse(l)
	return l.result, l.error
}
