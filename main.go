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

	// Request a greeting message.
	message, err := Greeting("Guilospanck")

	// if and error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
