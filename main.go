package main

import "fmt"

// range for
// Iterates over values received from a channel
// It automatically breaks, when the channel is closed
// range, does not return the second "ok" boolean value
func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
			ch <- i
		}

		// Close channel once goroutine has sent all its values
		close(ch)
	}()

	// in the main goroutin, range over channel to received values
	// once the channel is closed the range will exit
	for v := range ch {
		fmt.Println(v)
	}

}

/*
func main() {
	go func() {
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
		}
	}()
	// TODO: range over channel to recv values
}
*/
