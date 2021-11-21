# Concurrent Programming in Go

### Coroutines and Go Concurrency Model

1. Process is the basic unit of operating system for resource allocation and scheduling, whereas thread is the smallest unit of program execution flow, has independent stack, shared heap, and switched by the operating system.

2. Otherwise, a coroutine is a lighter-weight function than a thread. It is completely controlled by the program and executed in a **user mode**, so it has a significant performance improvement and does not consume resources like thread switching.

3. A thread can have multiple coroutines, which can switch between each other. These coroutines run in a **serial manner** and cannot take advantage of the multi-core capability of the CPU. The coroutine can be suspended somewhere, and can resume running at the suspended point.

4. Unlike the multithreaded shared memory model of Java and C++, Go uses CSP(Communicating Sequential Processes) as concurrency model. CSP does not communicate through shared memory, but shares memory through communication.

### Concurrency Impremented by Goroutine and Channel

1. `channel <- data` could send data to the certain channel, while `<- channel` could receive data from the channel. These two function should appear in pairs, otherwise, no matter the transmission or fetching, it will be blocked until another goroutine transmission or fetching.

2. The shared data type should be specified when declare the channel, which could be build-in, named, struct, reference type. At the same time, **only one** goroutine could send or receive data through the channel.

3. The sender's data transmission or the receiver's data acquisition is allowed to be in an asynchronous state when using a channel with a buffer, like `ch := make(chan int, 66)`. In this way, the sender will only be blocked when the buffer is full.

4. Go supports `select` keyword to deal with async I/O issues, like select() in UNIX, see: [1.channelSelect.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_7_go_concurrency/1.channelSelect.go)

### Concurrency Impremented using `sync` Package

1. Race conditions often cause unexpected problems in concurrent requests due to their non-atomic behavior. In order to avoid race conditions, Golang provides some solutions, such as channel blocking, mutual exclusion locks, etc.

2. Once `sync.Mutex` is obtained by a goroutine, other goroutine could only wait until the Mutex is released. Otherwise leads to a deadlock. See: [2.syncMutex.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_7_go_concurrency/2.syncMutex.go)

3. A `sync.RWMutex` is usually used in an application scenarios where read more and write less, see: [3.syncRWMutex.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_7_go_concurrency/2.syncRWMutex.go)

4. `sync.Once` is used to ensure that the function is only been called once, like close the channel, see: [4.syncOnce.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_7_go_concurrency/4.syncOnce.go)

5. A race detector could help find the potential concurrency issues, run `go run -race main.go`

### Go Concurrency Demo

1. Use the sieve method to obtain prime numbers, see: [5.primeNumFilter.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_7_go_concurrency/5.primeNumFilter.go)