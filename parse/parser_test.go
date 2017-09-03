package parse

import (
	"testing"

	"github.com/wreulicke/gojg/ast"
)

func TestParseString(t *testing.T) {
	if v, ok := Parse(`test`).(*ast.ValueNode); ok {
		t.Logf("value: %s", v)
	} else {
		t.Fatal("is not value node")
	}
}
