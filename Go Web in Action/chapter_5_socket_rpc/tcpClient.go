package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		fmt.Printf("error during dial tcp, err: %v \n", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		// get one-line input from client, send to server
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error during read line from stdin, err: %v \n", err)
		}
		fmt.Printf("get input from stdin, %s", line)
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("user exit client!")
			break
		}
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Printf("error during connection write, err: %v \n", err)
		}
		fmt.Printf("client has sent %d bytes date to the server \n", n)
	}
}

func main() {
	Client()
}
