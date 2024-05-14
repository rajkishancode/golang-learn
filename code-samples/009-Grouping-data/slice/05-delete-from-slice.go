package main

import "fmt"

// deleteing numbers from slice using append
func main() {
	
	x := []int{100,200,300,400,500,600,700} 
	fmt.Println(x)

	//remove 300 and 400 from above slice
	x = append(x[:2],x[4:]...)  // x[:2] -will give number on index 1,2 excluding 3
	fmt.Println(x)              // and x[4:] will give all number start from index 4 and ... will unpack all numbers .
}
