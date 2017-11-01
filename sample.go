package cdflib

import (
	"math/rand"
	"sort"
)

/*
An emperical distribution function.

This is a distribution function generated from a collection of sample data,
estimating the cumulative distribution function underlying the points
in the sample.

See:
https://en.wikipedia.org/wiki/Empirical_distribution_function
*/
type ECDF struct {
	discrete

	// The number of samples used to generate this emperical distribution.
	Size int
}

/*
Creates an emperical distribution from a set of samples.
Assumes the samples are sorted. If it detects that they are not, it will sort
the array and start over.
*/
func MakeECDF(samples []float64) *ECDF {
	n := len(samples)
	xs := make([]float64, 0, n)
	ps := make([]float64, 0, n)
	if n > 0 {
		total := float64(n)
		last := samples[0]
		for i := 1; i < n; i++ {
			x := samples[i]
			if x == last {
				continue
			}
			if x < last {
				// Not sorted! Sort the array and try again.
				xs = nil
				ps = nil
				sort.Float64s(samples)
				return MakeECDF(samples)
			}
			// Add an entry for the last value
			xs = append(xs, last)
			ps = append(ps, float64(i)/total)
			last = x
		}
		// Add the last entry
		xs = append(xs, last)
		ps = append(ps, 1.0)
	}
	return &ECDF{discrete{xs, ps}, n}
}

// Draws evenly-spaced samples of P(x) in the range (min, max), inclusive.
func UniformSamples(cdf CDF, min float64, max float64, n int) []float64 {
	if n <= 0 {
		return []float64{}
	}
	if n == 1 {
		return []float64{cdf.P(min)}
	}
	out := make([]float64, n)
	if max < min {
		min, max = max, min
	}
	step := (max - min) / float64(n-1)
	x := min
	for i := 0; i < (n - 1); i += 1 {
		out[i] = cdf.P(x)
		x += step
	}
	out[n-1] = cdf.P(max)
	return out
}

// Draws evenly-spaced samples of PDF(x) in the range (min, max), inclusive.
func UniformDensitySamples(cdf CDF, min float64, max float64, n int) []float64 {
	if n <= 0 {
		return []float64{}
	}
	if n == 1 {
		return []float64{cdf.P(min) - cdf.P(max)}
	}
	out := make([]float64, n)
	if max < min {
		min, max = max, min
	}
	step := (max - min) / float64(n+1)
	x := min
	p := cdf.P(min)
	for i := 1; i < (n - 1); i += 1 {
		x += step
		p2 := cdf.P(x)
		out[i] = (p2 - p) / step
		p = p2
	}
	p2 := cdf.P(max)
	out[n-1] = (p2 - p) / (max - x)
	return out
}

func randomSample(inv InverseCDF) float64 {
	var p float64
	for {
		p = rand.Float64()
		if p != 0.0 {
			return inv.Value(p)
		}
	}
}

// Draws a single random sample from a distribution.
func RandomSample(cdf CDF) float64 {
	return randomSample(cdf.Inverse())
}

// Draws random samples from a distribution and returns them as a slice.
func RandomSamples(cdf CDF, n int) []float64 {
	if n <= 0 {
		return []float64{}
	}
	inv := cdf.Inverse()
	out := make([]float64, n)
	for i := 0; i < n; i += 1 {
		out[i] = randomSample(inv)
	}
	return out
}
