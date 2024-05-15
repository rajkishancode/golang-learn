package main
import "fmt"

// - create a map with a key of TYPE string which is person's 	"last_first" name
// - and a value of TYPE []string which stores their favourite thing.Store three records in your map.
// - Print out all the values along with their index position in the slice

func main(){
	m := map[string][]string{ // here types of  key is string and  value is []string
			"bond_james":[]string{"Shaken not stirred","Martinis","Women"},
			"Monneypenny_miss":[]string{"James Bond","Literature","Computer Science"},
			"no_dr":[]string{"Being evil","Ice creams","Sunsets"},

	}
	fmt.Println(m)
	for k,v := range m{  // first range to loop over the map key and value
		fmt.Println("This is the record for",k)
		for i,v2 := range v{  // second range to loop over data of values
			fmt.Println("\t",i,v2)
		}
	}
}