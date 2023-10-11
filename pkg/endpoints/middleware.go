package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/luisfer-maze/go-kit-microservice/pkg/metrics"
	"time"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				level.Info(logger).Log("took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func metricsMiddleware(method string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func() {
				success := "true"
				if err != nil {
					success = "false"
				}
				metrics.RequestTotalCount.With([]string{"request", method, "success", success}...).Add(1)
			}()

			return next(ctx, request)
		}
	}
}
