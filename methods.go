/**
  GO doesn't have classes. The thing next to something as a class is a method.

  - A method is just a function with a receiver argument.
  - The receiver argument is a type, which can be a struct or not.
  - You can only declare a method with a receive whose type is defined in the same package.
  - The receiver can be a pointer (and usually is).
  - You should have either all receivers as values OR all receivers as pointers (not a mixture of both), because
*/

package main

import "fmt"

type SomeStruct struct {
	X, Y float64
}

// Example of a method function:
func (v SomeStruct) abs() float64 {
	var x float64 = v.X
	var y float64 = v.Y
	if v.X < 0 {
		x = -(v.X)
	}

	if v.Y < 0 {
		y = -(v.Y)
	}

	return x + y
}

// example of function that is not a method
func abs(v SomeStruct) float64 {
	var x float64 = v.X
	var y float64 = v.Y
	if v.X < 0 {
		x = -(v.X)
	}

	if v.Y < 0 {
		y = -(v.Y)
	}

	return x + y
}

// example of a method function that uses pointer as a receiver
func (v *SomeStruct) scale() {
	v.X *= 10
	v.Y *= 10
}

func MethodsFunction() {
	someStruct := SomeStruct{-3, 4}
	fmt.Println(someStruct.abs()) // see that we can only call abs method function from "SomeStruct" instance.
	fmt.Println(abs(someStruct))  // see that we can call the abs function because it is not a method.

	someStruct.scale()
	// see that we've changed the X and Y variables of "SomeStruct" inside the scale method function, but
	// as we were using pointers in scale, it changed the type SomeStruct struc itself, refleting on the abs() method function.
	fmt.Println(someStruct.abs())
	fmt.Println(abs(someStruct))

}
