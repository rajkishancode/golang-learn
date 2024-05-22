package main

import (
	"fmt"
	"runtime"
	"sync"
	
)
// WE ARE CREATING A RACE CONDITION

func main(){
	fmt.Println("Goroutines",runtime.NumGoroutine())
	fmt.Println("NumCPU",runtime.NumCPU())

	counter := 0

	const gs = 100 //go routines
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0;  i < gs ; i++{

		go func(){			  // start the go routine	
			v := counter      // reading the value
		//  time.Sleep(time.Second) // it will go to sleep for 1 second
			runtime.Gosched() // it says hey, go ahead and run something else if you have something else to run.
			v++			      // incrementing it
			counter = v       // writing value back to the shared variable 
			wg.Done() 		  // will done the go routine
			}()
			fmt.Println("Goroutines",runtime.NumGoroutine())
	}
	
	wg.Wait() // it waits until the counter in wait group comes down to ZERO.
	fmt.Println("Goroutines",runtime.NumGoroutine())
	fmt.Println("count",counter)
}


/*
func Gosched()
 - Gosched yields the processor, allowing other goroutines to run.
 - It does not suspend the current goroutine, so execution resumes automatically.
*/

/*
time.sleep() - sleep will sleep for  a duration
			 - one of constant it takes Second /micro-second/mini-second etc 
*/