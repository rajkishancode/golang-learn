package main

import "fmt"

//Channels are better way of synchronising and writing concurrent code
// it is little place where you can put data




func main(){
	c := make(chan int)  //heres a channel which takes a integer
	
	//'<-' this notation is used to put value in channel
	c <- 42

  fmt.Println(<-c) 
}
// it does not run - fatal error: all goroutines are asleep - deadlock! Because Channels BLOCK.

/* code run steps
  1.  start from package and goes to func main thats the entry point of our program
  2.  it makes the channel (line 12)
  3.  it comes here and try to put 42 on the channel (line 15)
  4.  And it BLOCKS because when you send and receive on a channel
      it really like racers races in a track race and
      have to pass the baton like a relay race(hand-to-hand)
  5.  So, the transaction cannot occur until both send and receive happens at the SAME TIME
  6.  If it does not happen at the same time it blocks the send and receive 
      until the receive is ready to pull it off.
  7. THIS IS ONE OF THE IMPORTANT THING YOU SHOULD KNOW ABOUT CHANNELS.
  8. CHANNELS BLOCK - remember this .
  9. after putting value in channel(line 15) there's nothing to pull it off , so it BLOCKS.
  10. You cant jump down to printline(line 17) because before print it need to pull if off.

*/















/* #Revision 

- We learn that ,we can write concurrent code
  that can have multiple process running at the same time if we have multiple CPU

- The concurrent design pattern allow us to write our code in such a way 
  that we can have that multipe goroutines runnning at the same time and
  can have parallel code if we have multiple CPUs

- we saw how to use packaga atomic,mutex , waitgroups to start coordinating how all
  those routines are running together

- as the control flow of the problem is not like one thread of execution going through different paths
  in your code and different decisions whether it is sequential or iterative or conditional 

- Now we have multiple players on the field just like in a ballet or an orchestra or football game
  or baseball game we have to coordinate the efforts of all those different peoples/operators/individuals


 */