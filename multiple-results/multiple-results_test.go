package multipleresults

import (
	"regexp"
	"testing"
)

// TestSwap will test the function multipleresults.swap.
// It will provide two strings and they must return in swapped orders.
func TestSwap(t *testing.T) {
	x := "Hello"
	y := "World"

	wantX := regexp.MustCompile(x)
	wantY := regexp.MustCompile(y)

	a, b := swap("Hello", "World")

	if !wantX.MatchString(b) || !wantY.MatchString(a) {
		// %q prints double-quotes strings; %s prints the string without the quotes
		// %d for int base-10
		// %v it's the 'joker'. Tells fmt to use the default format.
		t.Fatalf(`swap("Hello", "World") = %q, %q, want match for %q, %q`, a, b, wantY, wantX)
	}

}
