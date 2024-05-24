package main

import "fmt"

//person{} - this is composite literal ,sometimes people use struct-literal,map-literal,slice-literal
//just use the phrase composite literal - ex- person - type and curly braces - {}

//Anonymous Struct

type person struct { // person representing struct with some fields
	first string
	last  string
	age   int
}

// now we take above struct without person identifier and directly put below in place of person
func main() {
	//BEFORE CHANGE CODE
	// p1 := person{
	// 	first: "James",
	// 	last:  "Bond",
	// 	age:   33,
	// }
	//AFTER CODE CHANGE
	//NOW IT BECOMES ANONYMOUS STRUCT because it does't have a name
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   33,
	}

	fmt.Println(p1)
}

//why use anonmous struct
 // we use it to avoid code pollution
 // - to keep code lean and clean
 // - not having extraneous types and variables when you dont wanna need them
 // It you only want to use a struct in one little area YOU CAN MAKE A ANONYMOUS STRUCT