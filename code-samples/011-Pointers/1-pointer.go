package main

import "fmt"

//Pointers is like pointing to some location in some memory where value is stored


func main(){
     a:= 42
	 fmt.Println(a)
	 fmt.Println(&a) // Ampersand & gives you the address

	 fmt.Printf("%T\n",a) // type of a is int
	 fmt.Printf("%T\n",&a) //type of address is *int (Pointer to an int)
	 					   

	 b:= &a
	 fmt.Println(b)
	 fmt.Println(*b) // * gives you the value stored at an address when you have the address
	 fmt.Println(*&a)

	 *b = 43 // value at *b address set it to 43
	 println(a) // a value is now 43
	                 
}



