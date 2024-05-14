package main

import "fmt"

// create a program using switch statement with switch statement specified as a variable of TYPE string with the identifier "favsport"
func main() {
	favsport := "football"
	switch favsport{
	case "football":
		fmt.Println("play football")
	case "cricket":
		fmt.Println("play cricket")
	default:
		fmt.Println("nothing to play")
	}

}
