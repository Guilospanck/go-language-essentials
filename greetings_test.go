package main

import (
	"regexp"
	"testing"
)

// TestGreetingName calls Greeting with a name,
// checking for a valid return value.
func TestGreetingName(t *testing.T) {
	name := "Guilospanck"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting("Guilospanck")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greeting("Guilospanck") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestGreetingEmpty calls Greeting with an empty string,
// checking for an error.
func TestGreetingEmpty(t *testing.T) {
	msg, err := Greeting("")
	if msg != "" || err == nil {
		t.Fatalf(`Greeting("") = %q, %v, want "", error`, msg, err)
	}
}
