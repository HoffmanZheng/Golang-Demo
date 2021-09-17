package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

const expireInterval = 10

func ChatClient() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8099")
	if err != nil {
		fmt.Printf("error during net dial, err: %v \n", err)
	}
	Log(conn.RemoteAddr().String(), "connect successfully!")
	return conn
}

// send heartbeat to server periodically
func Sender(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	go func() { // ??? function must be invoked in go statement
		t := time.NewTicker(time.Second * 1)
		defer t.Stop()
		for {
			<-t.C //
			_, err := conn.Write([]byte("1"))
			if err != nil {
				fmt.Printf("error during writing in connection, err: %v \n", err)
				break // return
			}
		}
	}()

	name := ""
	fmt.Println("Please enter the chat name: ")
	fmt.Fscan(reader, &name) // don't block here!! it should
	fmt.Printf("get name: %s \n", name)
	msg := ""
	buf := make([]byte, 1024)
	_t := time.NewTicker(time.Second * expireInterval) // each time server send, fresh the time
	defer _t.Stop()
	go func() {
		<-_t.C
		fmt.Println("server failed, disconnect!")
		return
	}()

	for {
		go func() {
			for {
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("error during reading from connection, err: %v \n", err)
					return
				}
				_t.Reset(time.Second * expireInterval) // refresh the ticker
				if string(buf[0:n]) != "1" {
					fmt.Println(string(buf[0:n]))
				}
			}
		}()
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error during reading message, err: %v \n", err)
		}
		msg = line
		// fmt.Fscan(reader, &msg)
		i := time.Now().Format("2006-01-02 15:04:05")
		conn.Write([]byte(fmt.Sprintf("%s\n\t%s: %s", i, name, msg)))
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
	return
}

func main() {
	conn := ChatClient()
	Sender(conn)
	Log("End!")
}
