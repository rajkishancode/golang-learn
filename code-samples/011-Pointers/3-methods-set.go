package main

import (
	"fmt"
	"math"
)

// methods sets determine what methods attach to a TYPE
// It is exactly what the name says what is the set of methods for a given type?
// That is its methods set.

// > a NON POINTER RECEIVER - works with values that are POINTERS or NON-POINTERS

// > a POINTER RECEIVER only works with values that are POINTERS

// RECEIVERS     VALUES
// ---------------------------------------
//   (t T)          T and *T   -> receiver which is not a pointer works with both pointes and non-pointers value
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
