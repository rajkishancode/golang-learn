package main

import "fmt";

func main(){
	xi := []int{1,2,3,4,5,6,7,8,9,10};

	s := sum(xi...);  // unpacking/unfurling a  slice  [spread in javascript]
	fmt.Println(s);
}

func sum(x ...int)int{   // so variadic mean 0 or more , if you dont pass anything zero will be printed of type int
		fmt.Println(x)	
		fmt.Printf("%T\n",x)
		sum := 0
		for _,v:= range x{
			sum += v
		}	
		return sum;
}