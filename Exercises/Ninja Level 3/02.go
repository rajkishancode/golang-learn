package main

import "fmt"

//print every rune code point of the UPPERCASE alphabet  3 times

func main() {
	for i := 65; i <= 90; i++ {
			fmt.Println(i)
		for j := 1; j <= 3; j++ {

			fmt.Printf("\t%#U\n", i)
		}
	}
}
