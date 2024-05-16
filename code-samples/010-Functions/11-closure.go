package main

import "fmt"

//CLOSURE
//variables declare in the outer scope are accesible in the inner scopes
//closure helps in limiting the scope of the variable

func main() {
	//scope of x narrowed to func main
	var x int
	fmt.Println(x) //0
	x++

	{ //code block in code block
		y := 42
		fmt.Println(y) //42
	}
	//  fmt.Println(y)   - this gives error undefined: y as variable y is in another scope

	fmt.Println(x) //1

	a := incrementor() //incrementor runs it returns a function that is assigned to a variable
	b := incrementor()
	fmt.Println(a()) // here a is invoked and value of x variable is updated to 1
	fmt.Println(a()) // x = 2
	fmt.Println(a()) // x = 3
	fmt.Println(b()) //now for b scope is different so it start with x and update  x = 1
	fmt.Println(b()) // x  = 2

}

func incrementor() func() int { // incrementor returns a func that returns a int

	var x int
	return func() int {
		x++
		return x
	}
}
