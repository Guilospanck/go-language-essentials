/*
* Defers are used to wait for something to occur.
* - The deferred variable/func will be evaluated in the normal runtime, but will only execute after its surroundings functions return.
* - Deferred functions are executed in LAST IN, FIRST OUT order after the surrounding function returns
* - When there are named values, the deferred function will read/assign to the variable and return it
 */
package main

import "fmt"

// Here, we are going to print '0' because at the time the *defer fmt.Println(i)* was evaluated, the value of i was 0
func EvaluatedNormalRuntimeExecutedAfterReturn() {
	i := 0
	defer fmt.Println(i)
	i++
	// return
}

// In this, will print 3210
func LastInFirstOut() {
	defer fmt.Printf("\n")

	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// In this, as we have a named return value "i" and in our defer function we have this same "i" incrementing, it'll
// return 2 + 1 = 3 (we are returning 2, and then we know that our return is named "i" and in our deferred function
// we have i++, so the value of "i" before the deferred function is 2, and after the deferred function, which was called after
// the return statement, the value of "i" will be i = i+1 = 2+1 = 3)
func NamedReturnValues() (i int) {
	defer func() { i++ }()
	return 2
}

func NoNamedReturnValues() (j int) {
	i := 0
	defer func() { i++ }()
	return 1
}

func NoNamedReturnValues2() (i int) {
	j := 3
	defer func() { i++ }()
	return j
}
