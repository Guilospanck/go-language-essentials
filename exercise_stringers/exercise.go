package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// fmt package automatically searches for a method with this face: func (t type) String() string {} to
// print out anything related to that type

func (ip IPAddr) String() string {
	myString := make([]string, 0)
	for _, value := range ip {
		myString = append(myString, strconv.Itoa(int(value)))
	}

	return strings.Join(myString, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"Loopback": {127, 0, 0, 1},
		"Google":   {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip) // Here fmt will search for methods of String() string in order to plot
	}
}
