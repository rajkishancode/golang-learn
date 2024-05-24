package main

import "fmt"

func main() {
	c := make(chan int , 1) // we put 1 dl here
	
	c <- 42
	
	fmt.Println(<-c)
}

/*
  - This is a buffer channel
  - we put a dl here in line 6, second parameter
  - Buffer channel is a channel that allows certain value to sit in that channel
	regardless whether or not something there ready to pull if off
  - ok so i can say that my buffer channel will allow 1 one value to sit in there
  - so we get to this code(line8) like 42 is put on the channel
  - as i got a buffer of 1 , some room is there for it
  - And we can go past that line8 as 42 got put on the channel,its no longer blocking
    and will be able to pull it off.
  - Finally the code ran , this is the concept of Channel Block
*/