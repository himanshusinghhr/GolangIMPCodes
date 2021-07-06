package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"io/ioutil"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

var (
	sidekiqProcessed = setGaugeMetric("sidekiq_processed", "Sidekiq Processed", "lable", "lablevalue")
)

func setGaugeMetric(name string, help string, label string, labelvalue string) (prometheusGauge prometheus.Gauge) {
	var (
		gaugeMetric = prometheus.NewGauge(prometheus.GaugeOpts{
			Name:        name,
			Help:        help,
			ConstLabels: prometheus.Labels{label: labelvalue},
		})
	)

	return gaugeMetric
}

func getSidekiqProcessed() (sidekiq float64) {
	body := getContent("http://example.com/sidekiq/stats")
	processed := gjson.Get(body, "sidekiq.processed")
	conv, err := strconv.ParseFloat(processed.String(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return conv
}

func getContent(url string) (body string) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	res, err := httpClient.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func init() {
	prometheus.MustRegister(sidekiqProcessed)
}

func main() {
	sidekiqProcessed.Set(getSidekiqProcessed())

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
