package main

import "fmt"

// The channel owner is responsable for
// 1. creating a channel
// 2. writing to the channel
// 3 closing the channel

func main() {

	// Function that returns a receive only channel of type int
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			//Close channel when done
			defer close(ch)
			for i := 0; i < 5; i++ {
				// Send values over the channel
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	fmt.Println(ch)
	// Pass owner function to consumer
	consumer(ch)
}
