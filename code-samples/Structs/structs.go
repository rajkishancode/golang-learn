package main
import "fmt"
//Structs are data structure which is used stores values of different data types
// it aggregate together value of different data type

//we dont use words like object and class in go

type person struct { // we have a created a value of type person ,underline type is struct
	first string
	last  string
	age int
}

func main() {
	p1 := person{  // and here we created a value of type person and assign it to variable p1
		first: "James",
		last:  "Bond",
		age:32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:27,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first,p1.last,p1.age)
	fmt.Println(p2.first,p2.last,p2.age)
}
