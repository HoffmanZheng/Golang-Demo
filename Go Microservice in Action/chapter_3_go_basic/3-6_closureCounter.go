package main

import "fmt"

// 使用闭包的特性实现一个简单的计数器

func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	return func() int {
		initial++
		return initial
	}
}

func main() {
	c1 := createCounter(1)
	fmt.Printf("调用一号计数器：%d \n", c1())
	fmt.Printf("调用一号计数器：%d \n", c1())

	c2 := createCounter(10)
	fmt.Printf("调用二号计数器：%d \n", c2())
	fmt.Printf("调用二号计数器：%d \n", c2())
}
