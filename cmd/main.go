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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := log.NewJSONLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowAll())
	level.Info(logger).Log("msg", "starting app")

	//saver := storage.NewSaluteSaver(logger)
	saver := storage.NewAnotherStorage(logger)
	svc := service.NewService(saver, logger)
	end := endpoints.MakeEndpoints(logger, svc)
	grpcServer := transport.NewGRPCServer(end)

	cancelInterrupt := make(chan struct{})

	var g group.Group
	{
		grpcListener, err := net.Listen("tcp", ":9027")
		if err != nil {
			level.Error(logger).Log("err", "TCP error")
		}

		baseServer := grpc.NewServer()

		g.Add(func() error {
			level.Info(logger).Log("msg", "starting grpc server")
			model.RegisterHelloServiceServer(baseServer, grpcServer)

			return baseServer.Serve(grpcListener)
		}, func(error) {
			baseServer.GracefulStop()
			grpcListener.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp", ":8080")

		if err != nil {
			level.Error(logger).Log("err", err.Error())
			os.Exit(1)
		}

		g.Add(func() error {
			level.Info(logger).Log("msg", "starting metrics server")
			http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

			return http.Serve(httpListener, http.DefaultServeMux)
		}, func(err error) {
			httpListener.Close()
		})
	}
	{
		g.Add(func() error {
			c := make(chan os.Signal)

			signal.Notify(c, syscall.SIGINT, syscall.SIGALRM, syscall.SIGTERM)

			select {
			case sig := <-c:
				err := fmt.Errorf("signal received %s", sig)
				level.Error(logger).Log("err", err.Error())
				return err
			case <-cancelInterrupt:
				return nil
			}
		}, func(err error) {
		})

	}

	fmt.Println(g.Run())
}
