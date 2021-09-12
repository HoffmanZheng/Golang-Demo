package main

import (
	"fmt"
	"log"
	"net"
)

func Server() {

	l, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		fmt.Printf("error during listening net, err: %v \n", err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("error during accepting connection, err: %v \n", err)
		}
		fmt.Printf("connect from client: %v, ip: %v \n", conn, conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		fmt.Printf("server wait message from client: %s \n", c.RemoteAddr().String())
		buf := make([]byte, 1024)
		n, err := c.Read(buf) // block if no messenge received
		if err != nil {
			log.Fatal("error during read from connection, err: ", err)
			break
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	Server()
}
