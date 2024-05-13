package main
import "fmt"


//create TYPED AND UNTYPED CONSTANTS

const (
	a = 42
	b int = 43
)

func main(){
	fmt.Println(a,b)
}