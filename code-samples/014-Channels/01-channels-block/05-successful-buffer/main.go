package main

import "fmt"

func main() {
	c := make(chan int , 2) 
	
	c <- 42
	c <- 43
	
	// we pull all two values out of the channel
	fmt.Println(<-c)
	fmt.Println(<-c)
}