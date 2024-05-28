package main
import "fmt"
/*Two big pieces we have learn till now
 1.Channels Blocks
 2. Range clause which also Blocks until the channel close
 3. we will look on here tools like using a SELECT statement to work on channel 		
*/

//Now we will use SELECT STATEMENT TO PULL VALUES OFF

func main(){
	even := make(chan int)
	odd  := make(chan int)
	quit := make(chan int)

	//send - send values on to the channel
	go send(even,odd,quit)
	//receive
	receive(even,odd,quit)

	 fmt.Println("About to exit") //7. Finally , it ends here.
}

	func send(e,o,q chan <- int){ // here we sent on to the channel int [READING FROM LEFT TO RIGHT]
       for i := 0; i < 100; i++ {
			if i % 2 == 0 { // put value on even or odd channel based on whatever value i is?
				e <- i
			} else {
				o <- i
			}
	   }
	   q <- 0
	}

	func receive(e,o,q <- chan int){ // we will receive from channel that has value of type int
		for { // its a infinite loop till we get out of it
			select { // 1. then we have select statement that check from which channel we can get value of
					 // 2. it will get that value assign it to v and print it off.
					 // 3. then leave the select statement go back to above for loop
					 // 4. repeat until we get value of the quit channel 
				case v := <- e:
						fmt.Println("From the even channel",v) 	
				case v := <- o:
					fmt.Println("From the odd  channel", v)
				case v := <- q:
					fmt.Println("From the quit channel",v) // 5. In some case ,we get this value of quit channe
					return;                                // 6. And then returns and stops running 
			}
		}
	}

	/*
	*
	*
	*/