package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutex is used when we want "mutual exclusion" (only one goroutine can access a variable at a time to avoid conflicts)
*/

type SafeCounter struct {
	mu    sync.Mutex
	value map[string]int
}

func (sc *SafeCounter) Increment(key string) {
	defer sc.mu.Unlock()
	sc.mu.Lock()
	sc.value[key]++
}

func (sc *SafeCounter) RetrieveValue(key string) int {
	defer sc.mu.Unlock()
	sc.mu.Lock()
	return sc.value[key]
}

func main() {
	sc := SafeCounter{
		value: make(map[string]int),
	}

	for i := 0; i < 1000; i++ {
		go sc.Increment("someKey")
	}

	time.Sleep(time.Second)

	fmt.Println(sc.RetrieveValue("someKey"))

}
