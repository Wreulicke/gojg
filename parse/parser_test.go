package parse

import (
	"bufio"
	"strings"
	"testing"

	"github.com/wreulicke/gojg/ast"
)

type result struct {
	v ast.AST
	e error
}

func TestParseString(t *testing.T) {
	mustParse(t, `"{{test}}"`)
	mustParse(t, `'{{test}}'`)
}

var multilineStringTests = []struct {
	code     string
	expected string
}{
	{`"test"`, "test"},
	{`'"test'`, `"test`},
	{`'test\r\n'`, "test\r\n"},
	{`"\"test"`, `"test`},
	{`'\'test'`, `'test`},
	{`"\u305f"`, "た"},
	{`"\uD867\uDE3D"`, "𩸽"},
	{"`test\r\n\\hogehoge`", "test\r\n\\hogehoge"},
}

func TestParseMultilineString(t *testing.T) {
	for _, test := range multilineStringTests {
		r := mustParse(t, test.code)
		if r.e == nil {
			assertString(test.expected)(t, r.v)
		} else {
			t.Errorf("unexpected parse failed")
		}
	}
}

func TestParseStringWithFailure(t *testing.T) {
	mustFailToParse(t, `"test`)
}

func TestParseBool(t *testing.T) {
	mustParse(t, "true")
	mustParse(t, "false")
}

func TestParseNull(t *testing.T) {
	r := mustParse(t, "null")
	if _, ok := r.v.(*ast.NullNode); !ok {
		t.Errorf("unexpected node type. expected NullNode, but actual %s", getTypeName(r.v))
	}
}

func TestParseRawValueTemplate(t *testing.T) {
	r := mustParse(t, "{{test}}")
	if node, ok := r.v.(*ast.RawValueTemplateNode); !ok {
		t.Errorf("unexpected node type. expected RawValueTemplateNode, but actual %s", getTypeName(r.v))
	} else if node.ID.Name != "test" {
		t.Errorf("unexpected referenced id. expected %s, but actulal %s", "test", node.ID)
	}
}

func TestParseNumber(t *testing.T) {
	mustParse(t, "1")
	mustParse(t, "-1")
	mustParse(t, "4.5")
	mustFailToParse(t, "[5, 4.]")
	mustFailToParse(t, "[5, 4..5]")
}

func TestParseArray(t *testing.T) {
	mustParse(t, "[]")
	mustParse(t, `["test", 1]`)
	mustParse(t, `["test", 1, ]`)
	mustParse(t, `['test', 1]`)
	mustParse(t, "[1, {{test}}]")
	mustParse(t, "[1, {{test}}, ]")
	mustParse(t, "[{{test}}, -1]")
	mustParse(t, `["{{test}}"]`)
	mustParse(t, `["{{test}}", ]`)
}

func TestParseObject(t *testing.T) {
	mustParse(t, `{}`)
	mustParse(t, `{test: 1}`)
	mustParse(t, `{'test': 1}`)
	mustParse(t, "{'test': `hogehoge\r\nhogehoge`}")
	mustParse(t, `{'"test': 1}`)
	mustParse(t, `{'"test': 1, }`)
	mustParse(t, `{"test": 1}`)
	mustParse(t, `{"test": -1}`)
	mustParse(t, `{"test": {{test}}}`)
	mustParse(t, `{"test": "{{test}}"}`)
	mustParse(t, `{
		"fuga": "{{test}}",
		"hoge": "{{test}}",
	}`)
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
	return result{v: r, e: nil}
}
