package ast

type BoolTemplateNode struct {
	Id string
}

type NumberTemplateNode struct {
	Id string
}

type ValueNode struct {
	Value interface{}
	Id    string
}

type MemberNode struct {
	Name  string
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

func test() {
}
