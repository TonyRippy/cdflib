// Parser for the PX grammar.

// TODO: Add formal support to goyacc for %parse-param {} {}

%{

package lang

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
%}

%union {
	str string
  expr Expression
  list []Expression
  parg PlotArg
  pargs []PlotArg
  color Color
}

%type	<list> statements, args
%type <expr> statement
%type <expr> expression
%type <expr> primary_expression
%type <expr> postfix_expression
%type <expr> unary_expression
%type <expr> multiplicative_expression
%type <expr> additive_expression
%type <expr> shift_expression
%type <expr> relational_expression
%type <expr> equality_expression
%type <expr> logical_and_expression
%type <expr> logical_or_expression
%type <expr> conditional_expression
%type <expr> assignment_expression

%type	<parg> plotarg
%type	<pargs> plotargs
%type	<color> color

%token <str> NUM, COLOR, IDENTIFIER, PLOT
%token BRIAN

%token BREAK
%token COMMA
%token LPAREN
%token RPAREN
   
%token EQ_OP
%token NE_OP
%token LT_OP
%token LE_OP
%token GT_OP
%token GE_OP
%token AND_OP
%token OR_OP

%token IF
%token ELSE

%token ASSIGN_OP
%token ADD_OP
%token SUB_OP
%token MULT_OP
%token DIV_OP
%token MOD_OP
%token LEFT_OP
%token RIGHT_OP

%token LEFT_ASSIGN
%token RIGHT_ASSIGN
%token ADD_ASSIGN
%token SUB_ASSIGN
%token MULT_ASSIGN
%token DIV_ASSIGN
%token MOD_ASSIGN

%start top

%%

top
: statements { $! = $1 }
;

statements
: /* empty */                 { $$ = []Expression{} }
| statements BREAK statement  { $$ = append($1, $3) }
| statements BREAK            { $$ = $1 }
| statement                   { $$ = []Expression{$1} }
;

primary_expression
: IDENTIFIER   { $$ = Lookup($1) }
| BRIAN        { $$ = Brian() }
| NUM
  {
    f, _ := strconv.ParseFloat($1, 64)
    $$ = Float(f)
  }
| LPAREN expression RPAREN { $$ = $2 }
;

color
: COLOR      { $$ = ColorBySpec($1) }
| IDENTIFIER { $$ = ColorByName($1) }
;

plotarg
: LPAREN expression COMMA color COMMA NUM RPAREN
  {
    f, _ := strconv.ParseFloat($6, 64)
    $$ = PlotArg{$2, $4, f}
  }
| LPAREN expression COMMA color RPAREN
  {
    $$ = PlotArg{$2, $4, DefaultFillAlpha}
  }
| expression
  {
    $$ = PlotArg{$1, ColorByPosition(), DefaultFillAlpha}
  }
;

plotargs
: plotargs COMMA plotarg { $$ = append($1, $3) }
| plotarg                { $$ = []PlotArg{$1} }
;

args
: args COMMA expression { $$ = append($1, $3) }
| expression            { $$ = []Expression{$1} }
;

/*
So the problem here is how we treat functions in the language.
How are function names resolved? In the longer-term it would be nice to allow users to define
their own functions. So we need a function type, and the symbol table needs to support that.
We need to add a value type of FUNCTION or something like that to the assignment code.
This will require all built-in functions to be added to the symbol table on initialization.
*/
postfix_expression
: PLOT LPAREN RPAREN               { $$ = PlotNoArgs($1) }
| PLOT LPAREN plotargs RPAREN      { $$ = Plot($1, $3) }
| postfix_expression LPAREN RPAREN { $$ = Apply($1, []Expression{}) }
| postfix_expression LPAREN args RPAREN { $$ = Apply($1, $3) }
| primary_expression { $$ = $1 }
;
/*
| postfix_expression '[' expression ']'
| postfix_expression '.' IDENTIFIER
| postfix_expression PTR_OP IDENTIFIER
| postfix_expression INC_OP
| postfix_expression DEC_OP
;
*/

unary_expression
: postfix_expression { $$ = $1 }
| ADD_OP unary_expression { $$ = $2 } 
| SUB_OP unary_expression { $$ = Negate($2) }
/* | '!' unary_expression { $$ = Not($2) } */
;

multiplicative_expression
: unary_expression { $$ = $1 }
| multiplicative_expression MULT_OP unary_expression { $$ = Mult($1, $3) }
| multiplicative_expression DIV_OP unary_expression  { $$ = Div($1, $3) }
/* | multiplicative_expression MOD_OP unary_expression */
;

additive_expression
: multiplicative_expression { $$ = $1 }
| additive_expression ADD_OP multiplicative_expression { $$ = Add($1, $3) }
| additive_expression SUB_OP multiplicative_expression { $$ = Sub($1, $3) }
;

shift_expression
: additive_expression { $$ = $1 }
/*
| shift_expression LEFT_OP additive_expression
| shift_expression RIGHT_OP additive_expression
*/
;

relational_expression
: shift_expression { $$ = $1 }
/*
| relational_expression LT_OP shift_expression
| relational_expression GT_OP shift_expression
| relational_expression LE_OP shift_expression
| relational_expression GE_OP shift_expression
*/
;

equality_expression
: relational_expression { $$ = $1 }
/*
| equality_expression EQ_OP relational_expression
| equality_expression NE_OP relational_expression
*/
;

logical_and_expression
: equality_expression { $$ = $1 }
/*
| logical_and_expression AND_OP equality_expression
*/
;

logical_or_expression
: logical_and_expression { $$ = $1 }
/*
| logical_or_expression OR_OP logical_and_expression
*/
;

conditional_expression
: logical_or_expression { $$ = $1 }
/*
| logical_or_expression '?' expression ':' conditional_expression
*/
;

assignment_expression
: conditional_expression { $$ = $1 }
| IDENTIFIER ASSIGN_OP assignment_expression    { $$ = Assign($1, $3) }
| IDENTIFIER MULT_ASSIGN  assignment_expression { $$ = Assign($1, Mult(Lookup($1), $3)) }
| IDENTIFIER DIV_ASSIGN  assignment_expression  { $$ = Assign($1, Div(Lookup($1), $3)) }
/* | IDENTIFIER MOD_ASSIGN assignment_expression   { $$ = $3 } */
| IDENTIFIER ADD_ASSIGN assignment_expression   { $$ = Assign($1, Add(Lookup($1), $3)) }
| IDENTIFIER SUB_ASSIGN assignment_expression   { $$ = Assign($1, Sub(Lookup($1), $3)) }
/*
| IDENTIFIER LEFT_ASSIGN assignment_expression  { $$ = $3 }
| IDENTIFIER RIGHT_ASSIGN assignment_expression { $$ = $3 }
*/
;

expression
: assignment_expression { $$ = $1 }
;

statement
: expression { $$ = $1 }
;

%%
