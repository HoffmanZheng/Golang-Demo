# Go Advanced Network Programming

### Go Socket Programming

1. Socket is a connection between application layer and transport layer in computer network, which make it convenient for application to transport data.

2. `net` package in Golang abstracted and encapsulated the socket connection process. Therefore connection with any protocal could be established easily by `net.Dial()` function. Server could also easily listen to specific IP and port by `net.Listen()`. See: [1.tcpClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/1.tcpClient.go) and [2.tcpServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/2.tcpServer.go)

3. A simple chat program could be developed using go socket, which includes heartbeat detaction, see: [3.socketChatClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/3.socketChatClient.go) and [4.socketChatServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/4.socketChatServer.go)

### Go RPC Programming

1. `net/rpc` package uses Encoder and Decoder object in `encoding/gob` package, which encodes and decodes in GOB format. Cross-language calls are impossible in this way, due to the missing support of GOB encoding format.

2. A json format RPC could be impremented by using `net/rpc/jsonrpc`, see: [5.jsonRpcClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/5.jsonRpcClient.go) and [6.jsonRpcServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/6.jsonRpcServer.go)

3. A serve object should be defined and registered in rpc `rpc.Register(serveObj)`, the format of rpc method looks like as below:
```golang
func (t *T) MethodName(argType T1, replyType *T2) error
```

### Go Microservice

1. Microservice architecture is an evolution of SOA(Service-Oriented Architecture), which devides single application into several microservices. Microservice should obey single responsibility and service-oriented principle, encapsulate functions and provide services.

2. gRPC is a open-source, cross-platform, high performance RPC framework, it makes it easier to call a remote service just like call local method. It's easier to establish distributed serives by using gRPC, demo of gRPC, see: [7.gRpcClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/7.gRpcClient.go) and [8.gRpcServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_5_socket_rpc/8.gRpcServer.go)