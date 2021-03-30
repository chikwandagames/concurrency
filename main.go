package main

import (
	"fmt"
	"runtime"
	"sync"
)

// The goroutines in the for loops, 100 for withdrawals and 100 for deposites
// are running concurrently,
// the balance variable is shared between withdrals and deposits
// this can cause undetermistic results due to data race condition

func main() {

	// Tell the runtime to use 4 cores to run our goroutines
	runtime.GOMAXPROCS(4)

	// Shared variable
	var balance int
	var wg sync.WaitGroup

	// Here we use Mutex to gaurd access to the shared variable
	// sync.Mutex, provides exclusive access to a shared resource using
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	withdrawal := func(amount int) {
		mu.Lock()
		balance -= amount
		defer mu.Unlock()
	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.
	// run the program and check result.

	// TODO: fix the issue for consistent output.

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
			// fmt.Println(balance)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
