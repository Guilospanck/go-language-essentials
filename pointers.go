package main

import "fmt"

// A pointer holds the memory address of a value

// The type *T is a pointer to a T value. Its zero value is nil.
// => var p *int

// The & operator generates a pointer to its operand
// i := 42
// p = &i

func Pointers() {

	i, j := 40, 200

	p := &i         // p now is a pointer that holds the value of the address alocated for the variable "i"
	fmt.Println(p)  // prints the memory address of the pointer p
	fmt.Println(&i) // prints the memory address of i, which is the same  as the pointer p
	fmt.Println(&p) // prints the memory address of p (as variable), which is gonna be different from the two above

	fmt.Println(*p) // prints the value that the pointer p is appointing
	fmt.Println(i)  // prints the value of "i", which is gonna be the same as above

	p = &j         // pointer p now points to "j" memory address
	*p = *p / 2    // the value to which p points (the memory address of the "j" variable), is going to be changed
	fmt.Println(j) // prints 100 (200/2)

}
