package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func StructVertex() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Pointers to Structs
	p := &v
	p.X = 1e9 // if we were to use the correct notation of pointers, it'd be (*p).X = 1e9, but as this would be too cumbersome, the language let us do this instead;
	(*p).Y = 5
	fmt.Println(v)

	// Struct Literals
	v1 := Vertex{X: 1} // Y: 0 is implicit
	v2 := Vertex{}     // X: 0 and Y: 0 are implicit
	fmt.Println(v1, v2)

}
