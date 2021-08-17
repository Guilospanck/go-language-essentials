package main

import "fmt"

func RangeFunction() {
	var slicePow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// range returns an index and the element
	for i, v := range slicePow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// you can omit the second argument if you only want the index
	for i := range slicePow {
		fmt.Println(i)
	}

	fmt.Println()
	// you can skip the first and the second argument, if you're not going to use one of them
	for _, value := range slicePow {
		fmt.Println(value)
	}
}
