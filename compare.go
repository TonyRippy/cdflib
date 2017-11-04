package cdflib

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
