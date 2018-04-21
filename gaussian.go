package cdflib

import (
	math2 "github.com/TonyRippy/math"
	"math"
)

var (
	sqrt2Pi    = math.Sqrt(2.0 * math.Pi)
	invSqrt2Pi = 1.0 / sqrt2Pi
)

// The functions for the standard normal distribution.
// (μ=0 and σ=1)
func pdf(x float64) float64 {
	return math.Exp(-0.5*x*x) * invSqrt2Pi
}

func cdf(x float64) float64 {
	return 0.5 * (1.0 + math.Erf(x/math.Sqrt2))
}

func cdfInv(p float64) float64 {
	return math.Sqrt2 * math2.ErfInv(2*p-1)
}

type gaussian struct {
	mean   float64
	stddev float64
}

func (n *gaussian) P(x float64) float64 {
	return cdf((x - n.mean) / n.stddev)
}

func (n *gaussian) DX(x float64) float64 {
	return pdf((x-n.mean)/n.stddev) / n.stddev
}

func (n *gaussian) Inverse() InverseCDF {
	return &gaussianInverse{n}
}

type gaussianInverse struct {
	n *gaussian
}

func (i *gaussianInverse) Value(p float64) float64 {
	return i.n.mean + i.n.stddev*cdfInv(p)
}

func (i *gaussianInverse) Inverse() CDF {
	return i.n
}

// Returns the CDF for a Gaussian (or Normal) distribution.
func Gaussian(mean, stddev float64) DifferentiableCDF {
	return &gaussian{mean, stddev}
}

// Returns the CDF for a Normal (or Gaussian) distribution.
func Normal(mean, stddev float64) DifferentiableCDF {
	return Gaussian(mean, stddev)
}

// Shorthand for Normal(mean, stddev)
func N(mean, stddev float64) DifferentiableCDF {
	return Gaussian(mean, stddev)
}

type logNormal struct {
	mean   float64
	stddev float64
}

func (n *logNormal) P(x float64) float64 {
	if x <= 0 {
		return 0
	}
	return cdf((math.Log(x) - n.mean) / n.stddev)
}

func (n *logNormal) DX(x float64) float64 {
	return pdf((math.Log(x)-n.mean)/n.stddev) / (n.stddev * x)
}

func (n *logNormal) Inverse() InverseCDF {
	return &logNormalInverse{n}
}

type logNormalInverse struct {
	n *logNormal
}

func (i *logNormalInverse) Value(p float64) float64 {
	return math.Exp(i.n.mean + i.n.stddev*cdfInv(p))
}

func (i *logNormalInverse) Inverse() CDF {
	return i.n
}

// Returns the CDF for a LogNormal distribution.
// https://en.wikipedia.org/wiki/Log-normal_distribution
func LogNormal(mean, stddev float64) DifferentiableCDF {
	return &logNormal{mean, stddev}
}
