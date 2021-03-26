package main

func main() {

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
