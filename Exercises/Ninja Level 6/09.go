package main

import "fmt"
// a callback is when we pass a func into a func as an argument.For this exercise,
// pass a func into a func as an argument

func sum(n ...int)int{
	total := 0
	for _,v:= range n{
		total += v
	}
	return total
}
func oddSum(f func(xi ...int)int, n ...int)int{
	var yi []int
	for _,v:= range n{
		if(v % 2 != 0){
			yi = append(yi,v)
		}
	}
	return f(yi...)
}
func main(){

	 n := []int{1,2,3}
	 s := oddSum(sum,n...)
	 fmt.Println(s)	

}

// i created two function 1.sum and 2.oddSum
// ii pass sum function  as a callback in oddSum function  