package parse

import (
	"testing"

	"github.com/wreulicke/gojg/ast"
)

func TestParseString(t *testing.T) {
	r, err := Parse(`test`)
	if err == nil {
		t.Fatalf("cannot parsed %s", r)
	} else if v, ok := r.(*ast.ValueNode); ok {
		t.Logf("value: %s", v)
	} else {
		t.Fatal("is not value node")
	}
}
