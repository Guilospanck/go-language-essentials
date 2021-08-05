package main

import (
	"errors"
	"fmt"
)

func Greeting(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return a value that embeds the name in
	// a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
