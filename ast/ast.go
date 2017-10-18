package ast

type NumberNode struct {
	ID    *ID
	Raw   string
	Value float64
}

type StringNode struct {
	Value string
	ID    *ID
}

type BooleanNode struct {
	Value bool
	ID    *ID
}

type NullNode struct{}

type ID struct {
	Name string
}

type MemberNode struct {
	Name  AST
	Value AST
}

type ArrayNode struct {
	Value []AST
}

type ObjectNode struct {
	Members []AST
}

type AST interface {
}
