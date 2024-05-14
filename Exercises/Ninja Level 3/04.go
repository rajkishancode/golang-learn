package main
import "fmt"


//create a loop using this syntax
// for {}
// print out the years you have been alive

func main(){
  	bd := 1995
	for {
		if bd >= 2024 {
			break
		}
		bd++
		fmt.Println(bd)
	}
}