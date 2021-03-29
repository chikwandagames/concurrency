package main

import "fmt"

func main() {
	// All we need to do is add a capacity to make this a buffered channel
	// 6 is the capacity, which is the number of elements we want
	// to send without blocking, because we have 6 iterations in the for loop
	ch := make(chan int, 6)

	go func() {
		defer close(ch)

		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
