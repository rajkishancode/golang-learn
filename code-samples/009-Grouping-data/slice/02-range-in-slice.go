package main

import "fmt"


func main() {
	
	x := []int{42,43,55,56,67}

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range  x{ // print index - value using range on slice
		fmt.Println(i,v)
	}
   
}
