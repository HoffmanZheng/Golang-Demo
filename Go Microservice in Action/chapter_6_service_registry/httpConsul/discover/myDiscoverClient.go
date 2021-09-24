package discover

type InstanceInfo struct {
	Id                string                     `json:"Id"`
	Name              string                     `json:"Name"`
	Service           string                     `json:"Service,ommitempty"`
	Tags              []string                   `json:"tags,omitemptv"`
	Address           string                     `json:"Address"`
	Port              int                        `json:"Port"`
	Meta              map[string]string          `json:"Meta,omitempty"`
	EnableTagOverride bool                       `json:"EnableTagOverride"`
	Check             `json:"Check,omitempty"`   // 健康检查相关配置
	Weight            `json:"Weights,omitempty"` // 权重
}

type Check struct {
	DeregisterCriticalServiceAfter string   `json:"DeregisterCriticalServiceAfter"` // 多久之后注销服务
	Args                           []string `json:"Args,omitempty"`                 // 请求参数
	HTTP                           string   `json:"HTTP"`                           // 健康检查地址
	Interval                       string   `json:"Interval,omitempty"`             // Consul 主动检查间隔
	TTL                            string   `json:"TTL,omitempty"`                  // 服务实例主动维持心跳间隔，与Interval只存其一
}

type Weights struct {
	Passing int `json:"Passing"`
	Warning int `json:"Warning"`
}
