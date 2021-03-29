package main

import "fmt"

// Channel direction
// When using channels as function params, you can specify
// whether a channel is meant to only send or receive values
// this specification increases the type-safety of a program

func foo(in <-chan string, out chan<- string) {}

// Above
// "in"is a receive only channel
// "out"is a send only channel
// foo can only use "in" only to receive values
// foo can only use "out" only to send values

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and rMsg
	go addMsg(ch1)

	// Alternatively
	// go func(ch chan<- string) {
	// 	ch <- "another message"
	// }(ch1)

	go passOnMsg(ch1, ch2)

	// recv message on ch2
	v := <-ch2
	fmt.Println(v)

	// The goroutine that creates a channel, can write to that channel
	// and is also responsible for closing the channel and is the owner
	// of the channel
	// A goroutine that utiliese a channel has read only priviledges
	// to a channel
	defer close(ch1)
	defer close(ch2)
}

// Takes a send only channel
func addMsg(ch chan<- string) {
	// send message on ch1
	ch <- "message from sendMsg()"
}

// Takes a receive only and send only channel
func passOnMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	m := <-ch1
	// send it on ch2
	ch2 <- m

}
