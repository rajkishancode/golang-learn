package main

import "fmt"


//we have some channels going on here

func doSomething(x int) int {
	//Does something
	return x * 5
}

func main(){

	ch := make(chan int)

	go func(){
		// it retuns some value and we put that value on channel
		ch <- doSomething(5) // To launch it in go routine ,we wrap this in another anonymous self executing fn

	}()

	fmt.Println(<-ch) // print that value on channel
}