package promexporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PromExporter struct {
	ingestion prometheus.Histogram
	query     prometheus.Histogram
}

func init() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":9090", nil)
	log.Println("Enabled prometheus /metrics endpoint on port :9090")
}

func NewPromExporter(inbuckets float64, outbuckets float64) PromExporter {
	ingestion := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "output",
			Name:      "ingestion_latency",
			Help:      "Bucket used to measure output latency.",
			Buckets:   prometheus.ExponentialBucketsRange(300, outbuckets*1000, 10),
		})

	query := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "query",
			Name:      "query_latency",
			Help:      "Bucket used to measure query latency",
			Buckets:   prometheus.ExponentialBucketsRange(300, inbuckets*1000, 10),
		},
	)

	prometheus.MustRegister(ingestion)
	prometheus.MustRegister(query)

	return PromExporter{
		ingestion: ingestion,
		query:     query,
	}
}

func (pe PromExporter) Export(direction string, value float64) {
	switch direction {
	case "out":
		pe.ingestion.Observe(value)
	case "in":
		pe.query.Observe(value)
	}
}
