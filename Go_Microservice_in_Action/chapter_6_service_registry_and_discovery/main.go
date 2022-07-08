package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"service_registry_and_discovery/config"
	"service_registry_and_discovery/discover"
	"service_registry_and_discovery/endpoint"
	"service_registry_and_discovery/service"
	"service_registry_and_discovery/transport"
	"strconv"
	"syscall"

	uuid "github.com/satori/go.uuid"
)

func main() {
	//  consul　启动命令：consul agent -dev
	// 从命令行读取相关参数，没有时使用默认值
	var (
		servicePort = flag.Int("service.port", 10086, "service port")
		serviceHost = flag.String("service.host", "127.0.0.1", "service host")
		serviceName = flag.String("service.name", "SayHello", "service name")
		consulPort  = flag.Int("consul.port", 8500, "consul port")
		consulHost  = flag.String("consul.host", "127.0.0.1", "consul host")
	)
	flag.Parse()
	ctx := context.Background()
	errChan := make(chan error)

	// TODO: 没有初始化服务发现客户端
	var discoveryClient discover.DiscoveryClient
	discoveryClient = discover.NewMyDiscoverClient(*consulHost, *consulPort) // 省略了获取服务发现客户端失败的处理
	var svc = service.NewDiscoveryServiceImpl(discoveryClient)
	sayHelloEndPoint := endpoint.MakeSayHelloEndPoint(svc)
	discoveryEndpoint := endpoint.MakeDiscoveryEndPoint(svc)
	healthEndpoint := endpoint.MakeHealthCheckEndPoint(svc)
	endpts := endpoint.DiscoveryEndPoints{
		SayHelloEndPoint:    sayHelloEndPoint,
		DiscoveryEndPoint:   discoveryEndpoint,
		HealthCheckEndPoint: healthEndpoint,
	}
	// 创建 http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, config.KitLogger)
	instanceId := *serviceName + "-" + uuid.NewV4().String()

	// 启动 http server
	go func() {
		config.Logger.Println("HttpServer start at port:" + strconv.Itoa(*servicePort))
		if !discoveryClient.Register(*serviceName, instanceId, "/health",
			*serviceHost, *servicePort, nil, config.Logger) {
			config.Logger.Printf("string-service for service %s is failed", serviceName)
			os.Exit(-1)
		}
		handler := r
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), handler)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	discoveryClient.DeRegister(instanceId, config.Logger)
	config.Logger.Println(error)
}
