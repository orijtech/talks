package main

import (
	"log"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

func main() {
	sd, err := stackdriver.NewExporter(stackdriver.Options{ProjectID: "census-demos"})
	if err != nil {
		log.Fatalf("Failed to register Stackdriver Trace exporter: %v", err)
	}
	trace.RegisterExporter(sd)
}
