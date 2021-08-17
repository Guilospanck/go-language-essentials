package main

import (
	"fmt"
	"strings"
)

func Arrays() {
	// Arrays have fixed size.
	var a [3]string // declares an array of strings of 3 positions
	a[0] = "Hello"
	a[1] = "String"
	a[2] = " .How are you?"

	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	// Slices have dynamic size
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays. They dont store any data, just describes a section of an underlying array.
	// All data that is changed, even in the slices, will change everything that uses the array as reference
	array := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	slice1 := array[0:2]
	slice2 := array[1:3]

	fmt.Println(slice1) // [John Paul]
	fmt.Println(slice2) // [Paul George]
	fmt.Println(array)  // [John Paul George Ringo]

	slice2[0] = "Potato" // changing "Paul" to "Potato"
	fmt.Println(slice1)  // [John Potato]
	fmt.Println(slice2)  // [Potato George]
	fmt.Println(array)   // [John Potato George Ringo]

	// Slice defaults.
	// When slicing, you may omit the high and low bounds.
	arrayTest := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// those slices below are all equivalent.
	s1 := arrayTest[0:10]
	s2 := arrayTest[:10]
	s3 := arrayTest[0:]
	s4 := arrayTest[:]
	fmt.Println(s1, s2, s3, s4)

	// slice has both length and capacity
	// length derives from the number of elements it contains.
	// capacity derives from the number of elements its underlying array contains.
	underlyingArray := [5]int{1, 2, 3, 4, 5}

	sliceOfThatArray := underlyingArray[:2]
	fmt.Println(len(sliceOfThatArray)) // must print 2
	fmt.Println(cap(sliceOfThatArray)) // must print 5

	// reslicing an array
	sliceOfThatArray = sliceOfThatArray[:4]
	fmt.Println(len(sliceOfThatArray)) // must print 4
	fmt.Println(cap(sliceOfThatArray)) // must print 5

	// Nil slices (the zero value of a slice if a nil)
	var anotherSlice []int
	fmt.Println(anotherSlice, len(anotherSlice), cap(anotherSlice))
	if anotherSlice == nil {
		fmt.Println("It's a nil!")
	}

	// instantiating another slice with "make"
	b := make([]int, 0, 5) // len = 0, cap = 5
	b1 := b[:3]            // len = 3 cap = 5
	fmt.Println(len(b1))
	fmt.Println(cap(b1))
	fmt.Println(b1)
	b2 := b1[1:] // len 2 cap 4
	fmt.Println(b2)
	fmt.Println(len(b2))
	fmt.Println(cap(b2))

	// slices of slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println()
	// append to slices
	var appendingSlice []int
	appendingSlice = append(appendingSlice, 0)
	fmt.Println(appendingSlice)
	appendingSlice = append(appendingSlice, 1, 2, 3, 4, 5)
	fmt.Println(appendingSlice)
	fmt.Println(len(appendingSlice))
	fmt.Println(cap(appendingSlice))

}
