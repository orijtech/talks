package main

import (
	"context"
        "time"

        "go.opencensus.io/stats"
        "go.opencensus.io/stats/view"
        "go.opencensus.io/tag"

	"./demo-stats"
)

func main() {
	view.RegisterExporter(new(custom.CustomExporter))
	view.SetReportingPeriod(time.Millisecond)
	_ = view.Register(custom.RequestsView)
	ctx, _ := tag.New(context.Background(), tag.Insert(custom.KeyMethod, "search"))
	stats.Record(ctx, custom.MRequestsIn.M(23))
	<-time.After(2 * time.Millisecond)
}
