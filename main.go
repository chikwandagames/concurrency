package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}

	}()

	// TODO: if there is no value on channel, do not block.

	// In the first iteration there is no message received from the channel,
	// The main goroutine did not block, but is proceeds with othe computation
	// "processing...",
	// In the second iteration, the message has been received from the channel
	// therefore it is printed to the terminal

	// We are able to perform nonblocking comminication with the aid of default
	// case statement
	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Printf("%v: %v \n", i, m)
		default:
			fmt.Println("no message received")
		}

		// Do some processing..
		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)

	}
}
