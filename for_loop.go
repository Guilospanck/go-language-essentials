package main

// Example of a For Loop
func Loop() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	return sum
}

// Example using a for where the init and post statementes are optional
func LoopContinued() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	return sum
}

// Another way of writing "while" function in Go
func WhileFunction() int {
	sum := 1
	for sum < 1000 { // read: while sum < 1000 {}
		sum += sum
	}

	return sum
}

// For loop forever
func LoopForever() int {
	sum := 1
	for {
		if sum > 1000 {
			break
		}
		sum += sum
	}

	return sum
}
