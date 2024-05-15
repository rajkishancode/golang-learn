package main

import "fmt"

// 1.create your own type "person" which will have underlying type "struct" so that it can store
// the following data
// 2.first name last name the favourite icecream flavors
// 3. create two values of TYPE person Print out the values ranging over the elements in the slice	which store favourite flavors

type person struct{
	first string
	last string
	favFlavors []string
}

func main(){
	p1 := person{
		first: "James",
		last: "Bonds",
		favFlavors: []string{"vanilla","strawberry"},
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		favFlavors: []string{"chocolate","hazelnut"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i,v:= range p1.favFlavors{
		fmt.Println("\t",i,v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i,v:= range p2.favFlavors{
		fmt.Println("\t",i,v)
	}

}

