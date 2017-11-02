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
