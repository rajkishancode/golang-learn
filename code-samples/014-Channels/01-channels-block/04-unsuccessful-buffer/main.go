package main

import "fmt"

//Unsuccessful Buffer is one of the important concept

func main() {
	c := make(chan int , 2) 
	
	c <- 42 // till here fine to put on 42
	c <- 43 // but here oops we have only room for 1
			// we are blocked until we pull the value off
	
	fmt.Println(<-c) // we never get down here as code is blocked in line 11
}