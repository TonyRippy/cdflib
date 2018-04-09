package lang

import (
	"fmt"
)

type brianExpr struct {
}

func (i *brianExpr) Type() ExpressionType {
	return VOID
}

func (i *brianExpr) Eval(env Environment) (Expression, error) {
	fmt.Fprint(env.Out(), "Brian is totally awesome!");
	return i, nil
}

func (i *brianExpr) MimeData() MimeData {
	return MimeData{
		"text/latex": "$\\infty$",
	}
}

func (x *brianExpr) String() string {
	return "\u221E"
}

func Brian() Expression {
	return &brianExpr{}
}
