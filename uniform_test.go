package cdflib

import (
	"math"
	"testing"
)

// Test standard values of P(x)
func TestUniform(t *testing.T) {
	cdf := checkSanity(Uniform(2, 6), t)
	x := []float64{1, 2, 3, 4, 5, 6, 7}
	expected := []float64{0, 0, 0.25, 0.5, 0.75, 1, 1}
	for i := 0; i < len(x); i += 1 {
		actual := cdf.P(x[i])
		if math.Abs(actual-expected[i]) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x[i], expected[i], actual)
		}
	}
}

func TestUniformInverse(t *testing.T) {
	inv := checkInverseSanity(Uniform(2, 6), t)
	p := []float64{-1, 0, 0.25, 0.5, 0.75, 1, 2}
	expected := []float64{2, 2, 3, 4, 5, 6, 6}
	for i := 0; i < len(p); i += 1 {
		actual := inv.Value(p[i])
		if math.Abs(actual-expected[i]) > EPSILON {
			t.Errorf("Expected Value(%f) to be %f, was %f.", p[i], expected[i], actual)
		}
	}
}

// Tests that the code handles swapped parameters gracefully.
func TestUniformSwap(t *testing.T) {
	cdf := checkSanity(Uniform(6, 2), t)
	x := 3.0
	expected := 0.25
	actual := cdf.P(x)
	if math.Abs(actual-expected) > EPSILON {
		t.Errorf("Expected P(%f) to be %f, was %f.", x, expected, actual)
	}
}

// Tests a zero-length range, where a == b.
func TestUniformPoint(t *testing.T) {
	cdf := Uniform(1, 1)
	if cdf != ONE {
		t.Errorf("Expected Uniform(1,1) == ONE, but it wasn't.")
	}
}
