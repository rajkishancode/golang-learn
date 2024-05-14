package main
import "fmt"

// create a for loop using this syntax 
// for condition {}
// print out the years you have been alive

func main(){
	bd := 1995
	for bd <= 2024 {
		
		fmt.Println(bd)
		bd++
	}
}