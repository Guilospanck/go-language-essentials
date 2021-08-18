package main

import "fmt"

func compute(potato func(x, y float64) float64) float64 {
	return potato(4, 5)
}

func FunctionValues() {

	someFunction := func(x, y float64) float64 {
		return x + y
	}

	fmt.Println(compute(someFunction))

}
