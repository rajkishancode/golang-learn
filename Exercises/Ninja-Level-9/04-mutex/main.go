package main

import (
	"fmt"

	"sync"
)

//Hands-on-Exercise #4

// Fix the race condition you created in the previous exercise by using a mutex

// - we will use mutex - a mutual exclusion lock to fix the race condition
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	incrementer := 0
	const gs = 100 // goroutines
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock()
			v := incrementer
			// runtime.Gosched() // it makes sense to remove runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait() // wait till counter comes down to zero
	fmt.Println("incrementer end value", incrementer)

}
