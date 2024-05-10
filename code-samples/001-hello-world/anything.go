package main

import "fmt"

func main() {
	fmt.Println("Hello everyone!!")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar(){
	fmt.Println("And here we exit!")
}

//control flow;
// 1 sequence
// 2 loop; iterative
// 3 conditional
