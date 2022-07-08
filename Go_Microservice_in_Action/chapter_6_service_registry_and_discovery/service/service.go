package service

import (
	"context"
	"errors"
	"service_registry_and_discovery/config"
	"service_registry_and_discovery/discover"
)

type Service interface {
	HealthCheck() bool
	SayHello() string
	DiscoveryService(ctx context.Context, serviceName string) ([]interface{}, error)
}

var ErrNotServiceInstances = errors.New("instances are not existed")

type DiscoverServiceImpl struct {
	discoverClient discover.DiscoveryClient
}

func NewDiscoveryServiceImpl(discoverClient discover.DiscoveryClient) Service {
	return &DiscoverServiceImpl{discoverClient: discoverClient}
}

// 从 consul 中根据服务名获取对应的服务实例信息列表
func (impl *DiscoverServiceImpl) DiscoveryService(ctx context.Context,
	serviceName string) ([]interface{}, error) {
	instances := impl.discoverClient.DiscoverServices(serviceName, config.Logger)
	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}

func (impl *DiscoverServiceImpl) SayHello() string {
	return "hello world!"
}

func (impl *DiscoverServiceImpl) HealthCheck() bool {
	return true
}
