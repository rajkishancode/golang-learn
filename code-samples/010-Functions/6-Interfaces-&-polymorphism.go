package main

import "fmt"

//interfaces allow us to define behaviour and allow us to do POLYMORPHISM
// a value can be of more than one type

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// speak attach to type secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "the secretAgent speak")
}

// speak attach to type person
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "the person speak")
}

//JOKE
//INTERFACE SAYS : HEY BABY IF YOU GOT THIS METHOD THEN YOU ARE MY TYPE
// as shown above person and secretAgent have speak method so they are human TYPE
//And any type attach to speak , is also type human
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type){
	case person:
			fmt.Println("i was in the barrrrr",h.(person).first)
	case secretAgent:
			fmt.Println("I was in the barrrrr",h.(secretAgent).first)
	}
	fmt.Println("I was passed in the bar", h)
}

type hotdog int

func main() {
	//below value is of the type secretAgent but it has speak method attach to it ,so it belongs to human TYPE also
	// so value can be of more than one type ex - 1.secretAgent and 2. human
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	//create a new person
	p1 := person{
		first: "Dr",
		last:  "No",
	}
	fmt.Println(sa1)  //{{James Bond } true}
	sa1.speak()      // I am James Bond the secretAgent speak
	sa2.speak()		 // I am Miss Moneypenny the secretAgent speak

	fmt.Println(p1)   // {Dr No}

	bar(sa1) // here bar able to take both secretAgent and person type because both are type human and we passed human in bar
	bar(sa2) // and this is POLYMORPHISM POLY-MANY , MORPHISM -CHANGE.
	bar(p1) 
	//i was in the barrrrr James
	//i was in the barrrrr {{James Bond} true}
	//i was in the barrrrr Miss
	//i was in the barrrrr { {Miss MoneyPenny} true}
	//i was in the barrrrr Dr
	//I was passed in the bar {Dr No}

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

//interfaces allow value to be of more than one type
