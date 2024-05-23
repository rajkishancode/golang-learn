package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Hands-on-Exercise #5
// Fix the race condition you created in ex4 by using package atomic

func main() {
	var wg sync.WaitGroup

	var incrementer int64
	const gs = 100 // goroutines
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {

			atomic.AddInt64(&incrementer, 1) // just one line code to increment
			fmt.Println("incrementer-----", atomic.LoadInt64(&incrementer))

			wg.Done()
		}()

	}

	wg.Wait() // wait till counter comes down to zero
	fmt.Println("incrementer end value", incrementer)

}


// Thats Atomic to get rid of the race condition