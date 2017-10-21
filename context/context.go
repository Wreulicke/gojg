package context

import (
	"bufio"
	"errors"
	"os"

	"github.com/wreulicke/gojg/ast"
	"github.com/wreulicke/gojg/parse"
)

type Resolver func(context Context) interface{}
type Context map[string]Resolver

func resolveValue(context Context, node ast.AST) interface{} {
	switch n := node.(type) {
	case *ast.BooleanNode:
		return n.Value
	case *ast.RawValueTemplateNode:
		return context[n.ID.Name](context)
	case *ast.StringNode:
		if n.ID != nil {
			return context[n.ID.Name](context)
		}
		return n.Value
	default:
		return n
	}
}

func resolver(context Context, key string, node ast.AST) {
	context[key] = func(context Context) interface{} {
		return resolveValue(context, node)
	}
}

func CreateContext(contextMap *map[string]string, contextFile **os.File) (Context, error) {
	context := make(Context, len(*contextMap))
	if *contextFile != nil {
		r := bufio.NewReader(*contextFile)
		node, e := parse.Parse(r)
		if e != nil {
			return nil, e
		}
		switch t := node.(type) {
		case *ast.ObjectNode:
			for _, v := range t.Members {
				name := v.Name
				if name.ID != nil {
					key := "{{" + name.ID.Name + "}}"
					resolver(context, key, v.Value)
				} else {
					resolver(context, name.Value, v.Value)
				}
			}
		default:
			return nil, errors.New("unexpected node type")
		}
	}

	for k, v := range *contextMap {
		resolver(context, k, v)
	}
	return context, nil
}
