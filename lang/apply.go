package lang

import (
	"errors"
	"fmt"
	"github.com/TonyRippy/cdflib"
)

// TODO: How to handle varargs???

type argSpec struct {
	Name string
	Type ExpressionType
}

type funcFn func([]Expression) (Expression, error)

type funcSpec struct {
	Args []argSpec
	Fn funcFn
}

func builtinFunctions() map[string][]*funcSpec {
	funcs := make(map[string][]*funcSpec)
	var specs []*funcSpec
	
	// Normal distribution (Gaussian)
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{},
			func(args []Expression) (Expression, error) {
				cdf := cdflib.Gaussian(0, 1)
				return &cdfExpr{DIFFERENTIABLE, cdf}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"mean", FLOAT},
				argSpec{"stddev", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var mean, stddev float64
				var err error
				if mean, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Gaussian, parameter 'mean': %s", err)
				}
				if stddev, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Gaussian, parameter 'stddev': %s", err)
				}
				cdf := cdflib.Gaussian(mean, stddev)
				return &cdfExpr{DIFFERENTIABLE, cdf}, nil
			}},
	}
	funcs["Gaussian"] = specs
	funcs["Normal"] = specs
	funcs["N"] = specs

	// LogNormal
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{},
			func(args []Expression) (Expression, error) {
				cdf := cdflib.LogNormal(0, 1)
				return &cdfExpr{DIFFERENTIABLE, cdf}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"mean", FLOAT},
				argSpec{"stddev", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var mean, stddev float64
				var err error
				if mean, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Gaussian, parameter 'mean': %s", err)
				}
				if stddev, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Gaussian, parameter 'stddev': %s", err)
				}
				cdf := cdflib.LogNormal(mean, stddev)
				return &cdfExpr{DIFFERENTIABLE, cdf}, nil
			}},
	}
	funcs["LogNormal"] = specs
	funcs["LN"] = specs

	// Negate
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{
				argSpec{"value", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v float64
				var err error
				if v, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Negate, parameter 'value': %s", err)
				}
				return &floatExpr{-v}, nil
			}},
	}
	funcs["Neg"] = specs
	funcs["Negate"] = specs

	// Add
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{},
			func(args []Expression) (Expression, error) {
				return &floatExpr{0}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", ANY},
			},
			func(args []Expression) (Expression, error) {
				return args[0], nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 1: %s", err)
				}
				return &floatExpr{v0 + v1}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", CDF},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 cdflib.CDF
				var err error
				if v0, err = AsCDF(args[0]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.AddScalar(v0, v1)}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", CDF},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 0: %s", err)
				}
				var v1 cdflib.CDF
				if v1, err = AsCDF(args[1]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.AddScalar(v1, v0)}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", DIFFERENTIABLE},
				argSpec{"arg1", CDF},
			},
			func(args []Expression) (Expression, error) {
				var v0 cdflib.DifferentiableCDF
				var err error
				if v0, err = AsDifferentiable(args[0]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 0: %s", err)
				}
				var v1 cdflib.CDF
				if v1, err = AsCDF(args[1]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.Add(v0, v1)}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", CDF},
				argSpec{"arg1", DIFFERENTIABLE},
			},
			func(args []Expression) (Expression, error) {
				var v0 cdflib.DifferentiableCDF
				var err error
				if v0, err = AsDifferentiable(args[1]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 1: %s", err)
				}
				var v1 cdflib.CDF
				if v1, err = AsCDF(args[0]); err != nil {
					return nil, fmt.Errorf("'Add, parameter 0: %s", err)
				}
				return &cdfExpr{CDF, cdflib.Add(v0, v1)}, nil
			}},
	}
	funcs["Add"] = specs

	// Subtract
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{},
			func(args []Expression) (Expression, error) {
				return &floatExpr{0}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", ANY},
			},
			func(args []Expression) (Expression, error) {
				return args[0], nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Subtract, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Subtract, parameter 1: %s", err)
				}
				return &floatExpr{v0 - v1}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", CDF},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 cdflib.CDF
				var err error
				if v0, err = AsCDF(args[0]); err != nil {
					return nil, fmt.Errorf("'Subtract, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Subtract, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.SubtractScalar(v0, v1)}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", CDF},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Subtract, parameter 0: %s", err)
				}
				var v1 cdflib.CDF
				if v1, err = AsCDF(args[1]); err != nil {
					return nil, fmt.Errorf("'Subtract, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.AddScalar(cdflib.Neg(v1), v0)}, nil
			}},
	}
	funcs["Sub"] = specs
	funcs["Subtract"] = specs

	// Multiply
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{},
			func(args []Expression) (Expression, error) {
				return &floatExpr{1}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", ANY},
			},
			func(args []Expression) (Expression, error) {
				return args[0], nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Multiply, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Multiply, parameter 1: %s", err)
				}
				return &floatExpr{v0 * v1}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", CDF},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 cdflib.CDF
				var err error
				if v0, err = AsCDF(args[0]); err != nil {
					return nil, fmt.Errorf("'Multiply, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Multiply, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.MultiplyScalar(v0, v1)}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", CDF},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Multiply, parameter 0: %s", err)
				}
				var v1 cdflib.CDF
				if v1, err = AsCDF(args[1]); err != nil {
					return nil, fmt.Errorf("'Multiply, parameter 1: %s", err)
				}
				return &cdfExpr{CDF, cdflib.MultiplyScalar(v1, v0)}, nil
			}},
	}
	funcs["Mult"] = specs
	funcs["Multiply"] = specs

	// Divide
	specs = []*funcSpec{
		&funcSpec {
			[]argSpec{},
			func(args []Expression) (Expression, error) {
				return &floatExpr{1}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", ANY},
			},
			func(args []Expression) (Expression, error) {
				return args[0], nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", FLOAT},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 float64
				var err error
				if v0, err = AsFloat(args[0]); err != nil {
					return nil, fmt.Errorf("'Divide, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Divide, parameter 1: %s", err)
				}
				if v1 == 0 {
					return nil, errors.New("Divide by zero.")
				}
				return &floatExpr{v0 / v1}, nil
			}},
		&funcSpec {
			[]argSpec{
				argSpec{"arg0", CDF},
				argSpec{"arg1", FLOAT},
			},
			func(args []Expression) (Expression, error) {
				var v0 cdflib.CDF
				var err error
				if v0, err = AsCDF(args[0]); err != nil {
					return nil, fmt.Errorf("'Divide, parameter 0: %s", err)
				}
				var v1 float64
				if v1, err = AsFloat(args[1]); err != nil {
					return nil, fmt.Errorf("'Divide, parameter 1: %s", err)
				}
				if v1 == 0 {
					return nil, errors.New("Divide by zero.")
				}
				return &cdfExpr{CDF, cdflib.DivideScalar(v0, v1)}, nil
			}},
	}
	funcs["Div"] = specs
	funcs["Divide"] = specs

	return funcs
}

var (
	functions = builtinFunctions()
)

type applyExpr struct {
	name Expression
	args []Expression
}

func (x *applyExpr) Type() ExpressionType {
	return ANY
}

func matchSpec(spec *funcSpec, argTypes []ExpressionType) (funcFn, []castFn, int) {
	// NOTE: We'll need to change this if we want to support varargs.
	if len(spec.Args) != len(argTypes) {
		return nil, nil, -1
	}
	dsum := 0
	castFns := make([]castFn, len(argTypes))
	for i, from := range(argTypes) {
		var d int
		d, castFns[i] = canCast(from, spec.Args[i].Type)
		if d < 0 {
			return nil, nil, -1
		}
		dsum += d
	}
	return spec.Fn, castFns, dsum
}

func (x *applyExpr) Eval(env Environment) (Expression, error) {
	if x.name.Type() != SYMBOL {
		return nil, errors.New("Not a function, cannot apply.")
	}
	name := x.name.(*lookupExpr).Name

	// Try and find the function by name
	specs, ok := functions[name]
	if !ok {
		return nil, fmt.Errorf("Undefined function \"%s\".", name)
	}

	// Evaluate all the args
	args := make([]Expression, len(x.args))
	argTypes := make([]ExpressionType, len(x.args))
	var err error
	for i, expr := range(x.args) {
		if args[i], err = expr.Eval(env); err != nil {		
			return nil, fmt.Errorf("Unable to evaluate arg %d of function \"%s\".", i, name)
		}
		argTypes[i] = args[i].Type()
	}

	// Find the overload that best matches the evaluated args.
	bestScore := 1000000
	var bestFn funcFn = nil
	var bestCastFns []castFn = nil
	for _, spec := range(specs) {
		fn, castFns, score := matchSpec(spec, argTypes)
		if fn == nil {
			continue
		}
		if score < bestScore {
			bestFn, bestCastFns, bestScore = fn, castFns, score
		}
	}
	if bestFn == nil {
		return nil, fmt.Errorf("Arguments do not match any overloads of function \"%s\".", name)
	}

	// Cast the evaluated args to the required type.
	for i, expr := range(args) {
		castFn := bestCastFns[i]
		if castFn != nil {
			args[i] = castFn(expr)
		}
	}

	// Apply the function
	return bestFn(args)
}

func (x *applyExpr) MimeData() MimeData {
	return MimeData{}
}

func (x *applyExpr) String() string {
	var args string
	for i, expr := range(x.args) {
		if i > 0 {
			args += ", "
		}
		args += expr.String()
	}
	return fmt.Sprintf("%s(%s)", x.name, args)
}

func Apply(name Expression, args []Expression) Expression {	
	return &applyExpr{name, args}
}

func Negate(expr Expression) Expression {
	return &applyExpr{&lookupExpr{"Negate"}, []Expression{expr}}
}

func Add(a Expression, b Expression) Expression {
	return &applyExpr{&lookupExpr{"Add"}, []Expression{a, b}}
}

func Sub(a Expression, b Expression) Expression {
	return &applyExpr{&lookupExpr{"Sub"}, []Expression{a, b}}
}

func Mult(a Expression, b Expression) Expression {
	return &applyExpr{&lookupExpr{"Mult"}, []Expression{a, b}}
}

func Div(a Expression, b Expression) Expression {
	return &applyExpr{&lookupExpr{"Div"}, []Expression{a, b}}
}
