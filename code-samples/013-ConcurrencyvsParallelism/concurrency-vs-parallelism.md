Differnce between Concurrency vs Parallelism

- If you write code in go and you run that code on a machine that has only one CPU
- Your code is not gonna run in Parallel means 
               not gonna run at the same time 
               not gonna have multiple threads of code in Parallel running at the same time

               - it gonna run sequentially one instruction one after the other because you only have one CPU
               - If you have more than one CPU ,then you have the opportunity to run code in Parallel 
                 and that Parallelism.

                 Whats Concurrency ?
                 - Concurrency is a design pattern
                 - its a way that you write your code
                 - you can write concurrent code
                 - you can have code which has the possibility and potential to run in parallel
                 - So , if you have multiple cores that concurrent code can run in Parallel because you multiple 
                   cores
                 - So in parallel like you can have four threads of instructions executing at the sametime 
                    or executing in parallel
                
                 -But Concurrency which is a design pattern which is a way that you write code does not guarantee
                  that your code gonna run in parallel/ parallelism
                  
                 - whether you have more than one CPU determines your code run in parallel or not
                 - so if you have more than one CPU you ODDS are increased that you are able to run
                    your  concurrent code in parallel