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

// Greetings returns a map that associates each of the named people
// with a greeting message.
func Greetings(names []string) (map[string]string, error) {
	// a map to associate names with messages.
	//
	// syntax: messages := make(map[key-type]value-type)   which is equal to:  messages = map[key-type]value-type{}
	//
	// or, if you wanna initialize it with some data:
	//
	// messages := map[string]string{
	//		"Guilospanck": "Hello",
	//		"Gui": "Hey",
	// }

	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Greeting function to get a message for each name.
	// syntax: for key, value := range names  OR    for index, value := range names
	for _, name := range names {
		message, err := Greeting(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
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
