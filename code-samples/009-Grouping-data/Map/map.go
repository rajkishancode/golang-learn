package main
import "fmt"
// maps are powerful built-in data structure
// that store key of one type and value of other type key:value
// it store some value based upon some key
// and it allow super fast super efficient lookup
// ex we can store phone numbers of our friends by putting key as firstName and value as phone number
// And its an unordered list


func main(){

	m := map[string]int{  // So here,string is type of key and int is type of value 
		"James":32,
		"Mary":27,
	}

	fmt.Println(m["James"]) // accessing age of James using key 
	fmt.Println(m["random"]) // if its not present it will return 0 instead of error 

	v , ok := m["random"] // to check random is present or not
	fmt.Println(v) // value - 0
	fmt.Println(ok) // return false so not present

	// can use if condition to write some logic
	if v , ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT",v)
	}
}