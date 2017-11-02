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
			t.Errorf("Expected P(%f) == %f, was %f.", x, expected, actual)
		}
	}
}

func testUniformSampleSize(t *testing.T, cdf CDF, n int, expected []float64) {
	samples := UniformSamples(cdf, n)
	if len(samples) != len(expected) {
		t.Errorf("Expected %d samples, got %d.", len(expected), len(samples))
		return
	}
	for i, actual := range(samples) {
		if math.Abs(actual-expected[i]) > EPSILON {
			t.Errorf("Expected Sample(%d)[%d] == %f, was %f.", n, i, expected[i], actual)
		}
	}
}

func TestUniformSamples(t *testing.T) {
	cdf := Uniform(0, 10)
	testUniformSampleSize(t, cdf, 0, []float64{})
	testUniformSampleSize(t, cdf, 1, []float64{10})
	testUniformSampleSize(t, cdf, 2, []float64{0,10})
	testUniformSampleSize(t, cdf, 3, []float64{0,5,10})
	testUniformSampleSize(t, cdf, 4, []float64{0,10/3.0,20/3.0,10})
	testUniformSampleSize(t, cdf, 11, []float64{0,1,2,3,4,5,6,7,8,9,10})
}
