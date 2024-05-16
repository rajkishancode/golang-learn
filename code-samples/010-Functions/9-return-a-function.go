package main

import "fmt"
// here we are returning a function and a string
func main(){

	s1 := foo()
	fmt.Println(s1);

	x := boo()
	fmt.Println("value of x", x())
}

func foo()string{
	s :="Hello world."
	return s;
}

func boo() func() int{
	return func()int{
		return 451;
	}
}