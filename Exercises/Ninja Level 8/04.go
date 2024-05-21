package main

import (
	"fmt"
	"sort"
	
)

// Hands-on exercise 4 code-https://go.dev/play/p/H_q75mpmHW
// code-https://go.dev/play/p/H_q75mpmHW

// sort []int and []string for each person

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi

	
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}
