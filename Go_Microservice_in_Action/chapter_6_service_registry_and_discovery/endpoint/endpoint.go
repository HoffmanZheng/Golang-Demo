package endpoint

// 将请求转化为 service 接口可以处理的参数，
// 并将处理结果封装为 response 返回给 transport 层
import (
	"context"
	"service_registry_and_discovery/service"

	"github.com/go-kit/kit/endpoint"
)

type DiscoveryEndPoints struct {
	SayHelloEndPoint    endpoint.Endpoint
	DiscoveryEndPoint   endpoint.Endpoint
	HealthCheckEndPoint endpoint.Endpoint
}

type SayHelloRequest struct{}

type SayHelloResponse struct {
	Message string `json:"message"`
}

func MakeSayHelloEndPoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{},
		err error) {
		message := svc.SayHello()
		return SayHelloResponse{
			Message: message,
		}, nil
	}
}

type DiscoveryRequest struct {
	ServiceName string
}

type DiscoveryResponse struct {
	Instances []interface{} `json:"instances"`
	Error     string        `json:"error"`
}

func MakeDiscoveryEndPoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{},
		err error) {
		req := request.(DiscoveryRequest)
		instances, err := svc.DiscoveryService(ctx, req.ServiceName)
		var errString = ""
		if err != nil {
			errString = err.Error()
		}
		return DiscoveryResponse{
			Instances: instances,
			Error:     errString,
		}, nil
	}
}

type HealthRequest struct{}

type HealthResponse struct {
	Status bool `json:"status"`
}

func MakeHealthCheckEndPoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{},
		err error) {
		status := svc.HealthCheck()
		return HealthResponse{
			Status: status,
		}, nil
	}
}
