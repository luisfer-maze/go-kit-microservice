package endpoints

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/luisfer-maze/go-kit-microservice/model"
	"github.com/luisfer-maze/go-kit-microservice/pkg/service"
)

type Endpoints struct {
	SayHello endpoint.Endpoint
}

func MakeEndpoints(logger log.Logger, service service.Service) Endpoints {
	return Endpoints{
		SayHello: makeEndpointWithMiddleware(makeSayHelloEndpoint(service), logger, "SayHello"),
	}
}

func makeSayHelloEndpoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*model.HelloRequest)
		salute := service.ProcessSalute(fmt.Sprintf("Hello %s!!!", req.NameToSalute))
		return &model.HelloResponse{
			Salute: salute,
		}, nil
	}
}

func makeEndpointWithMiddleware(e endpoint.Endpoint, logger log.Logger, method string) endpoint.Endpoint {
	modEndpoint := e
	modEndpoint = metricsMiddleware(method)(modEndpoint)
	return loggingMiddleware(logger)(modEndpoint)
}
