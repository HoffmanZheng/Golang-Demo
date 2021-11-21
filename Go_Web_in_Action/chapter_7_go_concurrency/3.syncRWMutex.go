package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex = sync.RWMutex{}
var randomNum int
var wait = sync.WaitGroup{}

func main() {
	for i := 0; i < 30; i++ {
		wait.Add(1)
		if i%7 == 0 {
			go write(i)
		} else {
			go read(i)
		}
	}
	// time.Sleep(10 * time.Second)
	wait.Wait()
}

func read(i int) {
	fmt.Printf("goroutine %d, reading start... \n", i)
	rwMutex.RLock()
	fmt.Printf("goroutine %d, reading result: %d \n", i, randomNum)
	time.Sleep(time.Second)
	rwMutex.RUnlock()
	wait.Done()
	fmt.Printf("goroutine %d, reading end... \n", i)
}

func write(i int) {
	fmt.Printf("goroutine %d, writing start... \n", i)
	rwMutex.Lock()
	v := rand.Intn(10)
	randomNum = v
	fmt.Printf("goroutine %d, writing into: %d \n", i, randomNum)
	time.Sleep(time.Second)
	rwMutex.Unlock()
	wait.Done()
	fmt.Printf("goroutine %d, writing end... \n", i)
}
