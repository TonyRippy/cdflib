package cdflib

import (
	"fmt"
	"math"
	"sort"
)

// Round to nearest integer. Rounds half integers to the nearest even integer.
func nint(x float64) int {
	var i int
	if x >= 0 {
		i = int(x + 0.5)
		if (i&1) != 0 && (x+0.5) == float64(i) {
			i -= 1
		}
	} else {
		i = int(x - 0.5)
		if (i&1) != 0 && (x-0.5) == float64(i) {
			i += 1
		}
	}
	return i
}

/*
Calculates the Kolmogorov distribution function,
which gives the probability that Kolmogorov's test statistic will exceed
the value z assuming the null hypothesis. This gives a very powerful
test for comparing two one-dimensional distributions.
see, for example, Eadie et al, "statistocal Methods in Experimental
Physics', pp 269-270).

This function returns the confidence level for the null hypothesis, where:
  z  = dn*sqrt(n), and
  dn = is the maximum deviation between a hypothetical distribution
       function and an experimental distribution with
  n  = events

NOTE: To compare two experimental distributions with m and n events,
use z = sqrt(m*n/(m+n))*dn

Accuracy: The function is far too accurate for any imaginable application.
Probabilities less than 10^-15 are returned as zero.
However, remember that the formula is only valid for "large" n.
Theta function inversion formula is used for z <= 1

Ported from CERN's Root data analysis framework. (https://root.cern.ch/)
Specifically the TMath::KolmogorovProb() function, originally written in C++.
Source here: https://root.cern.ch/root/html/src/TMath.cxx.html
*/
func kprob(z float64) float64 {
	var p float64
	if z < 0.2 {
		p = 1
	} else if z < 0.755 {
		const w = 2.50662827
		// c1 - -pi**2/8, c2 = 9*c1, c3 = 25*c1
		const c1 = -1.2337005501361697
		const c2 = -11.103304951225528
		const c3 = -30.842513753404244
		v := 1.0 / (z * z)
		p = 1 - w*(math.Exp(c1*v)+math.Exp(c2*v)+math.Exp(c3*v))/z
	} else if z < 6.8116 {
		fj := [4]float64{-2, -8, -18, -32}
		r := [4]float64{0, 0, 0, 0}
		v := z * z
		maxj := nint(3.0 / z)
		if maxj < 1 {
			maxj = 1
		}
		for j := 0; j < maxj; j += 1 {
			r[j] = math.Exp(fj[j] * v)
		}
		p = 2 * (r[0] - r[1] + r[2] - r[3])
	} else {
		p = 0
	}
	return p
}

/*
Runs a Kolmogorov-Smirnov test for a given sample and reference distribution.

The returned value is the calculated confidence level, an estimate of the
likelihood that the given sample comes from the reference distribution.

See:
https://en.wikipedia.org/wiki/Kolmogorov%E2%80%93Smirnov_test
*/
func KolmogorovSmirnovTest(cdf CDF, sample []float64) float64 {
	// Find the maximum difference between the sample and the reference distribution.
	sort.Float64s(sample)
	n := float64(len(sample))
	max := 0.0
	ip := 0.0
	for i, x := range sample {
		p := cdf.P(x)
		diff := p - ip
		if diff < 0 {
			diff = -diff
		}
		if diff > max {
			max = diff
		}
		ip = float64(i+1) / n
		diff = ip - p
		if diff < 0 {
			diff = -diff
		}
		if diff > max {
			max = diff
		}
	}
	z := max * math.Sqrt(n)
	p := kprob(z)
	fmt.Printf("N: %d, Max: %f, K(%f) = %f\n", len(sample), max, z, p)
	return p
}

// Shorthand for KolmogorovSmirnovTest().
func KSTest(cdf CDF, sample []float64) float64 {
	return KolmogorovSmirnovTest(cdf, sample)
}
