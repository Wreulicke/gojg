//line parse/grammer.y:2
package parse

import __yyfmt__ "fmt"

//line parse/grammer.y:2
import "github.com/wreulicke/gojg/ast"

//line parse/grammer.y:8
type yySymType struct {
	yys    int
	ast    ast.AST
	values []ast.AST
	token  Token
}

const NUMBER = 57346
const FALSE = 57347
const NULL = 57348
const TRUE = 57349
const STRING = 57350
const TEMPLATE_BEGIN = 57351
const TEMPLATE_END = 57352
const BOOLEAN_PREFIX = 57353
const ID = 57354

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"FALSE",
	"NULL",
	"TRUE",
	"STRING",
	"TEMPLATE_BEGIN",
	"TEMPLATE_END",
	"BOOLEAN_PREFIX",
	"ID",
	"\"(\"",
	"\")\"",
	"\"{\"",
	"\"}\"",
	"\":\"",
	"\",\"",
	"\"[\"",
	"\"]\"",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse/grammer.y:119

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 44

var yyAct = [...]int{

	19, 22, 8, 4, 5, 6, 11, 13, 28, 12,
	23, 2, 33, 14, 29, 30, 27, 15, 21, 20,
	26, 8, 4, 5, 6, 11, 13, 18, 12, 16,
	24, 32, 14, 17, 34, 25, 15, 20, 31, 9,
	10, 7, 3, 1,
}
var yyPact = [...]int{

	17, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 16, 21, 11, -2, 18, 25, -1000, 4,
	-1, -1000, -12, -4, 1, -1000, -1000, 17, -1000, 17,
	-1000, -6, -1000, 29, -1000,
}
var yyPgo = [...]int{

	0, 43, 10, 42, 41, 40, 39, 1, 0,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 6, 6, 8, 8, 5, 5, 7, 7,
	4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 4, 2, 3, 3, 5, 2, 3, 1, 3,
	3,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 5, 6, 7, -4, 4, -6,
	-5, 8, 11, 9, 15, 19, 13, 12, 16, -8,
	8, 20, -7, -2, 12, 10, 16, 17, 20, 18,
	14, -2, -7, 18, -8,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 0, 0, 0, 0, 0, 0, 12, 0,
	0, 16, 0, 18, 0, 20, 13, 0, 17, 0,
	11, 14, 19, 0, 15,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	13, 14, 3, 3, 18, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 17, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 19, 3, 20, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 15, 3, 16,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:27
		{
			yylex.(*Lexer).result = yyDollar[1].ast
			yyVAL.ast = yyDollar[1].ast
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:33
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:36
		{
			yyVAL.ast = &ast.ValueNode{Value: false}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:39
		{
			yyVAL.ast = &ast.ValueNode{Value: nil}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:42
		{
			yyVAL.ast = &ast.ValueNode{Value: true}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:45
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:48
		{
			yyVAL.ast = &ast.ValueNode{Value: yyDollar[1].token.literal}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:51
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:54
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:57
		{
			text := yyDollar[1].token.literal
			yyVAL.ast = &ast.ValueNode{Value: text[1 : len(text)-1]}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse/grammer.y:63
		{
			yyVAL.ast = &ast.BoolTemplateNode{Id: yyDollar[3].token.literal}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse/grammer.y:68
		{
			yyVAL.ast = &ast.ObjectNode{Members: []ast.AST{}}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:73
		{
			yyVAL.ast = &ast.ObjectNode{Members: yyDollar[2].values}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:77
		{
			yyVAL.values = []ast.AST{ast.MemberNode{Name: yyDollar[1].token.literal, Value: yyDollar[3].ast}}
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse/grammer.y:80
		{
			size := len(yyDollar[5].values) + 1
			values := make([]ast.AST, size, size)
			values = append(values, ast.MemberNode{Name: yyDollar[1].token.literal, Value: yyDollar[3].ast})
			values = append(values, yyDollar[5].values...)
			yyVAL.values = values
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse/grammer.y:91
		{
			yyVAL.ast = &ast.ArrayNode{Value: []ast.AST{}}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:95
		{
			yyVAL.ast = &ast.ArrayNode{Value: yyDollar[2].values}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:101
		{
			yyVAL.values = []ast.AST{yyDollar[1].ast}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:105
		{
			size := len(yyDollar[3].values) + 1
			values := make([]ast.AST, size, size)
			values = append(values, yyDollar[1].ast)
			values = append(values, yyDollar[3].values...)
			yyVAL.values = values
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:115
		{
			yyVAL.ast = &ast.NumberTemplateNode{Id: yyDollar[2].token.literal}
		}
	}
	goto yystack /* stack new state and value */
}
