package cdflib

import (
	"math"
)

type discrete struct {
	// The discrete values for this random variable.
	x []float64
	// The cumulative probability for the discrete values.
	p []float64
}

func (cdf *discrete) P(x float64) float64 {
	// TODO: Use a more intellegent data structure, maybe a binary tree?
	last := 0.0
	for i, v := range cdf.x {
		if v > x {
			return last
		}
		last = cdf.p[i]
	}
	return 1.0
}

func (cdf *discrete) Inverse() InverseCDF {
	return &discreteInverse{cdf}
}

type discreteInverse struct {
	cdf *discrete
}

func (inv *discreteInverse) Value(p float64) float64 {
	last := math.Inf(+1)
	for i := len(inv.cdf.p) - 1; i >= 0; i -= 1 {
		if p > inv.cdf.p[i] {
			return last
		}
		last = inv.cdf.x[i]
	}
	if p > 0 {
		return last
	}
	return math.Inf(-1)
}

func (inv *discreteInverse) Inverse() CDF {
	return inv.cdf
}
