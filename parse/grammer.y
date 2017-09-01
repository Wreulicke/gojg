%{
package parse

import "github.com/wreulicke/gojg/ast"

%}

%union{
    ast ast.AST
    values []ast.AST
    token Token
}

%type<ast> json_template value boolean_template number_template array object
%type<token> stirngOrTemplate
%type<values> elements members
%token<token> NUMBER
%token<token> FALSE
%token<token> NULL
%token<token> TRUE
%token<token> STRING STRING_TEMPLATE
%token<token> TEMPLATE_BEGIN TEMPLATE_END
%token<> BOOLEAN_PREFIX
%token<token> ID

%%

json_template: value { 
    yylex.(*Lexer).result = $1
    $$ = $1
}

value: 
    boolean_template {
        $$ = $1
    }
    | FALSE {
        $$ = &ast.ValueNode{Value: false}
    }
    | NULL {
        $$ = &ast.ValueNode{Value: nil}
    }
    | TRUE {
        $$ = &ast.ValueNode{Value: true}
    }
    | number_template {
        $$ = $1
    }
    | NUMBER {
        $$ = &ast.ValueNode{Value: $1.literal}
    }
    | object  {
        $$ = $1
    }
    | array {
        $$ = $1
    }
    | stirngOrTemplate {
        text := $1.literal
        $$ = &ast.ValueNode{Value: text[1:len(text)-1]}
    }

stirngOrTemplate: STRING 
    {
        $$ = $1
    }
    | STRING_TEMPLATE {
        $$ = $1
    }

boolean_template: BOOLEAN_PREFIX "(" ID ")" 
    { 
        $$ = &ast.BoolTemplateNode{Id: $3.literal}
    }
object: 
    "{" "}" 
    {
        $$ = &ast.ObjectNode{Members: []ast.AST{}}
    }
    |
    "{" members"}" 
    {
        $$ = &ast.ObjectNode{Members: $2}
    }

members: 
    stirngOrTemplate ":" value {
        $$ = []ast.AST{ast.MemberNode{Name: $1.literal, Value: $3}}
    }
    | stirngOrTemplate ":" value "," members {
        size := len($5)+1
        values := make([]ast.AST, size, size)
        values = append(values, ast.MemberNode{Name: $1.literal, Value: $3})
        values = append(values, $5...)
        $$ = values
    }

array: 
    "[" "]" 
    {
        $$ = &ast.ArrayNode{Value: []ast.AST{}}
    }
    | "[" elements "]"
    {
        $$ = &ast.ArrayNode{Value: $2}
    }

elements: 
    value 
    { 
        $$ = []ast.AST{$1}
    }
    | value "," elements
    {   
        size := len($3)+1
        values := make([]ast.AST, size, size)
        values = append(values, $1)
        values = append(values, $3...)
        $$ = values
    }

number_template: 
    TEMPLATE_BEGIN ID TEMPLATE_END
    {
        $$ = &ast.NumberTemplateNode{Id: $2.literal}
    }

%%