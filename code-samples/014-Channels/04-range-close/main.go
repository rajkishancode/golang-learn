package main
import "fmt"

func main(){

	c := make(chan int)

	//send
	go func(){
		for i := 0 ; i < 5 ; i++ {
			c <- i //send value on channel
		}
		close(c) //close channel
				 // if not close , will wait for more values and (fatal error: all goroutines are asleep - deadlock!)
	}()

	//receive
	 for v := range c {  // range will loop over channel ,till channel will close
		fmt.Println(v)
	 }

	fmt.Println("About to exit")
}


   

// that how range workd, til pulls value from channel , till channel close