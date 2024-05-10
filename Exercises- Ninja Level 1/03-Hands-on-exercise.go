package main

import "fmt"


var x int
var y string
var z bool

func main(){
	x = 42
	y = "James Bond" 
	z = true

	s := fmt.Sprintf(x,y,z)
	fmt.Sprintf(s)
	

}

//start from here 13th may monday