package main

import "fmt"

// use a "defer" keyword to show that a defer func runs after the func containing it exits

func foo() {
	defer func(){
		fmt.Println("FOO DEFER RAN")
	}()
	fmt.Println(" foo ran")
}


// defer is the execution of  a function until the function containing it finishes executing
func main() {
	defer foo() // it will execute when the main function exits
	fmt.Println("Hello World.")
}
