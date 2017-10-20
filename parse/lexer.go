package parse

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/wreulicke/gojg/ast"
)

type Token struct {
	typ     int
	literal string
}

type Position struct {
	line   int
	column int
}

type Lexer struct {
	input    *bufio.Reader
	buffer   bytes.Buffer
	position *Position
	offset   int
	result   ast.AST
	error    error
}

const eof = -1

func (l *Lexer) Init(reader io.Reader) {
	l.input = bufio.NewReader(reader)
	l.position = &Position{}
}

//go:generate goyacc -o grammer.go grammer.y
func (l *Lexer) Error(e string) {
	message := fmt.Sprintf("%s in %d:%d.", e, (*l).position.line, (*l).position.column)
	err := errors.New(message)
	l.error = err
}

func (l *Lexer) parseFloat(str string) float64 {
	f64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		l.Error("unexpected number format error")
		return -1
	}
	return f64
}

func (l *Lexer) scanDigit(next rune) {
	if next == '0' {
		l.Error("unexpected digit '0'")
		return
	} else if isDigit(next) {
		next := l.Peek()
		for {
			if !isDigit(next) {
				break
			}
			l.Next()
			next = l.Peek()
		}
		next = l.Peek()
		if next == '.' {
			l.Next()
			next = l.Peek()
			if !isDigit(next) {
				l.Error("unexpected token: expected digits")
				return
			}
			for {
				if !isDigit(next) {
					break
				}
				l.Next()
				next = l.Peek()
			}
		}
		next = l.Peek()
		if next == 'e' || next == 'E' {
			l.Next()
			next := l.Peek()
			if next == '+' || next == '-' {
				l.Next()
			}
			next = l.Peek()
			if !isDigit(next) {
				l.Error("digit expected for number exponent")
			}
			l.Next()
			next = l.Peek()
			for {
				if !isDigit(next) {
					break
				}
				l.Next()
				next = l.Peek()
			}
		}
	} else {
		l.Error("error")
		return
	}
}

func isIdentifierPart(r rune) bool {
	return (unicode.IsLetter(r) || unicode.IsMark(r) || unicode.IsDigit(r) ||
		unicode.IsPunct(r)) && !strings.ContainsRune("{}[]():,", r)
}

func (l *Lexer) scanIdentifier() {
	next := l.Peek()
	if unicode.IsLetter(next) || next == '$' || next == '_' {
	} else {
		l.Error("expected identifier")
	}

	for next = l.Peek(); isIdentifierPart(next); {
		l.Next()
		text := l.buffer.String()
		if text == "bool" {
			return
		}
		next = l.Peek()
	}
}

func (l *Lexer) scanString() {
	for {
		switch next := l.Next(); {
		case next == '"':
			return
		case next == '\\':
			if strings.IndexRune(`"\/bfnrt`, l.Peek()) >= 0 {
				l.Next()
				break
			} else if r := l.Next(); r == 'u' {
				for i := 0; i < 4; i++ {
					if strings.IndexRune("0123456789ABDEFabcdef", l.Peek()) >= 0 {
						l.Next()
					} else {
						l.Error("expected 4 hexadecimal digits")
						return
					}
				}
			} else {
				l.Error("unsupported escape character")
				return
			}
		case unicode.IsControl(next):
			l.Error("cannot contain control characters in strings")
			return
		case next == eof:
			l.Error("unclosed string")
			return
		}
	}
}

func (l *Lexer) scanWhitespace() {
	ruNe := l.Peek()
	for unicode.IsSpace(ruNe) {
		l.Next()
		ruNe = l.Peek()
	}
	return
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) TokenText() string {
	return l.buffer.String()
}

func (l *Lexer) Scan() int {
retry:
	switch next := l.Next(); {
	case next == '"':
		l.scanString()
		text := l.TokenText()
		if len(text) < 2 {
			l.Error("expected 1 or more character")
		}
		if strings.HasPrefix(text, `"{{`) && strings.HasSuffix(text, `}}"`) {
			l.buffer.Reset()
			l.buffer.WriteString(text[3 : len(text)-3])
			return STRING_TEMPLATE
		}
		l.buffer.Reset()
		l.buffer.WriteString(text[1 : len(text)-1])
		return STRING
	case next == ',':
		return COMMA
	case next == ':':
		return COLON
	case next == '[':
		return ARRAY_BEGIN
	case next == ']':
		return ARRAY_END
	case next == '{':
		r := l.Peek()
		if r == '{' {
			l.Next()
			return TEMPLATE_BEGIN
		}
		return OBJECT_BEGIN
	case next == '}':
		r := l.Peek()
		if r == '}' {
			l.Next()
			return TEMPLATE_END
		}
		return OBJECT_END
	case next == '(':
		return BRACE_BEGIN
	case next == ')':
		return BRACE_END
	case next == '-':
		return MINUS
	default:
		if unicode.IsSpace(next) {
			l.scanWhitespace()
			l.buffer.Reset()
			goto retry
		} else if next == eof {
			return eof
		} else if isDigit(next) {
			l.scanDigit(next)
			return NUMBER
		}
		l.scanIdentifier()
		text := l.TokenText()
		if text == "bool" {
			return BOOLEAN_PREFIX
		} else if text == "false" {
			return FALSE
		} else if text == "null" {
			return NULL
		} else if text == "true" {
			return TRUE
		}
		return ID
	}
}

func (l *Lexer) Next() rune {
	r, w, err := l.input.ReadRune()
	if err == io.EOF {
		return eof
	}
	if r == '\n' {
		l.position = &Position{line: l.position.line + 1}
	}
	l.position.column += w
	l.offset += w
	l.buffer.WriteRune(r)
	return r
}

func (l *Lexer) Peek() rune {
	lead, err := l.input.Peek(1)
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.Error("unexpected input error")
		return 0
	}

	p, err := l.input.Peek(runeLen(lead[0]))

	if err == io.EOF {
		return eof
	} else if err != nil {
		l.Error("unexpected input error")
		return 0
	}

	ruNe, _ := utf8.DecodeRune(p)
	return ruNe
}

func runeLen(lead byte) int {
	if lead < 0xC0 {
		return 1
	} else if lead < 0xE0 {
		return 2
	} else if lead < 0xF0 {
		return 3
	}
	return 4
}

// Lex Create Lexer
func (l *Lexer) Lex(lval *yySymType) int {
	if l.error != nil {
		return -1
	}
	typ := l.Scan()
	text := l.TokenText()
	lval.token = Token{typ: typ, literal: text}
	l.buffer.Reset()
	return typ
}
