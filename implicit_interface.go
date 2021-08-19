/**
  A type implements an interface by implementing its methods (there is no "implements" keyword)

*/

package main

import "fmt"

type SomeInterface interface {
	M()
}

type SomeType struct {
	S string
}

// this methods means: the type "SomeType" is implementing the interface "SomeInterface"
func (t SomeType) M() {
	fmt.Println(t.S)
}

type AnotherSomeType struct {
	S string
}

func (tt AnotherSomeType) SomeFunction() {
	fmt.Println(tt.S)
}

func ImplicitInterfaceFunction() {
	var someInterface SomeInterface = SomeType{"Hello"}
	someInterface.M()

	// This is not going to work because type "AnotherSomeType" doesn't have M() implementation
	// var anotherSomeInterface SomeInterface = AnotherSomeType{"Hello"}
}
