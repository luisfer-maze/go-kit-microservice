package metrics

import "github.com/go-kit/kit/metrics/prometheus"
import stdprometheus "github.com/prometheus/client_golang/prometheus"

var (
	RequestTotalCount = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name:      "request_total",
		Namespace: "hello_service",
		Help:      "Total number of requests.",
	}, []string{"request", "success"})
)
