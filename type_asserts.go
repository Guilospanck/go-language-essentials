package main

import "fmt"

func TypeAssertion() {
	var i interface{} = "Hello" // this interface is a nil interface that can hold any value

	s := i.(string) // type assertion. In this case, if i holds a value of type "string", it'll return its underlying value (which is "Hello")
	fmt.Println(s)

	// if this type assertion exists, "str" is gonna have the underlying value (otherwise will be the zero value of the type)
	// and "ok" will be true (otherwise will be false)
	str, ok := i.(string)
	if ok {
		fmt.Println(str)
	}

	floatNumber, floatOK := i.(float64)
	if floatOK {
		fmt.Println(floatNumber)
	} else {
		fmt.Println("There is no float64 in this interface")
	}
}
