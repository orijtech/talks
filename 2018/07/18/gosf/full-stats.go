package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/zpages"
)

var MRequestsIn = stats.Int64("demo/requests", "The number of requests received", "1")
var MLatency = stats.Float64("demo/latency", "The number of seconds to process an entity", "ms")
var KeyMethod, _ = tag.NewKey("method")
var KeyIndex, _ = tag.NewKey("index")

func search(ctx context.Context, i int) {
	name := "search"
	if i%2 == 0 {
		name = "extract"
	}
	ctx, _ = tag.New(ctx, tag.Insert(KeyMethod, name))
	startTime := time.Now()
	defer func() {
		endTimeMs := float64(time.Since(startTime).Nanoseconds()) / 1e6
		// Record the stats here
		stats.Record(ctx, MRequestsIn.M(1), MLatency.M(endTimeMs))
	}()
	// Search logic below...
	sleep := time.Duration(rand.Intn(6700)) * time.Millisecond
	log.Printf("Sleeping for %s\n", sleep)
	<-time.After(sleep)
}

var LatencyView = &view.View{
	Measure: MLatency, Name: "demo/latency", Description: "Latency per call", TagKeys: []tag.Key{KeyMethod, KeyIndex},
	Aggregation: view.Distribution(0, 5, 10, 20, 50, 100, 200, 300, 500, 800),
}
var RequestsView = &view.View{
	Measure: MRequestsIn, Name: "demo/requests", Description: "Number of requests",
	Aggregation: view.Count(),
}

func main() {
	if err := view.Register(LatencyView, RequestsView); err != nil {
		log.Fatalf("Failed to register custom views: %v", err)
	}

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

	// Enable zPages
	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		if err := http.ListenAndServe(":7788", mux); err != nil {
			log.Fatalf("Failed to serve zPages: %v", err)
		}
	}()

	n := 12
	doneChan := make(chan bool)

	for i := 0; i < n; i++ {
		go func(id int) {
			defer func() {
				doneChan <- true
			}()
			max := rand.Intn(1e7)
			ctx, _ := tag.New(context.Background(), tag.Insert(KeyIndex, fmt.Sprintf("%d", id)))
			for i := 0; i < max; i++ {
				search(ctx, i)
			}
		}(i)
	}

	for i := 0; i < n; i++ {
		<-doneChan
	}

	log.Printf("Done!")
}
