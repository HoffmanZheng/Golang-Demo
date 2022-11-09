package discover

import (
	"github.com/hashicorp/consul/api/watch"
	"log"
	"strconv"
	"sync"

	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

type KitDiscoverClient struct {
	Host         string
	Port         int
	client       consul.Client
	config       *api.Config
	mutex        sync.Mutex  // atomic lock
	instancesMap sync.Map    // cache
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
		config: consulConfig,
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

	// try to load from cache
	instanceList, ok := consulClient.instancesMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	}

	consulClient.mutex.Lock()
	instanceList, ok = consulClient.instancesMap.Load(serviceName)
	if ok {
		return instanceList.([]interface{})
	} else {
		go func() {
			params := make(map[string]interface{})
			params["type"] = "service"
			params["service"] = serviceName
			plan, _ := watch.Parse(params)
			plan.Handler = func(u uint64, i interface{}) {
				if i == nil {
					return
				}
				v, ok := i.([]*api.ServiceEntry)
				if !ok {
					return // data error, ignore
				}
				if len(v) == 0 {   // no instance online
					consulClient.instancesMap.Store(serviceName, []interface{}{})
				}
				var healthServices []interface{}
				for _, service := range v {
					if service.Checks.AggregatedStatus() == api.HealthPassing {
						healthServices = append(healthServices, service.Service)
					}
				} // update the cache once instance list varies
				log.Println("fetch the service list from watching: ", healthServices)
				consulClient.instancesMap.Store(serviceName, healthServices)
			}
			defer plan.Stop()
			plan.Run(consulClient.config.Address)
		}()
	}
	defer consulClient.mutex.Unlock()

	entries, _, err := consulClient.client.Service(serviceName, "", false, nil)
	if err != nil {
		consulClient.instancesMap.Store(serviceName, []interface{}{})
		log.Println("Discover Service Error!")
		return nil
	}

	instances := make([]interface{}, len(entries))
	for i := 0; i < len(entries); i++ {
		instances[i] = entries[i].Service
	}
	consulClient.instancesMap.Store(serviceName, instances)
	return instances
}
