package main

import "fmt"

var y string
var z int

func main() {
	//DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	fmt.Println("printing the value of y", y, `there is nothing in place of y because it empty string like this ""`)
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
