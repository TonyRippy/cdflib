%{
// Lexer for CDFlib grammar.

package lang

import (
        "bufio"
        "bytes"
        "fmt"
        "io"
        "log"
)

// type yyLexer interface {
//   Lex(lval *yySymType) int
//   Error(e string)
// }

type yyLexerImpl struct {
  src *bufio.Reader
	buf *bytes.Buffer
	current byte
}

func makeLexer(src io.Reader) yyLexer {
	x := &yyLexerImpl{src: bufio.NewReader(src), buf: new(bytes.Buffer)}
  if b, err := x.src.ReadByte(); err == nil {
    x.current = b
  }
  return x
}

func (x *yyLexerImpl) Error(s string) {
	log.Printf("parse error: %s", s)
}

func (x *yyLexerImpl) snapshot(lval *yySymType) {
  lval.str = x.buf.String()
  //log.Printf("snapshot token: \"%s\"", lval.str)
}

func (x *yyLexerImpl) getc() byte {
  // The buffer should be one character behind current.
	if x.current != 0 {
		x.buf.WriteByte(x.current)
    x.current = 0
	}
	if b, err := x.src.ReadByte(); err == nil {
		x.current = b
	}
	return x.current
}

func (x *yyLexerImpl) Lex(lval *yySymType) int {
  defer x.snapshot(lval) 
  c := x.current  // init
%}

%yyc c
%yyn c = x.getc()

D  [0-9]+
E  [eE][-+]?{D}
F  {D}\.{D}?{E}?|{D}{E}?|\.{D}{E}?

H  [0-9a-fA-F]

%%
  x.buf.Reset()

[ \t\n\r]+    /* ignore */

[Bb]rian      return BRIAN

Plotd?        return PLOT

;             return BREAK

,             return COMMA

\(            return LPAREN 

\)            return RPAREN

if            return IF

else          return ELSE

\<\<=         return LEFT_ASSIGN

>>=           return RIGHT_ASSIGN

\/=           return DIV_ASSIGN

\+=           return ADD_ASSIGN

-=            return SUB_ASSIGN

\*=           return MULT_ASSIGN

%=            return MOD_ASSIGN

\+            return ADD_OP

-             return SUB_OP

\*            return MULT_OP

\/            return DIV_OP

==            return EQ_OP

!=            return NE_OP

\<=           return LE_OP

\<            return LT_OP

>=            return GE_OP

>             return GT_OP

=             return ASSIGN_OP

\<\<          return LEFT_OP

>>            return RIGHT_OP

&&            return AND_OP

\|\|          return OR_OP

{F}           return NUM

#{H}+         return COLOR

[a-zA-Z_][a-zA-Z0-9_]*  return IDENTIFIER

\0            return 0

%%
    // The golex generated scanner enters top of the user code section when
    // lexem recongition fails.
    panic(fmt.Sprintf("Unrecognized byte 0x%x after \"%s\"\n", x.current, x.buf.String()))
    return 0

}
