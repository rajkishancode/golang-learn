package main
import ("fmt"
		"sort")


type Person struct{
	Name string
	Age int
}
//ByName implements sort.Interface for []Person based on the Age field.
type ByName []Person
//bn denotes BYNAME
func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main(){
	p1 := Person{"James",32,}
	p2 := Person{"Moneypenny",27,}
	p3 := Person{"Q",64,}
	p4 := Person{"M",65,}

	people := []Person{p1,p2,p3,p4}
	
	fmt.Println(people)
	//SORT BY NAME
	sort.Sort(ByName(people))
	fmt.Println(people)
}