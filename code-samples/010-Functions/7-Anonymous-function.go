package main

import "fmt"

func main(){

	foo()
	
	func(){
		fmt.Println("This is Anonymous function without any parameter")
	}()

	func(x int){
		fmt.Println("The Anonymous function with parameter",x)
	}(42)
}

func foo(){
	fmt.Println("This is foo")
}