package cdflib

import (
	"math"
)

var (
	// A random variable that can only have a value of zero.
	ZERO = &discrete{[]float64{0}, []float64{1}}

	// A random variable that can only have a value of one.
	ONE = &discrete{[]float64{1}, []float64{1}}

	/*
  A special distribution that represents a nonsensical result.
  Functions only return a IEEE 754 “not-a-number” value.
	*/
	NaN = &nan{}
)

type nan struct {
}

func (s *nan) P(x float64) float64 {
	return math.NaN()
}

func (s *nan) Inverse() InverseCDF {
	return &nanInverse{s}
}

type nanInverse struct {
	cdf *nan
}

func (i *nanInverse) Value(p float64) float64 {
	return math.NaN()
}

func (i *nanInverse) Inverse() CDF {
	return i.cdf
}

/*
Creates a CDF that represents a constant real value.

Given the value x, the generated CDF has the following properties:
* P(v) ==>  { v < x: 0.0, v >= x: 1.0 }
* Inverse().Value(p) ==> { p <= 0: -Inf, 0 < p <= 1: x, 1 < p: +Inf }

Beware, because of its nature this function is not continuous!
*/
func Constant(x float64) CDF {
	// Look for special singletons
	if x == 0 {
		return ZERO
	} else if x == 1 {
		return ONE
	}
	return &discrete{[]float64{x}, []float64{1.0}}
}
