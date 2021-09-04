/**
* Concurrency X Parallelism
*
* - Concurrency: (Interruptability) => two lines of customers on the same cashier (lines taking turns ordering)
* - Parallelism: (Independentability) => two lines of customers ordering from two cashiers (each line has its own cashier)
*
*
*   ====> See exercise_equivalent_binary_trees for an implementation of goroutines with channels <====
*
*  To use goroutine, just add the "go" syntax before a method or function
*  Example:    go Fibonacci()
*
* => Channels are ways through which you can send and receive values from a routine.
*       They can be created as:
*       ch := make(chan int)
*
*       ch <- v    // Send v to channel ch.
*       v := <-ch  // Receive from ch, and assign value to v.
*
*       (The data flows in the direction of the arrow.)
*
*       v, ok := <-ch
*       // ok is false if there are no more values to receive and the channel is closed.
*
*       The loop
*                 for i := range ch
*       receives values from the channel repeatedly until it is closed.
*
*
**/
