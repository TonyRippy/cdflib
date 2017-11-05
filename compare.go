package cdflib

type min struct {
	cdfs []CDF
}

func (m *min) P(x float64) float64 {
	// What we really want is the probability that one or more values sampled from
	// the input distributions are <= x. The easiest way to calculate this is to
	// find the probability that *no* samples are less than x:
	p := 1.0
	for _, cdf := range m.cdfs {
		// P(x) is the probability that a sample from cdf is <= x.
		// 1 - P(x) is the probability it is > x.
		p *= 1.0 - cdf.P(x)
	}
	// ... and we return the probability that this is not true:
	return 1.0 - p
}

func (m *min) Inverse() InverseCDF {
	return &genericInverse{m}
}

/*
Calculates the minimum of a set of random variables.
*/
func Min(cdfs ...CDF) CDF {
	return &min{cdfs}
}

type max struct {
	cdfs []CDF
}

func (m *max) P(x float64) float64 {
	p := 1.0
	for _, cdf := range m.cdfs {
		p *= cdf.P(x)
	}
	return p
}

func (m *max) Inverse() InverseCDF {
	return &genericInverse{m}
}

/*
Calculates the maximum of a set of random variables.
*/
func Max(cdfs ...CDF) CDF {
	return &max{cdfs}
}
