package main

import (
	"fmt"
	"math"
)

// An "if" is not surrounded by parentheses, and the braces are necessary
func NormalIf(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}

	return lim
}

// "if" shorthanded.
// Note that "v" is only available inside the "if" scope
func IfShorthanded(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

// If shorthanded with else.
// Note that variables defined inside "if" are available inside "else"
func IfShorthandedWithElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// %g	%e for large exponents, %f otherwise.
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// can't use "v" here, though
	return lim
}
