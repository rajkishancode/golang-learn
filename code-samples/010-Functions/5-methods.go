package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//syntax for function ->
//func (r receiver) identifier(parameters) (return(s)) {code}
func (s secretAgent) speak() { // So When you HAVE a receiver(s secretAgent) it will attach this function speak()
	// to the type secretAgent
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	//one more example
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak();
	sa2.speak();
}


// Basically, using receiver you attach a method to any value of that type 
// and can access that value by chaining with a dot like in line 41,42