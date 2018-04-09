//go:generate golex -o lexer.go lang.l
//go:generate goyacc -o parser.go lang.y

package lang

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Interpreter interface {
	SetOut(out io.Writer) io.Writer
	SetErr(err io.Writer) io.Writer
	Eval(code string) (Expression, error)
}

type langImpl struct {
	out io.Writer
	err io.Writer
	vars map[string]Expression
}

func (i *langImpl) Out() io.Writer {
	return i.out
}

func (i *langImpl) SetOut(w io.Writer) io.Writer {
	old := i.out
	i.out = w
	return old
}

func (i *langImpl) Err() io.Writer {
	return i.err
}

func (i *langImpl) SetErr(w io.Writer) io.Writer {
	old := i.err
	i.err = w
	return old
}

func (i *langImpl) GetVar(name string) (Expression, bool) {
	expr, ok := i.vars[name]
	return expr, ok
}

func (i *langImpl) SetVar(name string, expr Expression) {
	i.vars[name] = expr
}

func Parse(code string) ([]Expression, error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		err = fmt.Errorf("Code failed to tokenize: %s", r)
	// 	}
	// }()
	reader := bytes.NewBufferString(code)
  parser := yyParserImpl{}
  rcode := parser.Parse(makeLexer(reader))
	if rcode != 0 {
		return nil, fmt.Errorf("Code failed to parse. (%d)", rcode)
	}
	return parser.userData.([]Expression), nil
}

func (i *langImpl) Eval(code string) (Expression, error) {
	exprs, err := Parse(code)
	if err != nil {
		return nil, err
	}
	last := Void()
	for _, expr := range(exprs) { 
		result, err := expr.Eval(i)
		if err != nil {
			return nil, err
		}
		last = result
	}
	return last, nil
}

func MakeInterpreter() Interpreter {
	return &langImpl{os.Stdout, os.Stderr, make(map[string]Expression)}
}
