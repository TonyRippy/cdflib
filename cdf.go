/*
Package cdflib provides constructs for manipulating cumulative distribution functions.
*/
package cdflib

/*
The cumulative distrubution function for a real-valued random variable.
https://en.wikipedia.org/wiki/Cumulative_distribution_function
*/
type CDF interface {
	/*
	  Returns the probability that the random variable will take a
	  value less than or equal to x.
	*/
	P(x float64) float64

	/*
	   Returns the inverse of this CDF.
	*/
	Inverse() InverseCDF
}

/*
The inverse of a cumulative distribution function.
https://en.wikipedia.org/wiki/Quantile_function
*/
type InverseCDF interface {

	/*
	  Returns the value at which the probability of the random variable
	  being less than or equal to that value is equal to the given probability.
	*/
	Value(p float64) float64

	Inverse() CDF
}
