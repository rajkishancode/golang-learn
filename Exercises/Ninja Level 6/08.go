package main

import "fmt"

//create a func which returns a func
// assign the returned func to a variable
//call the returned func

func greeting()func() string{
	return func() string {
		return "Hello world"
	}
}

func main(){

	f:= greeting()
	fmt.Println(f())
	
}