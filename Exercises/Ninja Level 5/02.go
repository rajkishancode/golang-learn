package main

import "fmt"

//
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
	// store values of p1 and p2 in map ,
	m := map[string]person{
		p1.last:p1,		// key is p1 last and value is whole p1
		p2.last:p2,		// key is p2 last and value is whole p2
	}

	for _,v:= range m{
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors{
				fmt.Println("\t",i,val)
		}
	}

}

