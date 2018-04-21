package lang

import (
	"errors"
	"fmt"
	"github.com/TonyRippy/cdflib"
	"strconv"
)

type ExpressionType int

/*
Notes on Types:
---------------
ANY instead of MIXED? A wild-card, any expression can be cast to any.
(??? Maybe except VOID? Unsure. ???)

SYMBOL is an unresolved identifier.

INTEGER, castable to FLOAT?

no BOOL... we will have constant DIFFERENTIABLE True "Always" and "Never" 

VOID is a placeholder for things that do not return a value. Cannot be converted to anything.
MIXED is an intermediate state, only possible before first call to Eval()?

CDF is a continuous CDF, meaning it's value is known for all points.

DISCRETE is a CDF with fixed values w/ probabilities.
Can always be downcast to a CDF.
Can be cast to a FLOAT if there is only one discrete value with P(x)=1.0

FLOAT is a scalar value, can be cast to DISCRETE.
Example: 1.0 -> discrete CDF that has P(x == 1.0) = 100%.

DIFFERENTIABLE is a CDF that has a known derivative.
Can be downcast to a CDF.

IMAGE is a raster image, cannot be cast to any other type.

LAMBDA is a function type, cannot be cast to any other type. 
*/

const (
	VOID ExpressionType = iota
	SYMBOL
	ANY
	FLOAT
	IMAGE
	CDF
	DISCRETE
	DIFFERENTIABLE
	LAMBDA
)

func TypeName(t ExpressionType) string {
	switch t {
	case VOID:
		return "VOID"
	case SYMBOL:
		return "SYMBOL"
	case FLOAT:
		return "NUMBER"
	case IMAGE:
		return "IMAGE"
	case CDF:
		return "CDF"
	case DISCRETE:
		return "DISCRETE"
	case DIFFERENTIABLE:
		return "DIFFERENTIABLE"
	case ANY:
		return "ANY"
	case LAMBDA:
		return "LAMBDA"
	}
	return "UNKNOWN"
}

type MimeData map[string]interface{}

type Expression interface {
	Type() ExpressionType
	Eval(env Environment) (Expression, error)
	MimeData() MimeData
	String() string
}

///////////////////////////////

type voidExpr struct {
}

func (x *voidExpr) Type() ExpressionType {
	return VOID
}

func (x *voidExpr) Eval(env Environment) (Expression, error) {
	return x, nil
}

func (x *voidExpr) MimeData() MimeData {
	return MimeData{
		"text/plain": x.String(),
	}
}

func (x *voidExpr) String() string {
	return "nil"
}

func Void() Expression {
	return &voidExpr{}
}

///////////////////////////////

type errorExpr struct {
	t ExpressionType
	err error
}

func (x *errorExpr) Type() ExpressionType {
	return x.t
}

func (x *errorExpr) Eval(env Environment) (Expression, error) {
	return nil, x.err
}

func (x *errorExpr) MimeData() MimeData {
	return MimeData{}
}

func (x *errorExpr) String() string {
	return fmt.Sprintf("<ERROR: %s>", x.err)
}

///////////////////////////////

type floatExpr struct {
	Value float64
}

func (x *floatExpr) Type() ExpressionType {
	return FLOAT
}

func (x *floatExpr) Eval(env Environment) (Expression, error) {
	return x, nil
}

func (x *floatExpr) MimeData() MimeData {
	return MimeData{
		"text/plain": x.String(),
	}
}

func (x *floatExpr) String() string {
	return strconv.FormatFloat(x.Value, 'g', -1, 64)
}

func Float(v float64) Expression {
	return &floatExpr{v}
}

func IsFloat(expr Expression) bool {	
	return expr.Type() == FLOAT
}

func AsFloat(expr Expression) (float64, error) {
	if expr.Type() == FLOAT {
		return expr.(*floatExpr).Value, nil
	}
	return 0, errors.New("not a numeric value")
}

func AsInt(expr Expression) (int, error) {
	f, err := AsFloat(expr)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

///////////////////////////////

type cdfExpr struct {
	t ExpressionType
	cdf cdflib.CDF
}

func (x *cdfExpr) Type() ExpressionType {
	return x.t
}

func (x *cdfExpr) Eval(env Environment) (Expression, error) {
	return x, nil
}

func (x *cdfExpr) MimeData() MimeData {
	return MimeData{}
}

func (x *cdfExpr) String() string {
	return "<CDF>"
}

func AsCDF(expr Expression) (cdflib.CDF, error) {
	x, ok := expr.(*cdfExpr)
	if !ok {
		return nil, errors.New("not a random variable")
	}
	return x.cdf, nil
}

func AsDifferentiable(expr Expression) (cdflib.DifferentiableCDF, error) {
	x, ok := expr.(*cdfExpr)
	if !ok {
		return nil, errors.New("not a random variable")
	}
	diff, ok := x.cdf.(cdflib.DifferentiableCDF)
	if !ok {
		return nil, errors.New("not a differentiable variable")
	}
	return diff, nil
}

///////////////////////////////

type assignExpr struct {
	Name string
	Value Expression
}

func (x *assignExpr) Type() ExpressionType {
	return x.Value.Type()
}

func (x *assignExpr) Eval(env Environment) (Expression, error) {
	v, err := x.Value.Eval(env)
	if err != nil {
		return nil, err
	}
	env.SetVar(x.Name, v)
	return v, nil
}

func (x *assignExpr) MimeData() MimeData {
	return MimeData{}
}

func (x *assignExpr) String() string {
	return fmt.Sprintf("%s = %s", x.Name, x.Value)
}

func Assign(name string, e Expression) Expression {
	return &assignExpr{name, e}
}

///////////////////////////////

type lookupExpr struct {
	Name string
}

func (x *lookupExpr) Type() ExpressionType {
	return SYMBOL
}

func (x *lookupExpr) Eval(env Environment) (Expression, error) {
	v, ok := env.GetVar(x.Name)
	if !ok {
		return nil, fmt.Errorf("Unknown symbol \"%s\".", x.Name)
	}
	return v, nil
}

func (x *lookupExpr) MimeData() MimeData {
	return MimeData{}
}

func (x *lookupExpr) String() string {
	return x.Name
}

func Lookup(name string) Expression {
	return &lookupExpr{name}
}
