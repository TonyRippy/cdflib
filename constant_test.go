package cdflib

import (
	"math"
	"testing"
)

func TestZero(t *testing.T) {
	if Constant(0) != ZERO {
		t.Errorf("Expected Constant(0) to return the singleton ZERO.")
	}
	checkSanity(ZERO, t)
}

func TestOne(t *testing.T) {
	if Constant(1) != ONE {
		t.Errorf("Expected Constant(1) to return the singleton ONE.")
	}
	checkSanity(ONE, t)
}

func TestConstant(t *testing.T) {
	cdf := checkSanity(Constant(2), t)
	input := []float64{-1, 0, 1, 2, 3}
	expected := []float64{0, 0, 0, 1, 1}
	for i, x := range input {
		actual := cdf.P(x)
		if math.Abs(expected[i]-actual) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x, expected[i], actual)
		}
	}
}

func TestConstantInverse(t *testing.T) {
	inv := checkInverseSanity(Constant(2), t)
	input := []float64{-1, 0, 0.5, 1, 1.5}
	expected := []float64{math.Inf(-1), math.Inf(-1), 2, 2, math.Inf(+1)}
	for i, p := range input {
		actual := inv.Value(p)
		if math.Abs(expected[i]-actual) > EPSILON {
			t.Errorf("Expected V(%f) to be %f, was %f.", p, expected[i], actual)
		}
	}
}
