package main
import "fmt"

//When to use Pointer
// - Lets say we have large chunk of data 
// - and you dont want pass the big chunk of data around through your program
// - You just passed the address where the data is stored
// - You need to change something at a certain location you can use a pointer.


func main(){
	x := 0
	foo(x)
	fmt.Println(x) //0
}

func foo(y int){
	fmt.Println(y) //0
	y = 43
	fmt.Println(y) //43
}




//NOTE - EVERYTHING IN GO is PASS BY VALUE	,
// DROP PHRASED AND CONCEPTS YOU MAY HAVE FROM OTHER LANGUAGES,
// pass by reference , pass by copy forget those values.