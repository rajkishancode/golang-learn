There are multiple concurrent design pattern
1.fan-in  - Taking values from many channels and putting those values onto one channel
 Pretty common concurrent design pattern
   - We have a chunk of work
   - i dont know how much it gonna be
   - let fan them out to as many  go routines as possible
   - all they work on that as much as they can
   - and when we get results
   - we fan those result back in to another channels
   - we will get a channel just for results
 2.fan-out design pattern
 