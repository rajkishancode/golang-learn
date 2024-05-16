package main

import "fmt"

//func expression - we will assign a function to a variable that is called func expression 

func main(){
	f := func(){
		fmt.Println("This is func expression")

	}
	f();

	g := func(x int){
		fmt.Println("Func expression with parameter",x)

	}
	g(42);
}


// IN Go functions are first class citizens which means functions are treated as any other variable
// we can 
//- assign function in a variable 
//- return function in other function
//- pass function as a argument