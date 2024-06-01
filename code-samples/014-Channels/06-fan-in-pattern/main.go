// Go playground link - https://go.dev/play/p/Sr0F_nUZCXy
// rob pike FAN IN - https://go.dev/play/p/buy30qw5MM
package main

import (
	"fmt"
	"sync"
)

/*
* Fan in - Taking values from many channels and putting those values onto one channel
*/

func main(){
   even  :=  make(chan int)
   odd   :=  make(chan int)
   fanin :=  make(chan int)

   go send(even ,odd)

   go receive(even , odd ,fanin) 

	for v := range fanin{   // It will print all value of fanin and then will close fanin channel finally at the end.
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}
//send channel		
func send(even, odd chan <- int){

	for i := 0 ; i < 10 ; i++{
		 if i % 2 == 0 {
			even <- i
		 }else{
			odd <- i
		 }
	}

	close(even)
	close(odd)
}
//receive channel
func receive(odd,even <- chan int,fanin chan <- int){// on even,odd channel  receive of value type int
								 					 // on fanin channel we  gonna send of value type int
													 // so, in receive fn we have both receive and send channel

	var wg sync.WaitGroup // here we are using waitgroups and launched 2 go functions
	wg.Add(2)			  // one is for even and one is for Odd.


	//here we take values from 2 channels(even,odd) and put it on 1 channel(fanin) - this is what fanin design pattern
	go func(){
		for v := range even { // code will block here and will range ,till that channel close
			fanin <- v        // value it will pull it off and put it on to fanin
		}
		wg.Done() 			  // these are waiting until code finished
	}()
		
	go func(){
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}


/*
	<- chan  (if  "<-"  is on left side its receive channel)
	chan <-  (if  "<-"  is on right side its send channel)
*/