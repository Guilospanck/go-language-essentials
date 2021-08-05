package multipleresults

import (
	"fmt"

	"rsc.io/quote"
)

// In Go the type is put after the variable and if two or more variables share the same type, you can omit from the first ones.
func swap(x, y string) (string, string) {
	return y, x
}

// the := (short assignment statement) can be used in place of a 'var' declaration with implicit type. Outside a function,
// every statement begins with a keyword (var, func, and so on), and so the := construct is not available.

// function named with Uppercase will be exported by default to those who import our package.
func MultipleResults() string {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	return quote.Glass()
}
