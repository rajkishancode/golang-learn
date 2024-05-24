package main
import "fmt"

func main(){
	c := make(chan int,2)  // this is sent-only channel - only send value to a channel

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-------------")
	fmt.Printf("%T\t",c)
}