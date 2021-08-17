package main

import "fmt"

// A map maps keys to values
// A zero-value of a map is Nil. A nil map has no keys, nor can keys be added.

// let's say we have this struct which holds latitude and longitude for places
type Position struct {
	Lat, Long float64
}

func MapFunction() {
	// And here I'm mapping string keys to a struct Position
	mapping := make(map[string]Position)
	mapping["Bahamas"] = Position{Lat: 40.68433, Long: -74.39967}

	// you can also instatiate like this:
	mapping2 := map[string]Position{
		"Bahamas": Position{
			Lat: 40.68433, Long: -74.39967,
		},
		"Google": Position{
			Lat: 37.42202, Long: -122.08408,
		},
	}

	// you can omit the type of values
	mapping3 := map[string]Position{
		"Bahamas": {Lat: 40.68433, Long: -74.39967},
	}

	fmt.Printf("The position of Bahamas is %+v\n", mapping["Bahamas"])
	fmt.Printf("The position of Google is %+v\n", mapping2["Google"])
	fmt.Printf("The position of Bahamas is %+v\n", mapping3["Bahamas"])

	/*
	 * Mutating maps
	 */
	mutatingMap := make(map[string]string)

	// Inserting or updating element
	mutatingMap["key1"] = "value1"
	mutatingMap["key2"] = "value2"

	// retrieve an element
	elem := mutatingMap["key1"]
	fmt.Println(elem)

	// delete an element
	delete(mutatingMap, "key2")

	// testing if a key is present in the map
	element, ok := mutatingMap["key1"]
	if ok {
		fmt.Printf("The key exists and the element is %v", element)
	} else {
		fmt.Println("The key doesn't exist.")
	}

}
