package main

import "fmt"

//create a function with the identifier foo that returns an int
func foo()int{
	return 42
}

//create a function with the identifier bar that returns an int and a string
func bar()(int,string){
	i := 45
	s := "Hello World."
	return i,s
}
func main(){
	//call both funcs 
	f := foo()
	b, c := bar()
	
	//print out their result
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(c)

}

