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

	fmt.Println("===================")
	// Prints for Loop result
	fmt.Println(Loop())
	fmt.Println(LoopContinued())
	fmt.Println(WhileFunction())
	fmt.Println(LoopForever())

	fmt.Println("===================")
	// Prints for If result
	fmt.Println(NormalIf(2, 5, 40))
	fmt.Println(IfShorthanded(2, 5, 40))
	fmt.Println(IfShorthandedWithElse(2, 5, 30))

	fmt.Println(" Square root: ===================")
	// Square root.go
	// fmt.Println(Sqrt(25))
	fmt.Println(SqrtAutomatic(25))

	fmt.Println(" Switch: ===================")
	// switch.go
	Switch(1)
	SwitchShorthanded()
	SwitchWithNoCondition()

	fmt.Println(" Defer: ===================")
	// defer.go
	EvaluatedNormalRuntimeExecutedAfterReturn()
	LastInFirstOut()
	fmt.Println(NamedReturnValues())
	fmt.Println(NoNamedReturnValues())
	fmt.Println(NoNamedReturnValues2())

	fmt.Println(" Pointers: ===================")
	// pointers.go
	Pointers()

	fmt.Println(" Structs: ===================")
	// structs.go
	StructVertex()

}
