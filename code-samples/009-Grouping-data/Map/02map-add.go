package main

import "fmt"

func main(){

	m := map[string]int{   
		"James":32,
		"Mary":27,
	}

	//adding new key value in map data structure
	m["Bob"]=30
	fmt.Println(m)

	// loop over map
	for k,v := range m {  // k= key, v= value , m is the map
		fmt.Println(k,v)
	}
}