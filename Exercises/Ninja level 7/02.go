package main
import "fmt"

//create a person struct
type person struct{
	name string	
		
}
//create a func called "changeMe" with a *person as a parameter (pointer person as parameter)
    // - change the value stored at the *person address
func changeMe(p *person){ //takes value of type and pointer to a person
		p.name = "Missmoney Penny"    //1st way
		// (*p).name = "Missmoney P"  //2nd way - deferencing the address (*P) and then .name and assign new value "Missmoney P" 	 	 
}

func main(){
	//create a value of type person
	p1 := person{
		name:"James Bond",
		}
	// Print out the value
	fmt.Println(p1)
	//call changeMe by passing the address
	 changeMe(&p1)
	//in func main ,print out the value
     fmt.Println(p1)
}

//IMPORTANT
	//- to deference a struct , use (*value).field
				// -   p1.name - you are referencing just the name field
				// - (*p1).name - you are referencing all of that person
	//- As an exception, if the type of x is a named pointer type and (*x).f is a valid
	//  selector expression denoting a field (but not a method), x.f is shorthand for (*x).f
	//-link - https://golang.org/ref/spec#Selectors