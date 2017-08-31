package parse

import (
	"errors"
	"text/scanner"

	"github.com/wreulicke/gojg/ast"
)

type Token struct {
	typ     int
	literal string
}

type Lexer struct {
	scanner.Scanner
	result ast.AST
	error  error
}

func (l *Lexer) Error(e string) {
	l.error = errors.New(e)
}

func (l *Lexer) Lex(lval *yySymType) int {
	ruNe := l.Scan()
	token := int(ruNe)
	if token == scanner.Int || token == scanner.Float {
		token = NUMBER
	} else if token == scanner.Ident {
		text := l.TokenText()
		if text == "bool" {
			token = BOOLEAN_PREFIX
		} else if text == "false" {
			token = FALSE
		} else if text == "null" {
			token = NULL
		} else if text == "true" {
			token = TRUE
		} else {
			token = ID
		}
	} else if token == scanner.String {
		token = STRING
	} else if ruNe == '{' {
		if l.Peek() == '{' {
			l.Next()
			token = TEMPLATE_BEGIN
		}
	} else if ruNe == '}' {
		if l.Peek() == '}' {
			l.Next()
			token = TEMPLATE_END
		}
	}
	lval.token = Token{typ: token, literal: l.TokenText()}
	// fmt.Println(lval.token)
	return token
}
