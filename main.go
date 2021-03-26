package main

import (
	"fmt"
	"sync"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

// The reason why the reference to the i variable is kept by the go routine
// is because the runtime is clever enough to see that the reference to
// the variable i is still being held by the goroutine, so it moves it
// from the stack to the heap so that the goroutine still has access
// to the variable even after the enclosing outer function returns

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Printf("return from function %v \n", i)
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}
