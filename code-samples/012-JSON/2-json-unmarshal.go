package main

import (
	"encoding/json"
	"fmt"
)

//func Unmarshal(data []byte, v any) error
// takes - slice of byte and the value pointed to by v and returns  an error
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main(){
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s) // convert json string  into slice of byte

	fmt.Printf("%T\n",s)   //string
	fmt.Printf("%T\n",bs)  // uint8 
						   // byte  -  alias for uint8
						   // uint8 -  the set of all unsigned  8-bit integers (0 to 255)
    
						 
	// people := []person{}  
	var people []person //more readable

	err := json.Unmarshal(bs,&people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Print out all the data",people)

	for i,v := range people{
		fmt.Println("\nPERSON NUMBER",i);
		fmt.Println(v.First,v.Last,v.Age);
	}
}