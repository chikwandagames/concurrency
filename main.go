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
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch

	select {

	case m := <-ch:
		fmt.Println(m)

	case <-time.After(3 * time.Second):
	case <-time.After(1 * time.Second):
		fmt.Println("1 second timeout")
	}

}
