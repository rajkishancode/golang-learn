package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Bond", "Money penny", "do no":
		fmt.Println("miss money or bond or dr no")

	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
		fallthrough // it takes you to the next case ,DONT use fallthrough(suggestion)
	default:
		fmt.Println("This is default case!")
	}
}
