package main

import "fmt"

// make is a built in function in go

func main() {
	
	x := make([]int,10,12) // make take slice type , len and capacity  
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x,3433)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x,3435)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	x = append(x,3436) // here capacity was 12 and you added 13th number so it will auto double the capacity and makes it 24
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}