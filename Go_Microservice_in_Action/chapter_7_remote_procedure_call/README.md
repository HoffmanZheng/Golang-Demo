# Remote Procedure Call RPC

### Principle and Implementation

1. Though HTTP could also provide remote call between microservices, RPC protocal is relatively more flexible and could be customized.

2. A client stub is responsible for conversion (marshalling) of parameters used in a function call and deconversion of results passed from the server after execution of the function. A server stub is responsible for deconversion of parameters passed by the client and conversion of the results after the execution of the function.

3. Both client and server should be aware of message format, and a standard is required for the communication between the heterogeneous systems.

4. Though the Highly customized private communication protocol lacks generality comparing to the public protocol, it could improve the performance and flexibility. Normally protocol contains payload and control command, the latter is generally heartbeat and function command.

5. If the RPC will retry when failed, the idempotent property of function should be ensured.

### Golang Native RPC

1. Golang native rpc lib `net/rpc` provides the access to object methods by `Register`, see: [/basic/server.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Microservice_in_Action/chapter_7_remote_procedure_call/basic/server.go). It would obtain all RPC suitable methods from object, and save them into a map (server stub). When receive a request from client, server would parse(through specified protocol) and get the method from the stub, and call the corresponding method by reflection.

2. The client could call synchronizely or asynchronizely, async call is implemented by goroutine, and the result is received(waited) by done channel, see: [/basic/client.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Microservice_in_Action/chapter_7_remote_procedure_call/basic/client.go). Each RPC request will generate a Call object(include method, param, return value), and be saved in a map (client stub). When client receive the response from server, it will fetch stub from map and deserialize the return value.

3. Every RPC request creates request and response struct repeatedly, which leads to some GC pressure. In order to reuse the request and response struct, an object pool was created, object will be freed after they're used.

4. Generally, Golang native RPC component is a basic version RPC framework, concise but high scalable. It only implements the basic network communication of RPC, other functions like registry and discovery,  timeout fuse are still missing. Thus it is not out-of-the-box, an inofficial intensified version could be found in [rpcx](https://github.com/smallnest/rpcx).

### High Performance gRPC

1. gRPC is an open-source, high performance, universal RPC framework, it supports multiple languages. Based on the HTTP/2 standard, it could realize the headers compression, request reusage, and client application communication.

2. Protocol compiler: `protoc --proto_path=pb --go_out=plugins=grpc:pb  --go_opt=paths=source_relative string.proto`. Server and client both depend on the generated [proto files](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_7_remote_procedure_call/pb), a gRPC demo located at [/grpc](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_7_remote_procedure_call/grpc)

3. Using stream transport could improve the performance,