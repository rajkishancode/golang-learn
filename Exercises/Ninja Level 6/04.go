package main

import "fmt"

//create a user defined struct with
   //- the identifier person
   //- the fields
     // - first , last , age

 type person struct{
	first  string
	last string
	age int
 }


// attach a method to type person with
    // the identifer speak
    // the method should have the person say their name and age

func (p person) speak(){
	fmt.Println("Hi My name is",p.first,p.last,"and i am",p.age,"years old.");
}

// create a value of type person
// call the method from the value of type person

func main(){
	p :=  person{
		first: "James",
		last: "Bond",
		age: 33,
	}
	p.speak()

}