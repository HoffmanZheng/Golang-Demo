# Remote Procedure Call RPC

### Principle and Implementation

1. Though HTTP could also provide remote call between microservices, RPC protocal is relatively more flexible and could be customized.

2. A client stub is responsible for conversion (marshalling) of parameters used in a function call and deconversion of results passed from the server after execution of the function. A server stub is responsible for deconversion of parameters passed by the client and conversion of the results after the execution of the function.

3. Both client and server should be aware of message format, and a standard is required for the communication between the heterogeneous systems.

4. 

5. 

### Golang Native RPC



### High Performance gRPC