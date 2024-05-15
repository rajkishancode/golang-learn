package main
import "fmt"



func main(){
	//USING A COMPOSITE LITERAL
	//create an ARRAY which holds 5 VALUES of TYPE int
	//assign values to each index position
	 x := [5]int{11,22,33,44,55}

	fmt.Println(x)


	//range over the array and print the values out
	for i ,v:= range x{
		fmt.Println(x[i]) //print using index
		fmt.Println(v)   // print direct value
	}

	fmt.Printf("%T",x)  //print type of array using format printing
}