# Concurrency 


## Concurrency vs Parallel 

1. Two Programs execute in parallel if they execute at the exact same time
2. At Time t, an instruction is being run by both P1 and P2 
3. Since each core runs one instruction at one time, we need replicated hardware

Why use parallel execution 
1. Tasks may be considered more quickly 
2. However, some tasks must be performed sequentially (Wash dish -> Dry dish) 
3. The key thing to note - some tasks are more parallelizable than others 

Why use concurrency 
1. Writing concurrent code is quite difficult 
2. Can we achieve speed up without parallelism 
    1. Design faster processors? 
    2. Get speed up without changing software 
3. In the past, perhaps we didn't need it. 
4. Von Neumman bottleneck: memory access time bottleneck
    1. Can we speed things up by putting data in the cache? 
    
5. Clock rate increase has been tapering off
6. Power wall, power use from a chip is too high. High power leads to high temperature.
7. Start and end time overlap 



Concurrent Programming 
1. Mapping tasks to the hardware 
2. OS and Go Runtime Scheduler 

Hiding latency
1. Programs can minimise waiting for I/O 