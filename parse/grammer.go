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
const STRING_TEMPLATE = 57351
const TEMPLATE_BEGIN = 57352
const TEMPLATE_END = 57353
const BOOLEAN_PREFIX = 57354
const ID = 57355

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"FALSE",
	"NULL",
	"TRUE",
	"STRING",
	"STRING_TEMPLATE",
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

//line parse/grammer.y:128

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 51

var yyAct = [...]int{

	11, 21, 24, 30, 8, 4, 5, 6, 16, 17,
	13, 35, 12, 25, 2, 22, 14, 31, 32, 29,
	15, 23, 28, 8, 4, 5, 6, 16, 17, 13,
	18, 12, 16, 17, 34, 14, 22, 36, 26, 15,
	19, 20, 27, 33, 16, 17, 9, 10, 7, 3,
	1,
}
var yyPact = [...]int{

	19, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 16, 27, 24, 0, -1000, -1000, 25, 31,
	-1000, 5, 1, -1000, -18, -2, 3, -1000, -1000, 19,
	-1000, 19, -1000, -8, -1000, 36, -1000,
}
var yyPgo = [...]int{

	0, 50, 13, 49, 48, 47, 46, 0, 2, 1,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 7, 7, 3, 6, 6, 9, 9, 5, 5,
	8, 8, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 2, 3, 3, 5, 2, 3,
	1, 3, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, 5, 6, 7, -4, 4, -6,
	-5, -7, 12, 10, 16, 20, 8, 9, 14, 13,
	17, -9, -7, 21, -8, -2, 13, 11, 17, 18,
	21, 19, 15, -2, -8, 19, -9,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 0, 0, 0, 0, 11, 12, 0, 0,
	14, 0, 0, 18, 0, 20, 0, 22, 15, 0,
	19, 0, 13, 16, 21, 0, 17,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 15, 3, 3, 19, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 18, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 20, 3, 21, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 16, 3, 17,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
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
		//line parse/grammer.y:28
		{
			yylex.(*Lexer).result = yyDollar[1].ast
			yyVAL.ast = yyDollar[1].ast
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:34
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:37
		{
			yyVAL.ast = &ast.ValueNode{Value: false}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:40
		{
			yyVAL.ast = &ast.ValueNode{Value: nil}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:43
		{
			yyVAL.ast = &ast.ValueNode{Value: true}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:46
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:49
		{
			yyVAL.ast = &ast.ValueNode{Value: yyDollar[1].token.literal}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:52
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:55
		{
			yyVAL.ast = yyDollar[1].ast
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:58
		{
			text := yyDollar[1].token.literal
			yyVAL.ast = &ast.ValueNode{Value: text[1 : len(text)-1]}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:64
		{
			yyVAL.token = yyDollar[1].token
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:67
		{
			yyVAL.token = yyDollar[1].token
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse/grammer.y:72
		{
			yyVAL.ast = &ast.BoolTemplateNode{Id: yyDollar[3].token.literal}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse/grammer.y:77
		{
			yyVAL.ast = &ast.ObjectNode{Members: []ast.AST{}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:82
		{
			yyVAL.ast = &ast.ObjectNode{Members: yyDollar[2].values}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:87
		{
			yyVAL.values = []ast.AST{ast.MemberNode{Name: yyDollar[1].token.literal, Value: yyDollar[3].ast}}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse/grammer.y:90
		{
			size := len(yyDollar[5].values) + 1
			values := make([]ast.AST, size, size)
			values = append(values, ast.MemberNode{Name: yyDollar[1].token.literal, Value: yyDollar[3].ast})
			values = append(values, yyDollar[5].values...)
			yyVAL.values = values
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse/grammer.y:100
		{
			yyVAL.ast = &ast.ArrayNode{Value: []ast.AST{}}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:104
		{
			yyVAL.ast = &ast.ArrayNode{Value: yyDollar[2].values}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse/grammer.y:110
		{
			yyVAL.values = []ast.AST{yyDollar[1].ast}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:114
		{
			size := len(yyDollar[3].values) + 1
			values := make([]ast.AST, size, size)
			values = append(values, yyDollar[1].ast)
			values = append(values, yyDollar[3].values...)
			yyVAL.values = values
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse/grammer.y:124
		{
			yyVAL.ast = &ast.NumberTemplateNode{Id: yyDollar[2].token.literal}
		}
	}
	goto yystack /* stack new state and value */
}
