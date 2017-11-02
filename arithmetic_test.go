package cdflib

import (
	"math"
	"testing"
)

func TestNegate(t *testing.T) {
	cdf := Gaussian(2, 5)
	neg := checkSanity(Neg(cdf), t)
	expected := 0.5
	actual := neg.P(-2)
	if (actual - expected) > EPSILON {
		t.Errorf("Expected P(mean) to be %f, was %f.", expected, actual)
	}
	expected = cdf.P(-3) // +2 - 5
	actual = neg.P(-7)   // -2 - 5
	if (actual - expected) > EPSILON {
		t.Errorf("Expected P(mean - stddev) to be %f, was %f.", expected, actual)
	}
	t.Logf("Test Neg(Neg(X)) == X")
	neg2 := checkSanity(Neg(neg), t)
	for x := -10.0; x <= +10.0; x += 1.0 {
		expected = cdf.P(x)
		actual := neg2.P(x)
		if math.Abs(actual-expected) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x, expected, actual)
		}
	}
}

func TestNegateInv(t *testing.T) {
	cdf := Neg(Gaussian(2, 5))
	inv := checkInverseSanity(cdf, t)
	p := 0.5
	expected := -2.0
	actual := inv.Value(p)
	if (actual - expected) > EPSILON {
		t.Errorf("Expected V(%f) to be %f, was %f.", p, expected, actual)
	}
}
