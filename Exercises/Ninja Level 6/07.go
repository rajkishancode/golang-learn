package main

import "fmt"
var x int = 7
var g func() = func(){fmt.Println("g from outside")}

//assign a func to a variable and call that func

func main(){
	f := func(){
		for i := 0; i<3; i++{
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n",f) // func() type
	fmt.Println("Done")

	fmt.Println(x)
	g()
}