package main

import (
	"fmt"
	"time"
)

func foo(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	foo("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go foo("function call")

	// goroutine with anonymous function
	go func() {
		foo("anonimous function")
	}()

	// goroutine with function value call
	fval := foo
	go fval("function value")

	// because main goroutine completes execution before other goroutines
	// we need wait for other goroutines to complete, so we can see their execution
	// results, so we sleep the main go routine
	fmt.Println("waiting for goroutines...")
	time.Sleep(10000 * time.Microsecond)

	fmt.Println("done..")
}

// GOUROUTINES
// We can think of goroutines as user space threads managed by
// the go runtime

// Goroutines are extremely lightweigth. they start with 2KB of stack,
// Which grows and shrinks as required

// Low CPU overhead - three instructions per function call
// This allows for creation of hundreds of thousands of goroutines in
// the same address space

// CHANNELS are used to for communication of data between goroutines
// so that sharing of memory can be avoided

// Context switching between goroutines is much cheaper than thread switching
// as goroutines have less state to store

// Go runtime
