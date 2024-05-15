package main
import "fmt"

// using the code from the previous example , DELETE a record FROM your map.
// Now print the map out using the "range" loop
func main(){
			m := map[string][]string{ // here types of  key is string and  value is []string
			"bond_james":[]string{"Shaken not stirred","Martinis","Women"},
			"Monneypenny_miss":[]string{"James Bond","Literature","Computer Science"},
			"no_dr":[]string{"Being evil","Ice creams","Sunsets"},
		}

		delete(m,"no_dr")
		fmt.Println(m)

		// print the map using range loop

		for k,v:= range m{
			fmt.Println(k)
			for j,v2:= range v {
				fmt.Println("\t",j,v2)
			}
		}

	}