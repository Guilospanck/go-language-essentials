package main

import "fmt"

func Sqrt(x float64) float64 {
	initialGuess := 1.0
	for i := 0; i < 10; i++ {
		initialGuess = initialGuess - (initialGuess*initialGuess-x)/(2*initialGuess) // Newton-Raphson method
		fmt.Println(initialGuess)
	}

	return initialGuess
}

func SqrtAutomatic(x float64) float64 {
	z := 1.0
	for {
		z1 := z
		z -= (z*z - x) / (2 * z)                  // Newton-Rhapson method
		if (z1-z) >= 0 && (z1-z) <= 0.000000005 { // will stop when the error is too small
			break
		}
	}

	return z
}
