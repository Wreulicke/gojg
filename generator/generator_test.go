package generator

import "testing"
import "github.com/wreulicke/gojg/ast"

import "bytes"
import "bufio"
import "github.com/wreulicke/gojg/context"

func TestGenerateNumber(t *testing.T) {
	mustGenerateAndTest(t, &ast.NumberNode{Value: 1}, "1")
	mustGenerateAndTest(t, &ast.NumberNode{Value: -1}, "-1")
	mustGenerateAndTest(t, &ast.NumberNode{Value: 1.1}, "1.1")
}

func TestGenerateObject(t *testing.T) {
	mustGenerateAndTest(t, new(ast.ObjectNode), "{}")
	mustGenerateAndTest(t, &ast.ObjectNode{Members: []ast.MemberNode{}}, `{}`)
	mustGenerateAndTest(t, &ast.ObjectNode{
		Members: []ast.MemberNode{
			ast.MemberNode{
				Name:  &ast.StringNode{Value: "hogehoge"},
				Value: &ast.StringNode{Value: "hogehoge"},
			},
		}}, `{"hogehoge":"hogehoge"}`)
	mustGenerateAndTest(t, &ast.ObjectNode{
		Members: []ast.MemberNode{
			ast.MemberNode{
				Name:  &ast.StringNode{Value: "hogehoge"},
				Value: &ast.StringNode{Value: "hogehoge"},
			},
			ast.MemberNode{
				Name:  &ast.StringNode{Value: "vvv"},
				Value: &ast.NumberNode{Value: 1.1},
			},
		}}, `{"hogehoge":"hogehoge","vvv":1.1}`)
}

func mustGenerateAndTest(t *testing.T, str ast.AST, expected string) {
	buffer := new(bytes.Buffer)
	writer := bufio.NewWriter(buffer)
	g := NewGenerator(context.Context{}, writer)
	err := g.Generate(str)
	if err != nil {
		t.Error(err)
		return
	}
	writer.Flush()

	actual := buffer.String()
	if actual != expected {
		t.Errorf(`
			actual: %s
			expected: %s
		`, actual, expected)
	}
}
