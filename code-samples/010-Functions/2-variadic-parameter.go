package main

import "fmt"

func main(){
    x := sum(2,3,4,5,6,7,8)
	fmt.Println("The total sum is ",x);
}


func sum(x ...int)int{   //here x store unlimited values of numbers, so it called variadic paramter
	fmt.Println(x)
	fmt.Printf("%T\n",x) // type is slice of int

	sum := 0
	for i,v:= range x{
		sum += v
		fmt.Println("For item in the index position",i,"we are now adding",v,"to the total which is now", sum);
	}
	return sum;
}

