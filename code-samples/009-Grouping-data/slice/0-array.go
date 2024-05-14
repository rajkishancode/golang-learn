package main

import "fmt"


func main() {
	
	var x [5]int
	fmt.Println(x)
	x[3] = 42 // adding 42 at index 3
	fmt.Println(x)
	fmt.Println(len(x)) // len to get size of array

   
}
