package main

import (
	"fmt"
	"time"
)

// If a goroutine G1 spawns other goroutines G2 and G3 to carry out some tasks
// The question in what order will the results be receiced from G1 and G2

// Select
// Select is like a switch statement
// Each case stetement specifies a send or receive on some channel
// Each case statement is not evaluated sequentially
// All case statements are considered simulteneosly to see if any is ready and
// each case has an equal chance of being selected
// Select waits until some case is ready to proceed
// when one channel is ready, that operation will proceed
// if multiple cases are ready it wil pick one at random

// Empty select will block forever  "select{}"
// Select on a nil channel will block forever "var ch chan string"
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		// time.Sleep(1 * time.Second)
		time.Sleep(3 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// TODO: multiplex recv on channel - ch1, ch2
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
