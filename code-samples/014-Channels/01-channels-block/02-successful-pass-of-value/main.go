package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// Now that gonna RUN
// Because
/* 1. made the channel (line 6)
   2. this channel launches of (line 8) ,goroutine launched off and its running
   3. We now have two(line6,8) goroutines running in our concurrent design pattern
   4. Even if we are running on a single CPU like the compiler,the go runtime
      knows it this is legitimate , we have done our concurrent design pattern correctly
      So from multiple CPUs it will run well ,even with a single CPU it will run well.
   5. But here(line 9) we put 42 and this code BLOCKS right there in this goroutine(line8-10)
   6. But because this goroutine has sent to run on its own the other main goroutine(line6)
	  comes down here(line8) fires off this one and then the flow continues	
   7. And then this one(line12) blocks until it takes the value off
   9. So this one(line8-10) ready to put it ON and this one(line12) ready to take it OFF
	  they running at the same time
	  they interlock
	  they pass the baton
	  And Done.
   10.Its one way to get a successful pass.
*/ 