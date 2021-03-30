package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// In this case we have multiple goroutines accessing a varible counter
// that are running in parallel
// This causes a data race, which will corrucpt the data
// a simple counter++ is not concurrent safe

// Atomic
// To increment the counter in a concurrent safe way, we can use atomic
// Atomic is a lockless operation
// The Atomic Add function can be called by multiple goroutines concurrently
// and access to the memory will be concurrent safe

func main() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	// TODO: implement concurrency safe counter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 10; c++ {
				// counter++ // This would produce random results
				atomic.AddUint64(&counter, 1)
				// Use atomic to imcreament counter by 1
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}
