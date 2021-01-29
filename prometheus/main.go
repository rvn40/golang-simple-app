package main

import (
  "log"
  "net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpReqCounters = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "webapp_counter",
		Help: "Root endpoint total http requests counter.",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(httpReqCounters)
}


func main() {
  http.Handle("/", promhttp.Handler())
	httpReqCounters.Inc()

  log.Fatal(http.ListenAndServe("0.0.0.0:9110", nil))
}
