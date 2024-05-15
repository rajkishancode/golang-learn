package main
import "fmt"


//embedded structs
// here we will learn how to take one type and embedded into another type

type person struct{
	first string
	last string
	age int
}
type secretAgent struct{
	person // here you dont need to give field name just put person 
	ltk bool
}
func main(){
	sa1 := secretAgent{
		person:person{ // but here you have to put the field name person and then its value,person is type & {} - composite literal
		first:"James",
		last:"Bond",
		age:33,
	},
		ltk :true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first,sa1.last,sa1.age)
}

// These first,last,second are inner-fields/inner-types of person and got promoted to type secretAgent 
// we dont have to do sa1.person.first etc
// we are directly accessing values sal.first 