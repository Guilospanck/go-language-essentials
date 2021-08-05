package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Greeting returns a greeting for the named person.
func Greeting(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return a value that embeds the name in
	// a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup,
// after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// a slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
