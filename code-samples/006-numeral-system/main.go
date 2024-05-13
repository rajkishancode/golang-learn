package main

import "fmt"

func main(){

	s := "H"
	fmt.Println(s)
	
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n",s)
	
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n",n) // print type
	fmt.Printf("%b\n",n) // print binary
	fmt.Printf("%#X\n",n) // print hexadecimal
	
}