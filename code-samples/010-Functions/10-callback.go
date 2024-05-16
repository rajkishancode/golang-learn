package main

import "fmt"

//CALLBACK
//passing a function as a argument
//functional programming that is not recommended in go , however it is good to be aware of callbacks
//idiomatic: write clear,simple, readable code

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(x...)

	fmt.Println("Total sum ", s)

	s2 := evenSum(sum, x...) // it take a sum function as an  argument and slice of int as second argument
	fmt.Println("Sum of even numbers", s2)

	s3 := oddSum(sum,x...)
	fmt.Println("Sum of odd numbers",s3);

}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

func evenSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func oddSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {

			yi = append(yi, v)
		}
	}
	return f(yi...);
}
