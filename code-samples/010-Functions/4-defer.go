package main

import "fmt"
//A "defer" function invokes a function whose execution is deferred to the moment of the surrounding functions returns,
//either because surrounding function executed a return statement , reached the end of its function body 
// or because the corresponding goroutine is panicking
func main(){
	defer foo()
	bar()
}

func foo(){
	fmt.Println("This is foo.")
}
func bar(){
	fmt.Println("This is bar.")
}