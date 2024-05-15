package main
import "fmt"
//USING A COMPOSITE LITERAL
	//create an SLICE of TYPE int
	//assign 10 values
	// print type of slice
func main(){

	x := []int{11,12,13,14,15,16,17,18,19,20,21}
	for i,v := range x{
		fmt.Println(i,v)
	}

	fmt.Printf("%T",x)
}