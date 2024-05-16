package main

import "fmt"

//create a new type vehicle,underlying type is struct and fields are doors and color
type vehicle struct {
	doors int
	color string
}

// create two new types truck and sedan , the underlying type of each of these new types is a struct
// embed the vehicle type in both truck and sedan
// give truck the field "fourWheel" which will be set to bool
// give sedan the field "luxury" which will be set to bool

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// using the vehicle , truck and sedan structs:
	// using a composite literal , create a value of type truck and assign values to the fields
	t := truck{
		vehicle: vehicle{
			
			doors: 2,
			color: "white",
		},
		
		fourWheel: true,
	}
	// using a composite literal , create a value of type sedan and assign values to the fields
	s := sedan{
		vehicle: vehicle{
			doors:4,
			color:"silver",
		},
		luxury:true,
	}
	//print out each of these values
	fmt.Println(t)
	fmt.Println(s)
	
	//print out a single field from each of these values
	fmt.Println(t.doors)
	fmt.Println(s.doors)
	// fmt.Println()



}
