package main
import "fmt"


func main(){
	x := 0
	fmt.Println("x before",&x) 
	fmt.Println("x before",x) 
	foo(&x)
	fmt.Println("x after",&x) 
	fmt.Println("x after",x) 
}

func foo(y *int){ // passing an address y pointing to int - *int
	fmt.Println("y before",y) //0
	fmt.Println("y before",*y) //0
	*y = 43 // dereferenced value at this address set it to 43
	fmt.Println("y after",y) //43
	fmt.Println("y after",*y) 
}
