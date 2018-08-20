package collector

// prometheus metrics collector
type PrometheusMetrics struct {
	Target string
	NamespacePrefix string // 匹配namespace前缀
}
