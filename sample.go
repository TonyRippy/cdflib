package cdflib

import (
	"math"
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


/*
Draws n samples from a distribution such that the samples are representative
of the original distribution.

Specifically, it uses the inverse CDF and samples V(p) where p is in the range (0,1).
We want the output samples to approximate the distribution, so this takes the following
approach to small sample sizes:
  0 samples = [ ]
  1 samples = [ V(1) ]
  2 samples = [ V(0), V(1) ]
  3 samples = [ V(0), V(1/2), V(1) ]
  4 samples = [ V(0), V(1/3), V(2/3), V(1) ]
...and so forth.

The one exception is that the boundary values V(0) & V(1) are not included
if they are +/- Inf.
*/
func UniformSamples(cdf CDF, n int) []float64 {
	vs := make([]float64, 0, n)
	if n <= 0 {
		return vs
	}
	inv := cdf.Inverse()
	max := inv.Value(1)
	if !math.IsInf(max, 1) {
		n -= 1
	}
	if n > 0 {
		min := inv.Value(0)
		if !math.IsInf(min, -1) {
			vs = append(vs, min)
			n -= 1
		}
	}
	if n > 0 {
		d := float64(n+1)
		for i := 1; i <= n; i += 1 {
			p := float64(i) / d
			v := inv.Value(p)
			vs = append(vs, v)
		}
	}
	if !math.IsInf(max, 1) {
		vs = append(vs, max)
	}
	return vs
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
