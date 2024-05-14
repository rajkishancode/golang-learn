package main

import "fmt"


func main() {
	
	x := []int{1,2,3,4} 
	fmt.Println(x[1:]) // this will give all the element from start index
	fmt.Println(x[1:3]) //  elements from start index but exclude the last index ,just like slice in js

	for i:= 0 ; i<len(x); i++{ // print all numbers in a slice
		fmt.Println(x[i])
	}
}