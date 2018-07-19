package main

import (
	"log"
	"net/http"

	"go.opencensus.io/exporter/prometheus"
)

func main() {
	pe, err := prometheus.NewExporter(prometheus.Options{Namespace: "excavate"})
	if err != nil {
		log.Fatalf("Failed to create Prometheus exporter: %v", err)
	}
	view.RegisterExporter(pe)
	go func() {
		if err := http.ListenAndServe(":9888", pe); err != nil {
			log.Fatalf("Failed to create Prometheus exporter: %v", err)
		}
	}()
}
