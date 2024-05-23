package main
import "fmt"
//Hands-on-Exercise #2

/*
	This exercise will reinforce our understanding of methods sets:
	- create a type person struct
	- attach a method speak to type person using a pointer receiver
		- *person
	- create a type human interface
		- to implicitly implement the interface , a human must have the speak method
	- create func "saySomething"
		- have it take in a human as a parameter
		- have it call the speak method
	- show the following in your code
		- you CAN pass value of type *person into saySomething
		- you CANNOT pass a value of type person into saySomething
*/

type person struct {
	first string
	age int
}

func (p *person) speak() {
	fmt.Printf("Hi i am %v , %d years old.", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}

func main(){
	p := person{first:"raj" , age:28}
	saySomething(&p) // pass value of pointer to that type

	// saySomething(p) does not work

	// p.speak() can call directly also
}
