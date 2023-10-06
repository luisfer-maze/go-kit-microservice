package transport

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/luisfer-maze/go-kit-microservice/model"
	"github.com/luisfer-maze/go-kit-microservice/pkg/endpoints"
)

type grpcServer struct {
	sayHello grpc.Handler
}

func NewGRPCServer(endpoint endpoints.Endpoints) model.HelloServiceServer {
	return grpcServer{
		sayHello: grpc.NewServer(endpoint.SayHello, decodeHelloRequest, encodeHelloResponse),
	}
}

func (g grpcServer) SayHello(ctx context.Context, request *model.HelloRequest) (*model.HelloResponse, error) {
	_, resp, err := g.sayHello.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*model.HelloResponse), err
}

func decodeHelloRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*model.HelloRequest)
	return &model.HelloRequest{
		NameToSalute: req.NameToSalute,
	}, nil
}

func encodeHelloResponse(ctx context.Context, response interface{}) (interface{}, error) {
	res := response.(*model.HelloResponse)

	return &model.HelloResponse{
		Salute: res.Salute,
	}, nil
}
