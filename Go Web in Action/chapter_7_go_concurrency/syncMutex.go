package main

import (
	"fmt"
	"sync"
	"time"
)

/** sync.Mutex 在同一个 goroutine 中对解锁前的 Mutex 再次加锁会造成死锁 */

func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	mutex.Lock()
	fmt.Println("lock in main thread")
	for i := 0; i < 5; i++ {
		go func(i int) {
			wait.Add(1)
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
