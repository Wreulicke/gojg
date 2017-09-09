package parse

import (
	"reflect"
	"testing"

	"github.com/wreulicke/gojg/ast"
)

func TestParseString(t *testing.T) {
	r := mustParse(t, `"test"`)
	if v, ok := r.(*ast.ValueNode); ok {
		t.Logf("value: %s", v)
	} else {
		t.Fatalf("%s(type: %s) is not value node", r, getTypeName(r))
	}
}

func TestParseStringTemplate(t *testing.T) {
	r := mustParse(t, `"{{test}}"`)
	if v, ok := r.(*ast.ValueNode); ok { // TODO more fluent type
		t.Logf("value: %s", v)
		if v.Id != "test" {
			t.Fatalf("expected: test, actual %s", v.Id)
		}
	} else {
		t.Fatalf("%s(type: %s) is not value node", r, getTypeName(r))
	}
}

func TestParseBool(t *testing.T) {
	mustParse(t, "bool(test)")
	mustParse(t, "true")
	mustParse(t, "false")
}

func TestParse(t *testing.T) {

}

func mustParse(t *testing.T, src string) ast.AST {
	r, err := Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func getTypeName(o interface{}) string {
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}
