package main

import (
	"fmt"
	"math"
)
// IMPORTANT - "The method set of a type determines the INTERFACES that the type implements"

// methods sets determine what methods attach to a TYPE
// It is exactly what the name says what is the set of methods for a given type?
// That is its methods set.

// > a NON POINTER RECEIVER - works with values that are POINTERS or NON-POINTERS

// > a POINTER RECEIVER only works with values that are POINTERS

// RECEIVERS     VALUES
// ---------------------------------------
//   (t T)          T and *T   -> receiver which is not a pointer works with both pointers and non-pointers value
//   (t *T)         *T         -> here you can see pointer receiver only works with value that are pointer

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c *circle) area() float64 { // POINTER RECEIVER
	return math.Pi * c.radius * c.radius
}
func info(s shape) {
	fmt.Println("area", s.area())
}

// POINTER RECEIVER AND POINTER VALUE
func main() {
	c := circle{radius: 5,}
	fmt.Printf("%T\n", &c)
	info(&c)// POINTER VALUE
}

//POINTER RECEIVER AND NON-POINTER VALUE - DOES NOT WORk- you can change and see

/*
in line 40 if you just pass value ex - info(c) without '&' that makes it pointer value
	you will get error := cannot use c (variable of type circle) as shape value
 	in argument to info: circle does not implement shape (method area has pointer receiver)

	BUT you can directly run also like fmt.Println(c.area());
*/
