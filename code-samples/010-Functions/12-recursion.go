package main

import "fmt"

//RECURSION - when a function call itself

func main() {
	f := factorial(4)
	fmt.Println("factorial using recursion", f)
	f2 := factorial2(4)
	fmt.Println("factorial using loop", f2)
	f3 := loopfact(4)
	fmt.Println("factorial using loop", f3)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// using loop

func factorial2(n int) int {

	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// using decrement loop

func loopfact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
