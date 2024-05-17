package main

import "fmt"
// 1.create a function with the identifier foo
  // - takes a variadic paramter of type int
  // - pass in value of type []int in your func (unfurl the []int)
  // - return the sum of all values of type int passed in

func foo(xi ...int)int{
	sum := 0
	for _,v:= range xi{
		sum += v
	}	
	return sum;	

  }

  // create a func with the identifier bar that
     //- takes in parameter type []int
    // - return the sum of all values of type int passed in
  func bar(xi []int)int{
	 sum := 0
	 for _,v:=range xi{
		sum += v
	 }
	 return sum;
  }

func main(){
	//1
	n := []int{1,2,3,4,5}
	f := foo(n...)
	fmt.Println(f)
	
	//2
	n2 := []int{1,2,3,4,5}
	f2 := bar(n2)
	fmt.Println(f2)
}