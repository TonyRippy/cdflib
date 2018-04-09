//line lang.y:6
package lang

import __yyfmt__ "fmt"

//line lang.y:7
import (
	"strconv"
)

/*

TODO: Get this working...

statement : IF condition LBRACE statements RBRACE ELSE LBRACE statements RBRACE { $$ = px.If($2, $4, $8) }
          | IF condition LBRACE statements RBRACE { $$ = px.If($2, $4, nil) }
          | expr  { $$ = $1 }
          ;

condition : LPAREN condition RPAREN  { $$ = $2 }
          | NOT condition { $$ = px.Void() }
          | condition AND condition { $$ = px.Void() }
          | condition OR condition { $$ = px.Void() }
          | expr EQ expr { $$ = px.Void() }
          | expr NE expr { $$ = px.Void() }
          | expr LT expr { $$ = px.Void() }
          | expr LE expr { $$ = px.Void() }
          | expr GT expr { $$ = px.Void() }
          | expr GE expr { $$ = px.Void() }
          ;
*/

//line lang.y:36
type yySymType struct {
	yys   int
	str   string
	expr  Expression
	list  []Expression
	parg  PlotArg
	pargs []PlotArg
	color Color
}

const NUM = 57346
const COLOR = 57347
const IDENTIFIER = 57348
const PLOT = 57349
const BRIAN = 57350
const BREAK = 57351
const COMMA = 57352
const LPAREN = 57353
const RPAREN = 57354
const EQ_OP = 57355
const NE_OP = 57356
const LT_OP = 57357
const LE_OP = 57358
const GT_OP = 57359
const GE_OP = 57360
const AND_OP = 57361
const OR_OP = 57362
const IF = 57363
const ELSE = 57364
const ASSIGN_OP = 57365
const ADD_OP = 57366
const SUB_OP = 57367
const MULT_OP = 57368
const DIV_OP = 57369
const MOD_OP = 57370
const LEFT_OP = 57371
const RIGHT_OP = 57372
const LEFT_ASSIGN = 57373
const RIGHT_ASSIGN = 57374
const ADD_ASSIGN = 57375
const SUB_ASSIGN = 57376
const MULT_ASSIGN = 57377
const DIV_ASSIGN = 57378
const MOD_ASSIGN = 57379

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"COLOR",
	"IDENTIFIER",
	"PLOT",
	"BRIAN",
	"BREAK",
	"COMMA",
	"LPAREN",
	"RPAREN",
	"EQ_OP",
	"NE_OP",
	"LT_OP",
	"LE_OP",
	"GT_OP",
	"GE_OP",
	"AND_OP",
	"OR_OP",
	"IF",
	"ELSE",
	"ASSIGN_OP",
	"ADD_OP",
	"SUB_OP",
	"MULT_OP",
	"DIV_OP",
	"MOD_OP",
	"LEFT_OP",
	"RIGHT_OP",
	"LEFT_ASSIGN",
	"RIGHT_ASSIGN",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MULT_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line lang.y:271

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 112

var yyAct = [...]int{

	4, 55, 25, 15, 5, 73, 32, 33, 58, 14,
	30, 31, 28, 29, 26, 27, 70, 3, 71, 38,
	34, 35, 37, 66, 39, 58, 61, 59, 62, 60,
	41, 42, 43, 44, 45, 52, 48, 49, 24, 57,
	46, 47, 40, 72, 22, 1, 7, 19, 21, 68,
	69, 56, 53, 67, 54, 6, 8, 63, 9, 10,
	64, 11, 57, 65, 17, 18, 22, 12, 7, 19,
	21, 13, 16, 23, 50, 22, 20, 7, 19, 21,
	51, 22, 56, 7, 19, 21, 17, 18, 23, 2,
	22, 0, 36, 19, 21, 17, 18, 23, 0, 0,
	0, 17, 18, 0, 0, 0, 0, 0, 0, 0,
	17, 18,
}
var yyPact = [...]int{

	77, -1000, 29, -1000, -1000, -1000, -1000, -21, -1000, -1000,
	-1000, -1000, -1000, -14, -20, -1000, 9, 86, 86, 8,
	-1000, -1000, -1000, 77, 77, 77, 77, 77, 77, 77,
	86, 86, 86, 86, 62, -1000, -1000, -1000, 40, -4,
	-1000, -1000, -1000, -1000, -1000, -1000, -20, -20, -1000, -1000,
	-1000, 17, -1000, -1000, 16, -1000, 77, -1000, -1000, 77,
	-1000, 71, -1000, 13, -1000, -1000, 44, 6, -1000, -1000,
	39, -1000, -7, -1000,
}
var yyPgo = [...]int{

	0, 89, 80, 17, 0, 76, 72, 3, 9, 71,
	67, 61, 59, 58, 56, 55, 4, 1, 54, 53,
	45,
}
var yyR1 = [...]int{

	0, 20, 1, 1, 1, 1, 5, 5, 5, 5,
	19, 19, 17, 17, 17, 18, 18, 2, 2, 6,
	6, 6, 6, 6, 7, 7, 7, 8, 8, 8,
	9, 9, 9, 10, 11, 12, 13, 14, 15, 16,
	16, 16, 16, 16, 16, 4, 3,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 2, 1, 1, 1, 1, 3,
	1, 1, 7, 5, 1, 3, 1, 3, 1, 3,
	4, 3, 4, 1, 1, 2, 2, 1, 3, 3,
	1, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 3, 3, 3, 1, 1,
}
var yyChk = [...]int{

	-1000, -20, -1, -3, -4, -16, -15, 6, -14, -13,
	-12, -11, -10, -9, -8, -7, -6, 24, 25, 7,
	-5, 8, 4, 11, 9, 23, 35, 36, 33, 34,
	24, 25, 26, 27, 11, -7, 6, -7, 11, -4,
	-3, -16, -16, -16, -16, -16, -8, -8, -7, -7,
	12, -2, -4, 12, -18, -17, 11, -4, 12, 10,
	12, 10, 12, -4, -4, -17, 10, -19, 5, 6,
	10, 12, 4, 12,
}
var yyDef = [...]int{

	2, -2, 1, 5, 46, 45, 39, 6, 38, 37,
	36, 35, 34, 33, 30, 27, 24, 0, 0, 0,
	23, 7, 8, 0, 4, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 25, 6, 26, 0, 0,
	3, 40, 41, 42, 43, 44, 31, 32, 28, 29,
	21, 0, 18, 19, 0, 16, 0, 14, 9, 0,
	22, 0, 20, 0, 17, 15, 0, 0, 10, 11,
	0, 13, 0, 12,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37,
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
	lval     yySymType
	stack    [yyInitialStackSize]yySymType
	char     int
	userData interface{}
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
		//line lang.y:107
		{
			yyrcvr.userData = yyDollar[1].list
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line lang.y:111
		{
			yyVAL.list = []Expression{}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:112
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:113
		{
			yyVAL.list = yyDollar[1].list
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:114
		{
			yyVAL.list = []Expression{yyDollar[1].expr}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:118
		{
			yyVAL.expr = Lookup(yyDollar[1].str)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:119
		{
			yyVAL.expr = Brian()
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:121
		{
			f, _ := strconv.ParseFloat(yyDollar[1].str, 64)
			yyVAL.expr = Float(f)
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:125
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:129
		{
			yyVAL.color = ColorBySpec(yyDollar[1].str)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:130
		{
			yyVAL.color = ColorByName(yyDollar[1].str)
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line lang.y:135
		{
			f, _ := strconv.ParseFloat(yyDollar[6].str, 64)
			yyVAL.parg = PlotArg{yyDollar[2].expr, yyDollar[4].color, f}
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line lang.y:140
		{
			yyVAL.parg = PlotArg{yyDollar[2].expr, yyDollar[4].color, DefaultFillAlpha}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:144
		{
			yyVAL.parg = PlotArg{yyDollar[1].expr, ColorByPosition(), DefaultFillAlpha}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:150
		{
			yyVAL.pargs = append(yyDollar[1].pargs, yyDollar[3].parg)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:151
		{
			yyVAL.pargs = []PlotArg{yyDollar[1].parg}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:155
		{
			yyVAL.list = append(yyDollar[1].list, yyDollar[3].expr)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:156
		{
			yyVAL.list = []Expression{yyDollar[1].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:167
		{
			yyVAL.expr = PlotNoArgs(yyDollar[1].str)
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:168
		{
			yyVAL.expr = Plot(yyDollar[1].str, yyDollar[3].pargs)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:169
		{
			yyVAL.expr = Apply(yyDollar[1].expr, []Expression{})
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line lang.y:170
		{
			yyVAL.expr = Apply(yyDollar[1].expr, yyDollar[3].list)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:171
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:183
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:184
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line lang.y:185
		{
			yyVAL.expr = Negate(yyDollar[2].expr)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:190
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:191
		{
			yyVAL.expr = Mult(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:192
		{
			yyVAL.expr = Div(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:197
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:198
		{
			yyVAL.expr = Add(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:199
		{
			yyVAL.expr = Sub(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:203
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:211
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:221
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:229
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:236
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:243
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:250
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:251
		{
			yyVAL.expr = Assign(yyDollar[1].str, yyDollar[3].expr)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:252
		{
			yyVAL.expr = Assign(yyDollar[1].str, Mult(Lookup(yyDollar[1].str), yyDollar[3].expr))
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:253
		{
			yyVAL.expr = Assign(yyDollar[1].str, Div(Lookup(yyDollar[1].str), yyDollar[3].expr))
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:255
		{
			yyVAL.expr = Assign(yyDollar[1].str, Add(Lookup(yyDollar[1].str), yyDollar[3].expr))
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line lang.y:256
		{
			yyVAL.expr = Assign(yyDollar[1].str, Sub(Lookup(yyDollar[1].str), yyDollar[3].expr))
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:264
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line lang.y:268
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
