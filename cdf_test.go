package cdflib

import (
	"math"
	"testing"
)

const (
	EPSILON = 1e-10
)

func checkSanity(cdf CDF, t *testing.T) CDF {
	// Test -Inf
	lastX := math.Inf(-1)
	lastP := cdf.P(lastX)
	if lastP != 0.0 {
		t.Errorf("Expected P(-Inf) to be 0, was %f.", lastP)
	}
	// Check that the CDF is monotonically increasing
	for x := -5.0; x <= +5.0; x += 0.1 {
		p := cdf.P(x)
		if p < 0.0 {
			t.Errorf("P(%f) < 0, was %f", x, p)
			break
		} else if p > 1.0 {
			t.Errorf("P(%f) > 1, was %f", x, p)
			break
		} else if p < lastP {
			t.Errorf(
				"CDF is not monotonically increasing! P(%.1f)=%f, P(%.1f)=%f",
				lastX, lastP, x, p)
			break
		}
		lastX = x
		lastP = p
	}
	// Test +Inf
	lastP = cdf.P(math.Inf(+1))
	if lastP != 1.0 {
		t.Errorf("Expected P(+Inf) to be 1, was %f.", lastP)
	}
	return cdf
}

func checkInverseSanity(cdf CDF, t *testing.T) InverseCDF {
	inv := cdf.Inverse()
	if inv.Inverse() != cdf {
		t.Errorf("Inverse() of Inverse() not equal to original CDF.")
	}
	return inv
}
