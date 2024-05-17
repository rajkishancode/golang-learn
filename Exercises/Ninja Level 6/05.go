package main

import (
	"fmt"
	"math"
)

// create a type SQUARE
type square struct {
	side float64
}

// create a type CIRCLE
type circle struct {
	radius float64
}

// attach a method to each that calculates AREA and returns it
func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// create a type SHAPE which defines an interface as anything which has the AREA method
type shape interface {
	area() float64
}

// create a func INFO which takes a type shape and then prints the AREA
func info(s shape) {
	fmt.Println("Total Area", s.area())
}

// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

func main() {

	c1 := circle{
		radius: 5,
	}

	s1 := square{
		side: 4,
	}

	info(c1)
	info(s1)


}
