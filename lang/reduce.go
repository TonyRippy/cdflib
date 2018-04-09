package lang

type reduceExpr struct {
	Init Expression
	Args []Expression
	ReduceOp func(Expression, Expression) (Expression, error)
}

func (x *reduceExpr) Type() ExpressionType {
	return ANY
}

func (x *reduceExpr) Eval(env Environment) (Expression, error) {
	var result Expression
	var v Expression
	var err error
	if result, err = x.Init.Eval(env); err != nil {
		return nil, err
	}
	for _, expr := range(x.Args) {
		if v, err = expr.Eval(env); err != nil {
			return nil, err
		}
		if result, err = x.ReduceOp(result, v) ; err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (x *reduceExpr) MimeData() MimeData {
	return MimeData{}
}
