package main
import "fmt"

func main(){
	 x:= []int{42,43,44,45,46,47,48,49,50,51}
	 //append 52

	 x = append(x, 52)
	 fmt.Println(x)
	
	 x = append(x, 53,54,55)
	 fmt.Println(x)

	 y := []int{56,57,58,59,60} // append to the slice this slice

	 x = append(x, y...) // print the updated slice and y... is variadic element mean unlimited no of elements of that type
	 fmt.Println(x)
}