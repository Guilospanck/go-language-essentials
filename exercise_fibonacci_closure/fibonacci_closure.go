package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...
// Arithmetic progression:
/*
  an = a1 + (n-1)d
  a0 = 0 + (0-1)d
  a1 = 0 + (1-1)d = 0
  a2 = 0 + (2-1)d = d
  a3 = 0 + (3-1)d = 2d
  a4 = 0 + (4-1)d = 3d

  an = a1 + (n-1)d
  a0 = 1 + (0-1)d => 0
  a1 = 1 + (1-1)d = 1 + 0d
  a2 = 1 + (2-1)d = 1 + 1d
  a3 = 1 + (3-1)d = 1 + 2d
  a4 = 1 + (4-1)d = 1 + 3d

  a0 = 0 // desconsidering
  a1 = 1 => an = a1 + (n-1)*d => a1 = a1 + (1-1)*d => 1 = 1
  a2 = 1 => an = a1 + (n-1)*d => a2 = a1 + (2-1)*d => 1 = 1 + 1*d
  a3 = 2 => an = a1 + (n-1)*d => a3 = a1 + (3-1)*d => 2 = 1 + 2*d
  a4 = 3 => an = a1 + (n-1)*d => a2 = a1 + (2-1)*d => 1 = 1 + 1*d
  a5 = 8
  a6 = 13
  a7 = 21
  a8 = 34

*/

/**
  Geometric Progression

  an = a1 * r ^(n-1)
  a1 = 1 * r^(1-1) = 1
  a2 = 1 * r^(2-1) = r
  a3 = 1 * r^(3-1) = r²

*/

/**
  a0 = 0
  a1 = 1
  a2 = 1
  a3 = 2
  a4 = 3
  a5 = 5
  a6 = 8
  a7 = 13
  a8 = 21
  a9 = 34

  ==> Fibonacci function:
  an+1 = an + an-1
  a2 = a1 + a0
*/
func fibonacci() func() int {

	// a0 := 0
	// a1 := 1
	previousValue := 0
	currentValue := 0
	nextValue := 0

	iterator := 0

	// an+1 = an + an-1
	return func() int {
		if iterator <= 1 {
			nextValue = iterator
			iterator++
			currentValue = 1
			previousValue = 0
		} else {
			nextValue = currentValue + previousValue
			previousValue = currentValue
			currentValue = nextValue
		}

		return nextValue
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
