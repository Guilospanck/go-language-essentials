package main

import "fmt"

func fibo(ch chan int) {
	i, j := 0, 1
	for {
		ch <- j
		i, j = j, i+j
	}
}

func main() {
	ch := make(chan int)
	defer close(ch)
	go fibo(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
