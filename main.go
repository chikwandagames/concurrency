package main

import (
	"fmt"
	"sync"
)

// Broadcast wakes up all goroutines waiting on a condition

// Sync.Cond "conditional variables"
// A conditional variable is a container of goroutines that are
// waiting for a certian condition
// Conditional vars are used to synchronise execution of goroutines

// Wait() suspends the execution of a goroutine
// Signal() wakes up one goroutine waiting on a condition
// Broadcast() wakes up all goroutines waition on a condition

// Here we have 2 goroutines waiting on different conditions

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
		// Before accessing the shared resource, lock
		cnd.L.Lock()
		for len(sharedRsc) < 1 {
			// If condition not met
			cnd.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		// After processing release the lock
		cnd.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cnd.L.Lock()
		for len(sharedRsc) < 2 {
			cnd.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		cnd.L.Unlock()
	}()

	// In the main goroutine Lock, (modify) populate the shared resource,
	// wake the other 2 goroutines using Broadcast(), then release the lock
	cnd.L.Lock()
	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	// Send a broadcast signal to all goroutines the condition they are waiting on
	// has been met
	cnd.Broadcast()
	cnd.L.Unlock()

	wg.Wait()
}
