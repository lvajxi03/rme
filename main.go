package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/lvajxi03/rme/metrics"
)

func main() {
	int_metrics := make([]*metrics.Metric[int], 0)
	float_metrics := make([]*metrics.Metric[float64], 0)
	m1 := metrics.CreateMetric[int]("pierwsza", "Pierwsze rzeczy pierwsze", metrics.MetricCounter)
	hn, err := os.Hostname()
	if err != nil {
		hn = "NieznanyHost"
	}
	m1.AddLabel("nazwa_hosta", hn)
	m1.AddLabel("system", runtime.GOOS)
	m1.AddLabel("architektura", runtime.GOARCH)
	int_metrics = append(int_metrics, m1)
	m2 := metrics.CreateMetric[float64]("kurs_dolara", "Zmiana kursu dolara", metrics.MetricGauge, 0.01, 6.5)
	m2.AddLabel("nazwa_hosta", hn)
	m2.AddLabel("system", runtime.GOOS)
	m2.AddLabel("architektura", runtime.GOARCH)
	m2.AddLabel("bank", "NBP")
	float_metrics = append(float_metrics, m2)
	m3 := metrics.CreateMetric[float64]("kurs_dolara", "Zmiana kursu dolara", metrics.MetricGauge, 0.01, 6.5)
	m3.AddLabel("bank", "PKOBP")
	m3.AddLabel("nazwa_hosta", hn)
	m3.AddLabel("system", runtime.GOOS)
	m3.AddLabel("architektura", runtime.GOARCH)
	float_metrics = append(float_metrics, m3)
	m4 := metrics.CreateMetric[float64]("kurs_dolara", "Zmiana kursu dolara", metrics.MetricGauge, 0.01, 6.5)
	m4.AddLabel("bank", "BNPParibas")
	m4.AddLabel("nazwa_hosta", hn)
	m4.AddLabel("system", runtime.GOOS)
	m4.AddLabel("architektura", runtime.GOARCH)
	float_metrics = append(float_metrics, m4)
	m5 := metrics.CreateMetric[int]("rzut_kostka", "Rzut kostką", metrics.MetricGauge, 1, 6)
	m5.AddLabel("nazwa_hosta", hn)
	m5.AddLabel("system", runtime.GOOS)
	m5.AddLabel("architektura", runtime.GOARCH)
	m5.AddLabel("kostka", "szescienna")
	int_metrics = append(int_metrics, m5)
	m6 := metrics.CreateMetric[int]("rzut_kostka", "Rzut kostką", metrics.MetricGauge, 1, 8)
	m6.AddLabel("nazwa_hosta", hn)
	m6.AddLabel("system", runtime.GOOS)
	m6.AddLabel("architektura", runtime.GOARCH)
	int_metrics = append(int_metrics, m6)
	m6.AddLabel("kostka", "osmioscienna")
	m7 := metrics.CreateMetric[int]("rzut_kostka", "Rzut kostką", metrics.MetricGauge, 1, 12)
	m7.AddLabel("nazwa_hosta", hn)
	m7.AddLabel("system", runtime.GOOS)
	m7.AddLabel("architektura", runtime.GOARCH)
	m7.AddLabel("kostka", "dwunastoscienna")
	int_metrics = append(int_metrics, m7)
	router := gin.Default()
	router.GET("/metrics", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("%s\n%s", metrics.RenderAllMetrics(int_metrics), metrics.RenderAllMetrics(float_metrics)))
	})
	router.Run(":9101")
}
