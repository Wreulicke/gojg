%{
package parse

import "github.com/wreulicke/gojg/ast"

%}

%union{
    ast     ast.AST
    values  []ast.AST
    token   Token
    member *ast.MemberNode
    members []ast.MemberNode
    string  *ast.StringNode
}

%type<ast> json_template value raw_value_template array object
%type<ast> number_literal
%type<string> string_or_template
%type<member> member
%type<members> members
%type<values> elements
%token<token> MINUS 
%token<token> NUMBER
%token<token> FALSE
%token<token> NULL
%token<token> TRUE
%token<token> COLON COMMA
%token<token> STRING STRING_TEMPLATE
%token<token> TEMPLATE_BEGIN TEMPLATE_END
%token<token> OBJECT_BEGIN OBJECT_END
%token<token> BRACE_BEGIN BRACE_END
%token<token> ARRAY_BEGIN ARRAY_END
%token<> BOOLEAN_PREFIX
%token<token> ID

%%

json_template: value { 
    yylex.(*Lexer).result = $1
    $$ = $1
}

value: 
    FALSE {
        $$ = &ast.BooleanNode{Value: false}
    }
    | NULL {
        $$ = &ast.NullNode{}
    }
    | TRUE {
        $$ = &ast.BooleanNode{Value: true}
    }
    | number_literal {
        $$ = $1
    }
    | raw_value_template {
        $$ = $1
    }
    | object  {
        $$ = $1
    }
    | array {
        $$ = $1
    }
    | string_or_template {
        $$ = $1
    }

string_or_template: 
    STRING {
        $$ = &ast.StringNode{Value: $1.literal}
    }
    | STRING_TEMPLATE {
        $$ = &ast.StringNode{ID: &ast.ID{$1.literal}}
    }

number_literal: 
    MINUS NUMBER {
        lex := yylex.(*Lexer)
        num := lex.parseFloat($1.literal + $2.literal)
        $$ = &ast.NumberNode{Value: num}
    }
    | NUMBER {
        lex := yylex.(*Lexer)
        num := lex.parseFloat($1.literal)
        $$ = &ast.NumberNode{Value: num}
    }

raw_value_template: 
    TEMPLATE_BEGIN ID TEMPLATE_END {
        $$ = &ast.RawValueTemplateNode{ID: &ast.ID{$2.literal}}
    }

object: 
    OBJECT_BEGIN OBJECT_END  {
        $$ = &ast.ObjectNode{Members: []ast.MemberNode{}}
    }
    |
    OBJECT_BEGIN members OBJECT_END {
        $$ = &ast.ObjectNode{Members: $2}
    }
member: 
    ID COLON value {
        $$ = &ast.MemberNode{Name: &ast.StringNode{Value: $1.literal}, Value: $3}
    }
    | string_or_template COLON value {
        $$ = &ast.MemberNode{Name: $1, Value: $3}
    }

members: 
    member {
        $$ = []ast.MemberNode{*$1}
    }
    | member COMMA{
        $$ = []ast.MemberNode{*$1}
    }
    | member COMMA members {
        size := len($3)+1
        values := make([]ast.MemberNode, 0, size)
        values = append(values, *$1)
        values = append(values, $3...)
        $$ = values
    }

array: 
    ARRAY_BEGIN ARRAY_END {
        $$ = &ast.ArrayNode{Value: []ast.AST{}}
    }
    | ARRAY_BEGIN elements ARRAY_END {
        $$ = &ast.ArrayNode{Value: $2}
    }

elements: 
    value { 
        $$ = []ast.AST{$1}
    }
    | value COMMA { 
        $$ = []ast.AST{$1}
    }
    | value COMMA elements {   
        size := len($3)+1
        values := make([]ast.AST, 0, size)
        values = append(values, $1)
        values = append(values, $3...)
        $$ = values
    }

%%