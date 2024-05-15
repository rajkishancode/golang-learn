//create your own type

package main

import ("fmt")


type hotdog int

var x hotdog

func main(){

   
	 
	fmt.Println(x) //print value
	
	fmt.Printf("%T\n",x) //print type
	
	x = 42
	fmt.Println(x) //print value
}