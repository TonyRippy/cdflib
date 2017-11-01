package cdflib

import (
	"math"
)

type generic struct {
	inv InverseCDF
}

func (g *generic) Inverse() InverseCDF {
	return g.inv
}

func (g *generic) P(x float64) float64 {
	minP := 0.0
	maxP := 1.0
	// Perform a binary search within the bounds.
	const EPSILON = 1e-12
	for {
		mid := (maxP + minP) / 2
		cx := g.inv.Value(mid)
		if math.Abs(cx-x) < EPSILON {
			return mid
		}
		if cx < x {
			minP = mid
		} else {
			maxP = mid
		}
	}
}

type genericInverse struct {
	cdf CDF
}

func (i *genericInverse) Inverse() CDF {
	return i.cdf
}

// Finds the inverse of a CDF using a binary search of the range.
func (i *genericInverse) Value(p float64) float64 {
	cp := i.cdf.P(0)
	// Do an exponential search out for the bounds to use for the search
	var minX, maxX float64
	if cp <= p {
		// Look for an upper bound where P(x) > p
		minX = 0
		maxX = 1
		for {
			cp = i.cdf.P(maxX)
			if cp > p {
				break
			}
			minX = maxX
			maxX *= 2
		}
	} else { // cp > p
		// Look for an lower bound where P(x) < p
		minX = -1
		maxX = 0
		for {
			cp = i.cdf.P(minX)
			if cp <= p {
				break
			}
			maxX = minX
			minX *= 2
		}
	}
	// Now perform a binary search within the bounds.
	const EPSILON = 1e-12
	for {
		mid := (maxX + minX) / 2
		cp = i.cdf.P(mid)
		if math.Abs(cp-p) < EPSILON {
			return mid
		}
		if cp < p {
			minX = mid
		} else {
			maxX = mid
		}
	}
}
