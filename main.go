package main

import (
	"fmt"
	"sync"
)

// Sync.Cond "conditional variables"
// A conditional variable is a container of goroutines that are
// waiting for a certian condition
// Conditional vars are used to synchronise execution of goroutines

// Wait() suspends the execution of a goroutine
// Signal() wakes up one goroutine waiting on a condition
// Broadcast() wakes up all goroutines waition on a condition

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	// Conditional variable
	cnd := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.

		// Before accessing the shared resource, we lock
		cnd.L.Lock()
		// While the sharedRsc is empty, suspend the goroutine
		for len(sharedRsc) == 0 {
			// Wait(), releases the lock and suspend the goroutine
			cnd.Wait()
			// fmt.Println("nothing here")
		}

		fmt.Println(sharedRsc["rsc1"])
		// Release the lock, after processing
		cnd.L.Unlock()
	}()

	cnd.L.Lock()
	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"
	// Send a signal to the goroutine, that the codition has been met
	cnd.Signal()
	// Then release the lock
	cnd.L.Unlock()

	wg.Wait()
}
