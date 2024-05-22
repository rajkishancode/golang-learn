package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// to avoid a race condition



func main(){
	fmt.Println("Goroutines",runtime.NumGoroutine())
	fmt.Println("NumCPU",runtime.NumCPU())
	
	var counter int64 // as Atomic package use int64
	
	const gs = 100 //go routines
	var wg sync.WaitGroup
	wg.Add(gs)

    

	for i := 0;  i < gs ; i++{

		go func(){			  		
			atomic.AddInt64(&counter,1)  // pass address of counter and value = 1 by which it change
			fmt.Println("Counter\t",atomic.LoadInt64(&counter)) // read from my counter
			runtime.Gosched() 			
			wg.Done() 		  
			}()
			fmt.Println("Goroutines",runtime.NumGoroutine())
	}
	
	wg.Wait()  // this point all go routines are finished
	fmt.Println("Goroutines",runtime.NumGoroutine())
	fmt.Println("count",counter)
}

/*
 atomic.AddInt64(&counter,1) -> add to the counter
 atomic.LoadInt64(&counter)  -> read from the counter
*/