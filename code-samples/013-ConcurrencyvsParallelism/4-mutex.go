package main

import (
	"runtime"
	"fmt"
	"sync"
	// "time"
)

//MUTEX - something which is mutually exclusive




func main(){
	fmt.Println("Goroutines start",runtime.NumGoroutine())
	fmt.Println("NumCPU",runtime.NumCPU())
	
	counter := 0
	
	const gs = 100 //go routines
	var wg sync.WaitGroup
	wg.Add(gs)

    var mutex sync.Mutex  // create a variable mutex and called mutex from sync package type Mutex

	for i := 0;  i < gs ; i++{

		go func(){			  // start the go routine	
			mutex.Lock()
			v := counter      // reading the value
		//  time.Sleep(time.Second) // it will go to sleep for 1 second
			runtime.Gosched() // it says hey, go ahead and run something else if you have something else to run.
			v++			      // incrementing it
			counter = v       // writing value back to the shared variable 
			mutex.Unlock()
			wg.Done() 		  // will done the go routine
			}()
			fmt.Println("Goroutines mid",runtime.NumGoroutine())
	}
	
	wg.Wait() // it waits until the counter in wait group comes down to ZERO.
	fmt.Println("Goroutines end",runtime.NumGoroutine())
	fmt.Println("count",counter)
}

/*
mutex.Lock()
- no matter who accessing the code like a thread goes on to grab the variable
- everything goes to halt and wait
*/

/*
go run  main.go
go run -race main.go - to check race condition
*/