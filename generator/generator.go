package generator

import "github.com/wreulicke/gojg/ast"

import "fmt"
import "bufio"
import "errors"

func Generate(node ast.AST, writer *bufio.Writer) error {
	switch t := node.(type) {
	case *ast.BooleanNode:
		return writeBoolean(writer, t)
	case *ast.NumberNode:
		return writeNumber(writer, t)
	case *ast.NullNode:
		_, err := writer.WriteString("null")
		return err
	case *ast.StringNode:
		return writeString(writer, t)
	case *ast.ArrayNode:
		return writeArray(writer, t)
	case *ast.RawValueTemplateNode:
		return errors.New("not implemented")
	case *ast.ObjectNode:
		return writeObject(writer, t)
	default:
		return errors.New("unexpected node type")
	}
}

func writeBoolean(writer *bufio.Writer, node *ast.BooleanNode) error {
	if node.ID != nil {
		fmt.Println("not implemented")
		return errors.New("not implemented")
	}
	_, err := writer.WriteString(fmt.Sprint(node.Value))
	return err
}

func writeNumber(writer *bufio.Writer, node *ast.NumberNode) error {
	// do not print raw value? TODO discussion
	_, err := writer.WriteString(fmt.Sprint(node.Value))
	return err
}

func writeString(writer *bufio.Writer, node *ast.StringNode) error {
	if node.ID != nil {
		fmt.Println("not implemented")
		return errors.New("not implemented")
	}

	var err error
	if _, err = writer.WriteRune('"'); err != nil {
		return err
	}

	if _, err = writer.WriteString(node.Value); err != nil {
		return err
	}

	if _, err = writer.WriteRune('"'); err != nil {
		return err
	}

	return err
}

func writeArray(writer *bufio.Writer, node *ast.ArrayNode) error {
	_, err := writer.WriteRune('[')
	if err != nil {
		return err
	}
	if len(node.Value) > 0 {
		err := Generate(node.Value[0], writer)

		if err != nil {
			return err
		}

		for v := range node.Value[1:] {
			_, err := writer.WriteRune(',')
			if err != nil {
				return err
			}

			err = Generate(v, writer)
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

func writeObject(writer *bufio.Writer, node *ast.ObjectNode) error {
	_, err := writer.WriteRune('{')
	if err != nil {
		return err
	}
	if len(node.Members) > 0 {
		err := writeMember(writer, node.Members[0])
		if err != nil {
			return err
		}
		for _, v := range node.Members[1:] {
			_, err := writer.WriteRune(',')
			if err != nil {
				return err
			}
			err = writeMember(writer, v)
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

func writeMember(writer *bufio.Writer, node ast.AST) error {
	if v, ok := node.(*ast.MemberNode); ok {
		err := Generate(v.Name, writer)
		if err != nil {
			return err
		}

		_, err = writer.WriteRune(':')
		if err != nil {
			return err
		}

		err = Generate(v.Value, writer)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("unexpected object member")
}
