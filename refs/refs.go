package refs

import (
	"github.com/wreulicke/gojg/ast"
)

func Refs(node ast.AST) []string {
	refsMap := make(map[string]struct{})
	refsInternal(node, &refsMap)

	keys := make([]string, 0, len(refsMap))
	for k := range refsMap {
		keys = append(keys, k)
	}
	return keys
}

func refsInternal(node ast.AST, refs *map[string]struct{}) {
	switch t := node.(type) {
	case *ast.MemberNode:
		if t.Name.ID != nil {
			(*refs)[t.Name.ID.Name] = struct{}{}
		}
		refsInternal(t.Value, refs)
	case *ast.StringNode:
		if t.ID != nil {
			(*refs)[t.ID.Name] = struct{}{}
		}
	case *ast.ArrayNode:
		for _, v := range t.Value {
			refsInternal(v, refs)
		}
	case *ast.RawValueTemplateNode:
		if t.ID != nil {
			(*refs)[t.ID.Name] = struct{}{}
		}
	case *ast.ObjectNode:
		for _, v := range t.Members {
			if v.Name.ID != nil {
				(*refs)[v.Name.ID.Name] = struct{}{}
			}
			refsInternal(v.Value, refs)
		}
	}
}
