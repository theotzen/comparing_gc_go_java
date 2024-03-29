package metrics

import (
	"fmt"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	alloc = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_alloc_bytesssss",
		Help: "Number of bytes allocated and still in use.",
	})
	totalAlloc = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_total_alloc_bytes",
		Help: "Total number of bytes allocated, even if freed.",
	})
	sys = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_sys_bytesss",
		Help: "Number of bytes obtained from system.",
	})
	heapAlloc = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_heap_alloc_bytesss",
		Help: "Number of heap bytes allocated and still in use.",
	})
	heapSys = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_heap_sys_bytesss",
		Help: "Number of heap bytes obtained from system.",
	})
	heapIdle = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_heap_idle_bytesss",
		Help: "Number of heap bytes waiting to be used.",
	})
	heapReleased = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_heap_released_bytessss",
		Help: "Number of heap bytes released to the OS.",
	})
	heapObjects = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_memstats_heap_objectssss",
		Help: "Number of allocated objects.",
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
	prometheus.MustRegister(alloc)
	prometheus.MustRegister(totalAlloc)
	prometheus.MustRegister(sys)
	prometheus.MustRegister(heapAlloc)
	prometheus.MustRegister(heapSys)
	prometheus.MustRegister(heapIdle)
	prometheus.MustRegister(heapReleased)
	prometheus.MustRegister(heapObjects)
	prometheus.MustRegister(ListsCreated)
}

func RecordMetrics() {
	var m runtime.MemStats
	fmt.Println("Starting collecting metrics")
	for {
		runtime.ReadMemStats(&m)
		gcCount.Set(float64(m.NumGC))
		alloc.Set(float64(m.Alloc))
		totalAlloc.Set(float64(m.TotalAlloc))
		sys.Set(float64(m.Sys))
		heapAlloc.Set(float64(m.HeapAlloc))
		heapSys.Set(float64(m.HeapSys))
		heapIdle.Set(float64(m.HeapIdle))
		heapReleased.Set(float64(m.HeapReleased))
		heapObjects.Set(float64(m.HeapObjects))
	}
}
