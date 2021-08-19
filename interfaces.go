/**
  An interface type holds any value that implements those methods.
  So, in this case, Abser can hold every value that implements "Abs() float64" but is going to get an error if it's a different signature.

*/
package main

import "fmt"

// interface
type Abser interface {
	Abs() float64
}

type VerType struct {
	X, Y float64
}

// method function of VerType
func (v VerType) Abs() float64 {
	return v.X + v.Y
}

type AnotherStruct struct {
	X, Y int
}

// method function of AnotherStruct
func (a AnotherStruct) Abs() int {
	return a.X + a.Y
}

func InterfaceFunction() {

	var abser Abser // instantiating interface
	vertype := VerType{3, 4}

	abser = vertype // now "abser" has the Abs() method from VerType (only because it implements the same method signature)

	fmt.Println(abser.Abs())

	// but, if we have a method which implements a different signature, like the AnotherStruct.Abs() which returns an int and not a float64
	// you won't have the possibility of using the "Abser"
	//
	// var anotherAbser Abser
	// anotherStruct := AnotherStruct{1, 2}
	//
	// anotherAbser = anotherStruct // error, because anotherStruct.Abs() returns int and not float64

}
