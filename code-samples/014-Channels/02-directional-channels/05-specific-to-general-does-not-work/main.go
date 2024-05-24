package main

import "fmt"

func main() {

	c := make(chan int)  // bidirectional channel
	cs := make(<-chan int) // receive
	cr := make(chan<- int) // send 

	fmt.Println("-------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//1.specific to general doesn't assign
	// c = cr
	// c = cs

	//2.general to specific assign
	// cr = c
	// cs = c
	
	//3.general to specific converts
	fmt.Println("--------")
	fmt.Printf("c\t%T\n",(<-chan int)(c))
	fmt.Printf("c\t%T\n",(<-chan int)(c))
	
}
