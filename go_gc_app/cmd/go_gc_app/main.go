package main

import (
	"go_gc_app/internal/metrics"
	"go_gc_app/internal/utils"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go metrics.RecordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)

	for {
		list := utils.GenerateList(1000000)
		time.Sleep(time.Millisecond * 100)
		_ = list
	}
}
