# Service Registry and Discovery

### Fundamental

1. Microservice registers it's metadata to centered component, and maintain heartbeat with it. The client would obtain the instance list info from centered component before send request to it's target microservice.

2. Consul is an out-of-the-box service discovery component based on the Raft algorithm, interactive client interface defined in [discover/discoverClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/discover/discoverClient.go)

3. The project interface defined in [service/service.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/service/service.go), which could fetch the instance list from consul by `DiscoveryService`.

4. [endpoint](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/endpoint) layer is charge of converting the request into arguments which could be processed by [service](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/service) interface, and encapsulate the response to [transport](https://github.com/HoffmanZheng/Golang-Demo/tree/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/transport) layer.

### Demo: Interact with Consul

1. In order to accomplish the service registry and discovery, we could call the HTTP API of Consul directly, see: [discover/myDiscoverClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/discover/myDiscoverClient.go)

2. [Go-kit](https://github.com/go-kit/kit) framework provides the implementation of some common service registry center, such as Consul, Zookeeper, Etcd, Eureka. It's quite convenient for developer to use, see: [discover/kitDiscoverClient.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Microservice_in_Action/chapter_6_service_registry_and_discovery/discover/kitDiscoverClient.go)