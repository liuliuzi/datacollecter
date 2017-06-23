package api
import (
	"k8s.io/heapster/metrics/core"
	"time"
)

type MetricsService struct {
	MetricsCache   map[string]Metric
	LiveSeconds    time.Duration
	Mode           int
	InfluxdbSink   core.DataSink
}

type Metric struct {
	MetricValue string
	Timestamp   time.Time
}



