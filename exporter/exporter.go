package exporter

import promexporter "log-generator/exporter/prometheus"

type Exporter interface {
	Export(direction string, value float64)
}

func NewExporter(inbuckets float64, outbuckets float64) Exporter {
	return promexporter.NewPromExporter(inbuckets, outbuckets)
}
