package main
import "fmt"

//create a value and assign it to a variable
//Print the address of that value

func main(){
	x := 42
	fmt.Println(x)
	fmt.Println("Address of a",&x)
}