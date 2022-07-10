package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"service_registry_and_discovery/config"
	"service_registry_and_discovery/stringService/endpoint"
	"service_registry_and_discovery/stringService/plugins"
	"service_registry_and_discovery/stringService/service"
	"service_registry_and_discovery/stringService/transport"
	"strconv"
	"syscall"

	"service_registry_and_discovery/discover"

	uuid "github.com/satori/go.uuid"
)

func main() {

	// 获取命令行参数
	var (
		servicePort = flag.Int("service.port", 10085, "service port")
		serviceHost = flag.String("service.host", "127.0.0.1", "service host")
		consulPort  = flag.Int("consul.port", 8500, "consul port")
		consulHost  = flag.String("consul.host", "127.0.0.1", "consul host")
		serviceName = flag.String("service.name", "string", "service name")
	)

	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)
	var discoveryClient discover.DiscoveryClient
	discoveryClient, err := discover.NewKitDiscoverClient(*consulHost, *consulPort)

	if err != nil {
		config.Logger.Println("Get Consul Client failed")
		os.Exit(-1)
	}

	var svc service.Service
	svc = service.StringService{}
	// add logging middleware
	svc = plugins.LoggingMiddleware(config.KitLogger)(svc)

	stringEndpoint := endpoint.MakeStringEndpoint(svc)
	healthEndpoint := endpoint.MakeHealthCheckEndpoint(svc)
	endpts := endpoint.StringEndpoints{
		StringEndpoint:      stringEndpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	r := transport.MakeHttpHandler(ctx, endpts, config.KitLogger)
	instanceId := *serviceName + "-" + uuid.NewV4().String()

	// http server
	go func() {
		config.Logger.Println("Http Server start at port:" + strconv.Itoa(*servicePort))
		if !discoveryClient.Register(*serviceName, instanceId, "/health", *serviceHost, *servicePort, nil, config.Logger) {
			config.Logger.Printf("string-service for service %s failed.", serviceName)
			os.Exit(-1)
		}
		handler := r
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	discoveryClient.DeRegister(instanceId, config.Logger)
	config.Logger.Println(error)
}
