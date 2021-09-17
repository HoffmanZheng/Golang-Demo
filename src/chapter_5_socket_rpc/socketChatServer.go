package main

import (
	"fmt"
	"net"
	"time"
)

type Heartbeat struct {
	expireTime int64
}

var ConnSlice map[net.Conn]*Heartbeat

const expireInterval = 10

func ChatServer() {
	l, err := net.Listen("tcp", "localhost:8099")
	if err != nil {
		fmt.Printf("error during listening tcp, err: %v \n", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("error during accepting from listener, err: %v \n", err)
		}
		fmt.Printf("received message from %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// init expireTime in heartbeat
		ConnSlice[conn] = &Heartbeat{
			expireTime: time.Now().Add(time.Second * expireInterval).Unix(),
		}
		fmt.Println("start to handle connection")
		go handlerConn(conn)
	}
}

func handlerConn(c net.Conn) {
	defer c.Close()
	fmt.Println("get Conn in handler!")
	buf := make([]byte, 1024)
	for {
		fmt.Printf("server wait message from client: %s \n", c.RemoteAddr().String())
		n, err := c.Read(buf)
		fmt.Printf("server received %d bytes from client: %v \n", n, c.RemoteAddr())
		// update expireTime in heartbeat or disconnect from client
		if ConnSlice[c].expireTime > time.Now().Unix() {
			ConnSlice[c].expireTime = time.Now().Add(time.Second * expireInterval).Unix()
		} else {
			fmt.Println("server disconnect from client due to longterm missing message!")
			return
		}
		if err != nil {
			fmt.Printf("error during reading connection, err: %v \n", err)
			return // client disconnect!
		}
		// if heartbeat check, skip the rest code
		if string(buf[0:n]) == "1" {
			c.Write([]byte("1"))
			continue
		}
		// iterate the conn in ConnSlice, delete the conn, which failed the heartbeat check
		for conn, heart := range ConnSlice {
			if conn == c { // current conn already been processed
				continue
			}
			if heart.expireTime < time.Now().Unix() {
				delete(ConnSlice, conn)
				conn.Close()
				fmt.Printf("delete conn: %v \n", conn.RemoteAddr())
				fmt.Printf("currently existing conn: %v \n", ConnSlice)
				continue
			}
			conn.Write(buf[0:n]) // other connection, send the message from current conn
		}
	}
}

func main() {
	ConnSlice = map[net.Conn]*Heartbeat{}
	ChatServer()
}
