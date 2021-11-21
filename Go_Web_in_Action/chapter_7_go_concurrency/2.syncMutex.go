package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{} // like countdownLatch in Java
	mutex.Lock()
	fmt.Println("lock in main thread")
	for i := 0; i < 5; i++ {
		wait.Add(1)
		go func(i int) {
			fmt.Printf("before lock, i: %d \n", i)
			mutex.Lock()
			fmt.Printf("lock, i: %d \n", i)
			time.Sleep(time.Second)
			fmt.Printf("before unlock, i: %d \n", i)
			mutex.Unlock()
			fmt.Printf("unlock, i: %d \n", i)
			wait.Done()
		}(i)
	}
	mutex.Unlock()
	fmt.Println("unlock in main thread")
	wait.Wait()
	fmt.Println("all tasks have been finished")
}
