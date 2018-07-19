package main

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

type customPrintExporter struct{}
func (ce *customPrintExporter) ExportView(vd *view.Data) { fmt.Printf("Send me to your metrics backend:\n%+v\n", vd)}

func main() {
	view.RegisterExporter(new(customPrintExporter)); view.SetReportingPeriod(time.Millisecond)
	_ = view.Register(RequestsView)
	ctx, _ := tag.New(context.Background(), tag.Insert(KeyMethod, "search"))
	stats.Record(ctx, MRequestsIn.M(23))
	<-time.After(2 * time.Millisecond)
}

var KeyMethod, _ = tag.NewKey("method")
var MRequestsIn, RequestsView = stats.Int64("demo/requests", "The number of requests received", "1"), &view.View{Measure: MRequestsIn, Name: "demo/requests", Description: "Number of requests", Aggregation: view.Count(), TagKeys: []tag.Key{KeyMethod}}
