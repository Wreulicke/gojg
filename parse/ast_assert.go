package parse

import (
	"reflect"
	"testing"

	"github.com/wreulicke/gojg/ast"
)

type assertion func(t *testing.T, actual ast.AST)

func assertString(expected string) assertion {
	return func(t *testing.T, actual ast.AST) {
		if node, ok := actual.(*ast.StringNode); !ok {
			t.Errorf("unexpected node type. expected StringNode, but actual %s", getTypeName(actual))
		} else if node.Value != expected {
			t.Errorf("error: expected %s, but actual %s", expected, node.Value)
		}
	}
}

func assertHasMember(memberName string) assertion {
	return func(t *testing.T, actual ast.AST) {
		node, ok := actual.(*ast.ObjectNode)
		if !ok {
			t.Errorf("unexpected node type. expected ObjectNode, but actual %s", getTypeName(actual))
		}
		for _, m := range node.Members {
			if m.Name != nil && m.Name.Value == memberName {
				return
			}
		}
		t.Errorf("not found member:%s", memberName)
	}
}

func assertObject(assertions []assertion) assertion {
	return func(t *testing.T, actual ast.AST) {
		node, ok := actual.(*ast.ObjectNode)
		if !ok {
			t.Errorf("unexpected node type. expected ObjectNode, but actual %s", getTypeName(actual))
		}
		if len(assertions) != len(node.Members) {
			t.Errorf("unexpected array size. expected %d, but actual %d", len(assertions), len(node.Members))
			return
		}
		for i, m := range node.Members {
			assertions[i](t, m)
		}
	}
}

func assertArray(assertions []assertion) assertion {
	return func(t *testing.T, actual ast.AST) {
		node, ok := actual.(*ast.ArrayNode)
		if !ok {
			t.Errorf("unexpected node type. expected ArrayNode, but actual %s", getTypeName(actual))
		}
		if len(assertions) != len(node.Value) {
			t.Errorf("unexpected array size. expected %d, but actual %d", len(assertions), len(node.Value))
			return
		}
		for i, v := range node.Value {
			assertions[i](t, v)
		}
	}
}

func getTypeName(o interface{}) string {
	t := reflect.TypeOf(o)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}
