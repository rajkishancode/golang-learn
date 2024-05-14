package main

import "fmt"

// append is a special built in function
func main() {
	
	x := []int{13,22,33,24}
	fmt.Println(x)

	x = append(x, 100,200,300) // append numbers in a slice
	fmt.Println(x)

	y := []int{1000,2000,3000}
	x = append(x,y...) // append a whole new slice-y to other slice-x and
	fmt.Println(x)					// after y... three dots are called variadic parameter , it will just unpack all elements just like spread in js

}