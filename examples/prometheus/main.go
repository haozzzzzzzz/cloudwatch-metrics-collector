package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/cloudwatch-metrics-collector/collector"
	"github.com/haozzzzzzzz/go-rapid-development/utils/uos"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func service() {
	engine := gin.Default()
	engine.GET("/metrics", func(context *gin.Context) {
		promhttp.Handler().ServeHTTP(context.Writer, context.Request)
	})
	engine.Run(":18111")
}

func main() {
	go func() {
		service()
	}()

	var err error
	promMetrics := collector.PrometheusMetrics{
		Target:          "http://127.0.0.1:18111/metrics",
		NamespacePrefix: "cloudwatch_",
	}
	metrics, err := promMetrics.Pull()
	if nil != err {
		logrus.Errorf("pull metrics failed. error: %s.", err)
		return
	}

	fmt.Println(metrics)

	uos.WaitSignal()
}
