package cdflib

type uniform struct {
	min, max float64
}

func (u *uniform) P(x float64) float64 {
	if x >= u.max {
		return 1
	}
	if x <= u.min {
		return 0
	}
	return (x - u.min) / (u.max - u.min)
}

func (u *uniform) Inverse() InverseCDF {
	return &uniformInverse{u}
}

type uniformInverse struct {
	u *uniform
}

func (i *uniformInverse) Value(p float64) float64 {
	if p <= 0.0 {
		return i.u.min
	}
	if p >= 1.0 {
		return i.u.max
	}
	return i.u.min + p*(i.u.max-i.u.min)
}

func (i *uniformInverse) Inverse() CDF {
	return i.u
}

/*
Creates a uniform distribution in the range (a, b).

In probability theory and statistics, the continuous uniform distribution or
rectangular distribution is a family of symmetric probability distributions such
that for each member of the family, all intervals of the same length on the
distribution's support are equally probable. The support is defined by the two
parameters, a and b, which are its minimum and maximum values. The distribution
is often abbreviated U(a,b).

Source:
https://en.wikipedia.org/wiki/Uniform_distribution_(continuous)
*/
func Uniform(a, b float64) CDF {
	if a == b {
		return Constant(a)
	}
	if a > b {
		a, b = b, a
	}
	return &uniform{a, b}
}

// Shorthand for Uniform(a, b)
func U(a, b float64) CDF {
	return Uniform(a, b)
}
