package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people) //[{James Bond 32} {Miss Moneypenny 27}] -this is go data ,now we need to marshal it to sent it to somebody
	bs , err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) // convert bs- byte slice to string
}

//In Documentation ,
//func Marshal(v interface{}) (byte[],err)
// it takes    (v interface{}) means value of any type
// and returns (byte[],err)    means slice of byte and an error


//Note initially  you get [{},{}]
// Make the first letter of field name to Capital like first - First everywhere 
//then it will works and return json and you can sent it  
//[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]