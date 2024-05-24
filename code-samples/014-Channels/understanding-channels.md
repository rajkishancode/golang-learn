# Understanding Channels
Channels Introduction 
GO specs          - https://go.dev/ref/spec#Channel_types
GO Effective Docs - https://go.dev/doc/effective_go#channels

- making a channel
     c := make(chan int)
- putting values on a channel
     c <- 42
- taking values off a channel
     <-c
- buffered channels
     c := make(chan int,4)
- channels block
    - they are like runners in a relay race
       - they are synchronized
       - they have to pass/receive the value at the same time
          - just like runners in a relay race have to pass/receive the
            baton to each other at the same time
                - one runner can't pass the baton at one moment
                - and then , later have the other runner receive the baton
                - the baton is passed/received by the runners at the
                  same time
          - the value is passed/received synchronously; at the same time

- channels allow us to pass values between goroutines

#Golang Spec
"
The capacity, in number of elements , sets the size of the buffer
 in the channel. If the capacity is zero or absent, the channel
 is unbuffered and communication succeeds only when both a sender
 and receiver are ready.

"