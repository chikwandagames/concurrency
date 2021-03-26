package main

import (
	"fmt"
	"sync"
)

// Goroutines operate on the current values of a variable at the
// time of execution
// If you want a goroutine operate on a specific value, then we need
// to pass that as input to the goroutine

// Basically, if you don't pass i as an argument, the for loop will run and
// whatever the current value of the i is at the time of spinniing up
// a new goroutine, is what the goroutine will increment

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// what is the output
// TODO: fix the issue.

// for i := 1; i <= 3; i++ {
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println(i)
// 	}()
// }
