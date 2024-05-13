//create your own type

package main

import ("fmt")


type hotdog int

var x hotdog
var y int
func main(){

	fmt.Println(x) 
	fmt.Printf("%T\n",x) 
	
	x = 42
	fmt.Println(x) 

	y = int(x)  // conveting type of x to int from hotdog
	fmt.Println(y) //print value
	fmt.Printf("%T",y) // print type
}