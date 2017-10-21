package generator

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/wreulicke/gojg/ast"
	"github.com/wreulicke/gojg/context"
)

type Generator interface {
	Generate(node ast.AST) error
}

type generatorImpl struct {
	context context.Context
	writer  *bufio.Writer
}

func (g *generatorImpl) Generate(node ast.AST) error {
	writer := g.writer
	switch t := node.(type) {
	case *ast.BooleanNode:
		return g.writeBoolean(t)
	case *ast.NumberNode:
		return g.writeNumber(t)
	case *ast.NullNode:
		_, err := writer.WriteString("null")
		return err
	case *ast.StringNode:
		return g.writeString(t)
	case *ast.ArrayNode:
		return g.writeArray(t)
	case *ast.RawValueTemplateNode:
		return g.writeRawValue(g.context, t)
	case *ast.ObjectNode:
		return g.writeObject(t)
	default:
		return errors.New("unexpected node type")
	}
}

func NewGenerator(context context.Context, writer *bufio.Writer) Generator {
	g := generatorImpl{context: context, writer: writer}
	return &g
}

func (g *generatorImpl) writeRawValue(context context.Context, node *ast.RawValueTemplateNode) error {
	if v, ok := context[node.ID.Name]; ok {
		value := v(context)
		err := g.Generate(value)
		if err != nil {
			_, e := g.writer.WriteString(fmt.Sprint(value))
			return e
		}
		return err
	}
	return errors.New("cannot resolve value: id=" + node.ID.Name)
}

func (g *generatorImpl) writeBoolean(node *ast.BooleanNode) error {
	_, err := g.writer.WriteString(fmt.Sprint(node.Value))
	return err
}

func (g *generatorImpl) writeNumber(node *ast.NumberNode) error {
	writer := g.writer
	// do not print raw value? TODO discussion
	_, err := writer.WriteString(fmt.Sprint(node.Value))
	return err
}

func (g *generatorImpl) writeString(node *ast.StringNode) error {
	writer := g.writer
	if node.ID != nil {
		if value, ok := g.context[node.ID.Name]; ok {
			return g.writeString(&ast.StringNode{Value: fmt.Sprint(value(g.context))})
		}
		return fmt.Errorf("value:%s is not found", node.ID.Name)
	}

	if _, err := writer.WriteString(fmt.Sprintf("%q", node.Value)); err != nil {
		return err
	}

	return nil
}

func (g *generatorImpl) writeArray(node *ast.ArrayNode) error {
	writer := g.writer
	_, err := writer.WriteRune('[')
	if err != nil {
		return err
	}
	if len(node.Value) > 0 {
		err := g.Generate(node.Value[0])

		if err != nil {
			return err
		}

		for v := range node.Value[1:] {
			_, err := writer.WriteRune(',')
			if err != nil {
				return err
			}

			err = g.Generate(v)
			if err != nil {
				return err
			}
		}
	}
	_, err = writer.WriteRune(']')
	if err != nil {
		return err
	}
	return nil
}

func (g *generatorImpl) writeObject(node *ast.ObjectNode) error {
	writer := g.writer
	_, err := writer.WriteRune('{')
	if err != nil {
		return err
	}
	if len(node.Members) > 0 {
		err := g.writeMember(node.Members[0])
		if err != nil {
			return err
		}
		for _, v := range node.Members[1:] {
			_, err := writer.WriteRune(',')
			if err != nil {
				return err
			}
			err = g.writeMember(v)
			if err != nil {
				return err
			}
		}
	}
	_, err = writer.WriteRune('}')
	if err != nil {
		return err
	}
	return nil
}

func (g *generatorImpl) writeMember(node ast.MemberNode) error {
	writer := g.writer
	err := g.Generate(node.Name)
	if err != nil {
		return err
	}

	_, err = writer.WriteRune(':')
	if err != nil {
		return err
	}
	err = g.Generate(node.Value)
	if err != nil {
		return err
	}

	return nil
}
