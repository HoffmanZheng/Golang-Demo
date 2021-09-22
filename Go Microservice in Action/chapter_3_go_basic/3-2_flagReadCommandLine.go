package main

import (
	"flag"
	"fmt"
)

// 使用 flag 从命令行读取参数
// go run 3-2_flagReadCommandLine.go -lastName=Zheng -firstName=Hoffman -id=3309
func main() {
	lastName := flag.String("lastName", "王", "your last name")
	var firstName string
	// 直接传入变量地址
	flag.StringVar(&firstName, "firstName", "David", "your first name")
	id := flag.Int("id", 0, "your Id")
	flag.Parse() // 解析命令行参数
	fmt.Printf("I am %s %s, and my id is %d \n", firstName, *lastName, *id)
}
