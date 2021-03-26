package main

import (
	"fmt"
	"sync"
)

// To create a join point to allow a forked goroutine to join the
// main goroutine after the forked goroutine has completed execution.
// We use sync WaitGroup to block the main goroutine

// WaitGroup is like a concurrent counter
// 1. Calling Add() increment the count by the int passed to it
// 2. Calling Done() decrement the count by one
// 3. Calling Wait() waits until the count is zero

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int

	var wg sync.WaitGroup

	// 1. Add the number of goroutines to be created
	wg.Add(1)

	go func() {
		// 2. call Done() inside the goroutine closure to indicate that teh
		//    goroutine is exitting, use defer to make sure Done() is called
		//    on all exit points of the function.
		//    Done() will be executed at the end of the goroutine
		defer wg.Done()
		data++
	}()

	// Wait will block the main goroutine until all the forked goroutines
	// have exited
	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)
	fmt.Println("Done..")
}
