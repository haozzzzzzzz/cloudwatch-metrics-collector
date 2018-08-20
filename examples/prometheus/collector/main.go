package main

import (
	"fmt"

	"github.com/haozzzzzzzz/cloudwatch-metrics-collector/collector"
	"github.com/haozzzzzzzz/go-rapid-development/utils/uos"
	"github.com/sirupsen/logrus"
)

func main() {
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
