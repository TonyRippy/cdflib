package cdflib

import (
	math2 "github.com/TonyRippy/math"
	"math"
)

type negate struct {
	cdf CDF
}

func (n *negate) P(x float64) float64 {
	return 1.0 - n.cdf.P(-x)
}

func (n *negate) Inverse() InverseCDF {
	return &negateInverse{n, n.cdf.Inverse()}
}

type negateInverse struct {
	n     *negate
	other InverseCDF
}

func (i *negateInverse) Value(p float64) float64 {
	return -i.other.Value(1.0 - p)
}

func (i *negateInverse) Inverse() CDF {
	return i.n
}

/*
Given a random variable X represented by a distribution, calculate -X.
*/
func Neg(cdf CDF) CDF {
	return &negate{cdf}
}

type shift struct {
	cdf CDF
	x   float64
}

func (s *shift) P(x float64) float64 {
	return s.cdf.P(x - s.x)
}

func (s *shift) Inverse() InverseCDF {
	return &shiftInverse{s, s.cdf.Inverse()}
}

type shiftInverse struct {
	s     *shift
	other InverseCDF
}

func (i *shiftInverse) Value(p float64) float64 {
	return i.other.Value(p) + i.s.x
}

func (i *shiftInverse) Inverse() CDF {
	return i.s
}

type scale struct {
	cdf CDF
	x   float64
}

func (s *scale) P(x float64) float64 {
	return s.cdf.P(x / s.x)
}

func (s *scale) Inverse() InverseCDF {
	return &scaleInverse{s, s.cdf.Inverse()}
}

type scaleInverse struct {
	s     *scale
	other InverseCDF
}

func (i *scaleInverse) Value(p float64) float64 {
	return i.other.Value(p) * i.s.x
}

func (i *scaleInverse) Inverse() CDF {
	return i.s
}

/*
Given a random variable X represented by a distribution and a scalar value v, calculate X + v.
*/
func AddScalar(cdf CDF, v float64) CDF {
	return &shift{cdf, v}
}

/*
Given a random variable X represented by a distribution and a scalar value v, calculate X - v.
*/
func SubtractScalar(cdf CDF, v float64) CDF {
	return &shift{cdf, -v}
}

/*
Given a random variable X represented by a distribution and a scalar value v, calculate X * v.
*/
func MultiplyScalar(cdf CDF, v float64) CDF {
	if v == 0 {
		return Constant(0)
	}
	if v < 0 {
		cdf = &negate{cdf}
		v = -v
	}
	return &scale{cdf, v}
}

/*
Given a random variable X represented by a distribution and a scalar value v, calculate X / v.
Returns NaN if v == 0.
*/
func DivideScalar(cdf CDF, v float64) CDF {
	if v == 0 {
		return NaN
	}
	if v < 0 {
		cdf = &negate{cdf}
		v = -v
	}
	return &scale{cdf, 1.0 / v}
}

type addc struct {
	a    DifferentiableCDF
	b    CDF
	minx float64
}

func (c *addc) P(x float64) float64 {
	const (
		EPS     = 1e-10
		SMALL_X = -1e30 // TODO: What if x < SMALL_X?
	)
	if math.IsInf(x, +1) {
		return 1
	}
	if math.IsInf(x, -1) {
		return 0
	}
	f := func(y float64) float64 {
		a := c.a.DX(y)
		if a == 0.0 {
			return 0.0
		}
		b := c.b.P(x - y)
		return a * b
	}
	var minp float64
	if x <= c.minx {
		minp, _ = math2.Romberg(math2.MidPointInf(f, SMALL_X, x), EPS)
		return minp
	}
	minp, _ = math2.Romberg(math2.MidPointInf(f, SMALL_X, c.minx), EPS)
	p, _ := math2.Romberg(math2.MidPoint(f, c.minx, x), EPS)
	return minp + p
}

func (c *addc) Inverse() InverseCDF {
	return &genericInverse{c}
}

/*
Given random variables A and B, return a random variable that models A + B.
*/
func Add(a DifferentiableCDF, b CDF) CDF {
	const (
		SMALL_P = 1e-6
	)
	minx := a.Inverse().Value(SMALL_P)
	x := b.Inverse().Value(SMALL_P)
	if x < minx {
		minx = x
	}
	if minx >= 0 {
		minx = math.Copysign(0, -1) // negative zero
	}
	return &addc{a, b, minx}
}
