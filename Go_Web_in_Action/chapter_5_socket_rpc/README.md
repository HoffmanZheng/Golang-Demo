# Go Advanced Network Programming

### Go Socket Programming

1. Socket is a connection between application layer and transport layer in computer network, which make it convenient for application to transport data.

2. `net` package in Golang abstracted and encapsulated the socket connection process. Therefore connection with any protocal could be established easily by `net.Dial()` function. Server could also easily listen to specific IP and port by `net.Listen()`. See: [1.tcpClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/1.tcpClient.go) and [2.tcpServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/2.tcpServer.go)

3. A simple chat program could be developed using go socket, which includes heartbeat detaction, see: [3.socketChatClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/3.socketChatClient.go) and [4.socketChatServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/4.socketChatServer.go)

### Go RPC Programming

### Go Microservice