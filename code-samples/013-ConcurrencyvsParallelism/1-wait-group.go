// func init(){} JUST FYI it runs once at the beginning of the program , we use it in webdev to set the template.

package main

import (
	"fmt"
	"runtime"
	"sync"
)


var wg sync.WaitGroup // wg is of type waitGroup from sync package 
					  // and wg variable is of package scope means anywhere inside this file we can access it.
func main(){

	fmt.Println("---------------------start here---------------------------------")
	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t",runtime.GOARCH)
	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("Goroutines",runtime.NumGoroutine())
	
	wg.Add(1) // 1. here we say that - for One thing we are waiting for and add that to the variable wg 
	go foo() // go keyword launch a new go-routine 
	bar()    
	
	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("Goroutines",runtime.NumGoroutine())
	wg.Wait() // 3.and when its done finally our program stops waiting and executes it and end.
			  
}


func foo(){
	for i := 0 ; i < 10 ; i++{
		fmt.Println("foo:",i)
	}
	wg.Done() // 2.this is basically remove that 1 thing added  in wg.Add(1)
}

func bar(){
	for i := 0 ; i < 10 ; i++{
		fmt.Println("bar:",i)
	}
	
}


/*
type WaitGroup
func (wg *WaitGroup) Add(delta int) // delta is symbol for change in mathematics and it takes a int , we are giving here 1
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()
*/

/*
steps what happened
1. foo get launched
2. bar runs
3. CPUs             4 - this run
4. Goroutines 2       - this run
5. normally our program exit but now program wait due to wg.Wait() on line 28
6. foo does its work and say okay i am done
7. then line 28 checks and says now nothing to wait for
8. And the program exits!
*/


/*
To Fix the race condition we will use mutex
*/