package main

import (
	"log"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/haozzzzzzzz/go-rapid-development/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var err error
	engine := gin.Default()

	testCounter, err := metrics.NewCounter(
		"cloudwatch_test_prometheus",
		"main",
		"test_counter",
		"测试",
	)
	if nil != err {
		log.Fatal(err)
		return
	}

	go func() {
		for {
			testCounter.Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	engine.GET("/metrics", func(context *gin.Context) {
		promhttp.Handler().ServeHTTP(context.Writer, context.Request)
	})
	engine.Run(":18111")
}
