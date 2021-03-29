package main

import "fmt"

func main() {

	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		// Send c into the channel
		ch <- c
	}(1, 2)

	// In the main goroutine, we get the value from the channel
	r := <-ch
	fmt.Printf("Computed value: %v \n", r)
}

/*
func main() {
	go func(a, b int) {
		c := a + b
	}(1, 2)
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
*/
