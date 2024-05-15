package main
import "fmt"

//To DELETE from a slice , we use append along with SLICING 
//Q use APPEND & SLICING to get below values:-
//[42,43,44,48,49,50,51]
func main(){
	x := []int{42,43,44,45,46,47,48,49,50,51}
	x = append(x[:3],x[6:]...)  
	fmt.Println(x)

}