package parse

import (
	"bufio"
	"strings"
	"testing"

	"github.com/wreulicke/gojg/ast"
)

type result struct {
	t *testing.T
	v ast.AST
	e error
}

func TestParseString(t *testing.T) {
	mustParse(t, `"test"`)
	mustParse(t, `"{{test}}"`)
}

func TestParseStringWithFailure(t *testing.T) {
	mustFailToParse(t, `"test`)
}

func TestParseBool(t *testing.T) {
	mustParse(t, "true")
	mustParse(t, "false")
}

func TestParseNull(t *testing.T) {
	mustParse(t, "null")
}

func TestParseNumber(t *testing.T) {
	mustParse(t, "{{test}}")
	mustParse(t, "1")
	mustParse(t, "-1")
	mustParse(t, "4.5")
	mustFailToParse(t, "[5, 4.]")
	mustFailToParse(t, "[5, 4..5]")
}

func TestParseArray(t *testing.T) {
	mustParse(t, "[]")
	mustParse(t, `["test", 1]`)
	mustParse(t, "[1, {{test}}]")
	mustParse(t, "[{{test}}, -1]")
	mustParse(t, `["{{test}}"]`)
}

func TestParseObject(t *testing.T) {
	mustParse(t, `{}`)
	mustParse(t, `{test: 1}`)
	mustParse(t, `{"test": 1}`)
	mustParse(t, `{"test": -1}`)
	mustParse(t, `{"test": {{test}}}`)
	mustParse(t, `{"test": "{{test}}"}`)
	mustParse(t, `{"{{xxx}}": "{{test}}"}`)
	mustParse(t, `{"test": [
		1,
		-1,
		3.5,
		{{test}},
		"test",
		"{{test}}",
		true,
		false,
		null
	]}`)
}

func mustFailToParse(t *testing.T, src string) (ast.AST, error) {
	reader := bufio.NewReader(strings.NewReader(src))
	r, err := Parse(reader)
	if err != nil {
		return r, err
	}
	t.Errorf("unexpected to parse successfully. result: %v, src: %s", r, src)
	return r, err
}

func mustParse(t *testing.T, src string) result {
	reader := bufio.NewReader(strings.NewReader(src))
	r, err := Parse(reader)
	if err != nil {
		t.Errorf("error: %v, src: %s", err, src)
		return result{e: err}
	}
	return result{t: t, v: r, e: nil}
}
