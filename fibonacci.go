// Fibonnaci function
// 1 => currentValue = 1, previousValue = 0 => 1+0 = 1
// 1 => currentValue = 1, previousValue = 1 => 1+1 = 2
// 2 => currentValue = 2, previousValue = 1 => 2+1 = 3
// 3 => currentValue = 3, previousValue = 2 => 3+2 = 5
//
// an+1 = an + an-1

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34 ...

package main

import "fmt"

func fibonacci() func() int {

	nextValue := 0
	currentValue := 0
	previousValue := 0

	iterator := 0

	return func() int {
		if iterator <= 1 {
			nextValue = iterator
			currentValue = 1
			previousValue = 0
		} else {
			nextValue = currentValue + previousValue
			previousValue = currentValue
			currentValue = nextValue
		}

		iterator++

		return nextValue

	}

}

func ReturnFibo() {
	fibo := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fibo())
	}

}
