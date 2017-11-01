package cdflib

import (
	"math"
	"testing"
)

func TestGeneric(t *testing.T) {
	cdf := Normal(2, 5)
	g := &generic{cdf.Inverse()}
	for x := -8.0; x < 12.0; x += 0.3 {
		expected := cdf.P(x)
		actual := g.P(x)
		if math.Abs(expected-actual) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x, expected, actual)
		}
	}
}

func TestGenericInverse(t *testing.T) {
	cdf := Normal(2, 5)
	ninv := cdf.Inverse()
	ginv := &genericInverse{cdf}
	for i := 1; i < 9; i += 1 {
		p := float64(i) / 10.0
		expected := ninv.Value(p)
		actual := ginv.Value(p)
		if math.Abs(expected-actual) > EPSILON {
			t.Errorf("Expected V(%f) to be %f, was %f.", p, expected, actual)
		}
	}
}
