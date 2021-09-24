package discover

// 与 Consul 交互的客户端接口定义 ———— 方便替换具体的 Consul 客户端的实现方式
import (
	"log"
)

/**
* @param healthCheckUrl 健康检查地址
* @param meta 服务实例元数据
 */
type DiscoveryClient interface {
	Register(serviceName, instanceId, healthCheckUrl, instanceHost string,
		instancePort int, meta map[string]string, logger *log.Logger) bool

	DeRegister(instanceId string, logger *log.Logger) bool

	DiscoverServices(serviceName string, logger *log.Logger) []interface{}
}
