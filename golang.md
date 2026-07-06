
1. Golang Garbage Collection (GC)


Golang Garbage collection (GC) is a process that automatically manages memory allocation and deallocation in Go programs. The garbage collector identifies and frees up memory that is no longer in use, preventing memory leaks and improving application performance. Go's garbage collector is designed to be efficient and minimize pause times, allowing for smooth execution of concurrent applications.


What is pprof()?
pprof() is a profiling tool in Go that allows developers to analyze the performance of their applications. It provides insights into CPU usage, memory allocation, and other performance metrics by generating profiles that can be visualized and analyzed. This helps identify bottlenecks and optimize code for better performance.


It uses tri color marking and concurrent sweeping to manage memory. The garbage collector runs in the background and can be triggered automatically or manually, depending on the application's needs. Developers can also tune the garbage collector's behavior using environment variables and runtime settings to optimize performance for specific workloads.

Stack does not need to be garbage collected, as it is automatically managed by the Go runtime. However, heap memory, which is allocated dynamically, is subject to garbage collection. Developers can monitor and analyze the performance of the garbage collector using tools like pprof() to ensure efficient memory management in their applications.

Escape Analysis is a technique used by the Go compiler to determine whether a variable can be allocated on the stack or if it needs to be allocated on the heap. If a variable "escapes" the function scope, it is allocated on the heap, which means it will be subject to garbage collection. Conversely, if a variable does not escape, it can be allocated on the stack, which is more efficient and does not require garbage collection.

sync.Pool() is a mechanism in Go that provides a way to reuse objects and reduce memory allocations. It maintains a pool of pre-allocated objects that can be acquired and released by goroutines, helping to minimize the overhead of frequent allocations and deallocations. It helpes to improve performance by reducing the pressure on the garbage collector, as objects in the pool can be reused instead of being garbage collected and reallocated.

GOGC is an environment variable in Go that controls the garbage collector's target percentage of heap growth. By adjusting GOGC, developers can influence how aggressively the garbage collector runs. A higher GOGC value means the garbage collector will run less frequently, allowing for more heap growth before triggering a collection cycle, while a lower value will cause the garbage collector to run more often, potentially reducing memory usage but increasing CPU overhead.

