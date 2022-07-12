# Remote Procedure Call RPC

### Principle and Implementation

1. Though HTTP could also provide remote call between microservices, RPC protocal is relatively more flexible and could be customized.

2. A client stub is responsible for conversion (marshalling) of parameters used in a function call and deconversion of results passed from the server after execution of the function. A server stub is responsible for deconversion of parameters passed by the client and conversion of the results after the execution of the function.

3. Both client and server should be aware of message format, and a standard is required for the communication between the heterogeneous systems.

4. Though the Highly customized private communication protocol lacks generality comparing to the public protocol, it could improve the performance and flexibility. Normally protocol contains payload and control command, the latter is generally heartbeat and function command.

5. If the RPC will retry when failed, the idempotent property of function should be ensured.

### Golang Native RPC

1. Golang native rpc lib `net/rpc` provides the access to object methods by `Register`, see: [/basic/server.go](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_7_remote_precedure_call/basic/server.go). It would obtain all RPC suitable methods from object, and save them into a map (stub).

2. The client could call synchronizely or asynchronizely, async call is implemented by goroutine, and the result is received(waited) by done channel, see: [/basic/client.go](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_7_remote_precedure_call/basic/client.go)

3. 

### High Performance gRPC