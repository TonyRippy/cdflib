package cdflib

import (
	"math"
	"testing"
)

func TestMakeCdfFromSamples(t *testing.T) {
	samples := make([]float64, 20)
	for i := 1; i <= 10; i += 1 {
		samples[i-1] = float64(i)
	}
	for i := 10; i > 0; i -= 1 {
		samples[20-i] = float64(i)
	}
	cdf := checkSanity(MakeECDF(samples), t)
	for x := 1.0; x <= 10.0; x += 1.0 {
		expected := x / 10.0;
		actual := cdf.P(x)
		if math.Abs(actual-expected) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x, expected, actual)
		}
	}
}
