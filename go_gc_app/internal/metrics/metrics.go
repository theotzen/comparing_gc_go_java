package metrics

import (
	"fmt"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalAlloc = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_total_alloc_bytes",
		Help: "Total number of bytes allocated, even if freed.",
	})
	gcCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_gc_operations_total",
		Help: "Total number of completed GC cycles",
	})
	ListsCreated = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_linked_list_created_total",
		Help: "Total linked list created",
	})
)

func init() {
	prometheus.MustRegister(gcCount)
	prometheus.MustRegister(totalAlloc)
	prometheus.MustRegister(ListsCreated)
}

func RecordMetrics() {
	var m runtime.MemStats
	fmt.Println("Starting collecting metrics")
	for {
		runtime.ReadMemStats(&m)
		gcCount.Set(float64(m.NumGC))
		totalAlloc.Set(float64(m.TotalAlloc))
	}
}
