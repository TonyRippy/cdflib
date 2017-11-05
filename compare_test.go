package cdflib

import (
	"testing"
)

func TestMin(t *testing.T) {
	cdf1 := Normal(2, 5)
	cdf2 := Normal(3, 1)
	m := Min(cdf1, cdf2)
	checkSanity(m, t)

	// Compute the result by sampling.
	min := func(a float64, b float64) float64 {
		if (a < b) {
			return a
		}
		return b
	}
	expected, _ := Apply2(cdf1, cdf2, min, 500)
	samples := UniformSamples(expected, 1000)
	
	// Null hypothesis is that the manually computed samples above
	// were drawn from Min(cdf1,cdf2).
	const a = 0.001
	p := 1.0 - KSTest(m, samples)
	if a < p {
		t.Errorf("Null hypothesis was rejected. a = %.2f, was %f.", a, p)
	}
}
	
func TestMax(t *testing.T) {
	cdf1 := Normal(2, 5)
	cdf2 := Normal(3, 1)
	m := Max(cdf1, cdf2)
	checkSanity(m, t)

	// Compute the result by sampling.
	max := func(a float64, b float64) float64 {
		if (a > b) {
			return a
		}
		return b
	}
	expected, _ := Apply2(cdf1, cdf2, max, 500)
	samples := UniformSamples(expected, 1000)
	
	// Null hypothesis is that the manually computed samples above
	// were drawn from Max(cdf1,cdf2).
	const a = 0.001
	p := 1.0 - KSTest(m, samples)
	if a < p {
		t.Errorf("Null hypothesis was rejected. a = %.2f, was %f.", a, p)
	}
}
