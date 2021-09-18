package main

import (
	"fmt"
	"sync"
)

// 保证只执行以此和并发安全

var wait = sync.WaitGroup{}
var once = sync.Once{}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wait.Add(3)
	go func1(ch1)
	go func2(ch1, ch2)
	go func2(ch1, ch2)
	wait.Wait()
	for j := range ch2 {
		fmt.Printf("%d \n", j)
	}
}

func func1(ch chan<- int) {
	defer wait.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func func2(ch1 <-chan int, ch2 chan<- int) {
	defer wait.Done()
	for {
		i, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- 2 * i
	}
	once.Do(func() { close(ch2) })
}
