package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Hands-on-Exercise #3

/*
	- Using goroutine create an incrementer program
		- have a variable to hold the incrementer value
		- launch a bunch of goroutines
			- each goroutine should
			  - read the incrementer value
			  		- store it in a new variable
			  - yield the processor with runtime.GOsched()
			  - increment the new variable
			  - write the value in the new variable back to the incrementer variable
	- Use waitgroups to wait for all your goroutines to finish
	- the above will create a race condition
	- Prove that it is a race condition by using the -race flag

*/

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	const gs = 100 // goroutines
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done() // it pops one off , each one we launch
		}()

	}

	wg.Wait() // wait till counter comes down to zero
	fmt.Println("incrementer end value", incrementer)

}
