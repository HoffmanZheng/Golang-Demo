package transport

import (
	"context"
	"encoding/json"
	"endpoint"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// 声明对外暴露的 HTTP 服务
// 将 endpoint 与对应的 HTTP 路径进行绑定

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func MakeHttpHandler(ctx context.Context, endpoints endpoint.DiscoveryEndPoints,
	logger log.Logger) http.Handler {
	r := mux.NewRouter()
	// 定义处理处理器
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	// SayHello 接口
	r.Methods("GET").Path("/sayHello").Handler(kithttp.NewServer{
		endpoints.SayHelloEndPoint,
		decodeSayHelloRequest,
		encodeJsonResponse,
		options...,
	})
	// 服务发现接口
	r.Methodes("GET").Path("/discovery").Handler(kithttp.NewServer{
		endendpoints.DiscoveryEndpoint,
		decodeDiscoveryRequest,
		encodeJsonResponse,
		options...,
	})
	// create health check handler
	r.Methods("GET").Path("/health").Handler(kithttp.NewServer{
		enendpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeJsonResponse,
		options...,
	})
	return
}

func decodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.SayHelloRequest{}, nil
}

func decodeDiscoveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	serviceName := r.URL.Query().Get("servicename")
	if serviceName == "" {
		return nil, ErrorBadRequest
	}
	return endpoint.DiscoveryRequest{
		serviceName: serviceName,
	}, nil
}

func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.HealthRequest{}, nil
}

func encodeJsonResponse(ctx context.Context, w http.ResponseWriter, 
	response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch err {
	default: 
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{} {
		"error" : err.Error(),
	})
}
