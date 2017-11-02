package cdflib

type negate struct {
	cdf CDF
}

func (n *negate) P(x float64) float64 {
	return 1.0 - n.cdf.P(-x)
}

func (n *negate) Inverse() InverseCDF {
	return &negateInverse{n, n.cdf.Inverse()}
}

type negateInverse struct {
	n     *negate
	other InverseCDF
}

func (i *negateInverse) Value(p float64) float64 {
	return -i.other.Value(1.0 - p)
}

func (i *negateInverse) Inverse() CDF {
	return i.n
}

func Neg(cdf CDF) CDF {
	return &negate{cdf}
}

type shift struct {
	cdf CDF
	x   float64
}

func (s *shift) P(x float64) float64 {
	return s.cdf.P(x - s.x)
}

func (s *shift) Inverse() InverseCDF {
	return &shiftInverse{s, s.cdf.Inverse()}
}

type shiftInverse struct {
	s     *shift
	other InverseCDF
}

func (i *shiftInverse) Value(p float64) float64 {
	return i.other.Value(p) + i.s.x
}

func (i *shiftInverse) Inverse() CDF {
	return i.s
}

type scale struct {
	cdf CDF
	x   float64
}

func (s *scale) P(x float64) float64 {
	return s.cdf.P(x / s.x)
}

func (s *scale) Inverse() InverseCDF {
	return &scaleInverse{s, s.cdf.Inverse()}
}

type scaleInverse struct {
	s     *scale
	other InverseCDF
}

func (i *scaleInverse) Value(p float64) float64 {
	return i.other.Value(p) * i.s.x
}

func (i *scaleInverse) Inverse() CDF {
	return i.s
}

func AddScalar(cdf CDF, x float64) CDF {
	return &shift{cdf, x}
}

func SubtractScalar(cdf CDF, x float64) CDF {
	return &shift{cdf, -x}
}

func MultiplyScalar(cdf CDF, x float64) CDF {
	if x == 0 {
		return Constant(0)
	}
	if x < 0 {
		cdf = &negate{cdf}
		x = -x
	}
	return &scale{cdf, x}
}

func DivideScalar(cdf CDF, x float64) CDF {
	if x == 0 {
		return NaN
	}
	if x < 0 {
		cdf = &negate{cdf}
		x = -x
	}
	return &scale{cdf, 1.0 / x}
}
