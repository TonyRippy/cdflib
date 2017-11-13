package cdflib

import (
	"math"
	"fmt"
)

/*
Applies a unary function to a given random variable using sampling.

The return value is an empirical distribution function, calculated by applying 
the given function to n uniform samples from the distribution.
Non-real samples such as +/-Inf and NaN are dropped.
 
This is a brute force but simple approach to applying functions to distributions.
There may be more direct ways to find the output distribution by using the underlying maths.
In general that approach should be preferred over this one.
*/
func Apply1(A CDF, f func(float64) float64, n int) (*ECDF, error) {
	if n < 2 {
		return nil, fmt.Errorf("You must request at least two samples, requested %d.", n)
	}
	samples := UniformSamples(A, n)
	vs := make([]float64, 0, n)
	for _, av := range(samples) {
		v := f(av)
		if !(math.IsInf(v, 0) || math.IsNaN(v)) {
			vs = append(vs, v)
		}
	}		
	return MakeECDF(vs), nil
}

/*
Applies a binary function f to two random variables A & B using sampling.

The return value is an empirical distribution function, calculated by applying 
the given function to the cross product of n uniform samples from each distribution.
Non-real samples such as +/-Inf and NaN are dropped.
 
Please note that the resulting ECDF may have up to O(n^2) samples.
You may choose to reduce this by using the ECDF's Downsample function.
 
This is a brute force but simple approach to applying functions to distributions.
There may be more direct ways to find the output distribution by using the underlying maths.
In general that approach should be preferred over this one.
*/
func Apply2(A CDF, B CDF, f func(float64, float64) float64, n int) (*ECDF, error) {
	if n < 2 {
		return nil, fmt.Errorf("You must request at least two samples, requested %d.", n)
	}
	as := UniformSamples(A, n)
	bs := UniformSamples(B, n)
	vs := make([]float64, 0, n*n)
	for _, a := range(as) {
		for _, b := range(bs) {
			v := f(a, b)
			if !(math.IsInf(v, 0) || math.IsNaN(v)) {
				vs = append(vs, v)
			}
		}
	}
	return MakeECDF(vs), nil
}

/*
Like Apply2, but for three random variables.

Please note that the resulting ECDF may have up to O(n^3) samples.
You may choose to reduce this by using the ECDF's Downsample function.
*/
func Apply3(A CDF, B CDF, C CDF, f func(float64, float64, float64) float64, n int) (*ECDF, error) {
	if n < 2 {
		return nil, fmt.Errorf("You must request at least two samples, requested %d.", n)
	}
	as := UniformSamples(A, n)
	bs := UniformSamples(B, n)
	cs := UniformSamples(C, n)
	vs := make([]float64, 0, n*n*n)
	for _, a := range(as) {
		for _, b := range(bs) {
			for _, c := range(cs) {
				v := f(a, b, c)
				if !(math.IsInf(v, 0) || math.IsNaN(v)) {
					vs = append(vs, v)
				}
			}
		}
	}
	return MakeECDF(vs), nil
}
