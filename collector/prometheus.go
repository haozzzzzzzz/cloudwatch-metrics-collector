package collector

import (
	"github.com/haozzzzzzzz/cloudwatch-metrics-collector/metric"
)

// prometheus metrics collector
type PrometheusMetrics struct {
	Target          string
	NamespacePrefix string // 匹配namespace前缀
}

func (m *PrometheusMetrics) Pull() (metrics []*metric.Metric, err error) {
	metrics = make([]*metric.Metric, 0)
	return
}
