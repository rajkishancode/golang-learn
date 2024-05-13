package main

import "fmt"

func main(){

	s := "Hello playground!";

	fmt.Println(s)
	fmt.Printf("%T\n",s)

	bs := []byte(s);
	fmt.Println(bs) // it will give array of all ascii codes of string
	fmt.Printf("%T\n",bs)
}