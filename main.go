package main

import (
	"fmt"
	"log"
)

func main() {
	// Set propertires of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number.
	log.SetPrefix("Greeting: ")
	log.SetFlags(0)

	// slice of names
	names := []string{
		"Guilospanck",
		"Gui",
		"Test",
	}

	// Request some greetings messages.
	messages, err := Greetings(names)

	// if and error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned messages
	// to the console.
	fmt.Println(messages)
}
