package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Hands-on-Exercise #1

/*
   1>> In addition to the main goroutine, launch two additional goroutines
     	- each additional goroutine should print something out
   2>> use waitgroups to make sure each goroutine finishes before your program exists.
*/


func main() {
	
	fmt.Println("Begin goroutines", runtime.NumGoroutine())
	fmt.Println("Begin CPUs", runtime.NumCPU())
	
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()

	}()

	fmt.Println("Mid goroutines", runtime.NumGoroutine())
	fmt.Println("Mid CPUs", runtime.NumCPU())

	wg.Wait()

	fmt.Println("End goroutines", runtime.NumGoroutine())
	fmt.Println("End CPUs", runtime.NumCPU())

}
