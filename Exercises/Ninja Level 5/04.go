package main

import "fmt"

//create and use an anonymous struct.
// slice s := []int{}
// map 	 m := map[string]int

func main() {

	s := struct {
		
		first []string
		friends map[string]int
		favDrinks []string
	}{
		 first:[]string{"John"},
		 friends: map[string]int{"bond":007,},
		 favDrinks: []string{"champange","water"},
	}
		
		fmt.Println(s.first)
		fmt.Println(s.friends)
		fmt.Println(s.favDrinks)

		for k,v:= range s.friends{
			fmt.Println(k,v)
		}
		for i,v:= range s.favDrinks{
			fmt.Println(i,v)
		}


	}


