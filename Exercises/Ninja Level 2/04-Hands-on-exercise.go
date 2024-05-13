package main
import "fmt"

func main(){
	a := 42 
	fmt.Printf("%d\n%b\n%#x\n",a,a,a)
	b := a << 1
	fmt.Printf("%d\n%b\n%#x",b,b,b)
}

// assign a int to a variable
// print that int in decimal , binary and hex
// shift the bits of that int over 1 position to the left , and assigns that to a variable
// print that int in decimal , binary and hex
