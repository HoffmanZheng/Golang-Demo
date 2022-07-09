package discover

import (
	"log"
	"strconv"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

type KitDiscoverClient struct {
	Host   string
	Port   int
	client consul.Client
}

func NewKitDiscoverClient(consulHost string, consulPort int) (DiscoveryClient, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulHost + ":" + strconv.Itoa(consulPort)
	apiClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}
	client := consul.NewClient(apiClient)
	return &KitDiscoverClient{
		Host:   consulHost,
		Port:   consulPort,
		client: client,
	}, err
}

func (consulClient KitDiscoverClient) Register(serviceName, instanceId, healthCheckUrl, instanceHost string, instancePort int, meta map[string]string, logger *log.Logger) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID:      instanceId,
		Name:    serviceName,
		Address: instanceHost,
		Port:    instancePort,
		Meta:    meta,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			HTTP:                           "http://" + instanceHost + ":" + strconv.Itoa(instancePort) + healthCheckUrl,
			Interval:                       "15s",
		},
	}

	err := consulClient.client.Register(serviceRegistration)

	if err != nil {
		log.Println("Register Service Error!")
		return false
	}
	log.Println("Register Service Success!")
	return true
}

func (consulClient KitDiscoverClient) DeRegister(instanceId string, logger *log.Logger) bool {
	serviceRegistration := &api.AgentServiceRegistration{
		ID: instanceId,
	}
	err := consulClient.client.Deregister(serviceRegistration)
	if err != nil {
		logger.Println("Deregister Service Error!")
		return false
	}
	log.Println("Deregister Service Success!")
	return true
}

func (consulClient KitDiscoverClient) DiscoverServices(serviceName string, logger *log.Logger) []interface{} {
	entries, _, err := consulClient.client.Service(serviceName, "", false, nil)
	if err != nil {
		log.Println("Discover Service Error!")
		return nil
	}

	instances := make([]interface{}, len(entries))
	for i := 0; i < len(entries); i++ {
		instances[i] = entries[i].Service
	}
	return instances
}
