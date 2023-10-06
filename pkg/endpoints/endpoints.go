package endpoints

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/luisfer-maze/go-kit-microservice/model"
	"github.com/luisfer-maze/go-kit-microservice/pkg/service"
)

type Endpoints struct {
	SayHello endpoint.Endpoint
}

func MakeEndpoints(service service.Service) Endpoints {
	return Endpoints{
		SayHello: makeSayHelloEndpoint(service),
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
