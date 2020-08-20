package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	labels = map[string]string{"group": "test_worker__pdname__test_service"}
	fx     = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "fx_step",
		Help:        "f(x) with x",
		ConstLabels: labels,
	})
)

func main() {
	prometheus.MustRegister(fx)
	fx.Set(0.0)

	go f_x()
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func f_x() {
	// step is time in milisecond
	x, step := 0.0, 100
	for {
		fx.Set(change_fx(x))
		time.Sleep(time.Duration(step) * time.Millisecond)
		x += 1
	}
}

func change_fx(x float64) float64 {
	if x >= 0 && x <= 1000 {
		return 1.0
	} else if x > 1000 && x <= 1650 {
		return 23.0
	}

	return 0.0
}
