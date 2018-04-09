package lang

import (
	"github.com/TonyRippy/cdflib"
)

type castFn func(Expression) Expression

type castDef struct {
	Type ExpressionType
	Fn castFn
}

var (
	// Casting (and auto-conversion?) is allowed to a more general type.
	// INTEGER --> FLOAT --> DISCRETE --+--> CDF
	//                  DIFERENTIABLE --|
	castDefs = map[ExpressionType]castDef {
		FLOAT: castDef{
			DISCRETE, func(expr Expression) Expression {
				v := expr.(*floatExpr).Value
				cdf := cdflib.Constant(v)
				return &cdfExpr{DISCRETE, cdf}
			}},
		DISCRETE: castDef{
			CDF, func(expr Expression) Expression {
				cdf := expr.(*cdfExpr).cdf
				return &cdfExpr{CDF, cdf}
			}},
		DIFFERENTIABLE: castDef{
			CDF, func(expr Expression) Expression {
				cdf := expr.(*cdfExpr).cdf
				return &cdfExpr{CDF, cdf}
			}},
	}
)

// Determines how many type conversions are necessary to turn a given expression
// into the desired type. It does not actually convert the type. Instead, it
// returns two pieces of information:
//   1) How many conversion steps are needed to reach the desired type.
//   2) A closure that is able to perform the conversion steps. 
// Returns -1 if the expression cannot be cast to the desired type.
func canCast(from ExpressionType, to ExpressionType) (int, castFn) {
	if from == to || to == ANY {
		return 0, nil
	}
	def, ok := castDefs[from]
	if ok {
		d, f := canCast(def.Type, to)
		if d == 0 {
			return 1, def.Fn
		} else if d > 0 {
			return d + 1, func(e Expression) Expression {
				return f(def.Fn(e))
			}
		}
	}
	return -1, nil
}
