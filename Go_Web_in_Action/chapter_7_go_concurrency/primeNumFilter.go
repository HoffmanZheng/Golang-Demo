package main

import "fmt"

// 使用筛法求素数

func intGenerator() chan int {
	var numCollect chan int = make(chan int)
	go func() {
		for i := 2; ; i++ {
			numCollect <- i
			// 从通道中取出一个，就再放入一个整数
		}
	}()
	return numCollect
}

func filterPrime(nums chan int, num int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-nums // 这边取，intGenerator 就放
			if i%num == 0 {
				continue
			} else {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	max := 100
	nums := intGenerator()
	first := <-nums
	for first <= max {
		fmt.Printf("%d \n", first)
		nums := filterPrime(nums, first)
		first = <-nums
	}
}
