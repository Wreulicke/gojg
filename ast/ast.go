package ast

type NumberNode struct {
	Raw   string
	Value float64
}

type RawValueTemplateNode struct {
	ID *ID
}

type StringNode struct {
	Value string
	ID    *ID
}

type BooleanNode struct {
	Value bool
}

type NullNode struct{}

type ID struct {
	Name string
}

type MemberNode struct {
	Name  *StringNode
	Value AST
}

type ArrayNode struct {
	Value []AST
}

type ObjectNode struct {
	Members []MemberNode
}

type AST interface {
}
