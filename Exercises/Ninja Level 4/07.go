package main
import "fmt"
//1.create a slice of slice of string ([][]string)
//2. store the following data in multidimensional slice
// -"James","Bond","Shaken not stirred"
// -"Miss","Money Penny","Hello James"
//3.Range over the record ,then range over the data in each record

func main(){

	x := [][]string{}
	// fmt.Printf("%T",x)

	y := []string{"James","Bond","Shaken not stirred",}
	z := []string{"Miss","Money Penny","Hello James"}
	
	x = [][]string{y,z}
	fmt.Println(x)

	//range over the record
	for i,v := range x{
		fmt.Println(i,v)
		// and then range over the data in each record
		for j,v2 := range v{
			fmt.Println(j,v2)	
		}
	}
}