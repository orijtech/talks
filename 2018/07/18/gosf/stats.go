package main

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

var MRequestsIn = stats.Int64("demo/requests", "The number of requests received", "1")
var MLatency = stats.Float64("demo/latency", "The number of seconds to process an entity", "ms")
var KeyMethod, _ = tag.NewKey("method")

func search(w http.ResponseWriter, r *http.Request) {
	ctx, _ = tag.New(ctx, tag.Insert(KeyMethod, "search"))
	startTime := time.Now()
	defer func() {
		endTimeMs := float64(time.Since(startTime).Nanoseconds()) / 1e6
		// Record the stats here
		stats.Record(ctx, MRequestsIn.M(1), MLatency.M(endTimeMs))
	}()
	// Search logic below...
}
