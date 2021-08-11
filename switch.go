package main

import (
	"fmt"
	"runtime"
	"time"
)

func Switch(num int) {
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Any")
	}
}

func SwitchShorthanded() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other")
	}
}

// Switch without a condition works like switch true
func SwitchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
