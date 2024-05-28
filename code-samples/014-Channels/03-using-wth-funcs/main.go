package main
import "fmt"

func main(){

	c := make(chan int)

	//send
	go foo(c)

	//receive
	 bar(c)

	fmt.Println("About to exit")
}


//sent
func foo(c chan<- int){  // sent only channel - we are sending int from this channel(<-chan).
						 // the scope of this c is only this functions scope and channels are reference type so they are referencing the same underying data
	c <- 42	//we pass 42 onto c			 
}    


//receive
func bar(c <-chan int){ // receive only channel - we receiving from this channel(<-chan)  int.
	fmt.Println(<-c)  // println is pulling the value as we are receiving the value here
}

/*
in funcs you can specify
 > receive channel
   - you can receive values from the channel
   - a receive channel parameter
   - in the func, you can only pull values from the channel
   - you can't close a receive channel
 > send channel
   - you can push values to the channel
   - you can't receive/pull/read from the channel
   - you can only push values to the channel

*/