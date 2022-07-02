package main

import (
	"fmt"
	"time"
)

/**
1. 线程是一个通过共享堆内存来通信的，由操作系统调度和分派的基本单位
2. 协程是完全由用户态的程序所控制的，更加轻量级的一种函数，它可以在某个地方被挂起，并且可以重新在挂起处继续运行
3. CSP Communicating sequential Processes 通过通信来共享内存
4. channel 不管传还是取必阻塞， 直到另外的 goroutine 传或取为止
5. goroutine 协同式的调度会在几个时刻进行切换：新创建 goroutine，I/O 阻塞，系统调用阻塞等
6. 开启 goroutine 的函数返回时，goroutine 也会自动结束
7. 同时只能有一个 goroutine 访问通道
8. 非阻塞方式从通道接收数据 data, ok := <- ch 未接收到的话 data 就是通道类型的零值，CPU 占用高，使用较少
*/

func main() {

	ch := make(chan string)

	go func() {
		ch <- "goroutine"
	}()

	fmt.Println(<-ch)

	ch1 := make(chan string)
	ch2 := make(chan string)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(1)
		ch2 <- "two"
	}()

	go func() { // set a timeout
		time.Sleep(5 * time.Second)
		timeout <- true
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("%d receive from ch1, %s \n", i, msg1)
		case msg2 := <-ch2:
			fmt.Printf("%d receive from ch2, %s \n", i, msg2)
		case msg3 := <-timeout:
			fmt.Printf("%d timeout...%v \n", i, msg3)
		}
	}
}
