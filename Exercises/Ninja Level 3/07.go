package main
import "fmt"
// use else if
func main(){
	marks:= 90

	if marks == 90 {
		fmt.Println("Congrats you got 90")
	}else if marks == 80 {
		fmt.Println("You got less than 80")
	}else{
		fmt.Println("You Passed")
	}
}