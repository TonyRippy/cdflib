package cdflib

import (
	"math"
	"testing"
)

const (
	ONE_SIGMA_CDF = 0.841344746069
	TWO_SIGMA_CDF = 0.977249868052
)

func TestGaussian(t *testing.T) {
	cdf := checkSanity(Gaussian(2, 5), t)
	xs := []float64{
		math.Inf(-1),
		-8.0,
		-3.0,
		2.0,
		7.0,
		12.0,
		math.Inf(+1)}
	ps := []float64{
		0.0,
		1.0 - TWO_SIGMA_CDF,
		1.0 - ONE_SIGMA_CDF,
		0.5,
		ONE_SIGMA_CDF,
		TWO_SIGMA_CDF,
		1.0}
	for i, x := range xs {
		expected := ps[i]
		actual := cdf.P(x)
		if math.Abs(actual-expected) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x, expected, actual)
		}
	}
}

func TestGaussianInverse(t *testing.T) {
	cdf := Gaussian(2, 5)
	inv := checkInverseSanity(cdf, t)
	for x := -10.0; x <= +10.0; x += 1.0 {
		result := inv.Value(cdf.P(x))
		if math.Abs(result-x) > EPSILON {
			t.Errorf("Expected V(P(x)) != x. Expected %f, was %f.", x, result)
		}
	}
}

func TestLogNormal(t *testing.T) {
	cdf := checkSanity(LogNormal(2, 5), t)
	xs := []float64{-1, 0, 1, 2, 3, 4, 5, 10, 20, 100}
	ps := []float64{
		0, 0,
		0.3445782583897,
		0.3969033774375,
		0.4284673427443,
		0.4511560474127,
		0.4688693148147,
		0.5241280690952,
		0.5789259083397,
		0.6988284697464}
	for i, x := range xs {
		expected := ps[i]
		actual := cdf.P(x)
		if math.Abs(actual-expected) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", x, expected, actual)
		}
	}
}

func TestLogNormalInverse(t *testing.T) {
	cdf := LogNormal(2, 5)
	inv := checkInverseSanity(cdf, t)
	for x := 0.0; x <= 10.0; x += 0.5 {
		result := inv.Value(cdf.P(x))
		if math.Abs(result-x) > EPSILON {
			t.Errorf("Expected V(P(x)) != x. Expected %f, was %f.", x, result)
		}
	}
}
