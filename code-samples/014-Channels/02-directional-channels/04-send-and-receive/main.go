package main
import "fmt"

func main(){
	c := make(chan int)
	cr := make(<- chan  int)//receive channel - receiving from a channel of int
	cs := make(chan <- int) //sent channel- sending on chan an int

	fmt.Println("-----Printing out the types------")
	fmt.Printf("c\t%T\n",c)
	fmt.Printf("cr\t%T\n",cr) 
	fmt.Printf("cs\t%T\n",cs) 
}