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

func TestAddScalar(t *testing.T) {
	cdf := Gaussian(2, 5)
	const b = 2.0
	t.Logf("Add(%.1f)", b)
	cdf2 := checkSanity(AddScalar(cdf, b), t)
	for x := -10.0; x <= +10.0; x += 1.0 {
		y := x + b
		expected := cdf.P(x)
		actual := cdf2.P(y)
		if math.Abs(actual-expected) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", y, expected, actual)
		}
	}
}

func TestAddScalarInv(t *testing.T) {
	cdf := AddScalar(Gaussian(0, 1), 2)
	inv := checkInverseSanity(cdf, t)
	p := 0.5
	expected := 2.0
	actual := inv.Value(p)
	if (actual - expected) > EPSILON {
		t.Errorf("Expected V(%f) to be %f, was %f.", p, expected, actual)
	}
}

func TestSubtractScalar(t *testing.T) {
	cdf := Gaussian(2.0, 5.0)
	const b = 3.0
	t.Logf("Subtract(%.1f)", b)
	cdf2 := checkSanity(SubtractScalar(cdf, b), t)
	for x := -10.0; x <= +10.0; x += 1.0 {
		y := x - b
		expected := cdf.P(x)
		actual := cdf2.P(y)
		if math.Abs(actual-expected) > EPSILON {
			t.Errorf("Expected P(%f) to be %f, was %f.", y, expected, actual)
		}
	}
}

func TestSubtractScalarInv(t *testing.T) {
	cdf := SubtractScalar(Gaussian(5, 1), 2)
	inv := checkInverseSanity(cdf, t)
	p := 0.5
	expected := 3.0
	actual := inv.Value(p)
	if (actual - expected) > EPSILON {
		t.Errorf("Expected V(%f) to be %f, was %f.", p, expected, actual)
	}
}

func TestMultiplyScalar(t *testing.T) {
	cdf := Gaussian(2, 5)
	for m := -3.0; m <= 3.0; m += 0.5 {
		t.Logf("Multiply(%.1f)", m)
		mult := checkSanity(MultiplyScalar(cdf, m), t)
		if m == 0 {
			// Special case: CDF * 0 = 0 for all x.
			if mult != ZERO {
				t.Errorf("Expected Mult(0) == ZERO, it wasn't.")
			}
		} else {
			expected := 0.5
			actual := mult.P(2 * m)
			if math.Abs(actual-expected) > EPSILON {
				t.Errorf("Expected P(mean) to be %f, was %f.", expected, actual)
			}
			expected = cdf.P(2 + 5)
			if m < 0 {
				expected = 1 - expected
			}
			actual = mult.P((2 + 5) * m)
			if math.Abs(actual-expected) > EPSILON {
				t.Errorf("Expected P(mean + stddev) to be %f, was %f.", expected, actual)
			}
		}
	}
}

func TestMultiplyScalarInv(t *testing.T) {
	cdf := MultiplyScalar(Gaussian(5, 10), 5)
	inv := checkInverseSanity(cdf, t)
	p := 0.5
	expected := 25.0
	actual := inv.Value(p)
	if (actual - expected) > EPSILON {
		t.Errorf("Expected V(%f) to be %f, was %f.", p, expected, actual)
	}
}

func TestDivideScalar(t *testing.T) {
	cdf := Gaussian(2, 5)
	for m := -3.0; m <= 3.0; m += 0.5 {
		t.Logf("Divide(%.1f)", m)
		div := DivideScalar(cdf, m)
		if m == 0 {
			// Special case: CDF / 0 is invalid, will return a nonsense result.
			if div != NaN {
				t.Errorf("Expected Mult(0) == NaN, it wasn't.")
			}
		} else {
			div := checkSanity(div, t)
			expected := 0.5
			actual := div.P(2 / m)
			if math.Abs(actual-expected) > EPSILON {
				t.Errorf("Expected P(mean) to be %f, was %f.", expected, actual)
			}
			expected = cdf.P(2 + 5)
			if m < 0 {
				expected = 1 - expected
			}
			actual = div.P((2 + 5) / m)
			if math.Abs(actual-expected) > EPSILON {
				t.Errorf("Expected P(mean + stddev) to be %f, was %f.", expected, actual)
			}
		}
	}
}

func TestDivideScalarInv(t *testing.T) {
	cdf := DivideScalar(Gaussian(10, 5), 5)
	inv := checkInverseSanity(cdf, t)
	p := 0.5
	expected := 2.0
	actual := inv.Value(p)
	if (actual - expected) > EPSILON {
		t.Errorf("Expected V(%f) to be %f, was %f.", p, expected, actual)
	}
}
