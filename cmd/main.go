package main

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/luisfer-maze/go-kit-microservice/model"
	"github.com/luisfer-maze/go-kit-microservice/pkg/endpoints"
	"github.com/luisfer-maze/go-kit-microservice/pkg/service"
	"github.com/luisfer-maze/go-kit-microservice/pkg/storage"
	"github.com/luisfer-maze/go-kit-microservice/pkg/transport"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {

	logger := log.NewJSONLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowAll())
	level.Info(logger).Log("msg", "starting app")

	//saver := storage.NewSaluteSaver(logger)
	saver := storage.NewAnotherStorage(logger)
	svc := service.NewService(saver, logger)
	end := endpoints.MakeEndpoints(svc)
	grpcServer := transport.NewGRPCServer(end)

	var g group.Group
	{
		grpcListener, err := net.Listen("tcp", ":9027")
		if err != nil {
			level.Error(logger).Log("err", "TCP error")
		}

		baseServer := grpc.NewServer()

		g.Add(func() error {
			model.RegisterHelloServiceServer(baseServer, grpcServer)

			return baseServer.Serve(grpcListener)
		}, func(error) {
			baseServer.GracefulStop()
			grpcListener.Close()
		})
	}

	fmt.Println(g.Run())
}
