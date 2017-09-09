package parse

import (
	"errors"
	"strings"
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

//go:generate goyacc -o grammer.go grammer.y
func (l *Lexer) Error(e string) {
	l.error = errors.New(e)
}

func (l *Lexer) Lex(lval *yySymType) int {
	ruNe := l.Scan()
	token := int(ruNe)
	text := l.TokenText()
	if token == scanner.Int || token == scanner.Float {
		token = NUMBER
	} else if token == scanner.Ident {
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
		text = text[1 : len(text)-1]
		if strings.HasPrefix(text, "{{") && strings.HasSuffix(text, "}}") {
			text = text[2 : len(text)-2]
			token = STRING_TEMPLATE
		}
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
	lval.token = Token{typ: token, literal: text}
	// fmt.Println(lval.token)
	return token
}
