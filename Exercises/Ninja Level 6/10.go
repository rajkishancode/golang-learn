package main

import "fmt"

//Closure is when we enclosed the scope of a variable in some code block. For this handon exercise
//create a func which enclosed the "scope" of a variable

func main() {

	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g := incrementor()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}

//incrementor is a great example of closure
func incrementor() func() int {
	num := 0
	return func() int {
		num++

		return num
	}
}
