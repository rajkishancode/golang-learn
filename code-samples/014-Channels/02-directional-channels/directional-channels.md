# Directional Channels

- Channel type. Read from left to right

code :
      - seeing the type
         - code from previous video
            - https://go.dev/play/p/a98otBr4eX

         - send & receive (bidirectional)
            - https://go.dev/play/p/di1mKkGF6Y

            - "send and receive" means "send and receive"
               https://go.dev/play/p/SHr3lpX4so
                 -already saw the above code

            - send means send
                - error:"invalid operation: <-cs (receive from send-only type chan<-int>)"
                https://go.dev/play/p/oB-p3KMiH6

            - receive means receive
                - error:"invalid operation: cr <-42 (send to receive-only type <-chan int)"
                https://go.dev/play/p/_DBRuelmEq
                
            - "A channel may be constrained only to send or only to receive [general to specific] by conversion or assignment" - Goland Spec
              - doesn't assign
                - specific to general
                   - https://go.dev/play/p/bG7H6l03VQ
                   - specific to specific 